package parking_test

import (
	"testing"

	"github.com/adityatresnobudi/parking-system/parking"
	"github.com/adityatresnobudi/parking-system/entity"
	"github.com/adityatresnobudi/parking-system/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAttendantPark(t *testing.T) {

	t.Run("should return ticket if park succeeded", func(t *testing.T) {
		p := parking.NewLot(2)
		a := parking.NewAttendant([]*parking.Lot{p})
		car := &entity.Car{PlateNumber: "T 3 ST"}

		result, _ := a.Park(car)

		assert.NotNil(t, result)
	})

	t.Run("should return error if car with same plate enter twice", func(t *testing.T) {
		p := parking.NewLot(2)
		a := parking.NewAttendant([]*parking.Lot{p})
		car := &entity.Car{PlateNumber: "T 3 ST"}
		car2 := &entity.Car{PlateNumber: "T 3 ST"}

		ticket1, err1 := a.Park(car)
		ticket2, err2 := a.Park(car2)

		assert.NotNil(t, ticket1)
		assert.Nil(t, err1)
		assert.Nil(t, ticket2)
		assert.ErrorIs(t, err2, parking.ErrParkedCarTwice)
	})

	t.Run("should return error if there is no available position", func(t *testing.T) {
		p := parking.NewLot(2)
		a := parking.NewAttendant([]*parking.Lot{p})
		car1 := &entity.Car{PlateNumber: "T 3 ST"}
		car2 := &entity.Car{PlateNumber: "P O LE"}
		car3 := &entity.Car{PlateNumber: "E 4 RR"}

		_, _ = a.Park(car1)
		ticket2, err2 := a.Park(car2)
		ticket3, err := a.Park(car3)

		assert.NotNil(t, ticket2)
		assert.Nil(t, err2)
		assert.Nil(t, ticket3)
		assert.ErrorIs(t, err, parking.ErrUnavailablePosition)
	})

	t.Run("should return ticket when park if lot1 is full but lot2 is still available", func(t *testing.T) {
		p1 := parking.NewLot(2)
		p2 := parking.NewLot(2)
		a := parking.NewAttendant([]*parking.Lot{p1, p2})
		car1 := &entity.Car{PlateNumber: "T 3 ST"}
		car2 := &entity.Car{PlateNumber: "P O LE"}
		car3 := &entity.Car{PlateNumber: "E 4 RR"}

		_, _ = a.Park(car1)
		_, _ = a.Park(car2)
		ticket3, err3 := a.Park(car3)

		assert.NotNil(t, ticket3)
		assert.Nil(t, err3)
	})

	t.Run("should return error when park if all lot is full", func(t *testing.T) {
		p1 := parking.NewLot(2)
		p2 := parking.NewLot(1)
		a := parking.NewAttendant([]*parking.Lot{p1, p2})
		car1 := &entity.Car{PlateNumber: "T 3 ST"}
		car2 := &entity.Car{PlateNumber: "P O LE"}
		car3 := &entity.Car{PlateNumber: "E 4 RR"}
		car4 := &entity.Car{PlateNumber: "P O ST"}

		_, _ = a.Park(car1)
		_, _ = a.Park(car2)
		_, _ = a.Park(car3)
		ticket4, err4 := a.Park(car4)

		assert.Nil(t, ticket4)
		assert.ErrorIs(t, err4, parking.ErrUnavailablePosition)
	})

	t.Run("should return error if car with same plate enter twice", func(t *testing.T) {
		p := parking.NewLot(2)
		a := parking.NewAttendant([]*parking.Lot{p})
		car1 := &entity.Car{PlateNumber: "T 3 ST"}
		car2 := &entity.Car{PlateNumber: "T 3 ST"}

		ticket1, err1 := a.Park(car1)
		ticket2, err2 := a.Park(car2)

		assert.NotNil(t, ticket1)
		assert.Nil(t, err1)
		assert.Nil(t, ticket2)
		assert.ErrorIs(t, err2, parking.ErrParkedCarTwice)
	})

	t.Run("should return error when car is already parked in another lot", func(t *testing.T) {
		p1 := parking.NewLot(2)
		p2 := parking.NewLot(1)
		a := parking.NewAttendant([]*parking.Lot{p1, p2})
		car1 := &entity.Car{PlateNumber: "T 3 ST"}
		car2 := &entity.Car{PlateNumber: "P O LE"}
		car3 := &entity.Car{PlateNumber: "T 3 ST"}

		_, _ = a.Park(car1)
		_, _ = a.Park(car2)
		ticket3, err3 := a.Park(car3)

		assert.Nil(t, ticket3)
		assert.ErrorIs(t, err3, parking.ErrParkedCarTwice)
	})

	t.Run("should return error when car is already parked in another lot", func(t *testing.T) {
		p1 := parking.NewLot(1)
		p2 := parking.NewLot(1)
		a := parking.NewAttendant([]*parking.Lot{p1, p2})
		car1 := &entity.Car{PlateNumber: "P O LE"}
		car2 := &entity.Car{PlateNumber: "T 3 ST"}
		car3 := &entity.Car{PlateNumber: "T 3 ST"}

		ticket1, _ := a.Park(car1)
		_, _ = a.Park(car2)
		_, _ = a.UnPark(ticket1)
		ticket3, err3 := a.Park(car3)

		assert.Nil(t, ticket3)
		assert.ErrorIs(t, err3, parking.ErrParkedCarTwice)
	})

	t.Run("should change parking style when changing attendant parking style", func(t *testing.T) {
		mockLotSelector := mocks.NewLotSelector(t)
		l1 := parking.NewLot(1)
		l2 := parking.NewLot(2)
		car := &entity.Car{PlateNumber: "T 3 ST"}
		expected := l1
		a := parking.NewAttendant([]*parking.Lot{l1, l2})
		a.ChangeStyle(mockLotSelector)

		mockLotSelector.On("SelectLot", a.GetAvailLots()).Return(expected)
		ticket, _ := a.Park(car)

		assert.NotNil(t, ticket)
	})

	t.Run("should park car in highest capacity lot when parking style is HighestCapacity", func(t *testing.T) {
		p1 := parking.NewLot(2)
		p2 := parking.NewLot(3)
		a := parking.NewAttendant([]*parking.Lot{p1, p2})
		car1 := &entity.Car{PlateNumber: "P O LE"}
		car2 := &entity.Car{PlateNumber: "T 3 ST"}
		expected := car2

		_, _ = a.Park(car1)
		a.ChangeStyle(&parking.HighestCapacity{})
		ticket2, _ := a.Park(car2)
		returnedCar2, _ := p2.UnPark(ticket2)

		assert.Same(t, expected, returnedCar2)
	})

	t.Run("should park car in highest free space lot when parking style is HighestFreeSpace", func(t *testing.T) {
		p1 := parking.NewLot(2)
		p2 := parking.NewLot(3)
		a := parking.NewAttendant([]*parking.Lot{p1, p2})
		car1 := &entity.Car{PlateNumber: "P O LE"}
		car2 := &entity.Car{PlateNumber: "T 3 ST"}
		car3 := &entity.Car{PlateNumber: "B 3 ST"}
		expected := car3

		_, _ = a.Park(car1)
		a.ChangeStyle(&parking.HighestCapacity{})
		_, _ = a.Park(car2)
		a.ChangeStyle(&parking.HighestFreeSpace{})
		ticket3, _ := a.Park(car3)
		returnedCar3, _ := p2.UnPark(ticket3)

		assert.Same(t, expected, returnedCar3)
	})
}

