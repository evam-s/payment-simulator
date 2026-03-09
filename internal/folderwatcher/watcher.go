package fw

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"payment-simulator/internal/kafka"
	"time"

	"github.com/fsnotify/fsnotify"
)

func InitializeFolderWatcher() {
	watchFolder("bucket/banks", "banks")
	watchFolder("bucket/accounts", "accounts")
	log.Println("All FW initialized.")
}

func watchFolder(path, topicName string) {
	cwd, _ := os.Getwd()
	path = filepath.Join(cwd, path)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}

	var debounce *time.Timer

	log.Println("Watching", path, "for changes...")
	go func() {
		defer watcher.Close()
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return // channel closed
				}
				log.Println("Event:", event)

				if event.Op.String() != "REMOVE" {
					if debounce != nil {
						debounce.Stop()
					}

					debounce = time.AfterFunc(time.Second, func() {
						log.Println("Event.name:", event.Name)
						readFile(event.Name, topicName)
					})
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Watcher error:", err)
			}
		}
	}()
}

func readFile(fileName, topic string) {
	if fileData, err := os.ReadFile(fileName); err != nil {
		log.Println("There was some error in reading file", fileName, ", Error:", err)
	} else {
		log.Println("filedata:", string(fileData))

		var raw []map[string]any
		json.Unmarshal(fileData, &raw)

		goLimiter := make(chan int, 10)
		for _, record := range raw {
			val, _ := json.Marshal(record)
			goLimiter <- 1
			go func(v []byte) {
				kafka.PublishToTopic(v, topic)
				<-goLimiter
			}(val)
		}
	}
}
