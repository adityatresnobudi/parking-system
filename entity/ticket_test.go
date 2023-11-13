package entity_test

import (
	"testing"

	. "github.com/adityatresnobudi/parking-system/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewTicket(t *testing.T) {
	t.Run("should return ticket with non empty ID", func(t *testing.T) {
		t1 := NewTicket()

		assert.NotEmpty(t, t1.ID)
	})

	t.Run("should return ticket with unique ID", func(t *testing.T) {
		t1 := NewTicket()
		t2 := NewTicket()
		t3 := NewTicket()

		assert.NotEqual(t, t1, t2)
		assert.NotEqual(t, t2, t3)
	})
}
