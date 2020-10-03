package handlers

import "github.com/labstack/echo/v4"

// Base Error
type Error struct {
	Code    int
	Message string
	Field string
}

func InternalError(c echo.Context, e string) error {
	result := &Error{Code: 500, Message: "An Internal Error has Occurred", Field: e}
	return c.JSON(result.Code, result)
}

func BadRequestError(c echo.Context, e string) error {
	result := &Error{Code: 400, Message: "The server could not understand the request due to invalid syntax." , Field: e}
	return c.JSON(result.Code, result)
}

