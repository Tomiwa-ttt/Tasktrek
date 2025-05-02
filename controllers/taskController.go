package controllers

import (
	"context"
	"net/http"
	"tasktrek/config"
	"tasktrek/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTask(c *gin.Context) {
	taskCollection := config.DB.Database("tasktrek").Collection("tasks")

	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task data"})
		return
	}

	// Associate task with user's email from context
	userEmail, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	newTask.UserEmail = userEmail.(string)
	newTask.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := taskCollection.InsertOne(ctx, newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully", "task": newTask})
}

func UpdateTask(c *gin.Context) {
	taskCollection := config.DB.Database("tasktrek").Collection("tasks")

	// Get task ID from URL params
	taskID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask models.Task
	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task data"})
		return
	}

	// Create update data
	updateData := bson.M{}
	if updatedTask.Title != "" {
		updateData["title"] = updatedTask.Title
	}
	if updatedTask.Description != "" {
		updateData["description"] = updatedTask.Description
	}
	updateData["completed"] = updatedTask.Completed

	// Update task
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = taskCollection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": updateData})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}
