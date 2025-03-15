package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// wsHandler is a handler for the WebSocket endpoint
// It reads alerts from Kafka and sends them to the client
// It also sends pings to the client to keep the connection alive
func wsHandler(kafkaBrokers, topic string, repo ObjectRepository, username, password *string) gin.HandlerFunc {
	return func(c *gin.Context) {
		groupId := c.Query("groupId")
		if groupId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "groupId not found"})
			return
		}

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.Error(err)
			return
		}
		defer conn.Close()

		// Configure the connection
		// 1. Set the pong handler to reset the read deadline when a pong is received
		setPongHandler(conn)

		// 2. Set initial read deadline - connection will time out if no pong is received within this period
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))

		// 3. Start a goroutine just to detect the connection closing
		// Even if we don't read application messages, we need a read loop to process control frames
		go readLoop(conn)

		// Start the Kafka consumer
		kafkaChannel := make(chan StampProbabilities)
		consumer := NewConsumer(kafkaBrokers, groupId, topic, username, password)
		go consume(kafkaChannel, consumer)

		// Main loop for sending data and pings
		pingTicker := time.NewTicker(30 * time.Second)
		defer pingTicker.Stop()
		mainLoop(conn, kafkaChannel, pingTicker, repo)
	}
}

// parseAlert converts a ZtfAlert to a GeoJSON feature used by d3-celestial
func parseAlert(alert StampProbabilities, object Object) (gin.H, error) {
	return gin.H{
		"type": "Feature",
		"id":   alert.Candid,
		"geometry": gin.H{
			"type":        "Point",
			"coordinates": []float64{object.Meanra, object.Meandec},
		},
		"properties": gin.H{
			"name": alert.Candid,
			"dim":  object.G_r_mean,
			"type": "snr",
		},
	}, nil
}

// readLoop reads control frames from the client
// This is needed to process pongs and other control frames
func readLoop(conn *websocket.Conn) {
	for {
		// ReadMessage is needed to process incoming control frames including pong responses
		// We don't care about the content, we just need to call Read to process control frames
		_, _, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err,
				websocket.CloseGoingAway,
				websocket.CloseNormalClosure,
				websocket.CloseNoStatusReceived) {
				slog.Error("WebSocket error", "error", err)
			}
			return // Connection is closed, exit the goroutine
		}
	}
}

// mainLoop sends data to the client and pings the client
// For each incoming alert, it sends the alert to the client
// Every 30 seconds, it sends a ping to the client
// If the client does not respond to the ping, the connection will be closed
func mainLoop(conn *websocket.Conn, kafkaChannel chan StampProbabilities, pingTicker *time.Ticker, repo ObjectRepository) {
	for {
		select {
		case alert := <-kafkaChannel:
			object, err := getObject(alert.ObjectID, repo)
			if err != nil {
				slog.Error("Error getting object", "error", err)
				continue
			}
			// Send data
			data, err := parseAlert(alert, object)
			if err != nil {
				slog.Error("Error parsing alert", "error", err)
				continue
			}
			if err := conn.WriteJSON(data); err != nil {
				slog.Error("Write error", "error", err)
				return
			}
			slog.Debug("Sent data to client")

		case <-pingTicker.C:
			// Send a ping
			slog.Debug("Sending ping to client")

			// WriteControl is non-blocking, so we set a write deadline
			if err := conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(5*time.Second)); err != nil {
				slog.Error("Ping error", "error", err)
				return
			}
		}
	}
}

// setPongHandler sets the pong handler for a connection
// The pong handler resets the read deadline when a pong is received
func setPongHandler(conn *websocket.Conn) {
	conn.SetPongHandler(func(string) error {
		// When we get a pong response, reset the read deadline
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		slog.Debug("Received pong from client")
		return nil
	})
}

func getObject(oid string, repo ObjectRepository) (Object, error) {
	object, err := repo.GetObject(oid)
	if err != nil {
		return Object{}, err
	}
	return object, nil
}
