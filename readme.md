**⚠️ Work in Progress**  
This project is currently under development. Features may be incomplete or subject to change. 🎅🎁

# Masked Santa

Masked Santa is a simple web application built in Go (Golang) that helps organize a Masked Santa gift exchange without requiring email addresses. Instead, participants are given a unique single-use URL to view their gift assignment. The app also supports exclusions, allowing you to specify pairs of participants who should not exchange gifts (e.g., couples or family members).

---

## Features
- **Room-based Gift Exchange**: Create a unique room for your group.
- **Participant Management**: Add participants by name and specify exclusions (optional).
- **Fisher-Yates Shuffle**: Assign participants fairly and randomly, considering exclusions.
- **Single-Use URLs**: Each participant gets a private link to view their assignment.
- **No Email Required**: Fully anonymous without the need for email addresses.
- **Simple UI**: Minimalistic and straightforward user interface.

---

## Installation

### Prerequisites
- [Go](https://golang.org/dl/) installed (version 1.18 or later)
- Git installed

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/secret-santa.git
   cd secret-santa
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

4. Open your browser and navigate to:
   ```
   http://localhost:8080
   ```

---

## Usage

### 1. Create a Room
- Go to the homepage and click "Create a Room."
- You will be redirected to a unique URL for your room.

### 2. Add Participants
- Enter the name of each participant.
- Optionally, specify exclusions for a participant (comma-separated names of others they should not be assigned to).

### 3. Assign Gifts
- Once all participants are added, click "Assign Gifts."
- The system will use the Fisher-Yates shuffle algorithm to assign gift recipients, respecting any exclusions.

### 4. View Assignments
- Each participant gets a unique URL to view their assigned gift recipient.
- Share the links with participants to let them see their assignments.

---

## Folder Structure

```
secret-santa/
├── main.go                 # Entry point of the application
├── models/
│   └── room.go             # Room and Participant data structures
├── handlers/
│   └── room_handler.go     # Business logic for room and participant management
└── templates/
    ├── index.html          # Homepage template
    ├── room.html           # Room management template
    └── participant.html    # Assignment page template
```

---

## Customization
You can customize the app by editing the templates in the `templates/` directory to fit your design preferences.

---

## Contributing
Contributions are welcome! If you’d like to improve the app, please:
1. Fork the repository.
2. Create a feature branch:
   ```bash
   git checkout -b feature-name
   ```
3. Commit your changes and push:
   ```bash
   git commit -m "Add feature name"
   git push origin feature-name
   ```
4. Create a pull request.

---

## License
This project is licensed under the [MIT License](LICENSE).

---

## Acknowledgments
- [Gin Web Framework](https://github.com/gin-gonic/gin) for building the server.
- Fisher-Yates Shuffle for fair randomization.
- Inspiration from the holiday tradition of Secret Santa! 🎅🎄
