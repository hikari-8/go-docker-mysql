package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hikari-8/go-docker-mysql/src/article"
)

func open(path string, count uint) (*sql.DB, error) {
	db, err := sql.Open("mysql", path)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		fmt.Printf("retry... count:%v\n", count)
		if (count > 0) {
			return open(path, count)
		} else {
			return nil, err
		}
	}
	return db, nil
}

func connectDB() (*sql.DB, error) {
	var path string = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true",
		os.Getenv("MYSQL_USER"), 
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"))

	return open(path, 10)
}

func main() {
	db, err := connectDB()
	if err != nil {
		fmt.Println(err)
	} else {
		return
	} 
	defer db.Close()
	
	fmt.Println("mysql Connected!")
	article.ReadAll(db)
}