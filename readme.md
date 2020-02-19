#  search用来进行文件夹读取等操作、后续可以来获取网页链接

#  获取目录下的文件和目录

```
// ExtractPath 查找目录下的所有文件和文件夹
// 返回文件名切片、文件夹名切片和error
func ExtractPath(mpath string) (files, paths []string, err error)
```