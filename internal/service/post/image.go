package post

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/gofrs/uuid"
)

const imgMaxSize = 5 << 20

const imagesDir = "images/"

func UploadPicture(file multipart.File, header *multipart.FileHeader) (string, string, error) {
	if file == nil {
		return "", "", http.ErrMissingFile
	}
	defer file.Close()

	name, path, fileType := hashImage(header.Filename)
	if err := ensureImagesDir(fileType); err != nil {
		return "", "", err
	}

	dst, err := os.Create(path)
	if err != nil {
		return "", "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", "", err
	}
	return name, fileType, nil
}
func hashImage(file string) (string, string, string) {
	hashImgName, err := uuid.NewV4()
	if err != nil {
		log.Println(err)
		return "", "", ""
	}

	fileName := strings.ReplaceAll(string(hashImgName.String()), "/", "@")
	fileName = strings.ReplaceAll(string(fileName), ".", "@")
	fileType := strings.Split(file, ".")
	typeF := fileType[len(fileType)-1]

	return fileName, fmt.Sprintf("%s/%s/%s.%s", imagesDir, typeF, fileName, typeF), typeF
}

func ensureImagesDir(typeF string) error {
	dirPath := fmt.Sprintf("%s/%s", imagesDir, typeF)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
