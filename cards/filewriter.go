// filewriter.go
package main

import "os"

// FileWriter defines the behavior for writing data to a file.
type FileWriter interface {
	WriteFile(filename string, data []byte, perm os.FileMode) error
}
