package main

import (
	"io/ioutil"
	"os"
	// "path/filepath"
	"encoding/json"
	"github.com/alistairfink/File-Organizer/Config"
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

	// for _, file := range files {
	// 	println(file.Name(), file.ModTime().UTC().String())
	// }

	// var difffiles []os.FileInfo
	// err = filepath.Walk(sourceFolder, func(path string, info os.FileInfo, err error) error {
	// 	difffiles = append(difffiles, info)
	// 	return nil
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// for _, file := range difffiles {
	// 	println(file.Name(), file.ModTime().UTC().String())
	// }
}

func exit() {
	println("Args: SourceFolder DestinationFolder Config")
	os.Exit(0)
}
