package search

import (
	"io/ioutil"
	"path/filepath"
)

// ExtractPath 查找目录下的所有文件和文件夹
// 返回文件名切片、文件夹名切片和error
func ExtractPath(mpath string) (files, paths []string, err error) {
	infos, err := ioutil.ReadDir(mpath)
	if err != nil {
		return
	}

	// sep := string(os.PathSeparator)
	for _, info := range infos {
		if info.IsDir() {
			paths = append(paths, filepath.Join(mpath, info.Name()))
		} else {
			files = append(files, filepath.Join(mpath, info.Name()))
		}
	}
	return
}
