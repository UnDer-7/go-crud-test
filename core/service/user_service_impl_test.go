package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"my-tracking-list-backend/core/app_error"
	"my-tracking-list-backend/core/domain"
	"my-tracking-list-backend/core/ports/driver"
	"my-tracking-list-backend/test_helpers"
	"my-tracking-list-backend/test_helpers/mock"
	"testing"
	"time"
)

func TestUserServiceImpl_SaveUser(t *testing.T) {
	mockCtx := context.TODO()

	t.Run("Should save user successfully", func(t *testing.T) {
		mockRepository, service, userSample, assert := setupService(t)

		mockRepository.
			EXPECT().
			ExistesByEmail(mockCtx, userSample.Email).
			Times(1).
			Return(false, nil)

		mockRepository.
			EXPECT().
			Persist(mockCtx, userSample).
			Times(1)

		_, err := service.SaveUser(mockCtx, userSample)
		assert.NotExpectedError(err)
	})

	t.Run("Should not save user with an ID", func(t *testing.T) {
		mockRepository, service, userSample, assert := setupService(t)
		userSample.ID = primitive.NewObjectID()

		mockRepository.
			EXPECT().
			ExistesByEmail(gomock.Any(), gomock.Any()).
			Times(0)
		mockRepository.
			EXPECT().
			Persist(gomock.Any(), gomock.Any()).
			Times(0)

		_, err := service.SaveUser(mockCtx, userSample)

		assert.ExpectedError(err)
		appError, ok := err.(*app_error.AppError)
		if !ok {
			t.Fatalf("Expected the error to of type app_error.AppError, but it is not\t%v", appError)
		}

		assert.ExpectedErrorStatusCode(422, appError)
	})

	t.Run("Should not save user that exists", func(t *testing.T) {
		mockRepository, service, userSample, assert := setupService(t)

		mockRepository.
			EXPECT().
			ExistesByEmail(mockCtx, gomock.Any()).
			Times(1).
			Return(true, nil)
		mockRepository.
			EXPECT().
			Persist(gomock.Any(), gomock.Any()).
			Times(0)

		_, err := service.SaveUser(mockCtx, userSample)
		assert.ExpectedError(err)

		appError, ok := err.(*app_error.AppError)
		if !ok {
			t.Fatalf("Expected the error to of type app_error.AppError, but it is not\t%v", appError)
		}

		assert.ExpectedErrorStatusCode(422, appError)
	})

	t.Run("Should return error from repository when verify if user exists", func(t *testing.T) {
		mockRepository, service, userSample, assert := setupService(t)
		expectedErr := errors.New("error expected")

		mockRepository.
			EXPECT().
			ExistesByEmail(mockCtx, gomock.Any()).
			Times(1).
			Return(true, expectedErr)
		mockRepository.
			EXPECT().
			Persist(gomock.Any(), gomock.Any()).
			Times(0)

		_, err := service.SaveUser(mockCtx, userSample)
		assert.ExpectedError(err)
		if !errors.Is(expectedErr, err) {
			t.Fatalf("Expected error to be type: %T but it is: %T", expectedErr, err)
		}
	})

	t.Run("Should return error from repository when persisting an user", func(t *testing.T) {
		mockRepository, service, userSample, assert := setupService(t)
		expectedErr := errors.New("error expected from save")

		mockRepository.
			EXPECT().
			ExistesByEmail(mockCtx, gomock.Any()).
			Times(1).
			Return(false, nil)
		mockRepository.
			EXPECT().
			Persist(gomock.Any(), gomock.Any()).
			Times(1).
			Return(userSample, expectedErr)

		_, err := service.SaveUser(mockCtx, userSample)

		assert.ExpectedError(err)
		if !errors.Is(expectedErr, err) {
			t.Fatalf("Expected error to be type: %T but it is: %T", expectedErr, err)
		}
	})
}

func TestUserService_FindByEmail(t *testing.T) {
	mockCtx := context.TODO()

	t.Run("Should find user by email successfully", func(t *testing.T) {
		mockRepository, service, userSample, assert := setupService(t)

		mockRepository.
			EXPECT().
			GetByEmail(mockCtx, userSample.Email).
			Times(1).
			Return(userSample, nil)

		usr, err := service.FindByEmail(mockCtx, userSample.Email)

		assert.NotExpectedError(err)
		if usr.Email != userSample.Email {
			t.Fatalf("Email didn't match expected email. Expected: %s - Actual: %s", userSample.Email, usr.Email)
		}
	})

	t.Run("Should return error from repository", func(t *testing.T) {
		mockRepository, service, userSample, assert := setupService(t)
		expectedErr := errors.New("error expected from find")

		mockRepository.
			EXPECT().
			GetByEmail(mockCtx, userSample.Email).
			Times(1).
			Return(userSample, expectedErr)

		_, err := service.FindByEmail(mockCtx, userSample.Email)

		assert.ExpectedError(err)
		if !errors.Is(expectedErr, err) {
			t.Fatalf("Expected error to be type: %T but it is: %T", expectedErr, err)
		}
	})
}

func TestUserService_UserExists(t *testing.T) {
	mockCtx := context.TODO()

	t.Run("Should test user existence successfully", func(t *testing.T) {
		mockRepository, service, userSample, assert := setupService(t)
		expectedResult := true

		mockRepository.
			EXPECT().
			ExistesByEmail(mockCtx, userSample.Email).
			Times(1).
			Return(expectedResult, nil)

		exists, err := service.UserExists(mockCtx, userSample.Email)

		assert.NotExpectedError(err)
		if exists != expectedResult {
			t.Fatalf("Expected result to be %v, but it is: %v", expectedResult, exists)
		}
	})

	t.Run("Should return error from repository", func(t *testing.T) {
		mockRepository, service, userSample, assert := setupService(t)
		expectedErr := errors.New("error expected from exists by email")

		mockRepository.
			EXPECT().
			ExistesByEmail(mockCtx, userSample.Email).
			Times(1).
			Return(false, expectedErr)

		_, err := service.UserExists(mockCtx, userSample.Email)

		assert.ExpectedError(err)
		if !errors.Is(expectedErr, err) {
			t.Fatalf("Expected error to be type: %T but it is: %T", expectedErr, err)
		}
	})
}

func setupService(t *testing.T) (*mocks.MockUserRepository, driver.UserService, domain.User, *test_helpers.Asserts) {
	assert := test_helpers.NewAssert(t)
	mockRepository := mocks.NewMockUserRepository(gomock.NewController(t))
	service := NewUserService(mockRepository)
	userSample := domain.User{
		ID:         primitive.NilObjectID,
		Email:      "tst@tst.tst",
		Name:       "Test",
		GivenName:  "Tester",
		FamilyName: "Test family",
		CreatedAt:  time.Now(),
		UpdatedAt:  nil,
	}
	return mockRepository, service, userSample, assert
}
