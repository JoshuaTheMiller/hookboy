package main

import (
	"crypto/sha256"
	"errors"
	"io"
	"log"
	"os"
)

func fileOrFolderExists(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

func isGrapplerAlreadyInstalled() (bool, error) {
	var gitFolder = ".git"

	if !fileOrFolderExists(gitFolder) {
		return false, errors.New("This must be ran at the top level of a git repository")
	}

	var grapplerInstallFile = ".git/hooks/grappler_install.yml"
	if !fileOrFolderExists(grapplerInstallFile) {
		return false, nil
	}

	return true, nil
}

func openFileAndGenerateConfigHash(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 30*1024)
	sha256 := sha256.New()

	for {
		n, err := file.Read(buf)
		if n > 0 {
			_, err := sha256.Write(buf[:n])

			if err != nil {
				log.Fatal(err)
			}
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Read %d bytes: %v", n, err)
			break
		}
	}

	sum := sha256.Sum(nil)
	return string(sum)
}
