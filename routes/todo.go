package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"todo-golang.com/models"
	"todo-golang.com/utils"
)



func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		utils.JsonResponse(w, utils.JsonObj{"message": "invalidInputs"}, http.StatusBadRequest)
		return
	}

	userId := r.Context().Value("userId").(int64)
	todo.UserId = userId
	todo.CreatedAt = time.Now()
	err = todo.Save()
	if err != nil {
		utils.JsonResponse(w, utils.JsonObj{"message": err.Error()}, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	todos, err := models.GetTodos()
	if err != nil {
		utils.JsonResponse(w, utils.JsonObj{"message": "couldn't fetch todos from the database."}, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todos)
}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := utils.GetIdParam(r.URL.Path)
	if err != nil {
		utils.JsonResponse(w, utils.JsonObj{"message": "Invalid id, please enter a valid todo id."}, http.StatusBadRequest)
		return
	}
	
	todo, err := models.GetTodo(id)
	if err != nil {
		utils.JsonResponse(w, utils.JsonObj{"message": "Couldn't feth todo."}, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := utils.GetIdParam(r.URL.Path)
	if err != nil {
		utils.JsonResponse(w, utils.JsonObj{"message": "Invalid id, please enter a valid todo id."}, http.StatusBadRequest)
		return
	}

	todo, err := models.GetTodo(id)
	if err != nil {
		utils.JsonResponse(w, utils.JsonObj{"message": "Could not fetch todo."}, http.StatusInternalServerError)
		return
	}

	userId := r.Context().Value("userId").(int64)
	if todo.UserId != userId {
		utils.JsonResponse(w, utils.JsonObj{"message": "Unauthorized."}, http.StatusUnauthorized)
		return
	}

	var bodyTodo models.Todo
	if err = json.NewDecoder(r.Body).Decode(&bodyTodo); err != nil {
		utils.JsonResponse(w, utils.JsonObj{"message": "Could not update todo."}, http.StatusInternalServerError)
		return
	}

	todo.Text = bodyTodo.Text
	err = todo.Update()
	if err != nil {
		utils.JsonResponse(w, utils.JsonObj{"message": "Could not update todo."}, http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, utils.JsonObj{"message": "Todo updated Successfully"}, http.StatusCreated)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := utils.GetIdParam(r.URL.Path)
	if err != nil {
		utils.JsonResponse(w, utils.JsonObj{"message": "Invalid id, please enter a valid todo id."}, http.StatusBadRequest)
		return
	}
	
	todo, err := models.GetTodo(id)
	if err != nil {
		utils.JsonResponse(w, utils.JsonObj{"message": "Could not fetch todo."}, http.StatusInternalServerError)
		return
	}

	userId := r.Context().Value("userId").(int64)
	if todo.UserId != userId {
			utils.JsonResponse(w, utils.JsonObj{"message": "Unauthorized."}, http.StatusUnauthorized)
			return
	}

	err = todo.Delete()
	if err != nil {
		utils.JsonResponse(w, utils.JsonObj{"message": "Couldn't delete todo"}, http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, utils.JsonObj{"message": "Todo deleted successfully"}, http.StatusOK)
}