package zip

import (
	"archive/zip"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type ZipFile struct {
	outPath string
	Files   []string
}

func (z *ZipFile) AddFile(filePath string) error {
	if _, err := os.Stat(filePath); err != nil {
		return err
	}

	z.Files = append(z.Files, filePath)
	return nil
}

func (z *ZipFile) SetOutPath(outPath string) {
	z.outPath = outPath
}

func (z *ZipFile) Write() error {
	if len(z.Files) == 0 {
		return errors.New("No files added to zip, not writing empty zip")
	}

	zipFile, err := os.Create(z.outPath)
	if err != nil {
		return err
	}

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, v := range z.Files {
		fileStat, err := os.Stat(v)
		if err != nil {
			return err
		}

		data, err := ioutil.ReadFile(v)
		if err != nil {
			return err
		}

		header := &zip.FileHeader{
			Name:           v,
			Modified:       fileStat.ModTime(),
			CreatorVersion: 3 << 8,
			Method:         zip.Deflate,
		}

		header.SetMode(os.FileMode(fileStat.Mode()))
		file, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		_, err = file.Write(data)
	}
	return nil
}

func NewZipFile(outPath string) (*ZipFile, error) {
	var zipFile *ZipFile

	if _, err := os.Stat(outPath); os.IsExist(err) {
		return zipFile, errors.New(fmt.Sprintf("File %s already exists", outPath))
	}

	zipFile = new(ZipFile)
	zipFile.SetOutPath(outPath)

	return zipFile, nil
}
