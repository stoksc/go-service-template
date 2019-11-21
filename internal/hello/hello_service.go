package hello

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

// Greeting describes the greeting request object
type Greeting struct {
	Name     string `json:"name" binding:"required" db:"name"`
	Greeting string `json:"greeting" binding:"required" db:"greeting"`
}

type HelloService struct {
	DB     *sql.DB
	Logger *zap.Logger
}

// GetBaseGreetingHandler handlers getting a base greeting
func (s *HelloService) GetBaseGreetingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"greeting": "fuck you",
	})
}

// GetGreetingHandler handlers getting a greeting
func (s *HelloService) GetGreetingHandler(c *gin.Context) {
	var greeting Greeting

	err := s.DB.QueryRow(`
		select name, greeting
		from greeting
		where name = $1
	`, c.Param("name")).Scan(&greeting.Name, &greeting.Greeting)

	if err == sql.ErrNoRows {
		c.Status(http.StatusNotFound)
		return
	}

	if err != nil {
		s.Logger.Error("Failed to perform query", zap.String("error", err.Error()))
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"greeting": greeting.Greeting + ", " + greeting.Name,
	})
}

// CreateGreetingHandler handles creating a greeting
func (s *HelloService) CreateGreetingHandler(c *gin.Context) {
	var json Greeting
	if err := c.ShouldBindJSON(&json); err != nil {
		s.Logger.Error("Failed to parse POST body", zap.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stmt, err := s.DB.Prepare("insert into greeting(name, greeting) values($1, $2)")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	_, err = stmt.Exec(json.Name, json.Greeting)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"greeting": json.Greeting + ", " + json.Name,
	})
}
