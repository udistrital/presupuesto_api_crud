package main

import (
	"fmt"
	"reflect"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
	_ "github.com/udistrital/presupuesto_crud/routers"

	"github.com/astaxie/beego"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://"+beego.AppConfig.String("PGuser")+":"+beego.AppConfig.String("PGpass")+"@"+beego.AppConfig.String("PGurls")+"/"+beego.AppConfig.String("PGdb")+"?sslmode=disable&search_path="+beego.AppConfig.String("PGschemas")+"")
	if beego.BConfig.RunMode == "dev" {
		// Database alias.
		name := "default"

		// Drop table and re-create.
		force := false

		// Print log.
		verbose := true

		// Error.
		err := orm.RunSyncdb(name, force, verbose)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func main() {

	beego.BConfig.RecoverFunc = func(ctx *context.Context) {
		type response struct {
			Code string
			Type string
			Body interface{}
		}
		out := response{}
		Body := ctx.Input.Data()["json"]
		defer func() {
			ctx.Output.JSON(out, true, false)

		}()
		if r := recover(); r != nil {
			beego.Error(r)
			ctx.ResponseWriter.WriteHeader(500)
			out.Body = r
			out.Code = ""
			out.Type = "error"
		} else {
			if reflect.ValueOf(Body).IsValid() {
				defer func() {
					if r := recover(); r != nil {
						beego.Error(r)
						out.Body = Body
						out.Type = "success"
						ctx.ResponseWriter.WriteHeader(200)
					}
				}()
				if reflect.ValueOf(Body).IsNil() {
					out.Body = nil
					out.Type = "No Data Found"
					ctx.ResponseWriter.WriteHeader(201)
				}

			} else {
				out.Body = Body
				out.Type = "success"
				ctx.ResponseWriter.WriteHeader(200)
			}
		}
	} //responseformat.GlobalResponseHandler

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		orm.Debug = true
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
