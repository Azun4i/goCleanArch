package apiserver

import (
	"bytes"
	"encoding/json"
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
		expectedCode int
	}{
		{
			name:         "valid",
			u:            &model.User{ID: "1", Firstname: "ivan", Lastname: "ivanov", Email: "test@test.com", Age: "18"},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid age",
			u:            &model.User{ID: "1", Firstname: "ivan", Lastname: "ivanov", Email: "test@test.com", Age: "1"},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name:         "invalid email",
			u:            &model.User{ID: "1", Firstname: "ivan", Lastname: "ivanov", Email: "test", Age: "18"},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockUseCaseLogic(ctrl)
			store.EXPECT().Create(gomock.Eq(tc.u)).Times(1)
			s := Newserver(store)

			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.u)

			req, _ := http.NewRequest(http.MethodPost, "/user", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
