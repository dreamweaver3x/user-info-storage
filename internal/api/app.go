package api

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"someproject/internal/models"
	"someproject/internal/repository"
)

type Application struct {
	repo *repository.UsersRepositoryPostgres
}

func NewApplication(repo *repository.UsersRepositoryPostgres) *Application {
	return &Application{repo: repo}
}

func (a *Application) Start(port string) {

	e := echo.New()

	u := e.Group("/users")
	u.GET("/", a.GetUserInfo)
	u.POST("/", a.AddUsersInfo)

	e.Logger.Fatal(e.Start(port))
}

func (a *Application) AddUsersInfo(c echo.Context) error {
	var users *[]models.User

	if err := c.Bind(users); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := a.repo.AddUsers(users)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}

func (a *Application) GetUserInfo(c echo.Context) error {
	var users *[]models.User

	if err := c.Bind(users); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	usersResponse, err := a.repo.GetUsers(users)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusAccepted, usersResponse)
}
