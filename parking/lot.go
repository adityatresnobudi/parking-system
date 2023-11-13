package parking

import (
	"errors"

	"github.com/adityatresnobudi/parking-system/entity"
)

var (
	ErrUnrecognizedParkingTicket = errors.New("unrecognized parking ticket")
	ErrUnavailablePosition       = errors.New("no available position")
	ErrParkedCarTwice            = errors.New("car already inside")
)

type Lot struct {
	parkedCars  map[string]*entity.Car
	subscribers []Subscriber
	capacity    int
}

type Subscriber interface {
	NotifyLotIsFull(*Lot)
	NotifyLotIsNotFull(*Lot)
}

type LotStatus struct {
	freeSpace  int
	parkedCars map[string]*entity.Car
}

func NewLot(capacity int) *Lot {
	return &Lot{
		parkedCars:  make(map[string]*entity.Car),
		subscribers: make([]Subscriber, 0),
		capacity:    capacity,
	}
}

func (l *Lot) Park(car *entity.Car) (*entity.Ticket, error) {
	if !l.IsNotFull() {
		return nil, ErrUnavailablePosition
	}
	if l.IsCarParked(car) {
		return nil, ErrParkedCarTwice
	}
	newTicket := entity.NewTicket()
	l.parkedCars[newTicket.ID] = car
	if !l.IsNotFull() {
		l.notifySubscibersFull()
	}
	return &newTicket, nil
}

func (l *Lot) UnPark(ticket *entity.Ticket) (*entity.Car, error) {
	if _, ok := l.parkedCars[ticket.ID]; !ok {
		return nil, ErrUnrecognizedParkingTicket
	}
	unparkedCar := l.parkedCars[ticket.ID]
	delete(l.parkedCars, ticket.ID)
	l.notifySubscibersNotFull()
	return unparkedCar, nil
}

func (l *Lot) IsCarParked(car *entity.Car) bool {
	for _, value := range l.parkedCars {
		if value.PlateNumber == car.PlateNumber {
			return true
		}
	}
	return false
}

func (l *Lot) IsNotFull() bool {
	return len(l.parkedCars) < l.capacity
}

func (l *Lot) Subscribe(sub Subscriber) {
	l.subscribers = append(l.subscribers, sub)
}

func (l *Lot) notifySubscibersFull() {
	for _, sub := range l.subscribers {
		sub.NotifyLotIsFull(l)
	}
}

func (l *Lot) notifySubscibersNotFull() {
	for _, sub := range l.subscribers {
		sub.NotifyLotIsNotFull(l)
	}
}

func (l *Lot) HasMoreCapacity(lot *Lot) bool {
	return l.capacity > lot.capacity
}

func (l *Lot) HasMoreFreeSpace(lot *Lot) bool {
	return l.countFreeSpace() > lot.countFreeSpace()
}

func (l *Lot) countFreeSpace() int {
	return l.capacity - len(l.parkedCars)
}

func (l *Lot) Status() LotStatus {
	return LotStatus{
		freeSpace:  l.countFreeSpace(),
		parkedCars: l.parkedCars,
	}
}
