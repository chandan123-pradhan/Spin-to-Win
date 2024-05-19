package helper
import (
	"encoding/json"
	"net/http"
	"school_management_app/models"
	"school_management_app/services"
	// "time"
	// "github.com/google/uuid"
)
func AddUserScore(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        var scoreRequest models.UpdateUserScore
        if err := json.NewDecoder(r.Body).Decode(&scoreRequest); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        // Call the services package to add the user score
        userScore, err := services.UpdateUserScore(scoreRequest.ID, int(scoreRequest.Score))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Respond with the user score model
        json.NewEncoder(w).Encode(userScore)
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}