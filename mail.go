package main

import (
    "io/ioutil"
    "log"
	"path/filepath"
	"os"
)

func main() {
	namePrefix := os.Args[1]
    files, err := ioutil.ReadDir(".")
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
		fileName := file.Name()
		if (!file.IsDir() && filepath.Ext(fileName) == ".mp3") {
			os.Rename(fileName, namePrefix + "_" + fileNameWithoutExtSliceNotation(fileName) + ".mp3")
		}
    }
}

func fileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}