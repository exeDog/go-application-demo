package forms

import (
	"github.com/jaswdr/faker"
	"testing"
)

func TestErrors_Add(t *testing.T) {
	e := errors{}

	fakerMock := faker.New()

	fieldName := fakerMock.Car().Maker()

	e.Add(fieldName, "some message")

	if e.Get(fieldName) != "some message" {
		t.Error("error message not found")
	}
}

func TestErrors_Get(t *testing.T) {
	e := errors{}

	if e.Get("somefield") != "" {
		t.Error("error message found")
	}

	fakerMock := faker.New()
	mockError := fakerMock.Car().Model()

	e.Add("somefield", mockError)

	if e.Get("somefield") != mockError {
		t.Error("error message not found")
	}
}
