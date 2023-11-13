package parking

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/adityatresnobudi/parking-system/entity"
)

var (
	ErrNoParkingLot = errors.New("parking lot haven't been setup")
	ErrInvalidInput = errors.New("invalid input")
)

func SetupHandler(arg string) (*Attendant, error) {
	if !isArgsValid(arg) {
		return nil, ErrInvalidInput
	}

	lots := make([]*Lot, 0)

	for _, v := range strings.Split(arg, ",") {
		capacity, err := strconv.Atoi(v)
		if err != nil {
			return nil, ErrInvalidInput
		}

		lots = append(lots, NewLot(capacity))
	}

	return NewAttendant(lots), nil
}

func ParkHandler(arg string, attendant *Attendant) (string, error) {
	if !isArgsValid(arg) {
		return "", ErrInvalidInput
	}

	if !isAttendantExist(attendant) {
		return "", ErrNoParkingLot
	}

	car := &entity.Car{PlateNumber: arg}
	ticket, err := attendant.Park(car)
	switch err {
	case ErrUnavailablePosition:
		return "", err
	case ErrParkedCarTwice:
		return "", err
	}
	return fmt.Sprintf("Car parked with ticket id %s", ticket.ID), nil
}

func UnParkHandler(arg string, attendant *Attendant) (string, error) {
	if !isArgsValid(arg) {
		return "", ErrInvalidInput
	}

	if !isAttendantExist(attendant) {
		return "", ErrNoParkingLot
	}

	ticket := &entity.Ticket{ID: arg}
	returnedCar, err := attendant.UnPark(ticket)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Car %s succesfully unparked!", returnedCar.PlateNumber), nil
}

func StatusHandler(attendant *Attendant) (string, error) {
	if !isAttendantExist(attendant) {
		return "", ErrNoParkingLot
	}

	res := "Parking Lot Status:\n"

	for i, v := range attendant.Status() {
		res += fmt.Sprintf("Lot #%d: %d spaces left\n", i+1, v.freeSpace)
		for ticket, car := range v.parkedCars {
			res += fmt.Sprintf("#%s %s\n", ticket, car.PlateNumber)
		}
	}

	return res, nil
}

func isAttendantExist(attendant *Attendant) bool {
	return attendant != nil
}

func isArgsValid(arg string) bool {
	return arg != ""
}
