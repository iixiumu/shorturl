package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

type Input struct {
	URL string `form:"url" binding:"required"`
}

type Output struct {
	URL string `form:"url" binding:"required"`
}

// search
func s2l(c *gin.Context) {
	var input Input
	var err error
	err = c.ShouldBind(&input)
	if err != nil {
		c.String(200, err.Error())
		return
	}

	shortUrl := input.URL
	var longUrl string

	redis := &RedisIncr{}
	longUrl, err = redis.Short2Long(shortUrl)
	if err != nil {
		myURL := &MyURL{}
		longUrl, err = myURL.Short2Long(shortUrl)
		if err != nil {
			id := redis.GenID()
			longUrl = trans(id)
			myURL.Save(shortUrl, longUrl)
		}
		redis.SetShort2Long(shortUrl, longUrl)
		redis.SetLong2Short(longUrl, shortUrl)
	} else {
		redis.Expire(longUrl)
	}

	output := &Output{
		URL: shortUrl,
	}
	c.JSON(200, output)
}

// search, if not exist insert
func l2s(c *gin.Context) {
	var input Input
	var err error
	err = c.ShouldBind(&input)
	if err != nil {
		c.String(200, err.Error())
		return
	}

	longUrl := input.URL
	var shortUrl string

	redis := &RedisIncr{}
	shortUrl, err = redis.Long2Short(longUrl)
	if err != nil {
		myURL := &MyURL{}
		shortUrl, err = myURL.Long2Short(longUrl)
		if err != nil {
			id := redis.GenID()
			shortUrl = trans(id)
			myURL.Save(shortUrl, longUrl)
		}
		redis.SetShort2Long(shortUrl, longUrl)
		redis.SetLong2Short(longUrl, shortUrl)
	} else {
		redis.Expire(shortUrl)
	}

	output := &Output{
		URL: shortUrl,
	}
	c.JSON(200, output)
}

func main() {
	go func() {
		log.Println(http.ListenAndServe(":8053", nil))
	}()

	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/s2l", s2l)
	r.POST("/l2s", l2s)

	r.Run(":8052")
}
