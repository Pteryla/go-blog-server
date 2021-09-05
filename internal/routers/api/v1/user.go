package v1

import "github.com/gin-gonic/gin"

type User struct{}

func NewUser() User {
	return User{}
}

func (u User) Get(c *gin.Context)    {}
func (u User) List(c *gin.Context)   {}
func (u User) Create(c *gin.Context) {}
func (u User) Update(c *gin.Context) {}
func (u User) Delete(c *gin.Context) {}
