package checkout

import (
	"testing"
	"time"
)

func TestValidateSession(t *testing.T) {
	c := NewTokenCache(5 * time.Second)
	if err := c.ValidateSession("tok-live-checkout-001"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
