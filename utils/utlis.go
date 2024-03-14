package utils

import (
	"reflect"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	StatuseCode int         `json:"status_code"`
	Message     string      `json:"message"`
	Error       string      `json:"error,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	Meta        Meta        `json:"meta,omitempty"` // Include pagination metadata
}

/// pagination
/* *****HOW TO USE***

1. define variable page and limit as below :
page, _ := strconv.Atoi(ctx.Query("page", "1"))
limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

2. created variable for data and for meta pagination in return response

ex : paginateddata, pagination := utils.GetPaginated(ctx, page, limit, Userlist)

*UserList is the data we want to process

return ctx.Status(200).JSON(utils.Response{
		StatuseCode: 200,
		Message:     "success",
		Data:        paginateddata,
		Meta:        utils.Meta{Pagination: pagination},
	})

*/

type Meta struct {
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Total       int   `json:"total"`
	Count       int   `json:"count"`
	PerPage     int   `json:"per_page"`
	CurrentPage int   `json:"current_page"`
	TotalPages  int   `json:"total_pages"`
	Links       Links `json:"links"`
}

type Links struct {
	Next string `json:"next"`
}

func GetPaginated(ctx *fiber.Ctx, page, limit int, data interface{}) (interface{}, Pagination) {

	// Check whether the data sent is sliced (partial data) or not
	sliceValue := reflect.ValueOf(data)
	if sliceValue.Kind() != reflect.Slice {
		return nil, Pagination{}
	}

	// get length data
	dataLength := sliceValue.Len()

	// get startindex and endindex
	startIndex := (page - 1) * limit
	endIndex := page * limit
	if endIndex > dataLength {
		endIndex = dataLength
	}

	// get data after paginated
	paginateddata := sliceValue.Slice(startIndex, endIndex).Interface()

	pagination := Pagination{
		Total:       dataLength,
		Count:       endIndex - startIndex,
		PerPage:     limit,
		CurrentPage: page,
		TotalPages:  (dataLength + limit - 1) / limit,
		Links: Links{
			Next: getNextPageURL(ctx, page, limit, dataLength),
		},
	}

	return paginateddata, pagination
}

func getNextPageURL(ctx *fiber.Ctx, currentPage, limit, totalItems int) string {
	if (currentPage * limit) >= totalItems {
		return ""
	}
	nextPage := currentPage + 1
	return ctx.BaseURL() + ctx.Path() + "?page=" + strconv.Itoa(nextPage) + "&limit=" + strconv.Itoa(limit)
}

////
