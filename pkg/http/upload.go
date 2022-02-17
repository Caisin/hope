package http

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/pborman/uuid"
	file2 "hope/pkg/file"
	"io"
	"mime/multipart"
	"os"
)

func RegisterUploadFileHTTPServer(s *http.Server) {
	r := s.Route("/")
	r.POST("/upload", UploadFile)
	r.POST("/uploads", Uploads)
}

// UploadFile 上传文件
func UploadFile(ctx http.Context) error {
	req := ctx.Request()
	fileName := req.FormValue("name")
	_, handler, err := req.FormFile("file")
	if err != nil {
		return err
	}
	ext := file2.GetExt(handler.Filename)
	err = SaveUploadedFile(handler, uuid.New()+"."+ext)
	if err != nil {
		return err
	}
	return ctx.String(200, "File "+fileName+" Uploaded successfully")
}

// Uploads 多文件上传
func Uploads(ctx http.Context) error {
	req := ctx.Request()
	form := req.MultipartForm
	files := form.File["upload[]"]
	//遍历数组进行处理
	for _, file := range files {
		uid := uuid.New()
		_ = SaveUploadedFile(file, uid)
	}
	return ctx.String(200, "successfully")
}

func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
