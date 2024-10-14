package controllers

import (
	"SolarInstaller/config"
	"SolarInstaller/models"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
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
	id := chi.URLParam(r, "id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		rnd.JSON(w, http.StatusBadRequest, map[string]interface{}{"message": "Invalid ID"})
		return
	}

	var t models.Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		rnd.JSON(w, http.StatusBadRequest, err)
		return
	}

	collection := config.DB.Collection("Solar")
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"title": t.Title, "completed": t.Completed}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		rnd.JSON(w, http.StatusInternalServerError, err)
		return
	}

	rnd.JSON(w, http.StatusOK, map[string]interface{}{"message": "Todo updated"})
}

// DeleteTodoHandler deletes a todo by ID
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
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
    // Get the ID from the URL
    id := chi.URLParam(r, "id")
	println(id)

    // Convert the ID string to MongoDB's ObjectID
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        rnd.JSON(w, http.StatusBadRequest, map[string]interface{}{
            "message": "Invalid ID format",
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
