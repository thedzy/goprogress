package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)

	// Walk the directory recursively and run Go files
	err := filepath.WalkDir(dirname, func(path string, file fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// if not dir and a go file
		if !file.IsDir() && filepath.Ext(path) == ".go" {
			if path == filename {
				// If it's this file, bail
				return nil
			}

			// Get the absolute path
			absPath, err := filepath.Abs(path)
			if err != nil {
				fmt.Println("Error:", err)
				return nil
			}

			// Print the file name, so you know which examople is running
			fmt.Println("File:", absPath)

			// Build the command to run the Go file
			cmd := exec.Command("go", "run", absPath)

			//  Output to stdout and err
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			// Start the command
			err = cmd.Start()
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}

			// Wait for the command to finish
			err = cmd.Wait()
			if err != nil {
				fmt.Println("Error:", err)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}

}
