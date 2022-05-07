package main

import (
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

func loadSamples() (formats []beep.Format, streamSeekers []beep.StreamSeeker, streamers []beep.Streamer) {

	for i := 0; i < 9; i++ {

		var f *os.File
		//samples
		if i == 0 {
			var err error
			f, err = os.Open("static/samples/organ.wav")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			var err error
			f, err = os.Open("static/samples/casio.wav")
			if err != nil {
				log.Fatal(err)
			}
		}

		streamer, format, err := wav.Decode(f)
		if err != nil {
			log.Fatal(err)
		}

		formats = append(formats, format)

		buffer := beep.NewBuffer(format)
		buffer.Append(streamer)
		streamer.Close()

		streamSeeker := buffer.Streamer(0, buffer.Len())

		streamSeekers = append(streamSeekers, streamSeeker)

		loop := beep.Loop(-1, streamSeeker)

		ctrl := &beep.Ctrl{Streamer: loop, Paused: true}
		streamers = append(streamers, ctrl)
	}

	return formats, streamSeekers, streamers
}

func createKeyMap() map[string]int {
	return map[string]int{
		"A": 0,
		"S": 1,
	}
}

func initSpeaker(formats []beep.Format, streamers []beep.Streamer) {
	//need to check all formats are the same after recording organ // TODO
	speaker.Init(formats[0].SampleRate, formats[0].SampleRate.N(time.Second/20)) //change 20 for latency/cpu usage
	//mixer doesn't close stream after a sample is played
	mixer := beep.Mix(streamers...)
	speaker.Play(mixer)
}

func keyEvents(w fyne.Window, streamSeekers []beep.StreamSeeker, streamers []beep.Streamer) {

	keyMap := createKeyMap()
	//desktop only key events
	if deskCanvas, ok := w.Canvas().(desktop.Canvas); ok {
		deskCanvas.SetOnKeyDown(func(key *fyne.KeyEvent) {

			//avoid race conditions using lock and unlock (at end)
			speaker.Lock()

			i := keyMap[string(key.Name)]
			//create new ctrl streamer that is paused --unable to get to pause variable from loops array
			streamers[i] = &beep.Ctrl{Streamer: beep.Loop(-1, streamSeekers[i]), Paused: false}

			speaker.Unlock()
		})

		deskCanvas.SetOnKeyUp(func(key *fyne.KeyEvent) {

			speaker.Lock()

			i := keyMap[string(key.Name)]
			streamSeekers[0].Seek(i)
			streamers[i] = &beep.Ctrl{Streamer: beep.Loop(-1, streamSeekers[i]), Paused: true}

			speaker.Unlock()
		})
	}
}

func setKeys(w fyne.Window) {

	//load audio
	formats, streamSeekers, streamers := loadSamples()

	//set beep library
	initSpeaker(formats, streamers)

	//listeners for keys
	keyEvents(w, streamSeekers, streamers)

}
