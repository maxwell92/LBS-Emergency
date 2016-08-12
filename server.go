package main

import (
	//"fmt"
	color "github.com/iris-contrib/color"
	iris "github.com/kataras/iris"
	//ia "irisapi"
	hc "httpclient"
	"log"
)

const (
	banner = `
               *
              | |
              | |
           - -   - -
          | | | | | |
          | | | | | |
          -----------  
(o゜▽゜)o  is listening on :10000
`

	logo = `
    ||        --------
    ||        ||-----
    ||        ||-----
    ||        ||-----
    ||------  ||-----
    --------  --------
	`
)

// why could this run even I didn't called it ?!
func init() {
	iris.StaticServe("./template", "/static")
	//iris.StaticWeb("./template", "/static", 1)
	color.Cyan(banner)
}

func Test(ctx *iris.Context) {
	client := hc.NewHttpClient("", "")
	response, err := client.Get("http://api.map.baidu.com/geocoder/v2/?ak=5RijeBzVjQ82uPx8gxGGfeNXlfRt7yH6&callback=renderReverse&location=39.983424,116.322987&output=json&pois=1")
	if err != nil {
		log.Println(err)
	}
	ctx.Write(string(response))
}

func Hello(ctx *iris.Context) {
	ctx.Write(logo)
	ctx.Write("This is LBS-Emergency")
	ctx.Write("You can use this to report a Emergency very fast")
}

func SOS(ctx *iris.Context) {
	ctx.Write("SOS send!")
	ctx.Write("They are on the way!!")
}

func main() {
	iris.Get("/test", Test)
	iris.Get("/hello", Hello)
	iris.Get("/sos", SOS)
	iris.Listen(":10000")
}
