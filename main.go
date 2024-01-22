package main

import (
	"log"

	"main/api"
	"main/websocket"

	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := bolt.Open("database.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("system_info"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.Static("/static", "./static")

	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/system-info", api.GetSystemInfoHandler)
	}

	r.GET("/ws", func(c *gin.Context) {
		websocket.HandleWebSocket(c.Writer, c.Request, db)
	})

	r.Run(":8081")
}
