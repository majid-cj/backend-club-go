package main

import (
	"github.com/kataras/iris"
)

//Auth ...
type Auth struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

var auths []Auth

func main() {
	app := iris.Default()

	app.Use(AuthenticateToken)

	app.Post("/", func(ctx iris.Context){
		var auth Auth
		ctx.ReadJSON(&auth)
		//do some validation things 
		auths = append(auths, auth)
	})


	app.Get("/test", func(ctx iris.Context){
		ctx.Write([]byte(ctx.RequestPath(true)))
	})


	app.Run(iris.Addr(":8080"))
}


//AuthenticateToken ...
func AuthenticateToken(ctx iris.Context) {
	if ctx.RequestPath(true) != "/"{
		if ctx.GetHeader("custom") == "im-hacker"{
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.Write([]byte("go away hacker, this is robust middleware"))
		}
	}else{ctx.Next()}
}

