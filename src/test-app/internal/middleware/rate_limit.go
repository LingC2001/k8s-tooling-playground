package middleware

import (
	"net"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
	visitors = make(map[string]*rate.Limiter)
	mu       sync.RWMutex
)

func getVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := visitors[ip]
	if !exists {
		limiter = rate.NewLimiter(10, 20) // 10 req/sec, burst of 20
		visitors[ip] = limiter
	}

	return limiter
}

func IPRateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)
		if err != nil {
			ip = c.Request.RemoteAddr
		}

		limiter := getVisitor(ip)
		if !limiter.Allow() {
			c.JSON(429, gin.H{
				"error": "rate limit exceeded",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
