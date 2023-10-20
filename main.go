package main

import (
	"crypto/rand"
	"os"
	"time"

	"github.com/oklog/ulid/v2"
)

const (
	version = "1.0.1"
)

func main() {
	if len(os.Args) == 2 && (os.Args[1] == "version" || os.Args[1] == "-v" || os.Args[1] == "--version") {
		print(version)
		return
	}

	if len(os.Args) != 1 {
		print("Usage: ulid")

		os.Exit(1)
	}

	id := ulid.MustNew(ulid.Timestamp(time.Now()), rand.Reader)

	print(id.String())
}
