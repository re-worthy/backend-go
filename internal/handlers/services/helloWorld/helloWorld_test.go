package helloworld

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/re-worthy/backend-go/internal/handlers/dto"
	handlers "github.com/re-worthy/backend-go/internal/handlers/types"
	"github.com/re-worthy/backend-go/pkg/utils"
)

func TestHelloWorldHandler(t *testing.T) {
	const (
		expectedName = "world"
	)

	ts := httptest.NewServer(handlers.Adapter(HelloWorldHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
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
