package main

import (
	"testing"
)

func TestGeneratePinCodeList(t *testing.T) {
	quantity := 100
	pinCodes := generatePinCodeList(quantity)

	if len(pinCodes) != quantity {
		t.Errorf("expected %d pincodes, but got %d", quantity, len(pinCodes))
	}

	seen := make(map[string]bool)
	for _, pin := range pinCodes {
		if seen[pin] {
			t.Errorf("duplicate pincode found: %s", pin)
		}
		seen[pin] = true
	}

	for _, pin := range pinCodes {
		if len(pin) != 4 {
			t.Errorf("invalid pincode: %s (4 digits expected)", pin)
		}
	}
}

func Test_convertPinToString(t *testing.T) {
	result := convertPinToString([]int{1234})
	if result != "1234" {
		t.Error("Fail to convert pincode to string")
	}
}
