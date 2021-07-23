package bod_test

import (
	"errors"
	"testing"

	"github.com/jadahbakar/asastarealty-backend/internal/master/bod"
	"github.com/jadahbakar/asastarealty-backend/mocks"
	"github.com/stretchr/testify/assert"
)

func setup() (mockBodRepo *mocks.BodRepository, srv bod.BodService) {
	mockBodRepo = new(mocks.BodRepository)
	srv = bod.NewBodService(mockBodRepo)
	return
}

func TestFindAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// setup
		repo, srv := setup()
		actual := []bod.Bod{}
		repo.On("SearchAll").Return(actual, nil)

		// testing
		data, err := srv.FindAll()

		// assertion
		assert.NoError(t, err)
		assert.Equal(t, actual, data)

		// end mock
		repo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		// setup
		repo, srv := setup()
		actual := []bod.Bod{}
		repo.On("SearchAll").Return(actual, errors.New("err"))

		// testing
		_, err := srv.FindAll()

		// assertion
		assert.Error(t, err)

		// end mock
		repo.AssertExpectations(t)
	})
}

func TestFindById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// setup
		repo, srv := setup()
		actual := bod.Bod{}
		repo.On("SearchById", 1).Return(actual, nil)

		// testing
		data, err := srv.FindById(1)

		// assertion
		assert.NoError(t, err)
		assert.Equal(t, actual, data)

		// end mock
		repo.AssertExpectations(t)
	})
	t.Run("Error", func(t *testing.T) {
		mockBodRepo := new(mocks.BodRepository)
		paramId := 1
		mockBodRepo.On("SearchById", paramId).Return(bod.Bod{}, errors.New("error"))
		srv := bod.NewBodService(mockBodRepo)

		_, err := srv.FindById(paramId)

		assert.Error(t, err)
		assert.Equal(t, "error", err.Error())
		mockBodRepo.AssertExpectations(t)
		mockBodRepo.AssertExpectations(t)

	})
}
