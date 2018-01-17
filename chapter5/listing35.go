// Sample program to show how a bytes.Buffer can also be used
// with the io.Copy function
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// main is the entry point
func main() {
	var b bytes.Buffer

	// Write a string to the buffer
	b.Write([]byte("Hello"))

	// User Fprintf to concatenate a string to the Buffer
	fmt.Fprintf(&b, " World!")

	// Write the content of the Buffer to stdout
	io.Copy(os.Stdout, &b)
}
