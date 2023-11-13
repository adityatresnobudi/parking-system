package parking_test

import (
	"fmt"
	"testing"

	"github.com/adityatresnobudi/parking-system/entity"
	"github.com/adityatresnobudi/parking-system/parking"
	"github.com/stretchr/testify/assert"
)

func TestUtility(t *testing.T) {

	t.Run("should return error when given invalid SetupHandler arguments", func(t *testing.T) {
		arg := ""

		attendant, err := parking.SetupHandler(arg)

		assert.ErrorIs(t, parking.ErrInvalidInput, err)
		assert.Nil(t, attendant)
	})

	t.Run("should return error when given invalid capacity list SetupHandler arguments", func(t *testing.T) {
		arg := "1,2,f"

		attendant, err := parking.SetupHandler(arg)

		assert.ErrorIs(t, parking.ErrInvalidInput, err)
		assert.Nil(t, attendant)
	})

	t.Run("should create new Attendant when given correct SetupHandler arguments", func(t *testing.T) {
		arg := "1,2,3"

		attendant, err := parking.SetupHandler(arg)

		assert.Nil(t, err)
		assert.NotNil(t, attendant)
	})

	t.Run("should return error when given invalid ParkHandler arguments", func(t *testing.T) {
		arg := ""
		expected := ""

		res, err := parking.ParkHandler(arg, nil)

		assert.ErrorIs(t, parking.ErrInvalidInput, err)
		assert.Equal(t, expected, res)
	})

	t.Run("should return error when Attendant is not initialize on ParkHandler", func(t *testing.T) {
		arg := "B 3 ST"
		expected := ""

		res, err := parking.ParkHandler(arg, nil)

		assert.ErrorIs(t, parking.ErrNoParkingLot, err)
		assert.Equal(t, expected, res)
	})

	t.Run("should return error output when given valid ParkHandler arguments and Park return ErrUnavailablePosition", func(t *testing.T) {
		attendant := parking.NewAttendant([]*parking.Lot{})
		arg := "B 3 ST"
		expected := ""

		res, err := parking.ParkHandler(arg, attendant)

		assert.NotNil(t, err)
		assert.Equal(t, expected, res)
	})

	t.Run("should return error output when given valid ParkHandler arguments and Park return error ErrParkedCarTwice", func(t *testing.T) {
		attendant := parking.NewAttendant([]*parking.Lot{parking.NewLot(1)})
		arg1 := "B 3 ST"
		arg2 := "B 3 ST"
		expected := ""

		_, _ = parking.ParkHandler(arg1, attendant)
		res, err := parking.ParkHandler(arg2, attendant)

		assert.NotNil(t, err)
		assert.Equal(t, expected, res)
	})

	t.Run("should return correct output when given valid ParkHandler arguments", func(t *testing.T) {
		attendant := parking.NewAttendant([]*parking.Lot{parking.NewLot(1)})
		arg := "B 3 ST"

		res, err := parking.ParkHandler(arg, attendant)

		assert.Nil(t, err)
		assert.Contains(t, res, "Car parked with ticket id")
	})

	t.Run("should return error when given invalid HandleUnpark arguments", func(t *testing.T) {
		arg := ""
		expected := ""

		res, err := parking.UnParkHandler(arg, nil)

		assert.ErrorIs(t, parking.ErrInvalidInput, err)
		assert.Equal(t, expected, res)
	})

	t.Run("should return error when Attendant is not initialize on UnParkHandler", func(t *testing.T) {
		arg := entity.NewTicket()
		expected := ""

		res, err := parking.UnParkHandler(arg.ID, nil)

		assert.ErrorIs(t, parking.ErrNoParkingLot, err)
		assert.Equal(t, expected, res)
	})

	t.Run("should return error output when given valid UnParkHandler arguments and Unpark return error", func(t *testing.T) {
		attendant := parking.NewAttendant([]*parking.Lot{})
		ticket := entity.NewTicket()
		expected := ""

		res, err := parking.UnParkHandler(ticket.ID, attendant)

		assert.NotNil(t, err)
		assert.Equal(t, expected, res)
	})

	t.Run("should return correct output when given valid UnParkHandler arguments", func(t *testing.T) {
		lot := parking.NewLot(1)
		attendant := parking.NewAttendant([]*parking.Lot{lot})
		car := &entity.Car{PlateNumber: "B 3 ST"}
		expected := "Car B 3 ST succesfully unparked!"
		ticket, _ := lot.Park(car)

		res, err := parking.UnParkHandler(ticket.ID, attendant)

		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})

	t.Run("should return error when Attendant is not initialize on StatusHandler", func(t *testing.T) {
		expected := ""

		res, err := parking.StatusHandler(nil)

		assert.ErrorIs(t, parking.ErrNoParkingLot, err)
		assert.Equal(t, expected, res)
	})

	t.Run("should return correct output when given valid StatusHandler arguments", func(t *testing.T) {
		lot := parking.NewLot(1)
		attendant := parking.NewAttendant([]*parking.Lot{lot})
		car := &entity.Car{PlateNumber: "B 3 ST"}
		ticket, _ := lot.Park(car)
		expected := fmt.Sprintf("Parking Lot Status:\nLot #1: 0 spaces left\n#%s %s\n", ticket.ID, car.PlateNumber)

		res, err := parking.StatusHandler(attendant)

		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})
}
