package study

import (
	"fmt"
	"github.com/disintegration/imaging"
	"io/ioutil"
)

func RecursionDirectory(dirPath string, ignore func(name string) bool) []string {//获取文件
	var fileList []string
	rd, _ := ioutil.ReadDir(dirPath)
	for _, f := range rd {
		if f.IsDir() {//文件夹
			subFileList := RecursionDirectory(fmt.Sprintf("%s/%s", dirPath, f.Name()), ignore)
			fileList = append(fileList, subFileList...)
		} else {//文件
			file := f.Name()
			if ((ignore != nil) && !ignore(file)) || ignore == nil {
				file = fmt.Sprintf("%s/%s", dirPath, file)
				fileList = append(fileList, file)
			}
		}
	}
	return fileList
}

func ResizeImage(filename string, newfilename string, width int, height int) error {//修改尺寸
	src, err := imaging.Open(filename)//打开文件
	if err != nil {
		return err
	}
	src = imaging.Resize(src, width, height, imaging.Lanczos)//改尺寸

	err = imaging.Save(src, newfilename)//保存
	return err
}
