package utils

import (
	"io/ioutil"
	"path"
	"strings"
)

func ReadSourcesFromDir(dirname string) ([][]byte, error) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	var sources [][]byte
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".go") && !strings.HasSuffix(f.Name(),"_test.go") {
			source, err := ioutil.ReadFile(path.Join(dirname, f.Name()))
			if err != nil {
				return nil, err
			}

			sources = append(sources, source)
		}
	}

	return sources, nil
}