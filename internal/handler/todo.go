package handler

import (
	"strconv"

	golang_ture "golang-ture"
	"golang-ture/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create Todo Item
// @Tags todo item
// @Description create todo item
// @ID create-todo-item
// @Accept  json
// @Produce  json
// @Param input body models.TodoItem true "todo item info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items [post]
func (h *Handler) createItem(c *gin.Context) {
	var input models.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.TodoItem.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get All Todo Items
// @Tags todo item
// @Description get all items
// @ID get-all-items
// @Accept  json
// @Produce  json
// @Success 200 {array} models.TodoItem
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items [get]
func (h *Handler) getAllItems(c *gin.Context) {
	items, err := h.service.TodoItem.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Summary Get Todo Item By Id
// @Tags todo item
// @Description get todo item by id
// @ID get-todo-item-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} models.TodoItem
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/:id [get]
func (h *Handler) getItemById(c *gin.Context) {
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	item, err := h.service.TodoItem.GetById(itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Update Todo Item By Id
// @Tags todo item
// @Description update todo item by id
// @ID update-todo-item-by-id
// @Accept  json
// @Produce  json
// @Param input body models.UpdateTodoItemInput true "todo item update info"
// @Success 200 {object} models.TodoItem
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/:id [put]
func (h *Handler) updateItem(c *gin.Context) {
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}
	var input models.UpdateTodoItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	item, err := h.service.TodoItem.Update(itemId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Delete Todo Item By Id
// @Tags todo item
// @Description delete todo item by id
// @ID delete-todo-item-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/:id [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.service.TodoItem.Delete(itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Get All Todo Items In Template
// @Tags todo item
// @Description get all todo items in template
// @ID get-all-todo-items-in-template
// @Accept  html
// @Produce  html
// @Success 200 {file} html
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/todo-list-page [get]
func (h *Handler) getItemsListPage(c *gin.Context) {
	tmpl, ok := golang_ture.Templates["todo_list_page.html"]
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "todo_list_page template not found")
		return
	}

	items, err := h.service.TodoItem.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := tmpl.Execute(c.Writer, items); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}
