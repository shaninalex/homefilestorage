package filemanager

import (
	_ "github.com/lib/pq"
)

type FileManager struct {
	ServiceStorage string
}

func Initialize(storage_service_url string) *FileManager {
	fm := &FileManager{}
	fm.ServiceStorage = storage_service_url

	return fm
}
