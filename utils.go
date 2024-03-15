package main

import (
	"bytes"
	"io"
	"math/rand"
	"mime/multipart"
	"strings"
)

func spawnUUID(length int) string {
	characters := "ABCDEFGHJKMNP9gqQRSToOLVvI1lWXYZabcdefhijkmnprstwxyz2345678"
	var result strings.Builder

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(characters))
		result.WriteByte(characters[randomIndex])
	}

	return result.String()
}

func prepareMultipartFormData(fileReader io.Reader, fileName, uuidValue, recModeValue string) (io.Reader, string) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// 注意：不再打开文件，直接使用提供的io.Reader
	fw, err := w.CreateFormFile("file", fileName) // 使用一个文件名
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(fw, fileReader); err != nil { // 直接复制io.Reader到表单
		panic(err)
	}

	w.WriteField("uuid", uuidValue)
	w.WriteField("rec_mode", recModeValue)

	w.Close()

	return &b, w.FormDataContentType()
}
