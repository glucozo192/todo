package controllers

import (
	"TOGO/configs"
	"TOGO/models"
	"TOGO/responses"
	"TOGO/untils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var taskCollection *mongo.Collection = configs.GetCollection(configs.DB, "tasks")

func CreateTask() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Go to CreateTask")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var task models.Task
		var check models.User
		var count int
		Id_User := r.Context().Value("Id_User").(string)
		objId, _ := primitive.ObjectIDFromHex(Id_User)
		defer cancel()
		//Get user by id user
		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&check)
		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		//Get tasks by id user
		results, _ := taskCollection.Find(ctx, bson.M{"id_user": objId})
		if results != nil {
			// Loop and check day
			for results.Next(ctx) {
				var singleTask models.Task
				if err = results.Decode(&singleTask); err != nil {
					untils.Error(rw, err.Error(), http.StatusInternalServerError)
					return
				}
				if singleTask.CreatedAt == time.Now().Format("01-02-2006") {
					count++
				}
			}
		} else {
			count = 0
		}
		//Get check
		if check.Limit < count {
			untils.Error(rw, " daily task limited, if you want add task, pls upgrate Vip", http.StatusBadRequest)
			return
		}

		// validate body
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			untils.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		//use the validate library to validate required fields
		if validationErr := validate.Struct(&task); validationErr != nil {
			untils.Error(rw, validationErr.Error(), http.StatusBadRequest)
			return
		}
		today := time.Now()
		newTask := models.Task{
			Id:        primitive.NewObjectID(),
			Id_user:   objId,
			Name:      task.Name,
			Content:   task.Content,
			CreatedAt: today.Format("01-02-2006"),
			Status:    "doing",
		}

		result, err := taskCollection.InsertOne(ctx, newTask)
		if err != nil {
			untils.Error(rw, "db not insert", http.StatusInternalServerError)
			return
		}

		responses.WriteResponse(rw, http.StatusCreated, result)

	}
}

func GetTask() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("go to GetTask")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var tasks []models.Task
		// var user models.User
		Id_User := r.Context().Value("Id_User").(string)
		objId, _ := primitive.ObjectIDFromHex(Id_User)
		defer cancel()

		//get tasks
		results, err := taskCollection.Find(ctx, bson.M{"id_user": objId})
		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleTask models.Task
			if err = results.Decode(&singleTask); err != nil {
				untils.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			tasks = append(tasks, singleTask)
		}
		responses.WriteResponse(rw, http.StatusOK, tasks)
	}
}

func GetAllTask() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("go to GetAllTask")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var tasks []models.Task
		defer cancel()
		results, err := taskCollection.Find(ctx, bson.M{})
		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleTask models.Task
			if err = results.Decode(&singleTask); err != nil {
				untils.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			tasks = append(tasks, singleTask)

		}
		responses.WriteResponse(rw, http.StatusOK, tasks)
	}
}

func DeleteTask() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Go to DeleteTask")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		taskId := params["id"]
		defer cancel()
		objIdTask, _ := primitive.ObjectIDFromHex(taskId)

		result, err := taskCollection.DeleteOne(ctx, bson.M{"id": objIdTask})
		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		if result.DeletedCount < 1 {
			untils.Error(rw, "Task with specified ID not found!", http.StatusNotFound)
			return
		}

		responses.WriteResponse(rw, http.StatusOK, "Delete completed")

	}
}

func GetOneTask() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Go to GetOneTask")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		taskId := params["id"]
		var task models.Task
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(taskId)
		//Get task
		err := taskCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&task)
		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		responses.WriteResponse(rw, http.StatusOK, task)

	}
}

func UpdateTask() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Go to UpdateTask")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		taskId := params["id"]
		var task models.Task
		var check models.Task
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(taskId)
		err := taskCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&check)
		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		//validate the request body
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			fmt.Println("erro from val")
			untils.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		if task.Name == "" {
			task.Name = check.Name
		}
		if task.Content == "" {
			task.Content = check.Content
		}
		fmt.Println(task)
		//use the validator library to validate required fields
		if validationErr := validate.Struct(&task); validationErr != nil {
			untils.Error(rw, validationErr.Error(), http.StatusBadRequest)
			return
		}
		update := bson.M{"name": task.Name, "content": task.Content}

		result, err := taskCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		//get updated task details
		var updatedTask models.Task
		if result.MatchedCount == 1 {
			err := taskCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedTask)
			if err != nil {
				untils.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		responses.WriteResponse(rw, http.StatusOK, updatedTask)
	}
}

func UpdateTaskStatus() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("GO to update task status")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		taskId := params["id"]
		objId, _ := primitive.ObjectIDFromHex(taskId)
		var task models.Task
		var check models.Task
		defer cancel()
		//validate the request body
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			fmt.Println("erro from val")
			untils.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		//Get Task
		err := taskCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&check)
		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		//Updatr Task
		update := bson.M{"status": task.Status}
		result, err := taskCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		if result.MatchedCount != 1 {
			untils.Error(rw, "update false", http.StatusInternalServerError)
		}
		check.Status = task.Status
		responses.WriteResponse(rw, http.StatusOK, check)
	}
}
