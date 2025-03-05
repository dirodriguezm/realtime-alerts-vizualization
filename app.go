package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var globalCount = 0

func start(kafkaBrokers string) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", func(c *gin.Context) {
		err := renderHome(c)
		if err != nil {
			c.Error(err)
			c.JSON(500, gin.H{
				"message": "Template rendering failed",
			})
		}
	})
	r.Static("/static", "./static")
	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.Error(err)
			return
		}
		defer conn.Close()

		// Configure the connection
		// 1. Set the pong handler to reset the read deadline when a pong is received
		conn.SetPongHandler(func(string) error {
			// When we get a pong response, reset the read deadline
			conn.SetReadDeadline(time.Now().Add(60 * time.Second))
			log.Println("Received pong from client")
			return nil
		})

		// 2. Set initial read deadline - connection will time out if no pong is received within this period
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))

		// 3. Start a goroutine just to detect the connection closing
		// Even if we don't read application messages, we need a read loop to process control frames
		go func() {
			for {
				// ReadMessage is needed to process incoming control frames including pong responses
				// We don't care about the content, we just need to call Read to process control frames
				_, _, err := conn.ReadMessage()
				if err != nil {
					if websocket.IsUnexpectedCloseError(err,
						websocket.CloseGoingAway,
						websocket.CloseNormalClosure,
						websocket.CloseNoStatusReceived) {
						log.Printf("WebSocket error: %v", err)
					}
					return // Connection is closed, exit the goroutine
				}
			}
		}()

		// Main loop for sending data and pings
		kafkaChannel := make(chan ZtfAlert)
		pingTicker := time.NewTicker(30 * time.Second)
		defer pingTicker.Stop()

		go consume(kafkaChannel, kafkaBrokers)
		for {
			select {
			case alert := <-kafkaChannel:
				// Send data
				if err := conn.WriteJSON(parseData(alert)); err != nil {
					log.Printf("Write error: %v", err)
					return
				}
				log.Println("Sent data successfully")

			case <-pingTicker.C:
				// Send a ping
				log.Println("Sending ping to client")

				// WriteControl is non-blocking, so we set a write deadline
				if err := conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(5*time.Second)); err != nil {
					log.Printf("Ping error: %v", err)
					return
				}
			}
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func renderHome(c *gin.Context) error {
	component := home()
	c.Status(200)
	return component.Render(c.Request.Context(), c.Writer)
}

func randomObjects(n int) []gin.H {
	features := make([]gin.H, n)
	for i := 0; i < n; i++ {
		features[i] = gin.H{
			"type": "Feature",
			"id":   fmt.Sprintf("%d", globalCount),
			"geometry": gin.H{
				"type":        "Point",
				"coordinates": []float64{rand.Float64()*360 - 180, rand.Float64()*180 - 90},
			},
			"properties": gin.H{
				"name": fmt.Sprintf("Object %d", globalCount),
				"dim":  rand.Float64() * 100,
				"type": "snr",
			},
		}
		globalCount++
	}

	return features
}

func parseData(alert ZtfAlert) gin.H {
	return gin.H{
		"type": "Feature",
		"id":   alert.Candid,
		"geometry": gin.H{
			"type":        "Point",
			"coordinates": []float64{alert.Candidate.Ra, alert.Candidate.Dec},
		},
		"properties": gin.H{
			"name": alert.Candid,
			"dim":  alert.Candidate.Magpsf,
			"type": "snr",
		},
	}
}
