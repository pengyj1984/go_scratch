/*
测试是否使用 connections pool 的性能差异。以及观察高并发时的行为。
测试了四种形式，使用gorm和普通方式，每种方式测试使用连接池和不使用连接池。
内容是读取数据库中的数据，然后修改数据更新到数据库。
但是这次测试每次运行只测试一种情况，不像之前测试同时跑几个情况，所以同样操作1000条数据，速度明显比之前快得多。
结果大致是使用gorm和不适用gorm性能差不多，主要还是看使用习惯。
使用池的时间大概 0.4s 到 2s 之间。
不适用池的时间为 3s 到 4s。
性能差距还是挺明显的。

另外，在同一时间并发连接很多的时候，会出现一些警告，比如连接被拒，还有sql语句执行较慢等。但是所有操作都还是成功了。
*/
package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type Role struct {
	Id            int    `gorm:"column:id;PRIMARY_KEY"`
	ActivateDays  int    `gorm:"column:activate_days;not null"`
	LastLoginTime int64  `gorm:"column:last_login_time;not null"`
	Data          string `gorm:"column:data;not null"`
}

func (Role) TableName() string {
	return "role"
}

var ch chan int
var dsn string

var _db *gorm.DB

func getDB() *gorm.DB {
	return _db
}

var _db2 *sql.DB

func getDB2() *sql.DB {
	return _db2
}

func main() {
	num := 0
	ch = make(chan int)
	defer close(ch)
	dsn = "root:Abc123..@(192.168.1.88:3306)/test"

	err := initDB()
	if err != nil {
		log.Fatalln("init db err:", err.Error())
	}
	defer closeDB()

	start := time.Now().UnixMilli()
	for i := 1; i <= 1000; i++ {
		go normalWithPool(i)
	}

	var success, failed int = 0, 0
	for {
		select {
		case r := <-ch:
			if r == 0 {
				success++
			} else {
				failed++
			}
			num++
		}

		if num >= 1000 {
			break
		}
	}

	diff := time.Now().UnixMilli() - start
	log.Println("success:", success)
	log.Println("failed:", failed)
	log.Println("finish, total cost:", diff)
}

func initDB() error {
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	_sqlDB, _ := _db.DB()
	_sqlDB.SetMaxOpenConns(100)

	_db2, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	_db2.SetMaxOpenConns(100)

	return nil
}

func closeDB() {
	_sqlDB, _ := _db.DB()
	_sqlDB.Close()

	_db2.Close()
}

func gormWithoutPool(id int) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		ch <- 1
		return
	}
	_sqlDB, _ := db.DB()
	defer _sqlDB.Close()

	var role Role
	tx := db.Session(&gorm.Session{
		PrepareStmt: true,
	})
	role.Id = id
	tx.Select("id, activate_days, last_login_time, data").Where("id = ?", id).Find(&role)
	now := time.Now().Unix()
	role.LastLoginTime = now
	role.ActivateDays++
	err = tx.Save(&role).Error
	if err != nil {
		ch <- 1
		return
	}

	ch <- 0
}

func gormWithPool(id int) {
	db := getDB()

	var role Role
	tx := db.Session(&gorm.Session{
		PrepareStmt: true,
	})
	role.Id = id
	tx.Select("id, activate_days, last_login_time, data").Where("id = ?", id).Find(&role)
	now := time.Now().Unix()
	role.LastLoginTime = now
	role.ActivateDays++
	err := tx.Save(&role).Error
	if err != nil {
		ch <- 1
		return
	}

	ch <- 0
}

func normalWithoutPool(id int) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		ch <- 1
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT `activate_days`, `last_login_time`, `data` FROM `role` WHERE `id` = ?")
	if err != nil {
		ch <- 1
		return
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	var activateDays sql.NullInt32
	var lastLoginTime sql.NullInt64
	var data sql.NullString
	err = row.Scan(&activateDays, &lastLoginTime, &data)
	if err != nil {
		ch <- 1
		return
	}

	now := time.Now().Unix()
	_, err = db.Exec("UPDATE `role` SET `activate_days` = ?, `last_login_time` = ? WHERE `id` = ?", activateDays.Int32+1, now, id)
	if err != nil {
		ch <- 1
		return
	}

	ch <- 0
}

func normalWithPool(id int) {
	db := getDB2()
	stmt, err := db.Prepare("SELECT `activate_days`, `last_login_time`, `data` FROM `role` WHERE `id` = ?")
	if err != nil {
		ch <- 1
		return
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	var activateDays sql.NullInt32
	var lastLoginTime sql.NullInt64
	var data sql.NullString
	err = row.Scan(&activateDays, &lastLoginTime, &data)
	if err != nil {
		ch <- 1
		return
	}

	now := time.Now().Unix()
	_, err = db.Exec("UPDATE `role` SET `activate_days` = ?, `last_login_time` = ? WHERE `id` = ?", activateDays.Int32+1, now, id)
	if err != nil {
		ch <- 1
		return
	}

	ch <- 0
}
