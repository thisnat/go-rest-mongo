package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/thisnat/go-rest-mongo/configs"
	"github.com/thisnat/go-rest-mongo/models"
	"github.com/thisnat/go-rest-mongo/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var noteCollection *mongo.Collection = configs.GetCollection(configs.DB, "notes")
var validate = validator.New()

func CreateNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var note models.Note
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&note); err != nil {
			c.JSON(http.StatusBadRequest, responses.NoteResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&note); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.NoteResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newNote := models.Note{
			Id:         primitive.NewObjectID(),
			UserId:     note.UserId,
			Content:    note.Content,
			CreateDate: primitive.NewDateTimeFromTime(time.Now()),
		}

		result, err := noteCollection.InsertOne(ctx, newNote)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.NoteResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.NoteResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetNotesByUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var notes []models.Note
		userId := c.Param("userId")

		results, err := noteCollection.Find(ctx, bson.M{"userId": userId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.NoteResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleNote models.Note
			if err = results.Decode(&singleNote); err != nil {
				c.JSON(http.StatusInternalServerError, responses.NoteResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			notes = append(notes, singleNote)
		}

		c.JSON(http.StatusOK,
			responses.NoteResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": notes}},
		)
	}
}

func GetAllNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var notes []models.Note
		results, err := noteCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.NoteResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleNote models.Note
			if err = results.Decode(&singleNote); err != nil {
				c.JSON(http.StatusInternalServerError, responses.NoteResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			notes = append(notes, singleNote)
		}

		c.JSON(http.StatusOK,
			responses.NoteResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": notes}},
		)
	}
}

func GetANote() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		noteId := c.Param("noteId")
		var note models.Note
		objId, _ := primitive.ObjectIDFromHex(noteId)

		err := noteCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&note)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.NoteResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.NoteResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": note}})
	}
}
