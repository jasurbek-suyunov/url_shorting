package helper

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func GenerateQrCode(text string) (string, error) {
	url := "https://api.qrserver.com/v1/create-qr-code/?size=150x150&data=" + text

	// request
	resp, err := http.Get(url)

	// check error
	if err != nil {
		return "", err
	}

	// close body
	defer resp.Body.Close()
	fmt.Println(resp.Body)
	path, err := SaveFile(resp.Body)
	if err != nil {
		return "", err
	}
	return path, nil
}

func SaveFile(file io.ReadCloser) (string, error) {

	dir := "uploads/" + pathGenerate()
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		fmt.Println(err.Error())
	}

	path := dir + uuid.New().String() + ".jpg"

	out, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer out.Close()

	out.ReadFrom(file)
	return path, err
}

func pathGenerate() string {
	return uuid.New().String()[0:3] + "/" + uuid.New().String() + "/"
}
