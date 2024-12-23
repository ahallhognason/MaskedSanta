package handlers

import (
	"net/http"
	"github.com/ahallhognason/MaskedSanta/models"
	"sync"
	"math/rand"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	rooms      = make(map[string]*models.Room)
	roomsMutex sync.Mutex
)

func ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateRoom(c *gin.Context) {
	roomsMutex.Lock()
	defer roomsMutex.Unlock()

	room := &models.Room{
		ID: models.GenerateID(),
	}
	rooms[room.ID] = room
	c.Redirect(http.StatusFound, "/room/"+room.ID)
}

func ShowRoom(c *gin.Context) {
	roomID := c.Param("roomID")

	roomsMutex.Lock()
	room, exists := rooms[roomID]
	roomsMutex.Unlock()

	if !exists {
		c.String(http.StatusNotFound, "Room not found")
		return
	}

	c.HTML(http.StatusOK, "room_page.html", gin.H{
		"roomID":       room.ID,
		"participants": room.Participants,
		"assigned":     room.Assigned,
	})
}

func RemoveParticipant(c *gin.Context) {
	roomID := c.Param("roomID")
	urlID := c.Param("urlID")

	roomsMutex.Lock()
	room, exists := rooms[roomID]
	roomsMutex.Unlock()

	if !exists {
		c.String(http.StatusNotFound, "Room not found")
		return
	}

	found := false
	newParticipants := make([]models.Participant, 0, len(room.Participants))
	var removedParticipantName string
	for _, participant := range room.Participants {
		if participant.URLID == urlID {
			found = true
			removedParticipantName = participant.Name
			continue
		}
		newParticipants = append(newParticipants, participant)
	}
	for i, participant := range newParticipants {
		filteredExclusions := make([]string, 0, len(participant.Exclusions))
		for _, exclusion := range participant.Exclusions {
			if exclusion != removedParticipantName {
				filteredExclusions = append(filteredExclusions, exclusion)
			}
		}

		participant.Exclusions = filteredExclusions
		newParticipants[i] = participant
	}

	if !found {
		c.String(http.StatusNotFound, "Participant not found")
		return
	}

	roomsMutex.Lock()
	room.Participants = newParticipants
	rooms[roomID] = room
	roomsMutex.Unlock()

	c.HTML(http.StatusOK, "room.html", gin.H{
		"roomID":       room.ID,
		"participants": room.Participants,
		"assigned":     room.Assigned,
	})
}

func AddParticipant(c *gin.Context) {
	roomID := c.Param("roomID")

	roomsMutex.Lock()
	room, exists := rooms[roomID]
	roomsMutex.Unlock()

	if !exists {
		c.String(http.StatusNotFound, "Room not found")
		return
	}

	name := c.PostForm("name")
	exclusions := c.PostFormArray("exclusions[]")

	urlID := models.GenerateID()
	participant := models.Participant{Name: name, Exclusions: exclusions, URLID: urlID}
	log.Println(participant, exclusions)

	roomsMutex.Lock()
	room.Participants = append(room.Participants, participant)
	roomsMutex.Unlock()

	// c.Redirect(http.StatusFound, "/room/"+roomID)
	c.HTML(http.StatusOK, "room.html", gin.H{
		"roomID":       room.ID,
		"participants": room.Participants,
		"assigned":     room.Assigned,
	})
}

func AssignGifts(c *gin.Context) {
	roomID := c.Param("roomID")

	roomsMutex.Lock()
	room, exists := rooms[roomID]
	roomsMutex.Unlock()

	if !exists || room.Assigned {
		c.String(http.StatusNotFound, "Room not found or already assigned")
		return
	}

	participants := room.Participants
	numParticipants := len(participants)
	if numParticipants < 2 {
		c.String(http.StatusBadRequest, "Not enough participants to assign gifts")
		return
	}

	// Fisher-Yates Shuffle with Exclusions
	shuffled := make([]models.Participant, numParticipants)
	copy(shuffled, participants)

	for {
		// Perform the shuffle
		for i := numParticipants - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
		}

		// Check for exclusions
		valid := true
		for i, giver := range participants {
			if contains(giver.Exclusions, shuffled[i].Name) || giver.Name == shuffled[i].Name {
				valid = false
				break
			}
		}

		if valid {
			break // The shuffle is valid
		}
	}

	// Assign gifts based on the shuffled list
	assignments := make(map[string]string)
	for i, giver := range participants {
		assignments[giver.Name] = shuffled[i].Name
	}

	// Update participants with their assignments
	roomsMutex.Lock()
	for i, participant := range participants {
		participant.AssignedTo = assignments[participant.Name]
		room.Participants[i] = participant
	}
	room.Assigned = true
	roomsMutex.Unlock()

	// c.Redirect(http.StatusFound, "/room/"+roomID)
	c.HTML(http.StatusOK, "room.html", gin.H{
		"roomID":       room.ID,
		"participants": room.Participants,
		"assigned":     room.Assigned,
	})
}

// Utility function to check if a string is in a list
func contains(list []string, str string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}
	return false
}

func ViewAssignment(c *gin.Context) {
	urlID := c.Param("urlID")

	roomsMutex.Lock()
	defer roomsMutex.Unlock()

	for _, room := range rooms {
		for _, p := range room.Participants {
			if p.URLID == urlID {
				c.HTML(http.StatusOK, "participant.html", gin.H{
					"name":       p.Name,
					"assignedTo": p.AssignedTo,
				})
				return
			}
		}
	}

	c.String(http.StatusNotFound, "Assignment not found")
}
