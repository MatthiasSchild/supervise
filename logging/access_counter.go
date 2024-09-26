package logging

import (
	"github.com/gin-gonic/gin"
)

const (
	accessCounterKey = "access-counter"
)

type AccessCounter map[string]int

func (a *AccessCounter) Increase(key string) {
	(*a)[key]++
}

func AccessCounterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(accessCounterKey, &AccessCounter{})
		c.Next()
	}
}

func GetAccessCounter(c *gin.Context) *AccessCounter {
	if accessCounter, ok := c.Get(accessCounterKey); ok {
		return accessCounter.(*AccessCounter)
	}
	return &AccessCounter{}
}
