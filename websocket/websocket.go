package websocket

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// We set our Read and Write buffer sizes
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

// The Upgrade function will take in an incoming request and upgrade the request
// into a websocket connection
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
    // this line allows other origin hosts to connect to our
    // websocket server
    upgrader.CheckOrigin = func(r *http.Request) bool { return true }

    // creates our websocket connection
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return ws, err
    }
    // returns our new websocket connection
    return ws, nil
}

func Writer(conn *websocket.Conn) {
    // we want to kick off a for loop that runs for the
    // duration of our websockets connection
    for {
        // we create a new ticker that ticks every 5 seconds
        ticker := time.NewTicker(5 * time.Second)

        // every time our ticker ticks
        for t := range ticker.C {
            // print out that we are updating the stats
            fmt.Printf("Updating Stats: %+v\n", t)

			output := fmt.Sprintf(t.Format("2006-01-02 15:04:05"))


            // and finally we write this JSON string to our WebSocket
            // connection and record any errors if there has been any
            if err := conn.WriteMessage(websocket.TextMessage, []byte(output)); err != nil {
                fmt.Println(err)
                return
            }
        }
    }
}