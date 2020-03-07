package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

//关键词const加上括号，有两个作用：
//1.变量值自增，类似于C++的枚举（这里用不到）
//2.全部设置为常量类型，比较方便
const (
	userName = "test"
	password = "asdfg12345"
	ip       = "cdb-axt937vt.gz.tencentcdb.com"
	port     = "10059"
	dbName   = "file"
)

var db *sql.DB

func init() {
	connectInfo := []string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}
	path := strings.Join(connectInfo, "")
	db, _ = sql.Open("mysql", path)

	if err := db.Ping(); err != nil {
		fmt.Printf("Open mysql error!")
		return
	}
	fmt.Println("Database Connect Success!")
}

//DBConn: 返回数据库连接对象
func DBConn() *sql.DB {
	return db
}
