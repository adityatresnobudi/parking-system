package parking

import (
	"sort"
)

type FirstAvailable struct {
}

func (fa *FirstAvailable) SelectLot(availableLots []*Lot) *Lot {
	return availableLots[0]
}

type HighestCapacity struct {
}

func (hc *HighestCapacity) SelectLot(availableLots []*Lot) *Lot {
	sort.Slice(availableLots, func(i int, j int) bool {
		return availableLots[i].HasMoreCapacity(availableLots[j])
	})
	return availableLots[0]
}

type HighestFreeSpace struct {
}

func (hf *HighestFreeSpace) SelectLot(availableLots []*Lot) *Lot {
	sort.Slice(availableLots, func(i int, j int) bool {
		return availableLots[i].HasMoreFreeSpace(availableLots[j])
	})
	return availableLots[0]
}
