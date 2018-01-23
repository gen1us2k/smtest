package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/looplab/fsm"
)

const (
	stateAvailable   = "available"
	stateDownloading = "downloading"
	stateStreaming   = "streaming"
	stateStreamed    = "streamed"
	stateDone        = "done"
)

// VideoInfo stores information about videos
type VideoInfo struct {
	URL      string
	FileName string
}

// Video represents struct for streamer
type Video struct {
	Info VideoInfo
	FSM  *fsm.FSM
}

// Videos is array of pointers to video to read states and manipulate with them
type Videos []*Video

// NewVideo creates a new video
func NewVideo(url, filename string) *Video {

	v := &Video{
		Info: VideoInfo{
			URL:      url,
			FileName: filename,
		},
	}

	v.FSM = fsm.NewFSM(
		stateAvailable,
		fsm.Events{
			{Name: stateDownloading, Src: []string{stateAvailable}, Dst: stateDownloading},
			{Name: stateStreaming, Src: []string{stateDownloading}, Dst: stateStreaming},
			{Name: stateStreamed, Src: []string{stateStreaming}, Dst: stateStreamed},
			{Name: stateDone, Src: []string{stateStreamed}, Dst: stateDone},
		},
		fsm.Callbacks{
			stateDone: func(e *fsm.Event) { fmt.Println("Item is in done state. Removing files") },
		},
	)

	return v
}

func main() {
	door := NewVideo("http://10.5.2.62", "smallVideo")
	d2 := NewVideo("http://10.2.1.2", "old_one")
	d3 := NewVideo("http://10.2.1.1", "done")
	var videos []*Video
	videos = append(videos, door)
	videos = append(videos, d2)
	videos = append(videos, d3)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			fmt.Println(len(videos))
			for i := 0; i < len(videos); i++ {
				video := videos[i]
				if video.FSM.Current() == stateDone {

					videos = append(videos[:i], videos[i+1:]...)
				}
			}
			time.Sleep(time.Second)
		}
		wg.Done()
	}()

	if err := door.FSM.Event(stateDownloading); err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	if err := door.FSM.Event(stateStreaming); err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	if err := door.FSM.Event(stateStreamed); err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	if err := door.FSM.Event(stateDone); err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	if err := d2.FSM.Event(stateDownloading); err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	if err := d2.FSM.Event(stateStreaming); err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	if err := d2.FSM.Event(stateStreamed); err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	if err := d2.FSM.Event(stateDone); err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	if err := d3.FSM.Event(stateDownloading); err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	if err := d3.FSM.Event(stateStreaming); err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	if err := d3.FSM.Event(stateStreamed); err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	if err := d3.FSM.Event(stateDone); err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	time.Sleep(10 * time.Second)
	door = NewVideo("http://10.5.2.62", "smallVideo")
	videos = append(videos, door)
	if err := door.FSM.Event(stateDownloading); err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	if err := door.FSM.Event(stateStreaming); err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	if err := door.FSM.Event(stateStreamed); err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	if err := door.FSM.Event(stateDone); err != nil {
		fmt.Println(err)
	}
	time.Sleep(10 * time.Second)
}
