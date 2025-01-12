# Go Recursive Function Bug

This repository demonstrates a common error in recursive functions written in Go: stack overflow due to excessive recursion and an incorrect base case.  The `bug.go` file contains the erroneous code. The `bugSolution.go` file provides a corrected version.

**Bug:** The original function `myFunc` incorrectly handles the base case (x == 0) and lacks a mechanism to prevent stack overflow when passed large numbers.

**Solution:** The solution addresses both issues by correcting the base case and adding a check to terminate recursion before the stack overflows. It also adds comments to help clarify the fix. 

This example highlights the importance of careful handling of base cases and stack overflow prevention in recursive algorithms.