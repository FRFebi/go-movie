package test

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/FRFebi/go-movie/config"
	"github.com/FRFebi/go-movie/controller"
	"github.com/FRFebi/go-movie/helper"
	"github.com/FRFebi/go-movie/middleware"
	"github.com/FRFebi/go-movie/model/domain"
	"github.com/FRFebi/go-movie/repository"
	"github.com/FRFebi/go-movie/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:1234567890@tcp(127.0.0.1:3306)/tht")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) *fiber.App {
	validate := validator.New()

	categoryRepository := repository.NewMovieRepository()
	categoryService := service.NewMovieService(categoryRepository, db, validate)
	categoryController := controller.NewMovieController(categoryService)
	router := config.NewRouter(categoryController)

	router.Use(middleware.Auth)

	return router
}

func truncateCategoryDB(db *sql.DB) {
	db.Exec("TRUNCATE movie")
}

func TestCreateMovieSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategoryDB(db)

	router := setupRouter(db)
	requestBody := strings.NewReader(`{
    "title": "Movie1",
    "description": "Movie nya keren",
    "rating": 8.8,
    "image": "http://localhost:8888/"
}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8888/api/v1/movies", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+base64.StdEncoding.EncodeToString([]byte("FikoRedhaFebiansyah")))

	resp, _ := router.Test(request, 30)
	body, _ := io.ReadAll(resp.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "Get HTTP Status 200")
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["message"])
	assert.Equal(t, "Movie1", responseBody["data"].(map[string]interface{})["title"])
}

func TestCreateMovieFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategoryDB(db)

	router := setupRouter(db)
	requestBody := strings.NewReader(`{
    "title": "",
    "description": "",
    "rating": null,
    "image": ""
}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8888/api/v1/movies", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+base64.StdEncoding.EncodeToString([]byte("FikoRedhaFebiansyah")))

	resp, _ := router.Test(request, 60)
	body, _ := io.ReadAll(resp.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode, "Get HTTP Status 200")
	assert.Equal(t, fiber.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, fiber.ErrBadRequest.Message, responseBody["message"])
}

func TestUpdateMovieSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategoryDB(db)

	tx, _ := db.Begin()
	movieRepository := repository.NewMovieRepository()
	movie := movieRepository.Save(context.Background(), tx, domain.Movie{Title: "Film Kereen", Description: "Recommended", Rating: 7.7, Image: "http://"})
	tx.Commit()

	router := setupRouter(db)
	requestBody := strings.NewReader(`{
    "title": "Film B Aja",
    "description": "Bikin Bosen",
    "rating": 3.5,
    "image": "www.google.com"
}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:8888/api/v1/movies/"+strconv.Itoa(movie.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+base64.StdEncoding.EncodeToString([]byte("FikoRedhaFebiansyah")))

	resp, _ := router.Test(request, 30)
	body, _ := io.ReadAll(resp.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "Get HTTP Status 200")
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["message"])
	assert.Equal(t, movie.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Film B Aja", responseBody["data"].(map[string]interface{})["title"])
}

func TestUpdateMovieFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategoryDB(db)

	tx, _ := db.Begin()
	movieRepository := repository.NewMovieRepository()
	movie := movieRepository.Save(context.Background(), tx, domain.Movie{Title: "Film Kereen", Description: "Recommended", Rating: 7.7, Image: "http://"})

	tx.Commit()

	router := setupRouter(db)
	requestBody := strings.NewReader(`{
    "title": "",
    "description": "",
    "rating": null,
    "image": ""
}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:8888/api/v1/movies/"+strconv.Itoa(movie.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+base64.StdEncoding.EncodeToString([]byte("FikoRedhaFebiansyah")))

	resp, _ := router.Test(request, 30)
	body, _ := io.ReadAll(resp.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, fiber.ErrBadRequest.Code, resp.StatusCode, "Get HTTP Status 200")
	assert.Equal(t, fiber.ErrBadRequest.Code, int(responseBody["code"].(float64)))
	assert.Equal(t, fiber.ErrBadRequest.Message, responseBody["message"])
}

func TestGetMovieSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategoryDB(db)

	tx, _ := db.Begin()
	movieRepository := repository.NewMovieRepository()
	movie := movieRepository.Save(context.Background(), tx, domain.Movie{Title: "Film Kereen", Description: "Recommended", Rating: 7.7, Image: "http://"})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8888/api/v1/movies/"+strconv.Itoa(movie.Id), nil)
	request.Header.Add("Authorization", "Bearer "+base64.StdEncoding.EncodeToString([]byte("FikoRedhaFebiansyah")))

	resp, _ := router.Test(request, 30)
	body, _ := io.ReadAll(resp.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "Get HTTP Status 200")
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["message"])
	assert.Equal(t, movie.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, movie.Title, responseBody["data"].(map[string]interface{})["title"])
}

func TestGetMovieFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategoryDB(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8888/api/v1/movies/404", nil)
	request.Header.Add("Authorization", "Bearer "+base64.StdEncoding.EncodeToString([]byte("FikoRedhaFebiansyah")))

	resp, _ := router.Test(request, 30)
	body, _ := io.ReadAll(resp.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)
	assert.Equal(t, fiber.ErrNotFound.Code, resp.StatusCode, "Get HTTP Status 200")
	assert.Equal(t, fiber.ErrNotFound.Code, int(responseBody["code"].(float64)))
	assert.Equal(t, fiber.ErrNotFound.Message, responseBody["message"])
}

func TestDeleteMovieSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategoryDB(db)

	tx, _ := db.Begin()
	movieRepository := repository.NewMovieRepository()
	movie := movieRepository.Save(context.Background(), tx, domain.Movie{Title: "Film Kereen", Description: "Recommended", Rating: 7.7, Image: "http://"})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8888/api/v1/movies/"+strconv.Itoa(movie.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+base64.StdEncoding.EncodeToString([]byte("FikoRedhaFebiansyah")))

	resp, _ := router.Test(request, 30)
	body, _ := io.ReadAll(resp.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "Get HTTP Status 200")
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["message"])
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategoryDB(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8888/api/v1/movies/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+base64.StdEncoding.EncodeToString([]byte("FikoRedhaFebiansyah")))

	resp, _ := router.Test(request, 30)
	body, _ := io.ReadAll(resp.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, fiber.ErrNotFound.Code, resp.StatusCode, "Get HTTP Status 200")
	assert.Equal(t, fiber.ErrNotFound.Code, int(responseBody["code"].(float64)))
	assert.Equal(t, fiber.ErrNotFound.Message, responseBody["message"])
}

func TestListCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategoryDB(db)

	tx, _ := db.Begin()
	movieRepository := repository.NewMovieRepository()
	movie1 := movieRepository.Save(context.Background(), tx, domain.Movie{Title: "Film Kereen", Description: "Recommended", Rating: 7.7, Image: "http://"})
	movie2 := movieRepository.Save(context.Background(), tx, domain.Movie{Title: "Film B Aja", Description: "Mending Nonton Jaipong", Rating: 4.3, Image: "http://"})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8888/api/v1/movies", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+base64.StdEncoding.EncodeToString([]byte("FikoRedhaFebiansyah")))

	resp, _ := router.Test(request, 30)
	body, _ := io.ReadAll(resp.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "Get HTTP Status 200")
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["message"])

	var movies = responseBody["data"].([]interface{})
	movies1 := movies[0].(map[string]interface{})
	movies2 := movies[1].(map[string]interface{})

	assert.Equal(t, movie1.Id, int(movies1["id"].(float64)))
	assert.Equal(t, movie1.Title, movies1["title"])
	assert.Equal(t, movie2.Id, int(movies2["id"].(float64)))
	assert.Equal(t, movie2.Title, movies2["title"])
}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	truncateCategoryDB(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8888/api/v1/movies", nil)
	request.Header.Add("Content-Type", "application/json")

	resp, _ := router.Test(request, 30)
	body, _ := io.ReadAll(resp.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, fiber.ErrUnauthorized.Code, resp.StatusCode, "Get HTTP Status 200")
	assert.Equal(t, fiber.ErrUnauthorized.Code, int(responseBody["code"].(float64)))
	assert.Equal(t, fiber.ErrUnauthorized.Message, responseBody["message"])
}
