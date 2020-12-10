package in

import (
	"context"
	"fmt"
)

// fanIn main struct
type fanIn struct {
	ctx          context.Context
	channelsList []<-chan string
	outChannel   chan string
}

// New Constructor of fanIn packet
func New(ctx context.Context, channelsList []<-chan string, outChannel chan string) *fanIn {
	return &fanIn{
		ctx:          ctx,
		channelsList: channelsList,
		outChannel:   outChannel,
	}
}

// Run get all list of channels, collect data and send it out
func (f *fanIn) Run() {
	for {
		for _, channel := range f.channelsList {
			select {
			case <-f.ctx.Done():
				fmt.Println("Run: Time to return")
				return
			case value, ok := <-channel:
				if !ok {
					return
				}
				f.outChannel <- value
			}
		}
	}
}

// Add new channel in list
func (f *fanIn) Add(channel <-chan string) {
	f.channelsList = append(f.channelsList, channel)
}

// GenerateChannel put words in generated channel and return it
func GenerateChannel(words []string) <-chan string {
	outChannel := make(chan string)

	go func() {
		defer close(outChannel)

		for _, word := range words {
			outChannel <- word
		}
	}()

	return outChannel
}
