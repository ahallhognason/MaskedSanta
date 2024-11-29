package models

import (
	"crypto/rand"
	"encoding/hex"
)

type Participant struct {
	Name       string
	Exclusions []string
	URLID      string
	AssignedTo string
}

type Room struct {
	ID           string
	Participants []Participant
	Assigned     bool
}

func GenerateID() string {
	bytes := make([]byte, 8)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
