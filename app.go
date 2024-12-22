package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) StartupCode() {
	go func() {
		// Specify a workspace directory and bind address
		cmd := exec.Command("code-server",
			"--auth", "none",
			"--bind-addr", "0.0.0.0:8080",
			"--without-connection-token",
			"--disable-window-controls",
		)

		// Capture output
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Start()
		if err != nil {
			println("Error starting code-server:", err.Error())
			return
		}

		// Wait for the process
		err = cmd.Wait()
		if err != nil {
			println("code-server process ended with error:", err.Error())
		}
	}()
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	fmt.Println("App startup!")
	a.StartupCode()
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
