package file

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	imgext "github.com/shamsher31/goimgext"
	"hope/pkg/util/str"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func PathCreate(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

// PathExist 判断目录是否存在
func PathExist(addr string) bool {
	s, err := os.Stat(addr)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}

func FileCreate(content *str.Buffer, name string) {
	file, err := os.Create(name)
	if err != nil {
		log.Println(err)
	}
	_, err = file.WriteString(content.String())
	if err != nil {
		log.Println(err)
	}
	file.Close()
}

type ReplaceHelper struct {
	Root    string //路径
	OldText string //需要替换的文本
	NewText string //新的文本
}

func (h *ReplaceHelper) DoWalk() error {

	return filepath.Walk(h.Root, h.walkCallback)

}

func (h ReplaceHelper) walkCallback(path string, f os.FileInfo, err error) error {

	if err != nil {
		return err
	}
	if f == nil {
		return nil
	}
	if f.IsDir() {
		log.Println("DIR:", path)
		return nil
	}

	//文件类型需要进行过滤

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		//err
		return err
	}
	content := string(buf)
	//替换
	newContent := strings.Replace(content, h.OldText, h.NewText, -1)

	//重新写入
	return ioutil.WriteFile(path, []byte(newContent), 0)
}

func FileMonitoringById(ctx context.Context, filePth string, id string, group string, hookfn func(context.Context, string, string, []byte)) {
	f, err := os.Open(filePth)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	f.Seek(0, 2)
	for {
		if ctx.Err() != nil {
			break
		}
		line, err := rd.ReadBytes('\n')
		// 如果是文件末尾不返回
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			log.Fatalln(err)
		}
		go hookfn(ctx, id, group, line)
	}
}

// GetFileSize 获取文件大小
func GetFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}

// GetCurrentPath 获取当前路径，比如：E:/abc/data/test
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// GetSize 获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// GetExt 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckExist 检查文件是否存在
func CheckExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// CheckPermission 检查文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir 检查文件夹是否存在
// 如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	if exist := !CheckExist(src); exist == false {
		if err := MakeDir(src); err != nil {
			return err
		}
	}

	return nil
}

// Open 打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// GetImgType 获取Img文件类型
func GetImgType(p string) (string, error) {
	file, err := os.Open(p)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	buff := make([]byte, 512)

	_, err = file.Read(buff)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	filetype := http.DetectContentType(buff)

	ext := imgext.Get()

	for i := 0; i < len(ext); i++ {
		if strings.Contains(ext[i], filetype[6:len(filetype)]) {
			return filetype, nil
		}
	}

	return "", errors.New("invalid image type")
}

// GetType 获取文件类型
func GetType(p string) (string, error) {
	file, err := os.Open(p)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	buff := make([]byte, 512)

	_, err = file.Read(buff)

	if err != nil {
		log.Println(err)
	}

	filetype := http.DetectContentType(buff)

	//ext := GetExt(p)
	//var list = strings.Split(filetype, "/")
	//filetype = list[0] + "/" + ext
	return filetype, nil
}

// ReadStr 文件读取
func ReadStr(filePath string) (string, error) {
	if contents, err := ioutil.ReadFile(filePath); err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		result := strings.Replace(string(contents), "\n", "", 1)
		//fmt.Println("Use ioutil.ReadFile to read a file:", result)
		return result, nil
	} else {
		return "", err
	}
}

// FormatFileSize 字节的单位转换 保留两位小数
func FormatFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		//return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}

func IsDir(name string) bool {
	if info, err := os.Stat(name); err == nil {
		return info.IsDir()
	}
	return false
}
func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}
func MakeDir(dir string) error {
	if !FileIsExisted(dir) {
		if err := os.MkdirAll(dir, 0777); err != nil { //os.ModePerm
			fmt.Println("MakeDir failed:", err)
			return err
		}
	}
	return nil
}

