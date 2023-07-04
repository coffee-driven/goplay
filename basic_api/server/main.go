package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	p := Person{
		Name: name,
		Age:  age,
	}
	return &p
}

func (p *Person) handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	if resp, err := json.Marshal(p); err != nil {
		fmt.Println("Error marshalling struct")
	} else {
		w.Write(resp)
	}
}

func (p *Person) getInfo() {
	fmt.Println(p)
}

func SetName(p Person, name string) Person {
	p.Name = name
	return p
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

func main() {
	my_person := NewPerson("Marek", 30)
	my_person.getInfo()

	server := NewServer()
	handler := http.HandlerFunc(my_person.handleRequest)

	http.Handle("/", handler)

	server.start()

}
