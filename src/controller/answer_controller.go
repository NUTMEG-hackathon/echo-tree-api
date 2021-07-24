package controller

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/mizu-ryo/echo-sam1/src/model"
)


func AnswerIndex(c echo.Context) error {
  answers := model.AnswerIndex()
  return c.JSON(http.StatusOK, answers)
}

