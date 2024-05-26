package service

import (
	"archive/zip"
	"io"
)

type ZipServiceStruct struct {
}

func NewZipService() *ZipServiceStruct {
	return &ZipServiceStruct{}
}

// Добавление файла в zip архив
func (c *ZipServiceStruct) AddFileToArchive(fileName string, srcFile io.Reader, zipw *zip.Writer) error {
	dstfile, err := zipw.Create(fileName)
	if err != nil {
		return err
	}

	if _, err := io.Copy(dstfile, srcFile); err != nil {
		return err
	}

	return nil
}
