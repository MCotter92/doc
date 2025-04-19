package api

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/MCotter92/doc/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

lock = lock.Mutex{}

func createDocument(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	d := utils.Document
}
