package testfytests_test

import (
	"errors"
	"testing"

	"github.com/guirialli/pos/test/gotests"
	"github.com/guirialli/pos/test/testfytests"
	"github.com/stretchr/testify/assert"
)

func TestCalculteTest(t *testing.T) {
	tax, error := gotests.CalculateTaxError(1000)
	assert.Nil(t, error)
	assert.Equal(t, 1100.00, tax)
}

func TestError(t *testing.T) {
	tax, error := gotests.CalculateTaxError(0)
	assert.NotNil(t, error)
	assert.Equal(t, 0.00, tax)
	assert.Equal(t, "invalid amount value has to be than 0", error.Error())
}

func TestCalculteTaxAndSave(t *testing.T) {
	repository := &testfytests.TaxRepositoryMock{}
	repository.On("Save", 1100.00).Return(nil)
	repository.On("Save", 0.00).Return(errors.New("error saving tax"))

	err := gotests.CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = gotests.CalculateTaxAndSave(0.0, repository)
	assert.NotNil(t, err)
	assert.Equal(t, "error saving tax", err.Error())
	repository.AssertExpectations(t)
}
