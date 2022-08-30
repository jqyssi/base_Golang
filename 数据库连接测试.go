package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)



func main() {

	db, err := sql.Open("mysql", "root:ranksin001@(127.0.0.1:3306)/ranksin")
	db.Ping()
	if err != nil {
		fmt.Println("连接失败")
		return
	}
	defer db.Close()

//
	stmt,err:=db.Prepare("insert into stu values (?,?,?)")
	if err!= nil{
			fmt.Println("预处理失败")
			fmt.Println(err)
			return
	}
	r,err:=stmt.Exec(18,"yuanke","长沙")
	if err!= nil{
			fmt.Println("执行失败")
			fmt.Println(err)
			return
	}
	count,err:=r.RowsAffected()
	if err!= nil{
		fmt.Println("结果获取失败")
		return
	}
	if count>0{
		fmt.Println("插入成功")
	}else{
		fmt.Println("插入失败")
	}
}