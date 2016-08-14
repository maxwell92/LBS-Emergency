package main

import (
	"fmt"
	color "github.com/iris-contrib/color"
	"github.com/iris-contrib/template/html"
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

type page struct {
	Title string
}

// why could this run even I didn't called it ?!
// if I named the function init()
func Init() {
	//iris.StaticServe("./template", "/static")
	//iris.StaticWeb("/", "./template", 0)

	iris.UseTemplate(html.New()).Directory("./template", ".html")
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

func About(ctx *iris.Context) {
	ctx.Write(logo)
	ctx.Write("This is LBS-Emergency")
	ctx.Write("You can use this to report a Emergency very fast")
}

func SOS(ctx *iris.Context) {
	ctx.Write("SOS send!")
	ctx.Write("They are on the way!!")
}

func Zoom(ctx *iris.Context) {
	ctx.MustRender("zoom.html", page{Title: "LBS-Emergency-Zoom"})
	fmt.Println("zoom")
}

func main() {
	Init()
	iris.Get("/", func(ctx *iris.Context) {
		ctx.MustRender("index.html", page{Title: "LBS-Emergency"})
	})

	iris.Get("/test", Test)
	iris.Get("/about", About)
	iris.Get("/sos", SOS)
	iris.Get("/zoom", Zoom)
	iris.Listen(":10000")
}
