package cmd

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "--- ", 0)
