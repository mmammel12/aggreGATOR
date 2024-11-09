package main

import (
	"github.com/mmammel12/aggreGATOR/internal/config"
	"github.com/mmammel12/aggreGATOR/internal/database"
)

// State -
type state struct {
	db  *database.Queries
	cfg *config.Config
}
