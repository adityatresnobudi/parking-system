package parking

import "github.com/adityatresnobudi/parking-system/entity"

type Attendant struct {
	lotList       []*Lot
	availableLots []*Lot
	parkingStyle  LotSelector
}

type LotSelector interface {
	SelectLot(lots []*Lot) *Lot
}

func NewAttendant(lots []*Lot) *Attendant {
	tLot := make([]*Lot, len(lots))
	copy(tLot, lots)
	a := &Attendant{
		lotList:       lots,
		availableLots: tLot,
		parkingStyle:  &FirstAvailable{},
	}
	a.SubsribeAllLot()
	return a
}

func (a *Attendant) Park(car *entity.Car) (*entity.Ticket, error) {
	if a.isCarParked(car) {
		return nil, ErrParkedCarTwice
	}
	if lot := a.findAvailableLot(a.lotList); lot != nil {
		selectedLot := a.parkingStyle.SelectLot(a.availableLots)
		return selectedLot.Park(car)
	}
	return nil, ErrUnavailablePosition
}

func (a *Attendant) UnPark(ticket *entity.Ticket) (*entity.Car, error) {
	if i := a.findTicket(ticket); i != -1 {
		return a.lotList[i].UnPark(ticket)
	}
	return nil, ErrUnrecognizedParkingTicket
}

func (a *Attendant) findAvailableLot(lotList []*Lot) *Lot {
	for _, val := range a.lotList {
		if val.IsNotFull() {
			return val
		}
	}
	return nil
}

func (a *Attendant) findTicket(ticket *entity.Ticket) int {
	for idx, val := range a.lotList {
		if _, ok := val.parkedCars[ticket.ID]; ok {
			return idx
		}
	}
	return -1
}

func (a *Attendant) isCarParked(car *entity.Car) bool {
	for _, lot := range a.lotList {
		if lot.IsCarParked(car) {
			return true
		}
	}
	return false
}

func (a *Attendant) SubsribeAllLot() {
	for _, l := range a.lotList {
		l.Subscribe(a)
	}
}

func (a *Attendant) NotifyLotIsFull(lot *Lot) {
	fullLotIdx := a.lotIdx(a.availableLots, lot)
	a.availableLots = append(a.availableLots[:fullLotIdx], a.availableLots[fullLotIdx+1:]...)
}

func (a *Attendant) NotifyLotIsNotFull(lot *Lot) {
	a.availableLots = append(a.availableLots, lot)
}

func (a *Attendant) lotIdx(lots []*Lot, lot *Lot) int {
	output := 0
	for idx, l := range lots {
		if lot == l {
			output = idx
			break
		}
	}
	return output
}

func (a *Attendant) ChangeStyle(style LotSelector) {
	a.parkingStyle = style
}

func (a *Attendant) GetAvailLots() []*Lot {
	return a.availableLots
}

func (a *Attendant) Status() []LotStatus {
	output := make([]LotStatus, 0)
	for _, lot := range a.lotList {
		output = append(output, lot.Status())
	}
	return output
}
