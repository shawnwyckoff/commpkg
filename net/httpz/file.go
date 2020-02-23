package httpz

import "github.com/mushroomsir/httpfile" // 大文件的上传下载，做得很好

func UploadFile(requestUrl, localFilename, remoteFilename string) {
	f := httpfile.NewReq(requestUrl, localFilename).SetHeader("filename", remoteFilename)
	f.UploadByStream()
}
