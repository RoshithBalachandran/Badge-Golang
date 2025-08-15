package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    store := cookie.NewStore([]byte("secret"))
    r.Use(sessions.Sessions("mysession", store))

   type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}	

r.POST("/login", func(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
        return
    }

    // Debug: print received values
    fmt.Println("Received:", req.Username, req.Password)

    if req.Username == "roshith" && req.Password == "123" {
        session := sessions.Default(c)
        session.Set("user", req.Username)
        session.Save()
        c.JSON(http.StatusOK, gin.H{"message": "login success"})
    } else {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
    }
})


    r.Run(":8080")
}
