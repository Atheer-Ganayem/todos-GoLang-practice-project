package routes

import (
	"encoding/json"
	"net/http"

	"todo-golang.com/models"
	"todo-golang.com/utils"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.JsonResponse(w, utils.JsonObj{"message": "Invalid inputs"}, http.StatusBadRequest)
		return 
	}

	if err := user.Save(); err != nil {
		utils.JsonResponse(w, utils.JsonObj{"message": "Couldn't create user"}, http.StatusInternalServerError)
		return 
	}

	utils.JsonResponse(w, utils.JsonObj{"message": "User created successfully"}, http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.JsonResponse(w, utils.JsonObj{"message": "Couldn't parse request body."}, http.StatusBadRequest)
		return 
	}

	if err := user.ValidateCredentials(); err != nil {
		utils.JsonResponse(w, utils.JsonObj{"message": "Invalid email or password"}, http.StatusBadRequest)
		return 
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		utils.JsonResponse(w, utils.JsonObj{"message": "Couldn't generate auth token."}, http.StatusInternalServerError)
		return 
	}

	utils.JsonResponse(w, utils.JsonObj{"message": "User logged on successfully", "token": token}, http.StatusOK)
}