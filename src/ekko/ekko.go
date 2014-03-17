package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

var logger = log.New(os.Stdout, "[ekko] ", log.LstdFlags)

type Config struct {
	Host string
	Port int
}

func NewConfig() *Config {
	c := &Config{}
	c.Load()

	return c
}

func (c *Config) Load() {
	c.Host = os.Getenv("HOST")

	// ignore the obvious error
	c.Port, _ = strconv.Atoi(os.Getenv("PORT"))
}

func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func main() {
	config := NewConfig()

	log.Printf("starting app with addr=%q", config.Addr())

	l, err := net.Listen("tcp", fmt.Sprintf("%s", config.Addr()))
	if err != nil {
		log.Fatalf("error starting listener on addr=%q err=%q", config.Addr(), err)
	}

	log.Printf("listening addr=%q", l.Addr())

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("error acceptiong err=%q", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// cheapest echo server ever
	io.Copy(conn, conn)
}
