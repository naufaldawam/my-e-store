package usecase

import (
	"errors"
	"project3/group3/domain"
	"project3/group3/mocks"
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddUser(t *testing.T) {
	repo := new(mocks.UserData)
	insertData := domain.User{
		ID:       1,
		Name:     "ivan",
		Email:    "ivan@gmail.com",
		Password: "123",
		Phone:    "081217076500",
		Role:     "user",
	}
	falseData := domain.User{
		ID:       0,
		Name:     "",
		Email:    "",
		Password: "",
		Phone:    "",
		Role:     "",
	}
	outputData := 1

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(outputData, nil).Once()

		srv := New(repo, validator.New())

		res, err := srv.AddUser(insertData)
		assert.NoError(t, err)
		assert.Equal(t, outputData, res)
		repo.AssertExpectations(t)
	})

	t.Run("Fill Error", func(t *testing.T) {
		repo.On("Fill Error", mock.Anything).Return(-1, errors.New("please make sure all fields are filled in correctly")).Once()

		srv := New(repo, validator.New())

		res, err := srv.AddUser(falseData)
		assert.NotNil(t, err)
		assert.Equal(t, -1, res)
	})

}
