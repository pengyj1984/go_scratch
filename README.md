# go_scratch
main.go.bak  go 和 channel
main.go.bak1 go 和 channel
main.go.bak2 golang interface模拟多态
main.go.bak3 golang go-sql-driver/mysql 使用 prepare 和 不适用 prepare 性能测试
main.go.bak4 访问mysql，测试gorm效率。并且也测试了普通方式使用stmt和不适用的区别。
main.go.bak5 测试是否使用 connections pool 的性能差异。以及观察高并发时的行为。

# 安装mysql插件
go get github.com/go-sql-driver/mysql

# 安装gorm插件
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql