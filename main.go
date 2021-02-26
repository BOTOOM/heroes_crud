package main

import (
	_ "github.com/BOTOOM/heroes_crud/routers"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/lib/pq"
)

func main() {
	schema, _ := beego.AppConfig.String("schema")
	db, _ := beego.AppConfig.String("db")
	user, _ := beego.AppConfig.String("user")
	pass, _ := beego.AppConfig.String("pass")
	host, _ := beego.AppConfig.String("host")
	port, _ := beego.AppConfig.String("port")
	// sqlconn, _ := beego.AppConfig.String("sqlconn")
	orm.RegisterDataBase("default", "postgres", "postgres://"+user+":"+pass+"@"+host+":"+port+"/"+db+"?sslmode=disable&search_path="+schema)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	beego.Run()
}
