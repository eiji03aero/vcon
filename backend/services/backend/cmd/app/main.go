package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	Id string
}

func getDBUrl () string {
	return fmt.Sprintf(
		"postgres://%s:%s/%s?user=%s&password=%s&sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
	)
}

func main() {
	db, err := sqlx.Connect("postgres", getDBUrl())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := gin.Default()

	r.GET("/test", func (c *gin.Context) {
		users := []User{}
		db.Select(&users, "SELECT * FROM users")
		c.JSON(200, gin.H{
			"users": users,
		})
	})

	r.POST("/test", func (c *gin.Context) {
		type Body struct {
			Name string `json:"name"`
		}
		body := Body{}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var id int
		rows, err := db.NamedQuery(
			`INSERT INTO users (name) VALUES (:name) RETURNING id`,
			map[string]interface{}{
				"name": body.Name,
			},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if rows.Next() {
			rows.Scan(&id)
		}

		c.JSON(http.StatusOK, gin.H{"id": id})
	})

	r.Run(":3100")
}
