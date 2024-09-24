package ptrutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToPointer(t *testing.T) {
	a := assert.New(t)
	t.Run("bool", func(t *testing.T) {
		tru := true
		fals := false

		a.Equal(&tru, ToPointer(true))
		a.Equal(&fals, ToPointer(false))
	})
}

func TestFromPointer(t *testing.T) {
	a := assert.New(t)

	t.Run("string", func(t *testing.T) {
		s2 := "test"
		a.Equal("test", FromPointer(&s2))
	})

	t.Run("Nil", func(t *testing.T) {
		var s *string // nil
		a.Equal("", FromPointer(s))
	})
}
