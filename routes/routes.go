package routes

import (
	"github.com/gin-gonic/gin"
    "todo-app/controllers"
)


func SetupRouter() *gin.Engine{
   r:= gin.Default()
   r.POST("/todos", controllers.CreateTodo)
   r.GET("/todos", controllers.GetTodos)
   r.PUT("/todos/:id",controllers.UpdateTodoHandler)
   r.DELETE("/todos/:id", controllers.DeleteTodoHandler)
   return r
}