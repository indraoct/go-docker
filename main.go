package main

import (
	"database/sql"
	"dockerinaja/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
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
	var cfg config.AppConfig
	
	//CONFIG APPS (OS level)
	err := envconfig.Process("indra",&cfg)
	if err != nil{
		log.Fatal(err.Error())
	}
	
	db, err := sql.Open("mysql", cfg.DbUser+":"+cfg.DbPass+"@tcp("+cfg.DbHost+":"+cfg.DbPort+")/"+cfg.DbName)
	
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
  e.Logger.Fatal(e.Start(":"+cfg.AppsPort))
}