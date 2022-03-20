package usecases_test

import (
	"github.com/stretchr/testify/assert"
	"goCleanArch/internal/model"
	"goCleanArch/internal/repository"
	"goCleanArch/internal/usecases"
	"testing"
)

var (
	databaseURL = "host=localhost port=5432 user=azunai password=0000 dbname=apicl_test sslmode=disable"
)

func TestUser_Validate(t *testing.T) {
	db, teardown := repository.TestDB(t, databaseURL)
	defer teardown("author")

	repo := repository.NewSqlstore(db)
	// получаем в use case repository
	us := usecases.NewUseCase(repo)

	testCases := []struct {
		name    string
		u       *model.User
		isValid bool
	}{
		{
			name:    "valid",
			u:       model.Testuser(t),
			isValid: true,
		},
		{
			name:    "empty email",
			u:       &model.User{ID: "2", Firstname: "test", Lastname: "testov", Email: ""},
			isValid: false,
		},
		{
			name:    "invalid email",
			u:       &model.User{ID: "2", Firstname: "test", Lastname: "testov", Email: "123fxv"},
			isValid: false,
		},
		{
			name:    "too yuong",
			u:       &model.User{ID: "2", Firstname: "test", Lastname: "testov", Email: "123fxv", Age: "1"},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, us.Validation(tc.u))
			} else {
				assert.Error(t, us.Validation(tc.u))
			}
		})
	}
}
