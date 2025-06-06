package serializer

import "go-cloud-disk/model"

// file serializer
type File struct {
	Uuid     string `json:"file_id"`
	FileName string `json:"filename"`
	FileType string `json:"filetype"`
	Size     int64  `json:"size"`
}

func BuildFile(file model.File) File {
	return File{
		Uuid:     file.Uuid,
		FileName: file.FileName,
		FileType: file.FilePostfix,
		Size:     file.Size,
	}
}

func BuildFiles(files []model.File) (FileSerializers []File) {
	for _, f := range files {
		FileSerializers = append(FileSerializers, BuildFile(f))
	}
	return
}
