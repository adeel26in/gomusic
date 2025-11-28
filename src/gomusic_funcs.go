package main

import (
	"log"
	"os"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/flac"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/gopxl/beep/v2/vorbis"
	"github.com/gopxl/beep/v2/wav"
)

func wavplayer(args string) {
	fileopen, err := os.Open(args)
	if err != nil {
		log.Println("Error opening file: ", err)
		return
	}

	defer fileopen.Close()

	streamer, format, err := wav.Decode(fileopen)
	if err != nil {
		log.Println("Error decoding WAV")
		return
	}

	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}

func mp3player(args string) {
	fileopen, err := os.Open(args)
	if err != nil {
		log.Println("Error opening file: ", err)
		return
	}

	defer fileopen.Close()

	streamer, format, err := mp3.Decode(fileopen)
	if err != nil {
		log.Println("Error decoding MP3")
		return
	}

	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}

func flacplayer(args string) {
	fileopen, err := os.Open(args)
	if err != nil {
		log.Println("Error opening file: ", err)
		return
	}

	defer fileopen.Close()

	streamer, format, err := flac.Decode(fileopen)
	if err != nil {
		log.Println("Error decoding FLAC")
		return
	}

	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}

func oggplayer(args string) {
	fileopen, err := os.Open(args)
	if err != nil {
		log.Println("Error opening file: ", err)
		return
	}

	defer fileopen.Close()

	streamer, format, err := vorbis.Decode(fileopen)
	if err != nil {
		log.Println("Error decoding Ogg Vorbis")
		return
	}

	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
