package controllers

import (
	"SolarInstaller/config"
	"SolarInstaller/models"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	// "github.com/go-chi/chi/v5"
	"github.com/thedevsaddam/renderer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Renderer instance
var rnd = renderer.New()

// FetchTodosHandler retrieves all todos
func FetchTodosHandler(w http.ResponseWriter, r *http.Request) {
	collection := config.DB.Collection("Solar")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		rnd.JSON(w, http.StatusInternalServerError, err)
		return
	}

	var todos []models.TodoModel
	if err = cursor.All(context.TODO(), &todos); err != nil {
		rnd.JSON(w, http.StatusInternalServerError, err)
		return
	}

	var todoList []models.Todo
	for _, t := range todos {
		todoList = append(todoList, models.Todo{
			ID:        t.ID.Hex(),
			Title:     t.Title,
			Completed: t.Completed,
			CreatedAt: t.CreatedAt,
		})
	}

	rnd.JSON(w, http.StatusOK, map[string]interface{}{"data": todoList})
}

// CreateTodoHandler adds a new todo
func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var t models.Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		rnd.JSON(w, http.StatusBadRequest, err)
		return
	}

	tm := models.TodoModel{
		Title:     t.Title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	collection := config.DB.Collection("Solar")
	result, err := collection.InsertOne(context.TODO(), tm)
	if err != nil {
		rnd.JSON(w, http.StatusInternalServerError, err)
		return
	}

	rnd.JSON(w, http.StatusCreated, map[string]interface{}{"message": "Todo created", "id": result.InsertedID})
}

// UpdateTodoHandler updates a todo by ID
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL path
	// Get the ID from the URL path
	pathParts := strings.Split(r.URL.Path, "/")

	// Get the last part of the path (the ID)
	id := pathParts[len(pathParts)-1]



	// Attempt to convert the ID to MongoDB's ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid ID format:", err)
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Debug: Log the incoming request body
	log.Println("Reading request body...")
	var updateData models.Todo
	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		log.Println("Failed to decode JSON:", err) // Debug: Log the error
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}

	// Debug: Log the parsed data
	log.Printf("Received Data: %+v\n", updateData)

	// Prepare the update document
	update := bson.M{
		"$set": bson.M{
			"title":     updateData.Title,
			"completed": updateData.Completed,
		},
	}

	// Perform the update in the database
	collection := config.DB.Collection("Solar")
	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		log.Println("Failed to update todo:", err)
		http.Error(w, "Failed to update todo", http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Send a success response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Todo updated successfully",
	})
}

// DeleteTodoHandler deletes a todo by ID
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL path
	pathParts := strings.Split(r.URL.Path, "/")

	// Get the last part of the path (the ID)
	id := pathParts[len(pathParts)-1]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		rnd.JSON(w, http.StatusBadRequest, map[string]interface{}{"message": "Invalid ID"})
		return
	}

	collection := config.DB.Collection("Solar")
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		rnd.JSON(w, http.StatusInternalServerError, err)
		return
	}

	rnd.JSON(w, http.StatusOK, map[string]interface{}{"message": "Todo deleted"})
}

func GetTodoByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL path
	pathParts := strings.Split(r.URL.Path, "/")

	// Get the last part of the path (the ID)
	id := pathParts[len(pathParts)-1]

	// Debug: Log the extracted ID
	log.Println("Extracted ID:", id)

	// Check if the ID is a valid 24-character hexadecimal string
	if len(id) != 24 {
		rnd.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid ID format: must be a 24-character hexadecimal string",
		})
		return
	}

	// Convert the ID string to MongoDB's ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		rnd.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to parse ID to ObjectID",
			"error":   err.Error(),
		})
		return
	}

	// Fetch the todo from the database
	var todo models.TodoModel
	collection := config.DB.Collection("Solar")
	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&todo)
	if err != nil {
		rnd.JSON(w, http.StatusNotFound, map[string]interface{}{
			"message": "Todo not found",
		})
		return
	}

	// Send the fetched todo as a response
	rnd.JSON(w, http.StatusOK, map[string]interface{}{
		"data": todo,
	})
}
