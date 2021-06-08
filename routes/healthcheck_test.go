package routes

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockDb struct {
	mock.Mock
}

func (m *mockDb) Up(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func TestHealthCheckResourceGet(t *testing.T) {
	assertions := assert.New(t)

	testCases := []struct {
		description    string
		mockDbUpReturn error
		status         string
	}{
		{"database is available", nil, "up"},
		{"database is unavailable", errors.New(""), "down"},
	}

	for _, c := range testCases {
		db := &mockDb{}
		db.On("Up", mock.Anything).Return(c.mockDbUpReturn)

		resource := NewHealthCheckResource(db)
		ts := httptest.NewServer(resource.Routes())
		defer ts.Close()

		t.Run(c.description, func(t *testing.T) {
			res, err := http.Get(ts.URL)
			assertions.Nil(err)

			status, err := io.ReadAll(res.Body)
			res.Body.Close()

			assertions.Nil(err)
			assertions.Equal([]byte(c.status), status)
		})
	}
}
