package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func toomanyargs() {
	err := errors.New("Max arguments are only 2")
	fmt.Println(err)
}

func unknowncommand() {
	err := errors.New("Unknown command given")
	fmt.Println(err)
}

func help() {
	fmt.Println("gomusic is a simple command line utility that plays music (ONLY WAV, MP3, FLAC and Ogg Vorbis)! Example: /home/<user>/Downloads/example.mp3")
}

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}

	if len(os.Args) > 2 {
		toomanyargs()
	}

	args := os.Args[1]

	if strings.HasSuffix(args, ".wav") {
		wavplayer(args)
	}

	if strings.HasSuffix(args, ".mp3") {
		mp3player(args)
	}

	if strings.HasSuffix(args, ".flac") {
		flacplayer(args)
	}

	if strings.HasSuffix(args, ".ogg") {
		oggplayer(args)
	}

	switch os.Args[1] {
	case "help", "--help", "-h":
		help()
	default:
		unknowncommand()
	}
}
