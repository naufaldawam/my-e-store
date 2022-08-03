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

func TestDeleteCase(t *testing.T) {
	repo := new(mocks.UserData)

	t.Run("Success Delete", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything).Return(1, nil)

		srv := New(repo, validator.New())
		delete, err := srv.DeleteCase(1)

		assert.NoError(t, err)
		assert.Equal(t, 1, delete)
	})
}

func TestUpdateCase(t *testing.T) {
	insertData := domain.User{
		Name:     "ivan",
		Email:    "ivan@gmail.com",
		Password: "123",
		Phone:    "081217076500",
	}

	repo := new(mocks.UserData)

	t.Run("Success Update", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, 1).Return(1, nil).Once()

		srv := New(repo, validator.New())
		update, err := srv.UpdateCase(insertData, 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, update)
	})

}

func TestGetProfile(t *testing.T) {
	insertData := domain.User{
		Name:     "ivan",
		Email:    "ivan@gmail.com",
		Password: "123",
		Phone:    "081217076500",
		Role:     "user",
	}

	repo := new(mocks.UserData)

	t.Run("Data Found", func(t *testing.T) {
		repo.On("GetSpecific", 1).Return(insertData, nil).Once()

		srv := New(repo, validator.New())
		myprofile, err := srv.GetProfile(1)
		assert.NoError(t, err)
		assert.Equal(t, insertData, myprofile)

	})

	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("GetSpecific", 0).Return(insertData, errors.New("error")).Once()

		srv := New(repo, validator.New())
		notfound, err := srv.GetProfile(0)

		assert.NotNil(t, err)
		assert.NotEqual(t, insertData, notfound)
	})
}
