package main

import (
	"log"
	"net/http"
	"os"

	db "todo-golang.com/DB"
	"todo-golang.com/middlewares"
	"todo-golang.com/routes"
)

func main() {
	db.InitDB()
	
	http.HandleFunc("GET /todos", routes.GetAllTodos)
	http.HandleFunc("GET /todos/{id}", routes.GetTodoById)
	http.Handle("POST /todos",  middlewares.IsAuth(http.HandlerFunc(routes.CreateTodo)))
	http.Handle("DELETE /todos/{id}",  middlewares.IsAuth(http.HandlerFunc(routes.DeleteTodo)))
	http.Handle("PUT /todos/{id}",  middlewares.IsAuth(http.HandlerFunc(routes.UpdateTodo)))

	http.HandleFunc("POST /signup", routes.Signup)
	http.HandleFunc("POST /login", routes.Login)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server listening on port:", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
