package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "database_name"
)

//连接数据库
func main()  {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

//查询
//func connectDB() *sql.DB{
//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//		"password=%s dbname=%s sslmode=disable",
//		host, port, user, password, dbname)
//
//	db, err := sql.Open("postgres", psqlInfo)
//	if err != nil {
//		panic(err)
//	}
//
//	err = db.Ping()
//	if err != nil {
//		panic(err)
//	}
//	return db
//}
//
//
//func query(db *sql.DB){
//	var id,name,price string
//
//
//	rows,err:=db.Query(" select * from table_name where name=$1","Proto")
//
//	if err!= nil{
//		fmt.Println(err)
//	}
//	defer rows.Close()
//
//	for rows.Next(){
//		err:= rows.Scan(&id,&name,&price)
//
//		if err!= nil{
//			fmt.Println(err)
//		}
//	}
//
//	err = rows.Err()
//	if err!= nil{
//		fmt.Println(err)
//	}
//
//	fmt.Println(id,name,price)
//}
//
//func main()  {
//	db:=connectDB()
//	query(db)
//
//}



