package save

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type Cate struct {
	gorm.Model
	Name string
	Pid  int
}

type Text struct {
	gorm.Model
	Knowledge string
	Cid       int
}

func (Cate) TableName() string {
	return "category"
}

func (Text) TableName() string {
	return "knowledge"
}

var db *gorm.DB = DbConn()

func DbConn() (db *gorm.DB) {
	username := "root"	//mysql用户名
	pwd := os.Getenv("MYSQL_PWD")	//mysql密码
	host := os.Getenv("MYSQL_HOST")	//mysql host
	database := "faq"	//mysql 数据库
	port := "3306"	//mysql 端口
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, pwd, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return
}

func SaveCate(content string, pid int) (id int) {
	cate := Cate{Name: content, Pid: pid}
	db.Create(&cate)
	//fmt.Println(result)
	return int(cate.ID)
}

func SaveText(content string, cid int) {
	text := Text{Knowledge: content, Cid: cid}
	db.Create(&text)
}

func Main() {
	SaveCate("根目录", -1)
}
