package apiserver

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"goCleanArch/internal/model"
	mockdb "goCleanArch/internal/usecases/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_handlCreateuser(t *testing.T) {
	testcases := []struct {
		name         string
		u            *model.User
		buildStubs   func(log *mockdb.MockUseCaseLogic, u *model.User)
		expectedCode int
	}{
		{
			name: "valid",
			u:    &model.User{ID: "1", Firstname: "ivan", Lastname: "ivanov", Email: "test@test.com", Age: "18"},
			buildStubs: func(log *mockdb.MockUseCaseLogic, u *model.User) {
				log.EXPECT().Create(gomock.Eq(u)).Times(1).Return(nil)
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "empty",
			u:    &model.User{},
			buildStubs: func(log *mockdb.MockUseCaseLogic, u *model.User) {
				log.EXPECT().Create(gomock.Eq(u)).Times(1).Return(errors.New(""))
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid age",
			u:    &model.User{ID: "1", Firstname: "ivan", Lastname: "ivanov", Email: "test@test.com", Age: "1"},
			buildStubs: func(log *mockdb.MockUseCaseLogic, u *model.User) {
				log.EXPECT().Create(gomock.Eq(u)).Times(1).Return(errors.New("too young"))
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid email",
			u:    &model.User{ID: "1", Firstname: "ivan", Lastname: "ivanov", Email: "test", Age: "18"},
			buildStubs: func(log *mockdb.MockUseCaseLogic, u *model.User) {
				log.EXPECT().Create(gomock.Eq(u)).Times(1).Return(errors.New(""))
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			logic := mockdb.NewMockUseCaseLogic(ctrl)
			tc.buildStubs(logic, tc.u)
			//logic.EXPECT().Create(gomock.Eq(tc.u)).Times(1).Return(errors.New(""))

			s := Newserver(logic)
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.u)

			req, _ := http.NewRequest(http.MethodPost, "/user", b)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
