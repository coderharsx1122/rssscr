package main

import (
	"net/http"

	"github.com/coderharsx1122/rssscr/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)
