package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/armon/go-socks5"
)

func main() {
	// Load credentials from file
	file, err := os.Open("users.conf")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	// Create a map to store credentials
	credentials := make(socks5.StaticCredentials)

	// Read in the credentials file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		if len(parts) != 2 {
			log.Fatalf("Invalid auth-pair in users.conf: %s", scanner.Text())
		}
		credentials[parts[0]] = parts[1]
	}
	// Check for errors reading the file
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// make sure we have at least one credential
	if len(credentials) == 0 {
		log.Fatal("No credentials provided")
	}

	// Create a SOCKS5 server
	authenticator := socks5.UserPassAuthenticator{Credentials: credentials}
	conf := &socks5.Config{
		AuthMethods: []socks5.Authenticator{authenticator},
	}
	server, err := socks5.New(conf)
	if err != nil {
		log.Fatal(err)
	}

	// Start listening and Serving
	log.Println("Socks5-Server running on :1080 (inside container!)")
	if err := server.ListenAndServe("tcp", ":1080"); err != nil {
		log.Fatal(err)
	}
}
