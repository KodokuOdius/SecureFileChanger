package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/KodokuOdius/SecureFileChanger/db"
	"github.com/gin-gonic/gin"
)

func (h *Handler) userCreate(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Println(err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode("Error parsing form")
		return
	}

	email := c.Request.Form.Get("email")
	if email == "" {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode("specify a user email please")
		return
	}

	password := c.Request.Form.Get("password")
	if password == "" {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode("specify a user pass please")
		return
	}

	user := db.NewUser(email, password, false)
	if user == nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode("User not create")
		return
	}

	j, err := json.Marshal(user)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode("Error occured during marshaling")
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(string(j))
	fmt.Println(user)
}
