package handler

import (
	"github.com/gin-gonic/gin"
	"golang-ture/internal/models"
	"html/template"
	"net/http"
)

const defaultTodoHtmlTemplatePath = "internal/templates/html/todo/"

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

func (h *Handler) getAllItems(c *gin.Context) {
	items, err := h.service.TodoItem.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) getItemById(c *gin.Context) {
	itemId := c.Param("id")

	item, err := h.service.TodoItem.GetById(itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) updateItem(c *gin.Context) {
	id := c.Param("id")
	var input models.UpdateTodoItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.TodoItem.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteItem(c *gin.Context) {
	itemId := c.Param("id")

	err := h.service.TodoItem.Delete(itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) getItemsListPage(c *gin.Context) {
	funcMap := map[string]interface{}{
		"Increment": func(i int) int {
			return i + 1
		},
	}
	fileName := defaultTodoHtmlTemplatePath + "todo_list_page.html"
	tmpl, err := template.New("todo_list_page.html").Funcs(funcMap).ParseFiles(fileName)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
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
