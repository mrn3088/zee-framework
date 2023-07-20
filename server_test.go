package web_framework

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHTTP_Start(t *testing.T) {
	h := NewHTTP(WithHTTPServerStop(nil))
	go func() {
		err := h.Start(":8080")
		if err != nil && err != http.ErrServerClosed {
			t.Fail()
		}
	}()
	fmt.Println("calling stop")
	err := h.Stop()
	if err != nil {
		t.Fail()
	}
}
