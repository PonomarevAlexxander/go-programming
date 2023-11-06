package yaml

import (
	"gopkg.in/yaml.v3"
)

type File struct {
	Path      string `yaml:"path"`
	Substring string `yaml:"substring"`
}

func Unmarshal(data []byte) ([]File, error) {
	var fileList struct {
		Files []File
	}

	err := yaml.Unmarshal(data, &fileList)
	if err != nil {
		return nil, err
	}
	return fileList.Files, nil
}
