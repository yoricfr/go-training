// Ouvre un fichier zip contenant du texte et l'affiche sur la console
package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	// Create a io.Reader from the file fichier zip
	reader, err := os.Open("text.zip")
	if err != nil {
		log.Fatal(err)
	}

	// Copy data from the reader (zip file) to the console
	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		log.Fatal(err)
	}
	reader.Close()

	// The result is gerbish... (compressed)
	// So, let's see how to use the zip package to actually read the zip file
	fmt.Println()

	// Get a pointer to a zip Reader that list all the files
	pReader, err := zip.OpenReader("text.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer pReader.Close()

	// Loop for each file in the zip archive
	for _, f := range pReader.File {
		fmt.Printf("Content of %s:\n", f.Name)

		// Get the io reader to access the file data
		ioReader, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer ioReader.Close()

		// Copy the data from the io reader to the console
		_, err = io.Copy(os.Stdout, ioReader)
		if err != nil {
			log.Fatal(err)
		}
	}

}
