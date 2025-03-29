// @FileName : 数据库连接最佳实践.go
// @Time : 2025/3/22 下午9:39
// @Author : luobozi

// 此处开始编写代码
package main

import (
	"database/sql"
	"fmt"
	// 引入 MySQL 驱动，注意这里使用了匿名导入
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 1. 定义数据库连接信息（DSN：Data Source Name）
	// 格式：username:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:123456@tcp(192.168.100.101:3306)/test1?charset=utf8mb4&parseTime=True&loc=Local"

	// 2. 连接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	// 程序退出前关闭数据库连接
	defer db.Close()

	// 3. 测试数据库连接是否成功
	if err = db.Ping(); err != nil {
		fmt.Println("数据库 Ping 失败:", err)
		return
	}
	fmt.Println("成功连接到数据库")

	// 4. 增：插入数据（使用参数化查询，防止SQL注入）
	insertSQL := "INSERT INTO account(id,name, money) VALUES(?,?, ?)"
	result, err := db.Exec(insertSQL, 0, "张三", 30)
	if err != nil {
		fmt.Println("插入数据失败:", err)
		return
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("获取插入数据ID失败:", err)
		return
	}
	fmt.Println("插入数据成功，ID:", lastInsertId)

	// 5. 查：查询数据
	querySQL := "SELECT id, name, money FROM account WHERE money > ?"
	rows, err := db.Query(querySQL, 20)
	if err != nil {
		fmt.Println("查询数据失败:", err)
		return
	}
	// 使用 defer 关闭 rows，避免资源泄露
	defer rows.Close()

	// 遍历查询结果
	for rows.Next() {
		var id int
		var name string
		var money int
		if err = rows.Scan(&id, &name, &money); err != nil {
			fmt.Println("读取数据失败:", err)
			return
		}
		fmt.Printf("查询结果: id=%d, name=%s, money=%d\n", id, name, money)
	}
	// 检查遍历过程中是否出错
	if err = rows.Err(); err != nil {
		fmt.Println("遍历结果时出错:", err)
		return
	}

	// 6. 改：更新数据
	updateSQL := "UPDATE account SET money = ? WHERE id = ?"
	result, err = db.Exec(updateSQL, 35, lastInsertId)
	if err != nil {
		fmt.Println("更新数据失败:", err)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("获取更新行数失败:", err)
		return
	}
	fmt.Println("更新数据成功，影响行数:", rowsAffected)

	// 7. 删：删除数据
	deleteSQL := "DELETE FROM account WHERE id = ?"
	result, err = db.Exec(deleteSQL, lastInsertId)
	if err != nil {
		fmt.Println("删除数据失败:", err)
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		fmt.Println("获取删除行数失败:", err)
		return
	}
	fmt.Println("删除数据成功，影响行数:", rowsAffected)
}
