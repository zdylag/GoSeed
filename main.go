// Application which greets you.
package main

import "fmt"

func main() {
	fmt.Println(NeedsAComment())
}

func NeedsAComment() string {
	return "Hi!"
}

// This isnt good enough.
func NeedsABetterComment() string {
	return "Hi!"
}
