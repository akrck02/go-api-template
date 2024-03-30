package http

import (
	"bytes"
	"io"
	"mime/multipart"

	"github.com/akrck02/valhalla-core-sdk/models"
	"github.com/gin-gonic/gin"
)

// Read body json as the given object
//
// [param] c | interface: gin context
// [param] obj | interface{}: object to read
//
// [return] error: error
func ReadBodyJson(params interface{}, obj interface{}) error {

	return nil
}

// MultipartToBytes reads a file from a multipart request
//
// [param] c | *gin.Context: gin context
// [param] key | string: key of the file
//
// [return] []byte: file bytes
func MultipartToBytes(c *gin.Context, key string) ([]byte, error) {

	fileheader, err := c.FormFile(key)
	if err != nil {
		return nil, err
	}

	// if empty file
	if fileheader == nil {
		return nil, nil
	}

	// open file
	var file multipart.File
	file, err = fileheader.Open()

	if err != nil {
		return nil, err
	}

	// read file bytes
	defer file.Close()
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// GetRequestMetadata returns the request metadata
//
// [param] c | *gin.Context: gin context
//
// [return] *models.Request: request metadata
func GetRequestMetadata(c *gin.Context) *models.Request {
	var request, exists = c.Get("request")

	if !exists {
		return nil
	}

	var casted models.Request = request.(models.Request)
	return &casted
}

const (
	HTTP_METHOD_GET     = 0
	HTTP_METHOD_POST    = 1
	HTTP_METHOD_PUT     = 2
	HTTP_METHOD_DELETE  = 3
	HTTP_METHOD_PATCH   = 4
	HTTP_METHOD_HEAD    = 5
	HTTP_METHOD_OPTIONS = 6
	HTTP_METHOD_TRACE   = 7
	HTTP_METHOD_CONNECT = 8
)
