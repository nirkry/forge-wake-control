package main

import (
	"context"
	"database/sql"
	"fmt"
	"forge-wake-control/db"
)

// App struct
type App struct {
	ctx context.Context
	db  *sql.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	conn, err := db.Connect()
	if err != nil {
		fmt.Println("DB init error:", err)
		return
	}

	// init table
	_, err = conn.Exec(`CREATE TABLE IF NOT EXISTS devices (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	)`)

	if err != nil {
		fmt.Println("DB table creation error:", err)
		return
	}

	// sample record insertion
	_, err = conn.Exec(`INSERT INTO devices (name) VALUES (?)`, "Sample Device")
	if err != nil {
		fmt.Println("DB record insertion error:", err)
		return
	}

	a.db = conn
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {
	if a.db != nil {
		a.db.Close()
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
