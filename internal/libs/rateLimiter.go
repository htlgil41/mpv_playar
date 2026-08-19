package libs

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Client struct {
	Limiter *rate.Limiter
}

var (
	Mu      sync.Mutex
	Clients = make(map[string]*Client)
)

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		Mu.Lock()
		if _, exists := Clients[ip]; !exists {
			Clients[ip] = &Client{Limiter: rate.NewLimiter(10, 5)}
		}
		cl := Clients[ip]
		Mu.Unlock()

		if !cl.Limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})
			return
		}
		c.Next()
	}
}
