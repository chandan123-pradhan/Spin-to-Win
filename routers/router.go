package routers

import (
	
	"net/http"
	// "time"
	// "github.com/google/uuid"
	"school_management_app/helper"
)





func SetupRoutes() {
	http.HandleFunc("/get_all_users", helper.HandleGetAllUsers)
	http.HandleFunc("/register", helper.HandleRegisteration)
	http.HandleFunc("/login",helper.HandleLogin)
	http.HandleFunc("/update_score",helper.AddUserScore)

}
