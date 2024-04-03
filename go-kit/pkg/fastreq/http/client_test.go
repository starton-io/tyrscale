package fastreq

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/starton-io/tyrscale/go-kit/pkg/fastreq/core"
)

// RequestTest defines the structure for each test case
type RequestTest struct {
	Method                     string
	URL                        string
	Body                       []byte
	Headers                    map[string]string
	CreateClientWithoutTimeout bool
	ExpectedStatus             int
	ExpectedBody               string
	ExpectedBodyMsg            string
}

func GenerateTestUserAgent(appName, appVersion string) string {
	return fmt.Sprintf("%s/%s (Test; Go)", appName, appVersion)
}

func MakeFakeHttpServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Respond based on the URL
		switch r.URL.Path {
		case "/test-get":
			if r.Method != "GET" {
				http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
				return
			}
			userAgent := r.Header.Get("User-Agent")
			fmt.Println(userAgent)
			if userAgent != "TestApp/1.0 (Test; Go)" {
				http.Error(w, "Invalid user agent", http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Mocked GET response"))
		case "/test-post":
			if r.Method != "POST" {
				http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "Mocked POST response"}`))
		case "/test-put":
			if r.Method != "PUT" {
				http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "Mocked PUT response"}`))
		case "/test-patch":
			if r.Method != "PATCH" {
				http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "Mocked PATCH response"}`))
		case "/test-delete":
			if r.Method != "DELETE" {
				http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "Not found", http.StatusNotFound)
		}
	}))
}

func TestFastHttpClientRequests(t *testing.T) {
	// Define your test cases
	tests := []RequestTest{
		{
			Method:                     "GET",
			URL:                        "/test-get",
			Headers:                    map[string]string{"Content-Type": ContentTypePlainText},
			ExpectedStatus:             http.StatusOK,
			ExpectedBody:               "",
			CreateClientWithoutTimeout: true,
		},
		{
			Method:          "POST",
			URL:             "/test-post",
			Body:            []byte("test data"),
			Headers:         map[string]string{"Content-Type": ContentTypeJson},
			ExpectedStatus:  http.StatusOK,
			ExpectedBodyMsg: "Mocked POST response",
		},
		{
			Method:          "PATCH",
			URL:             "/test-patch",
			Body:            []byte("test data"),
			Headers:         map[string]string{"Content-Type": ContentTypeJson},
			ExpectedStatus:  http.StatusOK,
			ExpectedBodyMsg: "Mocked PATCH response",
		},
		{
			Method:          "PUT",
			URL:             "/test-put",
			Headers:         map[string]string{"Content-Type": ContentTypeJson},
			ExpectedStatus:  http.StatusOK,
			ExpectedBodyMsg: "Mocked PUT response",
		},
		{
			Method:         "DELETE",
			URL:            "/test-delete",
			Headers:        map[string]string{"Content-Type": ContentTypeJson},
			ExpectedStatus: http.StatusNoContent,
			ExpectedBody:   "",
		},
		// Add more tests for PUT, DELETE, etc.
	}

	// Setup the test server
	server := MakeFakeHttpServer()
	defer server.Close()

	// Create an instance of your client
	client := NewBuilder().
		SetHeaders(map[string]string{
			"Content-Type": ContentTypePlainText,
		}).
		SetResponseTimeout(1000 * time.Millisecond).
		SetMaxConnections(100).
		SetUserAgent(GenerateTestUserAgent("TestApp", "1.0")).
		Build()

	// Run the tests
	for _, test := range tests {
		if test.CreateClientWithoutTimeout {
			client = NewBuilder().
				SetHeaders(map[string]string{
					"Content-Type": ContentTypePlainText,
				}).
				SetMaxConnections(100).
				SetUserAgent(GenerateTestUserAgent("TestApp", "1.0")).
				DisableTimeouts(true).
				Build()
		}
		switch test.Method {
		case "GET":
			resp, err := client.Get(server.URL+test.URL, nil)
			if err != nil {
				t.Log(resp)
				t.Fatalf("Failed to make request: %v", err)
			}
			if resp.StatusCode != test.ExpectedStatus {
				t.Fatalf("Expected status %d, got %d", test.ExpectedStatus, resp.StatusCode)
			}
		case "POST", "PUT", "PATCH":
			var resp *core.Response
			var err error
			if test.Method == "POST" {
				resp, err = client.Post(server.URL+test.URL, test.Body, test.Headers)
			} else if test.Method == "PUT" {
				resp, err = client.Put(server.URL+test.URL, test.Body, test.Headers)
			} else if test.Method == "PATCH" {
				resp, err = client.Patch(server.URL+test.URL, test.Body, test.Headers)
			}
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			if resp.StatusCode != test.ExpectedStatus {
				t.Fatalf("Expected status %d, got %d", test.ExpectedStatus, resp.StatusCode)
			}
			var res map[string]interface{}
			resp.Unmarshal(&res)
			if res["message"] != test.ExpectedBodyMsg {
				t.Fatalf("Expected body %s, got %s", test.ExpectedBodyMsg, string(resp.Body))
			}
		case "DELETE":
			resp, err := client.Delete(server.URL+test.URL, test.Headers)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			if resp.StatusCode != test.ExpectedStatus {
				t.Fatalf("Expected status %d, got %d", test.ExpectedStatus, resp.StatusCode)
			}
		case "OPTIONS":
			resp, err := client.Options(server.URL+test.URL, test.Headers)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			if resp.StatusCode != test.ExpectedStatus {
				t.Fatalf("Expected status %d, got %d", test.ExpectedStatus, resp.StatusCode)
			}
		}

	}
}
