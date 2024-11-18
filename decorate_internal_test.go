package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPluralize(t *testing.T) {
	assert.Equal(t, "cats", pluralize("cat"))
	assert.Equal(t, "buses", pluralize("bus"))
	assert.Equal(t, "boxes", pluralize("box"))
	assert.Equal(t, "parties", pluralize("party"))
	assert.Equal(t, "quizzes", pluralize("quiz"))
	assert.Equal(t, "heroes", pluralize("hero"))
	assert.Equal(t, "potatoes", pluralize("potato"))
	assert.Equal(t, "fishes", pluralize("fish"))
	assert.Equal(t, "leaves", pluralize("leaf"))
}
