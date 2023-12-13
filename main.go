package main

import (
	// "fmt"
	// "os"
	// "short-link/RandStrings"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"shortLink/HashString"
)


func main() {
	// fmt.Println("enter an interger number for generatin string:")
	// var i int
	// fmt.Scan(&i)
	// var p *int = &i
	// if *p < 350 {
	// 	fmt.Println("the number must be bigger than 3 \n please enter a number bigger than 350")
	// 	return
	// }
	// // fmt.Println(RandString.RandStringRunes(*p))
	// data := RandString.RandStringRunes(*p)
	// filePath := data[:8] + ".txt"
	// file, err := os.Create(filePath)
	// if err != nil {
	// 	fmt.Println("Error creating file:", err)
	// 	return
	// }
	// defer file.Close() // Make sure to close the file when you're done with it
	// _, err = file.Write([]byte(data))
	// if err != nil {
	// 	fmt.Println("Error writing to file:", err)
	// 	return
	// }
	dirPath := "./"
	files, err := getTxtFilesInDirectory(dirPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("List of .txt files:")
	for _, file := range files {
		fmt.Println(file)
	}
	for _, filePath := range files {
		content, err := readContentFromFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", filePath, err)
			continue
		}
		file, err := os.Create("hash-" + filepath.Base(filePath))
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		// Store the content in a new variable (you can process or use it as needed)
		newVariable := content
		data := HashString.HashString(content)
		_, err = file.Write([]byte(data))
		// Print or use the new variable as needed
		fmt.Printf("Content of %s:\n%s\n", filePath, newVariable)
	}
}

func getTxtFilesInDirectory(dirPath string) ([]string, error) {
	var txtFiles []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		// Check if the file has a ".txt" extension
		if filepath.Ext(path) == ".txt" {
			txtFiles = append(txtFiles, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return txtFiles, nil
}

func readContentFromFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
