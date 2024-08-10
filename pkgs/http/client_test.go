package http

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

type MockHttpClient struct {
	requests []string
}

// WARN: doesn't return a proper response
func (mock *MockHttpClient) Do(req *http.Request) (res *http.Response, err error) {
	mock.requests = append(mock.requests, fmt.Sprintf("%s %s", req.Method, req.URL.String()))
	return
}

func TestHttpClient(t *testing.T) {
	t.Run("Get() calls internal http client", func(t *testing.T) {
		mock_http := MockHttpClient{}
		client := HttpClient{Http: &mock_http}
		_, err := client.Get("https://test.deribit.com/api/v2/")
		if err != nil {
			t.Fatalf("unexpected error when performing GET request: %v", err)
		}

		have := mock_http.requests
		want := []string{"GET https://test.deribit.com/api/v2/"}
		if !reflect.DeepEqual(have, want) {
			t.Fatalf("unexpected request log - have %v, want %v", have, want)
		}
	})

	t.Run("Post() calls internal http client", func(t *testing.T) {
		mock_http := MockHttpClient{}
		client := HttpClient{Http: &mock_http}
		_, err := client.Post("https://test.deribit.com/api/v2/", "")
		if err != nil {
			t.Fatalf("unexpected error when performing POST request: %v", err)
		}

		have := mock_http.requests
		want := []string{"POST https://test.deribit.com/api/v2/"}
		if !reflect.DeepEqual(have, want) {
			t.Fatalf("unexpected request log - have %v, want %v", have, want)
		}
	})
}
