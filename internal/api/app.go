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
	u.POST("/", a.AddUser)
	u.POST("/comments", a.AddComments)
	u.POST("/likes", a.AddLikes)

	e.Logger.Fatal(e.Start(port))
}

func (a *Application) AddUser(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := a.repo.AddUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		log.Println(err)
		log.Println("1")
	}

	return c.NoContent(http.StatusCreated)
}

func (a *Application) AddLikes(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := a.repo.AddLikes(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}

func (a *Application) AddComments(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := a.repo.AddComments(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}

func (a *Application) GetUserInfo(c echo.Context) error {

	usersResponse, err := a.repo.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusAccepted, usersResponse)
}
