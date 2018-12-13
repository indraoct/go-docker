package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

type Users struct{
	Id uint64 `json:"id"`
	Username string `json:"username"`
	IsActive uint64 `json:"is_active"`
}

func main() {
	e := echo.New()
	var users Users
	var arr_users []Users
	
	db_user   := os.Getenv("DB_USER")
	db_pass   := os.Getenv("DB_PASS")
	db_host   := os.Getenv("DB_HOST")
	db_port   := os.Getenv("DB_PORT")
	apps_port := os.Getenv("PORT")
	
	db, err := sql.Open("mysql", db_user+":"+db_pass+"@tcp("+db_host+":"+db_port+")/dockeraja")
	
	defer db.Close()
	
	if err != nil{
		fmt.Println(err.Error())
	}
	
	e.GET("/", func(c echo.Context) error {
		
			return c.String(http.StatusOK, "Hello, World!")
	})
	
	e.GET("/user/getall",func(c echo.Context) error{
		arr_users = []Users{}
		res, err := db.Query("SELECT id,username,is_active FROM users order by id DESC")
		
		if err != nil{
			return c.String(http.StatusInternalServerError,err.Error())
		}
		
		for res.Next() {
			
				err = res.Scan(&users.Id,&users.Username,&users.IsActive)
			
			if err != nil{
				return c.String(http.StatusInternalServerError,"ERROR SCAN COOOK "+err.Error())
			}
			
			arr_users = append(arr_users,users)
		}
		
		return c.JSON(http.StatusOK,arr_users)
	})
  e.Logger.Fatal(e.Start(":"+apps_port))
}