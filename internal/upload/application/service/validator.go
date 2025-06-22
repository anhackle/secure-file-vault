package service

import "path/filepath"

var FileContentTypes = []string{
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	"application/vnd.ms-excel",
	"text/csv",
	"text/plain",
	"application/pdf",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	"application/msword",
	"application/vnd.openxmlformats-officedocument.presentationml.presentation",
	"application/vnd.ms-powerpoint",
	"image/jpeg",
	"image/png",
	"image/gif",
	"image/jpg",
}

var FileExtensions = []string{
	".xlsx",
	".xls",
	".csv",
	".txt",
	".pdf",
	".docx",
	".doc",
	".pptx",
	".ppt",
	".jpg",
	".jpeg",
	".png",
	".gif",
}

func GetFileExtension(fileName string) string {
	return filepath.Ext(filepath.Base(fileName))
}

func CheckContentType(contentTypeInput string) bool {
	for _, fileContentType := range FileContentTypes {
		if contentTypeInput == fileContentType {
			return true
		}
	}

	return false
}

func CheckFileExtension(fileExtensionInput string) bool {
	for _, fileExtension := range FileExtensions {
		if fileExtensionInput == fileExtension {
			return true
		}
	}

	return false
}
