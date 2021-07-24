package controller

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/mizu-ryo/echo-sam1/src/model"
)

func Hello(c echo.Context) error {
  return c.JSON(http.StatusOK, map[string]string{"message":"Hello, world!!!"})
}

func Seed(c echo.Context) error {
  model.UserSeed()
  model.TreeSeed()
  model.LogSeed()
  model.QuestionSeed()
  model.AnswerSeed()
  users := model.UserIndex()
  return c.JSON(http.StatusOK, users)
}

