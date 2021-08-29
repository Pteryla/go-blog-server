package v1

import "github.com/gin-gonic/gin"

type Collection struct{}

func NewCollection() Collection{
	return Collection{}
}

func(u Collection) Get(c *gin.Context){}
func(u Collection) List(c *gin.Context){}
func(u Collection) Create(c *gin.Context){}
func(u Collection) Update(c *gin.Context){}
func(u Collection) Delete(c *gin.Context){}