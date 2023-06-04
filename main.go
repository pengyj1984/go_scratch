/*
2023.6.4
试验mysql插件一些东西，包括:
multiStates, transaction
*/
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	dsn := "root:Abc123..@(localhost:3306)/chasinglight?multiStatements=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Open mysql error: ", err.Error())
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT `rid`, `account`, `gender`, `reg_time`, `reg_ip`, `reg_device` FROM `role` WHERE `rid` = ?; INSERT INTO `log` (`rid`, `time`, `op`, `log`) VALUES(?, ?, ?, ?)")
	if err != nil {
		fmt.Println("Prepare error: ", err.Error())
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(1, 1, time.Now().UnixMilli(), 0, "")
	var rid, gender int
	var reg_time int64
	var account, reg_ip, reg_device string
	err = row.Scan(&rid, &account, &gender, &reg_time, &reg_ip, &reg_device)
	if err != nil {
		fmt.Println("Query error: ", err.Error())
		return
	}

	fmt.Println("rid =", rid, ", account =", account, ", gender =", gender, ", reg_time =", reg_time, ", reg_ip =", reg_ip, "reg_device =", reg_device)
}
