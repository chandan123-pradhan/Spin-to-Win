package helper
import (
	"encoding/json"
	"fmt"
	"net/http"
	"school_management_app/models"
	"school_management_app/services"
	// "time"
	// "github.com/google/uuid"
)




func HandleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		todos, err := services.GetAllTodos()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(todos)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleRegisteration(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var todo models.RegistrationModel
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result, err := services.AddTodo(todo.Name, todo.Email, todo.Phone, todo.Photo, todo.Password)
		if err != nil {
			fmt.Println("error occurend")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond with success status
		w.WriteHeader(http.StatusCreated)
		user.ID = result
		user.Name=todo.Name;
		user.Email=todo.Email;
		user.Phone=todo.Phone;
		user.Photo=todo.Photo;
		
		json.NewEncoder(w).Encode(user)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var loginRequest models.LoginModel
		if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Call the services package to authenticate the user
		user, err := services.LoginUser(loginRequest.Email, loginRequest.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Respond with the authenticated user
		json.NewEncoder(w).Encode(user)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}