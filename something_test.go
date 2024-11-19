package core_test

import (
	"fmt"

	. "github.com/gildas/go-core"
	"github.com/google/uuid"
)

type Something interface {
	TypeCarrier
	GetData() string
}

type SomethingMore interface {
	Something
	fmt.Stringer
}

type Something1 struct {
	Data string `json:"data"`
}

func (s Something1) GetType() string {
	return "something1"
}

func (s Something1) GetData() string {
	return s.Data
}

func (something Something1) GetName() string {
	return something.Data
}

func (something Something2) String() string {
	return something.Data
}

type Something2 struct {
	Data string `json:"data"`
}

func (s Something2) GetType() string {
	return "something2"
}

func (s Something2) GetData() string {
	return s.Data
}

type Something3 struct {
	ID uuid.UUID `json:"id"`
}

func (s Something3) GetType() string {
	return "something3"
}

func (s Something3) GetID() uuid.UUID {
	return s.ID
}

type Something4 struct {
	ID string `json:"id"`
}

func (s Something4) GetType() string {
	return "something4"
}

func (s Something4) GetID() string {
	return s.ID
}
