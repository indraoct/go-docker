package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"net/http"
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
	
	db, err := sql.Open("mysql", "root:root@tcp(192.168.99.100:3306)/dockeraja")
	
	defer db.Close()
	
	if err != nil{
		fmt.Println(err.Error())
	}
	
	e.GET("/", func(c echo.Context) error {
		
			return c.String(http.StatusOK, "Hello, World!")
	})
	
	e.GET("/user/getall",func(c echo.Context) error{
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
  e.Logger.Fatal(e.Start(":1323"))
}