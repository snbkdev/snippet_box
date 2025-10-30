package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	mux.Handle("GET /snippet/create", dynamic.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", dynamic.ThenFunc(app.snippetCreatePost))
	// Для пользователя 
	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.userSignUp))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.userSignUpPost))
	mux.Handle("GET /user/login", dynamic.ThenFunc(app.userLogIn))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.userLogInPost))
	mux.Handle("POST /user/logout", dynamic.ThenFunc(app.userLogOutPost))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}