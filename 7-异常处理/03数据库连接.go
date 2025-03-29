// @FileName : 03数据库连接.go
// @Time : 2025/3/19 下午3:21
// @Author : luobozi

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 此处开始编写代码
func main() {
	//创建数据库连接
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.100.101:3306)/test1?charset=utf8")
	//书写格式  用户名:密码@tcp(主机地址:端口)/数据库名ll
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return
	}
	defer db.Close()
	sql_str := "select id,name from account where name =?"
	rows, e := db.Query(sql_str, "刘彪")
	if e != nil {
		fmt.Println("查询数据库错误！", e)
	}
	fmt.Println(rows)
	var id, username string
	for rows.Next() { //处理多条数据
		rows.Scan(&id, &username)
		fmt.Println("从数据库拿到的值为：", id, username)
	}
}
