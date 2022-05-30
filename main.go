package rhodium

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"text/template"
)

// App -
type App struct {
	router *mux.Router
}

// New -
func New() *App {
	r := mux.NewRouter()
	app := &App{
		router: r,
	}
	return app
}

type Context struct {
	Headers  http.Header
	request  *http.Request
	response http.ResponseWriter
	views    *template.Template
}

type RPCContext struct {
	Headers  http.Header
	request  *http.Request
	response http.ResponseWriter
}

// Body -
func (c *RPCContext) Body(body interface{}) error {
	return json.NewDecoder(c.request.Body).Decode(&body)
}

// Response -
func (c *RPCContext) Response(data map[string]interface{}) error {
	return json.NewEncoder(c.response).Encode(data)
}

// View -
func (ctx *Context) View(name string, data map[string]interface{}) error {
	tmpl, err := template.ParseFiles("public/base.html.tmpl", fmt.Sprintf("%s.html.tmpl", name))
	if err != nil {
		return err
	}
	return tmpl.ExecuteTemplate(ctx.response, "base", data)
}

type Handler func(ctx Context) error

func (app *App) wrapHandler(handler Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx := Context{
			Headers:  r.Header,
			request:  r,
			response: rw,
		}
		if err := handler(ctx); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			_, _ = rw.Write([]byte(err.Error()))
			return
		}
	}
}

func (app *App) wrapRPCHandler(handler RPCHandler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx := RPCContext{
			Headers:  r.Header,
			request:  r,
			response: rw,
		}
		response, err := handler(ctx)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			_, _ = rw.Write([]byte(err.Error()))
			return
		}

		rw.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(rw).Encode(response)
	}
}

func (app *App) Get(path string, handler Handler) {
	app.Route(Route{path, http.MethodGet, handler})
}

func (app *App) Post(path string, handler Handler) {
	app.Route(Route{path, http.MethodGet, handler})
}

type RPCHandler func(ctx RPCContext) (map[string]interface{}, error)

func (app *App) RPC(name string, handler RPCHandler) {
	app.RPCRoute(RPCRoute{
		Handler: handler,
		Name:    "/rpc/" + name,
	})
}

type RPCRoute struct {
	Name    string
	Handler RPCHandler
}

type Route struct {
	Path    string
	Method  string
	Handler Handler
}

func (app *App) Route(route Route) {
	app.router.HandleFunc(route.Path, app.wrapHandler(route.Handler)).Methods(route.Method)
}

func (app *App) RPCRoute(route RPCRoute) {
	app.router.HandleFunc("/rpc/"+route.Name, app.wrapRPCHandler(route.Handler)).Methods(http.MethodPost)
}

type Routes interface {
	Routes() []Route
}

type RPCRoutes interface {
	Routes() []RPCRoute
}

func (app *App) Routes(routes Routes) {
	for _, r := range routes.Routes() {
		app.Route(r)
	}
}

func (app *App) RPCRoutes(routes RPCRoutes) {
	for _, r := range routes.Routes() {
		app.RPCRoute(r)
	}
}

// Run -
func (app *App) Run(addr string) error {
	fs := http.FileServer(http.Dir("./public"))
	app.router.Handle("/", fs)
	err := http.ListenAndServe(addr, app.router)
	if err != nil {
		return errors.Wrap(err, "error starting server")
	}
	return nil
}
