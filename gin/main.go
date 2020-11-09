package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zeta-io/ginx"
	"github.com/zeta-io/zeta"
)

type Api struct {}

func (a *Api) list(context context.Context, c1 *context.Context, args struct{
	Name string `json:"name" param:"query,name"`
	Age *int `json:"age" param:"query,age"`
	NA  *string `json:"na" param:"query,na"`
	NB  *string `json:"nb" param:"query,nb"`
	B *bool `param:"query,b"`
	Version string `param:"path,version"`
	V1 *string `param:"path,version"`
	Size int  `param:"query,size" validate:"gte=0,lte=20"`
}) (string, error){
	fmt.Println(context)
	fmt.Println(c1)
	fmt.Println(args.Name)
	fmt.Println(args.Age == nil)
	fmt.Println(args.NA == nil)
	fmt.Println(args.NA)
	fmt.Println(args.NB == nil)
	fmt.Println(args.NB)
	fmt.Println("v1 is null ? ", args.V1 == nil)
	fmt.Println("version is ", args.Version)
	fmt.Println("size is", args.Size)
	return "hello nico", nil
}

var api = Api{}

func main() {
	router := zeta.Router("/api/:version/users")
	router.Get("", api.list)

	e := zeta.New(router, ginx.New(gin.New())).Run(":8080")
	if e != nil{
		panic(e)
	}
}