package homepg

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const logFmt = "\nexpected:\n%v\n got \n%v\n"

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
			expectedBody:   `{"first":"Eliot","last":"Easterling","middle":"D","phone":"234-703-9147","dob":"1982-10-04T11:30:00+04:00"}`,
		},
	}
	for _, test := range tests {
		test := test // test.name will be wrong when tests are run in parrallel, otherwise

		t.Run(test.name, func(t *testing.T) {
			h := New(nil)
			h.Handler(test.out, test.in)

			if test.out.Code != test.expectedStatus {
				t.Logf(logFmt, test.expectedBody, test.out.Code)
				t.Fail()
			}

			body := test.out.Body.String()
			body = strings.Trim(body, "\n\r\t ")

			if body != test.expectedBody {
				t.Logf(logFmt, test.expectedBody, body)
				t.Fail()
			}
		})
	}
}
