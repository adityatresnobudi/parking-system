package parking_test

import (
	"testing"

	"github.com/adityatresnobudi/parking-system/entity"
	"github.com/adityatresnobudi/parking-system/mocks"
	"github.com/adityatresnobudi/parking-system/parking"
	"github.com/stretchr/testify/assert"
)

func TestPark(t *testing.T) {

	t.Run("should return ticket if park succeeded", func(t *testing.T) {
		p := parking.NewLot(2)
		car := &entity.Car{PlateNumber: "T 3 ST"}

		result, _ := p.Park(car)

		assert.NotNil(t, result)
	})

	t.Run("should return error if car with same plate enter twice", func(t *testing.T) {
		car1 := &entity.Car{PlateNumber: "T 3 ST"}
		car2 := &entity.Car{PlateNumber: "T 3 ST"}
		p := parking.NewLot(2)

		ticket1, err1 := p.Park(car1)
		ticket2, err2 := p.Park(car2)

		assert.NotNil(t, ticket1)
		assert.Nil(t, err1)
		assert.Nil(t, ticket2)
		assert.ErrorIs(t, err2, parking.ErrParkedCarTwice)
	})

	t.Run("should return error if there is no available position", func(t *testing.T) {
		p := parking.NewLot(2)
		car1 := &entity.Car{PlateNumber: "T 3 ST"}
		car2 := &entity.Car{PlateNumber: "P O LE"}
		car3 := &entity.Car{PlateNumber: "E 4 RR"}

		_, _ = p.Park(car1)
		returnedCar2, err2 := p.Park(car2)
		ticket, err := p.Park(car3)

		assert.NotNil(t, returnedCar2)
		assert.Nil(t, err2)
		assert.Nil(t, ticket)
		assert.ErrorIs(t, err, parking.ErrUnavailablePosition)
	})

	t.Run("should call subscriber.notifyFull when parking lot is full", func(t *testing.T) {
		p := parking.NewLot(1)
		car := &entity.Car{PlateNumber: "T 3 ST"}
		a := parking.NewAttendant([]*parking.Lot{p})
		mockNotifySubs := new(mocks.Subscriber)
		p.Subscribe(mockNotifySubs)
		mockNotifySubs.On("NotifyLotIsFull", p).Return()

		result, _ := a.Park(car)

		assert.NotNil(t, result)
		mockNotifySubs.AssertNumberOfCalls(t, "NotifyLotIsFull", 1)
	})
}

func TestUnpark(t *testing.T) {

	t.Run("should return car when unpark if ticket exist", func(t *testing.T) {
		p := parking.NewLot(2)
		car := &entity.Car{PlateNumber: "T 3 ST"}
		ticketParkedCar, _ := p.Park(car)
		expected := car

		returnedCar, err := p.UnPark(ticketParkedCar)

		assert.Same(t, expected, returnedCar)
		assert.Nil(t, err)
	})

	t.Run("should return error when unpark if ticket does not exist", func(t *testing.T) {
		p := parking.NewLot(2)
		ticket := &entity.Ticket{ID: "ERR!"}

		returnedCar, err := p.UnPark(ticket)

		assert.Nil(t, returnedCar)
		assert.ErrorIs(t, err, parking.ErrUnrecognizedParkingTicket)
	})

	t.Run("should return error if unparked car does not have ticket", func(t *testing.T) {
		p := parking.NewLot(2)
		ticket := &entity.Ticket{}

		returnedCar, err := p.UnPark(ticket)

		assert.Nil(t, returnedCar)
		assert.ErrorIs(t, err, parking.ErrUnrecognizedParkingTicket)
	})

	t.Run("should return error if ticket has already been used", func(t *testing.T) {
		p := parking.NewLot(2)
		car1 := &entity.Car{PlateNumber: "T 3 ST"}
		ticketParkedCar, _ := p.Park(car1)

		_, _ = p.UnPark(ticketParkedCar)
		returnedCar2, err := p.UnPark(ticketParkedCar)

		assert.Nil(t, returnedCar2)
		assert.ErrorIs(t, err, parking.ErrUnrecognizedParkingTicket)
	})

	t.Run("should call subscriber.notifyNotFull when parking lot has available position", func(t *testing.T) {
		p := parking.NewLot(1)
		car := &entity.Car{PlateNumber: "T 3 ST"}
		a := parking.NewAttendant([]*parking.Lot{p})
		expected := car
		mockNotifySubs := new(mocks.Subscriber)
		p.Subscribe(mockNotifySubs)
		mockNotifySubs.On("NotifyLotIsFull", p).Return()
		mockNotifySubs.On("NotifyLotIsNotFull", p).Return()

		ticket, _ := a.Park(car)
		returnedCar, _ := a.UnPark(ticket)

		assert.NotNil(t, ticket)
		mockNotifySubs.AssertNumberOfCalls(t, "NotifyLotIsFull", 1)
		assert.Same(t, expected, returnedCar)
		mockNotifySubs.AssertNumberOfCalls(t, "NotifyLotIsNotFull", 1)
	})
}
