package controller

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/mizu-ryo/echo-sam1/src/model"
)


func LogIndex(c echo.Context) error {
  logs := model.LogIndex()
  return c.JSON(http.StatusOK, logs)
}

func AddLog(c echo.Context) error {
    log := new(model.Log)
    if err := c.Bind(log); err != nil {
        return err
    }

    model.CreateLog(log)

    return c.JSON(http.StatusCreated, log)
}

