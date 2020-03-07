package handler

import (
	"awesomeProject/meta"
	"awesomeProject/util"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//返回上传页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internal error server")
			return
		} else {
			io.WriteString(w, string(data))
		}
	} else if r.Method == "POST" {
		file, header, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data, err:%s\n", err.Error())
			return
		}
		fileMeta := meta.FileMeta{
			FileName:  header.Filename,
			FilePath:  "./storage/" + header.Filename,
			TimeStamp: time.Now().Format("2006-01-02 15:04:05"),
		}
		//本地创建文件句柄
		//注意格式，必须要保证存在storage文件夹，必须要加.号
		newFile, err := os.Create(fileMeta.FilePath)
		if err != nil {
			fmt.Printf("Failed to create file, err:%s", err.Error())
			return
		}
		defer file.Close()
		defer newFile.Close()

		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save data, err:%s\n", err.Error())
		}

		newFile.Seek(0, 0)
		fileMeta.FileMD5 = util.FileMD5(newFile)
		//meta.UpdateFileMeta(fileMeta)
		meta.UpdateFileMetaDB(fileMeta)
		//重定向到成功提示页面
		//重定向的时候不要写成file/upload/suc，不然会被定向到file/file/upload
		http.Redirect(w, r, "upload/suc", http.StatusFound)
	}

}

//提示上传已完成
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload Successful!")
}

//GetFileMetaHandler: 获取文件源信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filehash := r.Form["filehash"][0]
	//fMeta:=meta.GetdateFileMeta(filehash)
	fMeta, err := meta.GetdateFileMetaDB(filehash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(fMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

//DownloadHandler: 根据哈希下载文件
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmd5 := r.Form.Get("filehash")
	fm := meta.GetdateFileMeta(fmd5)

	//获取资源
	f, err := os.Open(fm.FilePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	//读取资源
	data, err := ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("content-disposition", "attachment;filename=\""+fm.FileName+"\"")
	w.Write(data)
}

//FileMetaUpdateHandler: 更新文件
func FileMetaUpdateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	opType := r.Form.Get("op")
	fmd5 := r.Form.Get("filehash")
	newFileName := r.Form.Get("filename")

	//有些不支持
	if opType != "0" {
		w.WriteHeader(http.StatusForbidden)
	}

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	curFileMeta := meta.GetdateFileMeta(fmd5)
	curFileMeta.FileName = newFileName
	meta.UpdateFileMeta(curFileMeta)
	//meta.UpdateFileMetaDB(curFileMeta)
	data, err := json.Marshal(curFileMeta)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

//FileDeleteHanlder: 删除文件
func FileDeleteHanlder(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmd5 := r.Form.Get("filehash")
	//这里只是祛除了map的记录
	//物理删除
	fMeta := meta.GetdateFileMeta(fmd5)
	err := os.Remove(fMeta.FilePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	meta.DeleteFileMeta(fmd5)

	w.WriteHeader(http.StatusOK)

}
