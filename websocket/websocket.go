package websocket

import (
	"log"
	"net/http"
	"runtime"
	"time"

	"main/database"
	"main/models"
	"main/utils"

	"github.com/boltdb/bolt"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// StartWebSocketServer starts the WebSocket server.
func StartWebSocketServer(db *bolt.DB) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		HandleWebSocket(w, r, db)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request, db *bolt.DB) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		data := GetSystemInfo()
		data.InsertionDate = time.Now().Format("2006-01-02 15:04:05")
		database.SaveSystemInfo(db, data)

		err := conn.WriteJSON(data)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("WebSocket connection error: %v", err)
			}
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func GetSystemInfo() models.SystemInfo {
	cpuUsage, _ := utils.GetCPUUsage()
	memInfo, _ := utils.GetMemoryInfo()
	diskInfo, _ := utils.GetDiskInfo()

	hostname, err := utils.GetHostname()
	if err != nil {
		// Handle the error, for example, log it and use a default value
		log.Printf("Error getting hostname: %v", err)
		hostname = "unknown"
	}

	return models.SystemInfo{
		Timestamp:     time.Now().Format(time.RFC3339),
		CPU:           cpuUsage,
		Memory:        memInfo.Used,
		OS:            utils.GetOS(),
		Hostname:      hostname,
		NumCPU:        runtime.NumCPU(),
		TotalMemory:   memInfo.Total,
		FreeMemory:    memInfo.Free,
		TotalStorage:  diskInfo.Total,
		FreeStorage:   diskInfo.Free,
		StorageDevice: diskInfo.Path, // Use the Path field as the device name
	}
}


