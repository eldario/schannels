package out

import (
	"context"
	"fmt"
)

// fanOut main struct
type fanOut struct {
	ctx         context.Context
	inChannel   chan string
	OutChannels []chan string
}

// New Constructor of fanOut packet
func New(ctx context.Context, inChannel chan string, outChannels []chan string) *fanOut {
	return &fanOut{
		ctx:         ctx,
		inChannel:   inChannel,
		OutChannels: outChannels,
	}
}

// Run get from one channel and send it to channels list
func (f *fanOut) Run() {
	for {
		select {
		case <-f.ctx.Done():
			fmt.Println("Run: Time to return")
			return
		case value, ok := <-f.inChannel:
			if !ok {
				return
			}
			for _, channel := range f.OutChannels {
				channel <- value
			}
		}

	}
}

// Add new channel in list
func (f *fanOut) Add(channel chan string) {
	f.OutChannels = append(f.OutChannels, channel)
}

// InsertWordInChannel insert a new word in channel
func InsertWordInChannel(words []string, inChannel chan<- string) {
	go func() {
		defer close(inChannel)

		for _, word := range words {
			inChannel <- word
		}
	}()
}
