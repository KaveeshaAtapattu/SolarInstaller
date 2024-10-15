package controllers

import (
	"SolarInstaller/config"
	"SolarInstaller/models"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "github.com/thedevsaddam/renderer"
)

// Renderer instance
// var rnd = renderer.New()

// CreateProjectHandler adds a new project
func CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Project
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		rnd.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body",
		})
		return
	}

	projectModel := models.ProjectModel{
		ProjectName: p.ProjectName,
		Location:    p.Location,
		DueDate:     p.DueDate,
		Status:      p.Status,
	}

	collection := config.DB.Collection("projects")
	result, err := collection.InsertOne(context.TODO(), projectModel)
	if err != nil {
		rnd.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create project",
			"error":   err.Error(),
		})
		return
	}

	rnd.JSON(w, http.StatusCreated, map[string]interface{}{
		"message": "Project created successfully",
		"id":      result.InsertedID,
	})
}

// GetProjectsHandler retrieves all projects
func GetProjectsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("get projects")
	collection := config.DB.Collection("projects")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		rnd.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve projects",
			"error":   err.Error(),
		})
		return
	}

	var projects []models.ProjectModel
	if err = cursor.All(context.TODO(), &projects); err != nil {
		rnd.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to parse projects",
			"error":   err.Error(),
		})
		return
	}

	var projectList []models.Project
	for _, p := range projects {
		projectList = append(projectList, models.Project{
			ID:          p.ID.Hex(),
			ProjectName: p.ProjectName,
			Location:    p.Location,
			DueDate:     p.DueDate,
			Status:      p.Status,
		})
	}

	rnd.JSON(w, http.StatusOK, map[string]interface{}{"data": projectList})
}

// GetProjectByIDHandler retrieves a project by ID
func GetProjectByIDHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("get projects by id")
	id := chi.URLParam(r, "id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		rnd.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid ID format",
		})
		return
	}

	var project models.ProjectModel
	collection := config.DB.Collection("projects")
	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&project)
	if err != nil {
		rnd.JSON(w, http.StatusNotFound, map[string]interface{}{
			"message": "Project not found",
		})
		return
	}

	rnd.JSON(w, http.StatusOK, map[string]interface{}{"data": project})
}

// UpdateProjectHandler updates a project by ID
func UpdateProjectHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		rnd.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid ID format",
		})
		return
	}

	var updateData models.Project
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		rnd.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body",
		})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"projectName": updateData.ProjectName,
			"location":    updateData.Location,
			"dueDate":     updateData.DueDate,
			"status":      updateData.Status,
		},
	}

	collection := config.DB.Collection("projects")
	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		rnd.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to update project",
			"error":   err.Error(),
		})
		return
	}

	if result.MatchedCount == 0 {
		rnd.JSON(w, http.StatusNotFound, map[string]interface{}{
			"message": "Project not found",
		})
		return
	}

	rnd.JSON(w, http.StatusOK, map[string]interface{}{
		"message": "Project updated successfully",
	})
}

// DeleteProjectHandler deletes a project by ID
func DeleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		rnd.JSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid ID format",
		})
		return
	}

	collection := config.DB.Collection("projects")
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		rnd.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to delete project",
			"error":   err.Error(),
		})
		return
	}

	rnd.JSON(w, http.StatusOK, map[string]interface{}{
		"message": "Project deleted successfully",
	})
}
