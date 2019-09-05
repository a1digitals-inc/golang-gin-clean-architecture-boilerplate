package controllers

import (
	"golang-boilerplates/api/database"
	"golang-boilerplates/api/helper"
	"golang-boilerplates/api/repositories"
	"golang-boilerplates/api/repositories/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Paging struct {
	Limit string
	Page  string
	Field string
}

// type

func GetProducts(c *gin.Context) {

	var res = c.Writer
	// var req = c.Request

	limit := c.DefaultPostForm("limit", "10")
	page := c.DefaultPostForm("page", "1")
	field := c.DefaultPostForm("field", "*")
	if field != "*" {
		field = helper.ToStringField(field)
	}
	var arr = make(map[string]string, 0)

	arr["limit"] = limit
	arr["page"] = page
	arr["field"] = field

	db, err := database.Connect()
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	repository := repo.NewRepositoryProducts(db)

	func(productRepository repositories.ProductRepository) {
		products, err := productRepository.FindAll(arr)
		if err != nil {
			helper.ErrorCustomStatus(res, http.StatusUnprocessableEntity, err.Error())
			return
		}
		helper.Responses(res, http.StatusOK, "Success", products)
	}(repository)
}
