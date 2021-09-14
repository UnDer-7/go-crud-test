package service

import (
	"my-tracking-list-backend/core/domain"
	"my-tracking-list-backend/core/ports/driver"
	mockDriven "my-tracking-list-backend/test_helpers/mock"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUserServiceImpl_SaveUser(t *testing.T) {
	t.Run("Should save user successfully", func(t *testing.T) {
		mockRepository, service, userSample := setupService(t)

		mockRepository.
			EXPECT().
			Persist(gomock.Eq(userSample)).
			Times(1)

		_, err := service.SaveUser(userSample)

		if err != nil {
			t.Fatal("Got an error but didn't want one")
		}
	})

	t.Run("Should handle repositories errors when saving user", func(t *testing.T) {
		mockRepository, service, userSample := setupService(t)
		expectedError := "generic error"
		mockRepository.
			EXPECT().
			Persist(gomock.Eq(userSample)).
			Times(1).
			Return(domain.User{}, errors.New(expectedError))

		_, err := service.SaveUser(userSample)

		if err == nil {
			t.Fatal("Expect an error to occur")
		}

		if err.Error() != expectedError {
			t.Errorf("expect error msg: %s, actual error msg: %s", expectedError, err.Error())
		}
	})
}

func TestUserServiceImpl_FindById(t *testing.T) {
	t.Run("Should find an user by id successfully", func(t *testing.T) {
		mockRepository, service, userSample := setupService(t)
		expectedId := 42
		userSample.ID = expectedId

		mockRepository.
			EXPECT().
			GetById(gomock.Eq(expectedId)).
			Times(1).
			Return(userSample, nil)

		_, err := service.FindById(expectedId)

		if err != nil {
			t.Fatalf("didn't expect an erro %v", err)
		}
	})

	t.Run("Should handle repositories errors when finding an user by id", func(t *testing.T) {
		mockRepository, service, _ := setupService(t)
		expectedError := "generic error"
		expectedId := 42
		mockRepository.
			EXPECT().
			GetById(gomock.Eq(expectedId)).
			Times(1).
			Return(domain.User{}, errors.New(expectedError))

		_, err := service.FindById(expectedId)

		if err == nil {
			t.Fatal("Expect an error to occur")
		}

		if err.Error() != expectedError {
			t.Errorf("expect error msg: %s, actual error msg: %s", expectedError, err.Error())
		}
	})
}

func setupService(t *testing.T) (*mockDriven.MockUserRepository, driver.UserService, domain.User) {
	mockRepository := mockDriven.NewMockUserRepository(gomock.NewController(t))
	service := NewUserService(mockRepository)
	userSample := domain.User{
		ID:       0,
		Email:    "abc@acb.abc",
		Password: "1234",
	}
	return mockRepository, service, userSample
}
