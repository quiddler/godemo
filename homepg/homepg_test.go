package homepg

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const logFmt = "expected %v got %v\n"

func Test(t *testing.T) {
	tests := []struct {
		name           string
		in             *http.Request
		out            *httptest.ResponseRecorder
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "home-route",
			in:             httptest.NewRequest("GET", "/", nil),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			expectedBody:   "whell i do love me some testin' dare boi",
		},
	}
	for _, test := range tests {
		test := test // test.name will be wrong when tests are run in parrallel totherwise

		t.Run(test.name, func(t *testing.T) {
			h := New(nil)
			h.Handler(test.out, test.in)

			if test.out.Code != test.expectedStatus {
				t.Logf(logFmt, test.expectedBody, test.out.Code)
				t.Fail()
			}

			body := test.out.Body.String()
			if body != test.expectedBody {
				t.Logf(logFmt, test.expectedBody, body)
				t.Fail()
			}
		})
	}
}
