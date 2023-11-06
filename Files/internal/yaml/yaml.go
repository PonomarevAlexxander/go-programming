package yaml

import (
	"go/course/files/internal/file"

	"gopkg.in/yaml.v3"
)

func Unmarshal(data []byte) ([]file.File, error) {
	var fileList struct {
		Files []file.File
	}

	err := yaml.Unmarshal(data, &fileList)
	if err != nil {
		return nil, err
	}
	files := fileList.Files
	return files, nil
}
