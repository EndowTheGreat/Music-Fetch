package fetch

import (
	"fmt"
	"log"
	"strconv"

	"github.com/fhs/gompd/mpd"
)

func parseTime(length string, status mpd.Attrs) string {
	if status["state"] == "stop" {
		return ""
	}
	converted, err := strconv.Atoi(length)
	if err != nil {
		log.Fatal("Could not parse song length.")
	}
	return fmt.Sprintf("%vm", strconv.Itoa(converted/60))
}
