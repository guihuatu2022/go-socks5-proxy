package main

import (
	"testing"

	"github.com/armon/go-socks5"
)

func TestAddUser(t *testing.T) {
	credentials := make(socks5.StaticCredentials)

	// Test adding a new user
	addUser(credentials, "testuser", "testpass")
	if credentials["testuser"] != "testpass" {
		t.Errorf("Expected testuser password to be testpass, got %s", credentials["testuser"])
	}

	// Test adding same user with same password (should not duplicate)
	addUser(credentials, "testuser", "testpass")
	if len(credentials) != 1 {
		t.Errorf("Expected 1 credential, got %d", len(credentials))
	}

	// Test adding same user with different password (should create new entry)
	addUser(credentials, "testuser", "newpass")
	if len(credentials) != 2 {
		t.Errorf("Expected 2 credentials, got %d", len(credentials))
	}
}

func TestVersion(t *testing.T) {
	if Version == "" {
		t.Error("Version should not be empty")
	}
}
