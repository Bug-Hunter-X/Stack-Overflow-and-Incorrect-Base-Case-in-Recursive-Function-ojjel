func myFunc(x int) int {
    if x == 0 {
        return 1 // This is incorrect. Should be 0
    }
    return x * myFunc(x-1) // Stack overflow for large x
}