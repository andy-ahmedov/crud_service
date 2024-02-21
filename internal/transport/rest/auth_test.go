package rest

import (
	"bytes"
	"context"
	"net/http/httptest"
	"testing"

	"github.com/andy-ahmedov/crud_service/internal/domain"
	"github.com/andy-ahmedov/crud_service/internal/service"
	mock_service "github.com/andy-ahmedov/crud_service/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestRest_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUserStorage, user domain.User)

	testTable := []struct {
		name               string
		inputBody          string
		inputUser          domain.User
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test", "email":"test@gmail.com", "password":"qwerty"}`,
			inputUser: domain.User{
				Name:     "Test",
				Email:    "test@gmail.com",
				Password: "qwerty",
			},
			mockBehavior: func(s *mock_service.MockUserStorage, user domain.User) {
				s.EXPECT().CreateUser(context.Background(), user).Return(nil)
			},
			expectedStatusCode: 200,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockUserStorage(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Users{Repo: auth}

			handler := NewHandler(nil, services)

			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}
