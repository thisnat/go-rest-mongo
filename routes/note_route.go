package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thisnat/go-rest-mongo/controllers"
)

func NoteRoute(router *gin.Engine) {
	router.POST("/note/create", controllers.CreateNote())
	router.GET("/note", controllers.GetAllNotes())
	router.GET("/note/by/:userId", controllers.GetNotesByUserId())
	router.GET("/note/:noteId", controllers.GetANote())
}
