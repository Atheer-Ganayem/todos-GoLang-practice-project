package middlewares

import (
	"context"
	"net/http"

	"todo-golang.com/utils"
)

func IsAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result, ok := r.Header["Authorization"]
		if !ok {
			utils.JsonResponse(w, utils.JsonObj{"message": "Unauthorized"}, http.StatusUnauthorized)
			return
		}

		userId, err := utils.VerifyToken(result[0])
		if err != nil {
			utils.JsonResponse(w, utils.JsonObj{"message": "Couldn't parse token"}, http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), "userId", userId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}