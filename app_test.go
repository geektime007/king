package king

import (
	"testing"
	"time"

	"github.com/geektime007/king/transport/http"
)

func TestApp(t *testing.T) {
	hs := http.NewServer()
	app := New(
		Name("geektime007"),
		Version("v0.0.1"),
		Server(hs),
	)
	time.AfterFunc(time.Second, func() {
		app.Stop()
	})
	if err := app.Run(); err != nil {
		t.Fatal(err)
	}
}
