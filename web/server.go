package main

import (
	"bytes"
	"log"
	"net/http"
	"strconv"

	"git.andrewcsellers.com/acsellers/card_sharp/config"
	"git.andrewcsellers.com/acsellers/card_sharp/lobby"
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
	r.Many(FrontGameCtrl{NewRenderableCtrl("desktop.html"), nil})
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
	*lobby.Lobby
}

func (FrontGameCtrl) Path() string {
	return "games"
}

func (fgc *FrontGameCtrl) PreItem() router.Result {
	var value string
	if cookie, err := fgc.Request.Cookie("party-lobby"); err != nil {
		fgc.Log.Printf("Could not retrieve Cookie: %s\n", err.Error())
		return router.NotAllowed{
			Request:  fgc.Request,
			Fallback: "/games/new",
		}
	} else {
		value = cookie.Value
	}
	var lobbyid string
	if err := config.Cookie.Decode("party-lobby", value, &lobbyid); err != nil {
		fgc.Log.Printf("Could not decode Cookie: %s\n", err.Error())
		return router.NotAllowed{
			Request:  fgc.Request,
			Fallback: "/games/new",
		}
	}
	l := lobby.Find(lobbyid)
	if l == nil {
		fgc.Log.Println("Could not retrieve lobby from cookie")
		return router.NotAllowed{
			Request:  fgc.Request,
			Fallback: "/games/new",
		}
		return nil
	}
	fgc.Lobby = l
	fgc.Context["Lobby"] = l
	return nil
}
func (fgc FrontGameCtrl) New() router.Result {
	fgc.Template = "new_game_lobby.html"
	return fgc.Render()

}

func (fgc FrontGameCtrl) Create() router.Result {
	fgc.Request.ParseForm()
	gid := fgc.Request.Form.Get("game_id")
	if gid == "" {
		return router.Redirect{
			Request: fgc.Request,
			URL:     "/games/new",
		}
	}

	id, err := strconv.Atoi(gid)
	if err != nil {
		return router.Redirect{
			Request: fgc.Request,
			URL:     "/games/new",
		}
	}

	g, err := config.Conn.Deck.Find(id)
	if err != nil {
		return router.Redirect{
			Request: fgc.Request,
			URL:     "/games/new",
		}
	}
	l := lobby.Create(g)
	en, err := config.Cookie.Encode("party-lobby", l.ID)
	if err != nil {
		return router.Redirect{
			Request: fgc.Request,
			URL:     "/games/new",
		}
	}

	http.SetCookie(fgc.Out, &http.Cookie{Name: "party-lobby", Value: en, Path: "/"})
	return router.Redirect{
		Request: fgc.Request,
		URL:     "/games/" + l.ID,
	}
}

func (fgc FrontGameCtrl) Show() router.Result {
	fgc.Layout = "lobby.html"
	fgc.Template = "game_lobby.html"
	return fgc.Render()

}
func (fgc FrontGameCtrl) Edit() router.Result {
	fgc.Template = "edit_game_lobby.html"
	return fgc.Render()
}

func (fgc FrontGameCtrl) Update() router.Result {
	return router.Redirect{
		Request: fgc.Request,
		URL:     "/",
	}
}

func (fgc FrontGameCtrl) Join() router.Result {
	fgc.Template = "join_game.html"
	return fgc.Render()
}

func (fgc FrontGameCtrl) OtherBase(sr *router.SubRoute) {
	sr.Get("join").Action("Join")
}