func TestAttendantUnPark(t *testing.T) {

	t.Run("should return car when unpark if ticket exist when unpark", func(t *testing.T) {
		p := parking.NewLot(2)
		a := parking.NewAttendant([]*parking.Lot{p})
		car := &entity.Car{PlateNumber: "T 3 ST"}
		ticket, _ := a.Park(car)
		expected := car

		returnedCar, err := a.UnPark(ticket)

		assert.Same(t, expected, returnedCar)
		assert.Nil(t, err)
	})

	t.Run("should return error when unpark if ticket does not exist", func(t *testing.T) {
		p := parking.NewLot(2)
		a := parking.NewAttendant([]*parking.Lot{p})
		ticket := &entity.Ticket{ID: "ERR!!!"}

		returnedCar, err := a.UnPark(ticket)

		assert.Nil(t, returnedCar)
		assert.ErrorIs(t, err, parking.ErrUnrecognizedParkingTicket)
	})

	t.Run("should return error when unpark if unparked car does not have ticket", func(t *testing.T) {
		p := parking.NewLot(2)
		a := parking.NewAttendant([]*parking.Lot{p})
		ticket := &entity.Ticket{}

		returnedCar, err := a.UnPark(ticket)

		assert.Nil(t, returnedCar)
		assert.ErrorIs(t, err, parking.ErrUnrecognizedParkingTicket)
	})

	t.Run("should assert error if ticket has already been used", func(t *testing.T) {
		p := parking.NewLot(2)
		a := parking.NewAttendant([]*parking.Lot{p})
		car1 := &entity.Car{PlateNumber: "T 3 ST"}
		ticket, _ := a.Park(car1)

		_, _ = a.UnPark(ticket)
		ticket2, err := a.UnPark(ticket)

		assert.Nil(t, ticket2)
		assert.ErrorIs(t, err, parking.ErrUnrecognizedParkingTicket)
	})
}