//使用io.Copy
func CopyFile(src, des string) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	//获取源文件的权限
	fi, _ := srcFile.Stat()
	perm := fi.Mode()

	//desFile, err := os.Create(des)  //无法复制源文件的所有权限
	desFile, err := os.OpenFile(des, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm) //复制源文件的所有权限
	if err != nil {
		return 0, err
	}
	defer desFile.Close()

	return io.Copy(desFile, srcFile)
}

//使用ioutil.WriteFile()和ioutil.ReadFile()
func CopyFile2(src, des string) (written int64, err error) {
	//获取源文件的权限
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	fi, _ := srcFile.Stat()
	perm := fi.Mode()
	srcFile.Close()

	input, err := ioutil.ReadFile(src)
	if err != nil {
		return 0, err
	}

	err = ioutil.WriteFile(des, input, perm)
	if err != nil {
		return 0, err
	}

	return int64(len(input)), nil
}

//使用os.Read()和os.Write()
func CopyFile3(src, des string, bufSize int) (written int64, err error) {
	if bufSize <= 0 {
		bufSize = 1 * 1024 * 1024 //1M
	}
	buf := make([]byte, bufSize)

	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	//获取源文件的权限
	fi, _ := srcFile.Stat()
	perm := fi.Mode()

	desFile, err := os.OpenFile(des, os.O_CREATE|os.O_RDWR|os.O_TRUNC, perm)
	if err != nil {
		return 0, err
	}
	defer desFile.Close()

	count := 0
	for {
		n, err := srcFile.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		if n == 0 {
			break
		}

		if wn, err := desFile.Write(buf[:n]); err != nil {
			return 0, err
		} else {
			count += wn
		}
	}

	return int64(count), nil
}

func CopyDir(srcPath, desPath string) error {
	//检查目录是否正确
	if srcInfo, err := os.Stat(srcPath); err != nil {
		return err
	} else {
		if !srcInfo.IsDir() {
			return errors.New("源路径不是一个正确的目录！")
		}
	}

	if desInfo, err := os.Stat(desPath); err != nil {
		return err
	} else {
		if !desInfo.IsDir() {
			return errors.New("目标路径不是一个正确的目录！")
		}
	}

	if strings.TrimSpace(srcPath) == strings.TrimSpace(desPath) {
		return errors.New("源路径与目标路径不能相同！")
	}

	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		//复制目录是将源目录中的子目录复制到目标路径中，不包含源目录本身
		if path == srcPath {
			return nil
		}

		//生成新路径
		destNewPath := strings.Replace(strings.ReplaceAll(path, "\\", "/"), srcPath, desPath, -1)

		if !f.IsDir() {
			file, err := CopyFile(path, destNewPath)
			if err != nil {
				fmt.Println(err.Error())
			}
			print(file)
		} else {
			if !FileIsExisted(destNewPath) {
				return MakeDir(destNewPath)
			}
		}

		return nil
	})

	return err
}

/* 获取指定路径下的所有文件，只搜索当前路径，不进入下一级目录，可匹配后缀过滤（suffix为空则不过滤）*/
func ListDir(dir, suffix string) (files []string, err error) {
	files = []string{}

	_dir, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	suffix = strings.ToLower(suffix) //匹配后缀

	for _, _file := range _dir {
		if _file.IsDir() {
			continue //忽略目录
		}
		if len(suffix) == 0 || strings.HasSuffix(strings.ToLower(_file.Name()), suffix) {
			//文件后缀匹配
			files = append(files, path.Join(dir, _file.Name()))
		}
	}

	return files, nil
}

// WalkDir /* 获取指定路径下以及所有子目录下的所有文件，可匹配后缀过滤（suffix为空则不过滤）*/
func WalkDir(dir, suffix string) (files []string, err error) {
	files = []string{}

	err = filepath.Walk(dir, func(fname string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			//忽略目录
			return nil
		}

		if len(suffix) == 0 || strings.HasSuffix(strings.ToLower(fi.Name()), suffix) {
			//文件后缀匹配
			files = append(files, fname)
		}

		return nil
	})

	return files, err
}

func Remove(fileName string) bool {
	return os.Remove(fileName) == nil
}

func RemoveAll(path string) bool {
	return os.RemoveAll(path) == nil
}
