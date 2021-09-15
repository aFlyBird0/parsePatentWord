package save

import (
	"fmt"
	"os"
	"testing"
)

func TestDbConn(t *testing.T) {
	Main()
}

func TestDbCombine(t *testing.T)  {
	username := "root"	//mysql用户名
	pwd := os.Getenv("MYSQL_PWD")	//mysql密码
	host := os.Getenv("MYSQL_HOST")	//mysql host
	database := "faq"	//mysql 数据库
	port := "3306"	//mysql 端口
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, pwd, host, port, database)
	fmt.Println(dsn)
	//if dsn == "这里是拼接成功后的数据库连接语句" {
	//	fmt.Println("拼接成功")
	//}else {
	//	fmt.Println("拼接失败")
	//}
}
