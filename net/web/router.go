package web

import "github.com/gin-gonic/gin"

const (
	GET Method = "GET"
	POST Method = "POST"
	PUT Method = "PUT"
	DELETE Method = "DELETE"
)

type (
	Router struct{
		ng *gin.Engine
	}

 	HandlerFunc func(*Ctx)

	Method string
)

func NewRouter() *Router {
	return &Router{ng:gin.Default()}
}

func (r *Router) Handle(m Method, relativePath string, fn HandlerFunc) {
	fn2 := func(c *gin.Context) {
		fn(&Ctx{ctx:c})
	}
	r.ng.Handle(string(m), relativePath, fn2)
}

func (r *Router) Serve(addr string) error {
	return r.ng.Run(addr) // listen and serve on 0.0.0.0:8080
}