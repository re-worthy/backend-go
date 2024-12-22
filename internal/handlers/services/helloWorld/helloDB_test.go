package helloworld

import (
	"net/http"
	"net/http/httptest"
	"testing"

	db_init "github.com/re-worthy/backend-go/internal/db/init"
	"github.com/re-worthy/backend-go/internal/handlers/dto"
	"github.com/re-worthy/backend-go/internal/handlers/tests"
	handlers "github.com/re-worthy/backend-go/internal/handlers/types"
	"github.com/re-worthy/backend-go/pkg/utils"
)

func TestHelloDB(t *testing.T) {
	baseHandler, onclose, err := tests.NewTestBaseHandler()
	if err != nil {
		t.Fatal(err)
	}
	defer onclose()

	db_init.InitDB(baseHandler.DB)

	ts := httptest.NewServer(handlers.Adapter(HelloDBHandler, baseHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Received http code %d, but expected %d", resp.StatusCode, http.StatusOK)
	}

	parsed, err := utils.ValidateJson[dto.THelloDB](&resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if parsed.Counter != 1 {
		t.Fatalf("Invalid counter value, got %d, but expected %d", parsed.Counter, 1)
	}
}

func TestHelloDBMultiple(t *testing.T) {
	const (
		cntRepeatCoutnter = 5
	)
	baseHandler, onclose, err := tests.NewTestBaseHandler()
	if err != nil {
		t.Fatal(err)
	}
	defer onclose()

	db_init.InitDB(baseHandler.DB)

	ts := httptest.NewServer(handlers.Adapter(HelloDBHandler, baseHandler))
	defer ts.Close()

	for i := 1; i <= cntRepeatCoutnter; i++ {
		resp, err := http.Get(ts.URL)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Received http code %d, but expected %d", resp.StatusCode, http.StatusOK)
		}

		parsed, err := utils.ValidateJson[dto.THelloDB](&resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		if parsed.Counter != i {
			t.Fatalf("Invalid counter value, got %d, but expected %d", parsed.Counter, i)
		}
	}
}
