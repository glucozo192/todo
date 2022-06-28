package controllers

import (
	"TOGO/configs"
	"TOGO/responses"
	"TOGO/untils"

	//"TOGO/middleware"
	"TOGO/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"github.com/gorilla/context"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

// var userCollection *dbiface.ConllectionAPI = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func GetAUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Go to GetAUser")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		userId := params["userId"]
		var user models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		responses.WriteResponse(rw, http.StatusOK, user)
	}
}

func EditAUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		userId := params["userId"]
		var user models.User
		var check models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)
		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&check)
		//validate the request body

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			untils.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		if user.Name == "" {
		}
		fmt.Println(user)
		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			untils.Error(rw, validationErr.Error(), http.StatusBadRequest)
			return
		}

		update := bson.M{"name": user.Name, "vip": user.Vip}

		result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		//get updated user details
		var updatedUser models.User
		if result.MatchedCount == 1 {
			err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)

			if err != nil {
				untils.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		responses.WriteResponse(rw, http.StatusOK, updatedUser)
	}
}

func DeleteAUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		userId := params["userId"]
		objId, _ := primitive.ObjectIDFromHex(userId)
		defer cancel()
		// detele task of user
		results, err := taskCollection.Find(ctx, bson.M{"id_user": objId})
		if err != nil {
			untils.Error(rw, "canot find task of user", http.StatusInternalServerError)
		}
		for results.Next(ctx) {
			var singleTask models.Task
			if err = results.Decode(&singleTask); err != nil {
				untils.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			task, err := taskCollection.DeleteOne(ctx, bson.M{"id": singleTask.Id})
			if err != nil {
				untils.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			if task.DeletedCount < 1 {
				untils.Error(rw, "Task with specified ID not found!", http.StatusNotFound)
				return
			}
		}

		// Delete user
		result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})

		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		if result.DeletedCount < 1 {
			untils.Error(rw, "User with specified ID not found!", http.StatusNotFound)
			return
		}

		responses.WriteResponse(rw, http.StatusOK, "Delete completed")
	}
}

func GetAllUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("go to GetAllUser")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []models.User
		defer cancel()
		results, err := userCollection.Find(ctx, bson.M{})

		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		time.Sleep(time.Second * 2)
		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUser models.User
			if err = results.Decode(&singleUser); err != nil {
				untils.Error(rw, err.Error(), http.StatusBadRequest)
				return
			}

			users = append(users, singleUser)
		}

		responses.WriteResponse(rw, http.StatusOK, users)
	}
}

func GetMe() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("go to Getme")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		Id_User := r.Context().Value("Id_User").(string)
		objId, _ := primitive.ObjectIDFromHex(Id_User)
		defer cancel()

		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			untils.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		newUser := models.User{
			Id:       user.Id,
			Username: user.Username,
			Name:     user.Name,
			Vip:      user.Vip,
			Limit:    user.Limit,
		}
		responses.WriteResponse(rw, http.StatusOK, newUser)

	}
}
