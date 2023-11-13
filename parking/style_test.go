package parking_test

import (
	"testing"

	"github.com/adityatresnobudi/parking-system/entity"
	"github.com/adityatresnobudi/parking-system/parking"
	"github.com/stretchr/testify/assert"
)

func TestStyleAlgo(t *testing.T) {

	t.Run("should return first available lot when using FirstAvailable strategy", func(t *testing.T) {
		p1 := parking.NewLot(2)
		p2 := parking.NewLot(2)
		expected := p1

		result := parking.LotSelector.SelectLot(&parking.FirstAvailable{}, []*parking.Lot{p1, p2})

		assert.Equal(t, expected, result)
	})

	t.Run("should return highest capacity lot when using HighestCapacity strategy", func(t *testing.T) {
		p1 := parking.NewLot(2)
		p2 := parking.NewLot(4)
		expected := p2

		result := parking.LotSelector.SelectLot(&parking.HighestCapacity{}, []*parking.Lot{p1, p2})

		assert.Equal(t, expected, result)
	})

	t.Run("should return highest free space lot when using HighestFreeSpace strategy", func(t *testing.T) {
		p1 := parking.NewLot(2)
		p2 := parking.NewLot(2)
		a := parking.NewAttendant([]*parking.Lot{p1, p2})
		car := &entity.Car{PlateNumber: "T 3 ST"}
		expected := p2

		ticket, _ := a.Park(car)
		result := parking.LotSelector.SelectLot(&parking.HighestFreeSpace{}, []*parking.Lot{p1, p2})

		assert.NotNil(t, ticket)
		assert.Equal(t, expected, result)
	})
}
