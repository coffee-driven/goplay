package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

type MessageManager struct {
	NewMessage Data
}

func NewMessageManager(message Data) *MessageManager {
	return &MessageManager{
		message,
	}
}

func (mm *MessageManager) manage() {
	message := &DataWithInternalTimestamp{
		mm.NewMessage,
		time.Now(),
	}

	message.save()
}

type Message interface {
	save()
}

type DataWithInternalTimestamp struct {
	Data
	TimestampPrc time.Time
}

type Data struct {
	Metadata     string `json: metadata`
	TimestampMsg time.Time
	Value        string `json: value`
}

func NewData(d Data) *DataWithInternalTimestamp {
	dt := DataWithInternalTimestamp{
		d,
		time.Now(),
	}
	return &dt
}

func (d *Data) getInfo() {
	fmt.Println(d)
}

func (d *DataWithInternalTimestamp) save() {
	fmt.Println("saved")
}

type Server struct {
	ip   net.IP
	port int
}

func NewServer() *Server {
	return &Server{
		ip:   net.ParseIP("127.0.0.1"),
		port: 8080,
	}
}

func (s *Server) start() {
	port := strconv.Itoa(s.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", s.ip, port), nil))
}

func HandleMessage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	receivedData := Data{}
	data := r.Body
	if err := json.NewDecoder(data).Decode(&receivedData); err != nil {
		fmt.Println(err)
	}

	mm := NewMessageManager(receivedData)
	mm.manage()

	// check json key must exist and has value
	// sanitize
}

func main() {
	data := Data{}
	data.getInfo()

	server := NewServer()
	handler := http.HandlerFunc(HandleMessage)

	http.Handle("/", handler)

	server.start()

}
