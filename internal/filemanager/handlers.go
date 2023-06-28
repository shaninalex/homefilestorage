package filemanager

import "os"

type FileResponse struct {
	Name       string `json:"name"`
	MimeType   string `json:"mime_type"`
	Size       uint   `json:"size"`
	SystemPath string `json:"system_path"`
	Hash       string `json:"hash"`
}

func (filemanager *FileManager) SaveFile(filename string, d []byte) (*FileResponse, error) {
	// TODO: use filemanager storage path
	tmpfile, err := os.Create("./" + filename)
	defer tmpfile.Close()
	if err != nil {
		return nil, err
	}
	tmpfile.Write(d)

	// TODO: create FileResponse object and fill the fields
	return nil, nil
}
