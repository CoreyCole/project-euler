package multiples

// IsMultipleOfThree returns true if passed int n is a multiple of 3
func IsMultipleOfThree(n int) bool {
    // if n is negative, make it positive
    if (n < 0) {
        n *= -1
    }
    
    if (n == 0) {
        return true
    } else if (n == 1) {
        return false
    } else {
        oddPosition := true
        odds := 0
        evens := 0
        for n != 0 {
            bitSet := n & 1
            if (oddPosition && bitSet == 1) {
                odds++
            } else if (bitSet == 1) {
                evens++
            }
            n = n >> 1
            oddPosition = !oddPosition
        }
        return IsMultipleOfThree(odds - evens)
    }
}