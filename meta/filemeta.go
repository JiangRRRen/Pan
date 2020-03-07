package meta

import (
	mydb "awesomeProject/db"
)

//FileMeta: 定义文件结构属性
type FileMeta struct {
	FileMD5   string //MD5作为文件ID
	FileName  string //文件名
	FileSize  int64  //大小
	FilePath  string //本地存储路径
	TimeStamp string //时间戳
}

//目前元信息都存储在内存中，程序结束就会丢失，一般都会保存在数据库中
var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta) //初始化
}

//UpdateFileMeta: 更新文件信息
func UpdateFileMeta(f FileMeta) {
	fileMetas[f.FileMD5] = f
}

//UpdateFileMetaDB：更新文件到数据库
func UpdateFileMetaDB(f FileMeta) {
	mydb.OnFileUploadFinished(f.FileMD5, f.FileName, f.FileSize, f.FilePath)
}

//GetdateFileMeta: 根据MD5值获取文件信息
func GetdateFileMeta(md5 string) FileMeta {
	return fileMetas[md5]
}

//GetdateFileMetaDB: 从数据库获取信息
func GetdateFileMetaDB(md5 string) (FileMeta, error) {
	tfile, err := mydb.GetFileMeta(md5)
	if err != nil {
		return FileMeta{}, err
	}
	fmeta := FileMeta{
		FileMD5:  tfile.FileHash,
		FileName: tfile.FileName.String,
		FileSize: tfile.FileSize.Int64,
		FilePath: tfile.FileAddr.String,
	}
	return fmeta, nil
}

//DeleteFileMeta: 简单的删除，线程不安全
func DeleteFileMeta(md5 string) {
	delete(fileMetas, md5)
}
