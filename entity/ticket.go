package entity

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	minID = 1000
	maxID = 9999
)

type Ticket struct {
	ID string
}

func NewTicket() Ticket {
	rand.Seed(time.Now().UnixNano())
	unique := rand.Intn(maxID-minID+1) + minID
	return Ticket{
		ID: fmt.Sprint(unique),
	}
}
