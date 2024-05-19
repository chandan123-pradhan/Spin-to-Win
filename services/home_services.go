package services

import (
	"log"
	"school_management_app/controllers"
	"school_management_app/models"
)
func UpdateUserScore(userID int64, score int) (models.UpdateUserScore, error) {
    // Check if a record exists for the given userID
    var count int
    err := controllers.DB.QueryRow("SELECT COUNT(*) FROM user_score WHERE user_id = ?", userID).Scan(&count)
    if err != nil {
        log.Println("Error checking if record exists:", err)
        return models.UpdateUserScore{}, err
    }

    if count == 0 {
        // If no record exists, insert a new record
        _, err = controllers.DB.Exec("INSERT INTO user_score (user_id, score) VALUES (?, ?)", userID, score)
        if err != nil {
            log.Println("Error inserting new record:", err)
            return models.UpdateUserScore{}, err
        }
    } else {
        // If a record exists, update the score
        _, err = controllers.DB.Exec("UPDATE user_score SET score = ? WHERE user_id = ?", score, userID)
        if err != nil {
            log.Println("Error updating record:", err)
            return models.UpdateUserScore{}, err
        }
    }

	
    // Return the updated user score
    return models.UpdateUserScore{ID: userID, Score: score}, nil
}

