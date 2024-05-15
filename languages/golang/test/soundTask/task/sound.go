package task

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

func PlaySound(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	var streamer beep.StreamSeekCloser
	var format beep.Format
	extName := filepath.Ext(filename)
	extUpper := strings.ToUpper(extName)
	switch extUpper {
	case ".WAV":
		streamer, format, err = wav.Decode(f)
	case ".MP3":
		streamer, format, err = mp3.Decode(f)

	}

	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
