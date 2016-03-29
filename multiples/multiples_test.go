package multiples

import (
    "testing"
    "strconv"
)

// TestIsMultipleOfThree validates functionality of isMultipleOfThree
func TestIsMultipleOfThree(t *testing.T) {
    for i := 0; i < 300; i++ {
        if (IsMultipleOfThree(i) != (i%3 == 0)) {
            t.Errorf(strconv.Itoa(i) + " is a multiple of 3 but function returned false")
        }
    }
}