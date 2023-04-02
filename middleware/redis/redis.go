package redis

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Lyianu/wen/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

var conn redis.Conn
var mu sync.Mutex

type RedisResponseWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w RedisResponseWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func init() {
	if setting.RedisHost == "" {
		return
	}
	var err error
	conn, err = redis.Dial("tcp", setting.RedisHost)
	if err != nil {
		log.Fatalf("Failed to connect to redis: %s", err)
	}
	if setting.RunMode == "debug" {
		fmt.Printf("Connected to redis: %s", setting.RedisHost)
	}
}

func Redis() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodGet {
			c.Next()
			return
		}

		mu.Lock()
		if setting.RunMode == "debug" {
			fmt.Printf("Redis: GET %s ", c.Request.URL.Path)
		}
		reply, err := redis.String(conn.Do("GET", c.Request.URL.Path))
		if setting.RunMode == "debug" {
			fmt.Printf("=> %s\n", reply)
		}
		mu.Unlock()
		if err == nil {
			c.String(200, reply)
			c.Abort()
			return
		}
		c.Writer = &RedisResponseWriter{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}

		c.Next()

		// TODO: use singleflight to avoid cache stampede
		mu.Lock()
		conn.Do("SETEX", c.Request.URL.Path, 30, c.Writer.(*RedisResponseWriter).Body.String())
		mu.Unlock()
	}
}
