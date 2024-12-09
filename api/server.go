package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/MCotter92/doc/utils"
	"github.com/djherbis/times"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func makeDocument(file string, Keywords []string) utils.DocumentStore {

	loc := filepath.Dir(file)
	if loc == "." {
		var err error
		loc, err = os.Getwd()
		if err != nil {
			fmt.Printf("Cannot retrieve file location: %v\n", err)
		}
	}

	ext := filepath.Ext(file)

	title := filepath.Base(file)

	times, err := times.Stat("file")
	if err != nil {
		fmt.Printf("Cannot retreive fiel stats", err)
	}

	newDocument := utils.DocumentStore{
		UUID:             uuid.NewString(),
		CreatedDate:      times.BirthTime(),
		LastModifiedDate: times.ChangeTime(),
		Keywords:         Keywords,
		Document: utils.Document{
			Title:     title,
			Extension: ext,
			Location:  loc,
		},
	}
	return newDocument
}

func storeDocument(c echo.Context) error {

	file := c.QueryParam("file")
	keywords := c.QueryParam("Keywords")

	if file == "" || keywords == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Both file and keywords are required"})
	}

	keywordsList := strings.Split(keywords, ",")
	newDocument := makeDocument(file, keywordsList)

	docJson, err := json.MarshalIndent(newDocument, "", "\t")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error marshalling document"})
	}

	// Open the Json file
	globalJson, err := os.OpenFile("data/global.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Cannot open file"})
	}
	defer globalJson.Close()

	// write the struct turned json to global.json
	if _, err := globalJson.Write(docJson); err != nil {
		globalJson.Close()
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Cannot write to docJson"})
	}

	return c.JSON(http.StatusOK, newDocument)
}

func api_PUT() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("doc/data/golbal.json", storeDocument)
	e.Logger.Fatal(e.Start(":1323"))
}
