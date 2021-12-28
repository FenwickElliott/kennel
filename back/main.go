package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Dog struct {
	Collar  string `json:"collar"`
	Alive bool   `json:"alive"`
	Toes  int    `json:"toes"`
}

var (
	allDogs = map[string]Dog{}
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-Forwarded-For"},
	}))
	r.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"})})
	r.Use(gin.Logger())

	r.GET("/dogs", func(c *gin.Context) {
		var dogs []Dog
		for _, dog := range allDogs {
			dogs = append(dogs, dog)
		}
		c.JSON(200, dogs)
	})
	r.GET("/dog/:collar", func(c *gin.Context) {
		if dog, ok := allDogs[c.Param("collar")]; ok {
			c.JSON(200, dog)
		} else {
			c.AbortWithStatusJSON(404, gin.H{"message": fmt.Sprintf("dog '%s' not found", c.Param("collar"))})
		} 
	})
	r.DELETE("/dog/:collar", func(c *gin.Context) {
		if _, ok := allDogs[c.Param("collar")]; ok {
			delete(allDogs, c.Param("collar"))
			c.JSON(200, gin.H{"message": fmt.Sprintf("dog '%s' deleted", c.Param("collar"))})
		} else {
			c.AbortWithStatusJSON(404, gin.H{"message": fmt.Sprintf("dog '%s' not found", c.Param("collar"))})
		} 
	})
	r.POST("/dogs", func(c *gin.Context) {
		var dog Dog
		err := c.Bind(&dog)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(400, gin.H{"message": "failed to deserialize", "error": err.Error()})
			return
		}

		log.Println("new dog:", dog)

		allDogs[dog.Collar] = dog
	})

	port := ":1066"
	log.Println("kennel-backend listening on", port)
	fatal(r.Run(port))
}

func basic() {
	http.HandleFunc("/dogs", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)
		switch r.Method {
		case http.MethodGet:

			var dogs []Dog
			for _, dog := range allDogs {
				dogs = append(dogs, dog)
			}

			body, err := json.Marshal(dogs)
			fatal(err)
	
			// log.Println(string(body))
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Headers", "*")
			w.Write(body)
		case http.MethodPost:
			var dog Dog
			err := json.NewDecoder(r.Body).Decode(&dog)
			if err != nil {
				log.Println(err)
				return
			}

			log.Println("new dog:", dog)

			allDogs[dog.Collar] = dog
		case http.MethodOptions:
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Headers", "*")
		default:
			log.Fatalf("unexpected case: %s, %s", r.Method, r.URL)
		}
	})

	

	port := ":1066"
	log.Println("kennel-backend listening on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func init() {
	log.SetFlags(log.LstdFlags|log.Lshortfile)
	gin.SetMode(gin.ReleaseMode)
}

func fatal(err error) {
	if err != nil {
		_, f, l, _ := runtime.Caller(1)
		log.Fatalf("fatal error throw by: '%s:%d, error: '%s'", f, l, err)
	}
}