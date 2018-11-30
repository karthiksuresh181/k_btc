package main

import (
	"fmt"
	_ "sample/models"
	_ "sample/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "smartparking1:karthik!@tcp(den1.mysql6.gear.host:3306)/smartparking1?charset=utf8")
	orm.RunCommand()
}

func main() {
	o := orm.NewOrm()
	name := "default"
	force := false
	verbose := true
	err := orm.RunSyncdb(name, force, verbose) //this is to create/drops the tables
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error in SyncDB")
	}
	o.Using("default")
	beego.Run()
}
