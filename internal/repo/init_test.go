package repo

import (
	"testing"
)

func TestNewIpConfigModule(t *testing.T) {
	// Call the New function to create an IpConfigModule instance
	ipConfigModule := New()

	// Assert that the IpConfigModule instance is not nil
	if ipConfigModule == nil {
		t.Error("Expected non-nil IpConfigModule instance, got nil")
	}
}
