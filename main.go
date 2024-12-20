package main

import (
	"embed"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Embed the skeleton files into the Go binary
//

//go:embed skeleton_files/*
var skeletonFiles embed.FS

// Function to initialize the project
func initProject() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		os.Exit(1)
	}

	// Define the folder where the skeleton files will be copied
	skeletonDir := "skeleton_files"
	err = copySkeletonFiles(skeletonDir, dir)
	if err != nil {
		fmt.Println("Error copying skeleton files:", err)
		os.Exit(1)
	}

	fmt.Println("Project skeleton initialized!")
}

// Function to copy the skeleton files from the embedded FS to the current directory
func copySkeletonFiles(srcDir, destDir string) error {
	entries, err := skeletonFiles.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(srcDir, entry.Name())
		destPath := filepath.Join(destDir, entry.Name())

		// Check if the entry is a directory or a file
		if entry.IsDir() {
			// Create the directory in the destination
			err := os.MkdirAll(destPath, 0755)
			if err != nil {
				return err
			}
			// Recursively copy the contents of the directory
			err = copySkeletonFiles(srcPath, destPath)
			if err != nil {
				return err
			}
		} else {
			// Copy the file to the destination
			content, err := skeletonFiles.ReadFile(srcPath)
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(destPath, content, 0644)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: skeleton init")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		initProject()
	default:
		fmt.Println("Unknown command. Use 'skeleton init'.")
	}
}
