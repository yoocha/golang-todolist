package app

import (
	"fmt"

	"strconv"

	"net/http"

	"encoding/json"

	"io/ioutil"

	"github.com/gin-gonic/gin"

	service "example.com/todolist/service"

	entity "example.com/todolist/entity"
)

var todoService *service.TodoService

func init() {
	todoService = service.CreateTodoService()
}

// @Summary save Todo
// @Schemes
// @Description save Todo
// @Tags Todo API
// @Accept json
// @Produce json
// @Param request body string true "{}"
// @Success 200 {string} Helloworld
// @Router /todos [post]
func saveTodo(c *gin.Context) {

	todo := requestBody(c, entity.Todo{})

	saveTodo := todoService.SaveTodo(todo)

	c.JSON(http.StatusOK, saveTodo)
}

// @Summary modify Todo
// @Schemes
// @Description modify Todo
// @Tags Todo API
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param request body string true "{}"
// @Success 200 {string} Helloworld
// @Router /todos/{id} [put]
func modifyTodo(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Params.ByName("id"), 10, 64)

	todo := requestBody(c, entity.Todo{})

	saveTodo := todoService.UpdateTodo(id, todo)

	c.JSON(http.StatusOK, saveTodo)
}

// @Summary modify Todo
// @Schemes
// @Description modify Todo
// @Tags Todo API
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {string} Helloworld
// @Router /todos/{id}/done [put]
func completeTodo(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Params.ByName("id"), 10, 64)

	saveTodo := todoService.CompleteTodo(id)

	c.JSON(http.StatusOK, saveTodo)
}

// @Summary delete Todo
// @Schemes
// @Description delete Todo
// @Tags Todo API
// @Param id path int true "Todo ID"
// @Success 200 {string} Helloworld
// @Router /todos/{id} [delete]
func deleteTodo(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Params.ByName("id"), 10, 64)

	todoService.DeleteTodo(id)

	c.Status(http.StatusOK)
}

// @Summary get Todo
// @Schemes
// @Description get Todo
// @Tags Todo API
// @Param id path int true "Todo ID"
// @Success 200 {string} Helloworld
// @Router /todos/{id} [get]
func getTodo(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Params.ByName("id"), 10, 64)

	todo := todoService.GetTodo(id)

	c.JSON(http.StatusOK, todo)
}

// @Summary get Todo list
// @Schemes
// @Description get Todo list
// @Tags Todo API
// @Success 200 {string} Helloworld
// @Router /todos [get]
func getTodos(c *gin.Context) {

	todos := todoService.GetTodos()

	c.JSON(http.StatusOK, todos)
}

func TodoAppRoute(r *gin.RouterGroup) {

	todoApi := r.Group("/todos")

	todoApi.GET("/:id", getTodo)
	todoApi.GET("", getTodos)
	todoApi.POST("", saveTodo)
	todoApi.PUT("/:id", modifyTodo)
	todoApi.PUT("/:id/done", completeTodo)
	todoApi.DELETE("/:id", deleteTodo)
}

func requestBody[T interface{}](c *gin.Context, types T) T {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	json.Unmarshal(value, &types)

	return types
}

func toJsonString[T interface{}](types T) []byte {
	resp, _ := json.Marshal(types)

	return resp
}
