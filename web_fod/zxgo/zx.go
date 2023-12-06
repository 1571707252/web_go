package zxgo

import (
	"fmt"
	"log"
	"net/http"
)

const ANY = "ANY"

type HandlerFunc func(ctx *Context)

type routerGroup struct {
	name             string
	handleFuncMap    map[string]map[string]HandlerFunc
	handlerMethodMap map[string][]string
}

//	func (r *routerGroup) Add(name string, handlerFunc HandlerFunc) {
//		r.handleFuncMap[name] = handlerFunc
//	}
func (r *routerGroup) handle(name string, method string, handlerFunc HandlerFunc) {
	_, ok := r.handleFuncMap[name]
	if !ok {
		r.handleFuncMap[name] = make(map[string]HandlerFunc)
	}
	_, ok = r.handleFuncMap[name][method]
	if ok {
		panic("有重复的路由")
	}
	r.handleFuncMap[name][method] = handlerFunc
}
func (r *routerGroup) Any(name string, handlerFunc HandlerFunc) {
	r.handle(name, ANY, handlerFunc)
}

func (r *routerGroup) Get(name string, handlerFunc HandlerFunc) {
	r.handle(name, http.MethodGet, handlerFunc)
}

// 同一目录下支持会覆盖
func (r *routerGroup) Post(name string, handlerFunc HandlerFunc) {
	r.handle(name, http.MethodPost, handlerFunc)
}

type router struct {
	routerGroups []*routerGroup
}

func (r *router) Group(name string) *routerGroup {
	routerGroup := &routerGroup{
		name:             name,
		handleFuncMap:    make(map[string]map[string]HandlerFunc),
		handlerMethodMap: make(map[string][]string),
	}
	r.routerGroups = append(r.routerGroups, routerGroup)
	return routerGroup
}

type Engine struct {
	router
}

func New() *Engine {
	return &Engine{
		router{},
	}
}
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	for _, group := range e.routerGroups {
		for name, methodHandle := range group.handleFuncMap {
			url := "/" + group.name + name
			if r.RequestURI == url {
				ctx := &Context{
					W: w,
					R: r,
				}
				handle, ok := methodHandle[ANY]
				if ok {
					handle(ctx)
					return
				}
				handle, ok = methodHandle[method]
				if ok {
					handle(ctx)
					return
				}
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, "%s %s not allowed \n", r.RequestURI, method)
				return
			}
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "%s not found \n", r.RequestURI)

}
func (e *Engine) Run() {
	//for _, group := range e.routerGroups {
	//	for key, value := range group.handlerMap {
	//		http.HandleFunc("/"+group.name+key, value)
	//	}
	//}
	http.Handle("/", e)
	err := http.ListenAndServe(":8111", nil)
	if err != nil {
		log.Fatal(err)
	}
}
