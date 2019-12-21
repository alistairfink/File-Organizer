package Config

import ()

type Config struct {
	Folders []Folder
}

type Folder struct {
	FolderName string
	Files      []string
}
