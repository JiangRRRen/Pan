package db

import (
	mydb "awesomeProject/db/mysql"
	"database/sql"
	"fmt"
)

type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

// OnFileUploadFinished: 文件上传完成
func OnFileUploadFinished(filehash string, filename string, filesize int64,
	filepath string) bool {
	//所谓prepared，即带有占位符的sql语句，这个语句可以根据不同的参数多次调用
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_file(`file_md5`,`file_name`,`file_size`," +
			"`file_addr`,`status`) values(?,?,?,?,1)")
	if err != nil {
		fmt.Printf("Failed to prepare statement, err: %s", err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(filehash, filename, filesize, filepath)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if rf, err := ret.RowsAffected(); err != nil {
		if rf <= 0 { //sql语句执行成功，但没有产生新的记录
			// 警告一下
			fmt.Printf("File with hash:%s has been uploaded before", filehash)
			return true
		}
		return false
	}
	return true
}

// GetFileMeta: 从数据库获取文件源信息
func GetFileMeta(filehash string) (*TableFile, error) {
	stmt, err := mydb.DBConn().Prepare("SELECT file_md5, file_addr,file_name,file_size FROM tbl_file WHERE file_md5=? AND status=1 limit 1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer stmt.Close()

	tfile := TableFile{}
	err = stmt.QueryRow(filehash).Scan(&tfile.FileHash, &tfile.FileAddr, &tfile.FileName, &tfile.FileSize)
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	return &tfile, nil
}
