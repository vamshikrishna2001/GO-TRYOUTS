package main

import (
	"io"
	"os"
)

// func printer(file *os.File, s string) {
// 	file.Write([]byte(s))
// 	file.Close()
// }

func printer(file io.Writer, s string) {
	file.Write([]byte(s))
	// file.Close() we cannot use this ... which is what we want the intended purpose of printer function is to print not the closing stuff
}

func main() {
	printer(os.Stdout, "hii")

}
