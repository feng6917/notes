package sound

import (
	"bytes"
	"io"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/sirupsen/logrus"
)

type SoundType int

const (
	SOUND_TYPE_MP3 SoundType = iota
	SOUND_TYPE_WAV
)

func PlaySound(st SoundType, data []byte) {
	var streamer beep.StreamSeekCloser
	var format beep.Format
	var err error
	switch st {
	case SOUND_TYPE_WAV:
		streamer, format, err = wav.Decode(bytes.NewBuffer(data))
	case SOUND_TYPE_MP3:
		streamer, format, err = mp3.Decode(io.NopCloser(bytes.NewBuffer(data)))
	}

	if err != nil {
		logrus.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
