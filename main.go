package main

import (
	"encoding/json"
	"github.com/alistairfink/File-Organizer/Config"
	"io/ioutil"
	"os"
)

func main() {
	args := (os.Args)[1:]
	if len(args) == 0 {
		println("Missing Source Folder")
		exit()
	}

	if len(args) == 1 {
		println("Missing Destination Folder")
		exit()
	}

	if len(args) == 2 {
		println("Missing Config JSON")
	}

	sourceFolder := args[0]
	files, err := ioutil.ReadDir(sourceFolder)
	if err != nil {
		println("Invalid Source:", sourceFolder)
		exit()
	}

	destinationFolder := args[1]
	err = os.MkdirAll(destinationFolder, os.ModePerm)
	if err != nil {
		println("Invalid Destination:", destinationFolder)
		exit()
	}

	configPath := args[2]
	var config Config.Config
	configFile, err := os.Open(configPath)
	defer configFile.Close()
	if err != nil {
		println("Invalid Config File")
		exit()
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	fileMap := make(map[string]os.FileInfo, len(files))
	for _, file := range files {
		if !file.IsDir() {
			fileMap[file.Name()] = file
		}
	}

	for _, folder := range config.Folders {
		rootFolder := destinationFolder + "/" + folder.FolderName
		err = os.MkdirAll(rootFolder, os.ModePerm)
		if err != nil {
			println("Error Creating Folder:", rootFolder)
		}

		for _, file := range folder.Files {
			if _, exists := fileMap[file]; !exists {
				println("Invalid File:", sourceFolder+"/"+file)
			} else {
				oldPath := sourceFolder + "/" + file
				newPath := rootFolder + "/" + file
				err := os.Rename(oldPath, newPath)
				if err != nil {
					println("Error Copying File:", file)
				}
			}
		}
	}
}

func exit() {
	println("Args: SourceFolder DestinationFolder Config")
	os.Exit(0)
}
