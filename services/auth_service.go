package services

import (
	// "fmt"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"school_management_app/controllers"
	"school_management_app/models"

	"github.com/go-sql-driver/mysql"
)


func GetAllTodos() ([]models.User, error) {
	var todos []models.User
	rows, err := controllers.DB.Query("SELECT id, name, email, phone, photo FROM User_list")
	if err != nil {
		log.Println("Error querying database:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.User
		if err := rows.Scan(&todo.ID, &todo.Name, &todo.Email,&todo.Phone, &todo.Photo, ); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error after rows iteration:", err)
		return nil, err
	}
	return todos, nil
}
func AddTodo(name string, email string, phone string, photo string, password string) (int64, error) {
hashedPassword := HashPassword(password)
    result, err := controllers.DB.Exec("INSERT INTO User_list (name, email, phone, photo, password) VALUES (?, ?, ?, ?, ?)", name, email, phone, photo,hashedPassword)
    if err != nil {
		if mysqlError, ok := err.(*mysql.MySQLError); ok && mysqlError.Number == 1062 {
            // MySQL error number 1062 represents a unique constraint violation
            // Return custom JSON response for unique constraint violation
            errorMessage := map[string]string{
				"status_code":"101",
                "error": "Duplicate entry",
                "message": "Email or phone number already exists",
            }
            jsonError, _ := json.Marshal(errorMessage)
            return 0,errors.New(string(jsonError))
		}
    }

    lastInsertedID, err := result.LastInsertId()
    if err != nil {
        log.Println("Error getting last insert ID:", err)
        return 0, err
    }

    log.Printf("New user added successfully with ID: %d\n", lastInsertedID)
    return lastInsertedID, nil
}

func LoginUser(email, password string) (models.User, error) {
    var user models.User
    // Hash the input password
    hashedPassword := HashPassword(password)

    // Query the database for the user with the provided email and hashed password
    err := controllers.DB.QueryRow("SELECT id, name, email, phone, photo FROM User_list WHERE email = ? AND password = ?", email, hashedPassword).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Photo)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            // If no rows were found, return a custom error indicating invalid credentials
            return models.User{}, errors.New("invalid email or password")
        }
        log.Println("Error querying database:", err)
        return models.User{}, err
    }
    // Return the user model if authentication is successful
    return user, nil
}


func HashPassword(password string) string {
    hasher := md5.New()
    hasher.Write([]byte(password))
    return hex.EncodeToString(hasher.Sum(nil))
}