package main

import (
	"log"

	"github.com/AungKyawPhyo1142/go-cli-app/config"
	"github.com/AungKyawPhyo1142/go-cli-app/migrations"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	config.ConnectDB()
	migrations.Migrate()

	m := NewModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatalf("Error running the application: %v", err)
	}
}
