package main

import (
	"bytes"
	"log"
	"net/http"

	"git.andrewcsellers.com/acsellers/card_sharp/config"
	"github.com/acsellers/multitemplate"
	"github.com/acsellers/platform/controllers"
	"github.com/acsellers/platform/router"
)

func main() {
	log.Fatal(http.ListenAndServe(config.WebPort(), Router()))
}

func Router() http.Handler {
	r := router.NewRouter()
	r.Many(PageCtrl{NewRenderableCtrl("desktop.html")})
	r.Many(FrontGameCtrl{NewRenderableCtrl("desktop.html")})
	r.Mount(controllers.AssetModule{
		AssetLocation: "public",
	})

	return r
}

type RenderableCtrl struct {
	*router.BaseController
	Template, Layout string
}

func NewRenderableCtrl(layout string) RenderableCtrl {
	return RenderableCtrl{
		&router.BaseController{},
		"", layout,
	}
}

func (rc RenderableCtrl) Render() router.Result {
	if *config.Dev {
		config.CompileTemplates()
	}

	ctx := &multitemplate.Context{
		Main:   rc.Template,
		Layout: rc.Layout,
		Dot:    rc.Context,
	}
	buf := &bytes.Buffer{}
	err := config.Tmpl.ExecuteContext(buf, ctx)
	if err != nil {
		return router.InternalError{err}
	} else {
		return router.Rendered{Content: buf}
	}
}

type PageCtrl struct {
	RenderableCtrl
}

func (PageCtrl) Path() string {
	return ""
}

func (pc PageCtrl) Index() router.Result {
	pc.Template = "front.html"
	return pc.Render()
}

func (pc PageCtrl) Show() router.Result {
	pc.Template = "front.html"
	return pc.Render()
}

type FrontGameCtrl struct {
	RenderableCtrl
}

func (FrontGameCtrl) Path() string {
	return "games"
}
