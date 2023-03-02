package pluginmanager

import (
	"io/ioutil"
	"os"
)

type scanCallback func(filePath string) error

func scanPath(path string, callBack scanCallback) error {
	files, err := getAllFiles(path)
	if err != nil {
		return err
	}
	for _, f := range files {
		err := callBack(f)
		if err != nil {
			return err
		}
	}
	return nil
}

func getAllFiles(path string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	sep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() {
			dirs = append(dirs, path+sep+fi.Name())
			getAllFiles(path + sep + fi.Name())
		} else {
			files = append(files, path+sep+fi.Name())
		}
	}

	// 子文件夹
	for _, subDir := range dirs {
		temp, _ := getAllFiles(subDir)
		files = append(files, temp...)
	}
	return files, nil
}
