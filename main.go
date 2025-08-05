package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/simoncdn/gator/internal/command"
	"github.com/simoncdn/gator/internal/config"
	"github.com/simoncdn/gator/internal/database"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("Can't connect to the database: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &command.State{
		Cfg: &cfg,
		DB:  dbQueries,
	}

	commands := command.Commands{
		CommandsList: map[string]func(*command.State, command.Command) error{},
	}
	commands.Register("login", command.HandlerLogin)
	commands.Register("register", command.HandlerRegister)
	commands.Register("reset", command.HandlerReset)
	commands.Register("users", command.HandlerUsers)
	commands.Register("agg", command.HandlerAgg)
	commands.Register("addfeed", command.HandlerAddFeed)

	if len(os.Args) < 2 {
		fmt.Printf("argumenst error\n")
		os.Exit(1)
	}

	cmd := command.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = commands.Run(programState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
