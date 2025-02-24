package service_struct

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockDataService struct {
	mock.Mock
	real DataService
}

func (m *MockDataService) GetData(id int) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

func TestMockDataService(t *testing.T) {
	mockService := new(MockDataService)

	// Define expected behavior
	mockService.On("GetData", 1).Return("Mocked Data", nil)
	mockService.On("GetData", 2).Return("", errors.New("not found"))

	// Case #1 - Test successful response
	result, err := mockService.GetData(1)
	assert.NoError(t, err)
	assert.Equal(t, "Mocked Data", result)

	// Case #2 - Test error response
	result, err = mockService.GetData(2)
	assert.Error(t, err)
	assert.Equal(t, "not found", err.Error())
	assert.Empty(t, result)

	// Optional: verify expectations
	mockService.AssertExpectations(t)
}
