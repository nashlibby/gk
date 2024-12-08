package gk

type Oss interface {
	// 上传文件
	UploadFile(remotePath string, file []byte) (string, error)
	// 上传远程文件
	TransformFile(remotePath string, url string) (string, error)
	// 获取object文件
	GetObjectFile(objectKey string) ([]byte, error)
}
