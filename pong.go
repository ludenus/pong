package main

import "os"
import "github.com/kataras/iris"

func main() {

	addr := addr()

	app := iris.Default()

	app.Logger().SetLevel("warn")

	app.Get("/ping/{param:string}", func(ctx iris.Context) {

		param := ctx.Params().Get("param")

		ctx.JSON(iris.Map{
			"pong": param,
		})
	})

	app.Run(iris.Addr(addr))
}

func addr() string {

	addr := ":80"
	fromEnv := os.Getenv("PONG_LISTENING_ADDRESS")

	if fromEnv != "" {
		addr = fromEnv
	}

	if len(os.Args) > 1 {
		addr = os.Args[1]
	}

	return addr
}
