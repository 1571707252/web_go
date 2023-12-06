package main

import (
	"fmt"
	"github.com/zhixu/zxgo"
)

func main() {
	//http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(writer, "%s 欢迎来到web框架教程", "mszlu.com")
	//})
	//err := http.ListenAndServe(":8111", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}

	engine := zxgo.New()
	g := engine.Group("user")
	g.Get("/hello", func(ctx *zxgo.Context) {
		fmt.Fprintf(ctx.W, "%s get欢迎来到web框架教程", "mszlu.com")
	})
	g.Post("/hello", func(ctx *zxgo.Context) {
		fmt.Fprintf(ctx.W, "%s post欢迎来到web框架教程", "mszlu.com")
	})
	g.Post("/info", func(ctx *zxgo.Context) {
		fmt.Fprintf(ctx.W, "%s info", "mszlu.com")
	})
	g.Any("/any", func(ctx *zxgo.Context) {
		fmt.Fprintf(ctx.W, "%s any", "mszlu.com")
	})
	//o := engine.Group("order")
	//o.Add("/get", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "%s 查询订单", "mszlu.com")
	//})
	engine.Run()
}
