package util

import (
	"bytes"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/yaml"
	"path/filepath"
	"strings"
)

//首字母大写
func FirstUpper(org string) string {
	return strings.ToUpper(org[:1]) + org[1:]
}

//yaml to 字节数组
func Yamls2Bytes(rootPath string, files []string) ([][]byte,error) {
	yamls := make([][]byte, len(files))
	for i, name := range files {
		yamlBytes, err := ioutil.ReadFile(filepath.Join(rootPath, name))
		if err != nil {
			return nil,err
		}
		yamls[i] = yamlBytes

	}
	return yamls,nil
}

//yaml to json
func Yamls2Jsons(data [][]byte) ([][]byte,error) {
	jsons := make([][]byte, 0)
	for _, yamlBytes := range data {
		yamls := bytes.Split(yamlBytes, []byte("---"))
		for _, v := range yamls {
			if len(v) == 0 {
				continue
			}
			obj, err := yaml.ToJSON(v)
			if err != nil {
				return nil,err
			}
			jsons = append(jsons, obj)
		}
	}
	return jsons,nil
}