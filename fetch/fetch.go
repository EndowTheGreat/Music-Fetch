package fetch

import (
	"fmt"

	"github.com/fhs/gompd/mpd"
)

func PrintArt(song mpd.Attrs, status mpd.Attrs, theme string) {
	length := parseTime(song["Time"], status)
	playingStatus := ""
	switch status["state"] {
	case "play":
		playingStatus = "Playing"
	case "stop":
		playingStatus = "None"
	case "pause":
		playingStatus = "Paused"
	}
	switch theme { // I'll add more at some point, or you can feel free to submit a PR.
	case "gruvbox":
		fmt.Printf("\t\033[38;2;69;133;136m      ;\n\t\033[38;2;69;133;136m      ;;\n\t\033[38;2;69;133;136m      ;';.\n\t\033[38;2;69;133;136m      ;  ;;       Title: \033[38;2;235;219;178m %v\n\t\033[38;2;69;133;136m      ;   ;;      Artist: \033[38;2;235;219;178m%v\n\t\033[38;2;69;133;136m      ;    ;;     Length: \033[38;2;235;219;178m%v\n\t\033[38;2;69;133;136m      ;    ;;     Status: \033[38;2;235;219;178m%v\n\t\033[38;2;69;133;136m      ;   ;'\n\t\033[38;2;69;133;136m      ;  '\n\t\033[38;2;69;133;136m ,;;;,;\n\t\033[38;2;69;133;136m ;;;;;;\n\t\033[38;2;69;133;136m ';;;;'\n\n", song["Title"], song["Artist"], length, playingStatus)
	case "nord":
		fmt.Printf("\t\033[38;2;129;161;193m      ;\n\t\033[38;2;129;161;193m      ;;\n\t\033[38;2;129;161;193m      ;';.\n\t\033[38;2;129;161;193m      ;  ;;       Title: \033[38;2;229;233;240m %v\n\t\033[38;2;129;161;193m      ;   ;;      Artist: \033[38;2;229;233;240m%v\n\t\033[38;2;129;161;193m      ;    ;;     Length: \033[38;2;229;233;240m%v\n\t\033[38;2;129;161;193m      ;    ;;     Status: \033[38;2;229;233;240m%v\n\t\033[38;2;129;161;193m      ;   ;'\n\t\033[38;2;129;161;193m      ;  '\n\t\033[38;2;129;161;193m ,;;;,;\n\t\033[38;2;129;161;193m ;;;;;;\n\t\033[38;2;129;161;193m ';;;;'\n\n", song["Title"], song["Artist"], length, playingStatus)
	}
}
