package main

import (
	goflag "flag"
	"log"

	"github.com/fhs/gompd/mpd"
	flag "github.com/spf13/pflag"
	"gitlab.com/EndowTheGreat/music-fetch/fetch"
)

var host *string = flag.String("host", ":6600", "MPD host to connect to.")

func init() {
	if host != nil {
		flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
		flag.Parse()
	}
}

func main() {
	conn, err := mpd.Dial("tcp", *host)
	if err != nil {
		log.Fatalf("Could not connect to mpd service: %v", err)
	}
	song, err := conn.CurrentSong()
	if err != nil {
		log.Fatal("Could not get current song.")
	}
	status, err := conn.Status()
	if err != nil {
		log.Fatal("Could not get current status.")
	}
	fetch.PrintArt(song, status, "gruvbox")
}
