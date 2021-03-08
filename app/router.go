package app

import (
	"cach-plan.go/app/handler"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/projects", a.handleRequest(handler.GetAllProjects))
	a.Post("/projects", a.handleRequest(handler.CreateProject))
	a.Get("/projects/{title}", a.handleRequest(handler.GetProject))
	a.Put("/projects/{title}", a.handleRequest(handler.UpdateProject))
	a.Delete("/projects/{title}", a.handleRequest(handler.DeleteProject))
	a.Put("/projects/{title}/archive", a.handleRequest(handler.ArchiveProject))
	a.Delete("/projects/{title}/archive", a.handleRequest(handler.RestoreProject))

	// Routing for handling the tasks
	a.Get("/projects/{title}/tasks", a.handleRequest(handler.GetAllTasks))
	a.Post("/projects/{title}/tasks", a.handleRequest(handler.CreateTask))
	a.Get("/projects/{title}/tasks/{id:[0-9]+}", a.handleRequest(handler.GetTask))
	a.Put("/projects/{title}/tasks/{id:[0-9]+}", a.handleRequest(handler.UpdateTask))
	a.Delete("/projects/{title}/tasks/{id:[0-9]+}", a.handleRequest(handler.DeleteTask))
	a.Put("/projects/{title}/tasks/{id:[0-9]+}/complete", a.handleRequest(handler.CompleteTask))
	a.Delete("/projects/{title}/tasks/{id:[0-9]+}/complete", a.handleRequest(handler.UndoTask))
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%v] %v\n", r.Method, r.RequestURI)
		handler(a.DB, w, r)
	}
}
