func myFunc(x int) int {
    if x == 0 {
        return 0 // Correct base case
    }
    if x < 0 {
        return -1 // Handle negative numbers 
    }
    if x > 1000 { // Prevent stack overflow
        return -2
    }
    return x * myFunc(x-1)
}