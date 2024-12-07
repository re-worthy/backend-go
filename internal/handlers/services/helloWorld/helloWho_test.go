package helloworld

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/re-worthy/backend-go/internal/handlers/dto"
	handlers "github.com/re-worthy/backend-go/internal/handlers/types"
	"github.com/re-worthy/backend-go/pkg/utils"
)

func TestValidHelloWhoHandler(t *testing.T) {
  const (
    expectedName = "123"
    jsonBody     = `{"name":"` + expectedName + `"}`
  )

	ts := httptest.NewServer(handlers.Adapter(HelloWhoHandler))
	defer ts.Close()

	resp, err := http.Post(ts.URL, "application/json", strings.NewReader(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Received http code %d, but expected %d", resp.StatusCode, http.StatusOK)
	}

	parsed, err := utils.ValidateJson[dto.THelloWorld](&resp.Body)
	if err != nil {
		t.Fatal(err)
	}

  if parsed.Hello != expectedName {
    t.Fatalf("Invalid gretting subject, got %s, but expected %s", parsed.Hello, expectedName)
  }
}

func TestBrokenJSONHelloWhoHandler(t *testing.T) {
  const (
    expectedName = "123"
    brokenJsonBody     = `{name":"` + expectedName + `"}`
  )

	ts := httptest.NewServer(handlers.Adapter(HelloWhoHandler))
	defer ts.Close()

	resp, err := http.Post(ts.URL, "application/json", strings.NewReader(brokenJsonBody))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("Received http code %d, but expected %d", resp.StatusCode, http.StatusOK)
	}

	parsed, err := utils.ValidateJson[dto.THelloWorld](&resp.Body)
	if err == nil {
		t.Fatalf("Expected to receive JSON parsing error, but got <nil>. Parsed = %v", parsed)
	}
}

func TestInvalidJSONHelloWhoHandler(t *testing.T) {
  const (
    expectedName = "123"
    brokenJsonBody     = `{"notname":"` + expectedName + `"}`
  )

	ts := httptest.NewServer(handlers.Adapter(HelloWhoHandler))
	defer ts.Close()

	resp, err := http.Post(ts.URL, "application/json", strings.NewReader(brokenJsonBody))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("Received http code %d, but expected %d", resp.StatusCode, http.StatusUnprocessableEntity)
	}

	parsed, err := utils.ValidateJson[dto.THelloWorld](&resp.Body)
	if err == nil {
		t.Fatalf("Expected to receive JSON parsing error, but got <nil>. Parsed = %v", parsed)
	}
}

func TestInvalidJSONNumberHelloWhoHandler(t *testing.T) {
  const (
    expectedName = "123"
    brokenJsonBody     = `{"name":` + expectedName + `}`
  )

	ts := httptest.NewServer(handlers.Adapter(HelloWhoHandler))
	defer ts.Close()

	resp, err := http.Post(ts.URL, "application/json", strings.NewReader(brokenJsonBody))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("Received http code %d, but expected %d", resp.StatusCode, http.StatusUnprocessableEntity)
	}

	parsed, err := utils.ValidateJson[dto.THelloWorld](&resp.Body)
	if err == nil {
		t.Fatalf("Expected to receive JSON parsing error, but got <nil>. Parsed = %v", parsed)
	}
}

