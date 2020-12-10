package buffered

import (
	"context"
	"fmt"
)

// bufferedChan structure for bufferedChan packet
type bufferedChan struct {
	ctx        context.Context
	inChannel  <-chan string
	bufferSize int
	OutChannel chan string
}

// New Constructor
func New(ctx context.Context, inChannel <-chan string, bufferSize int) *bufferedChan {
	return &bufferedChan{
		ctx:        ctx,
		inChannel:  inChannel,
		bufferSize: bufferSize,
		OutChannel: make(chan string, bufferSize),
	}
}

// Run put values into the channel that fit
func (b *bufferedChan) Run() {
	defer close(b.OutChannel)
	for value := range b.inChannel {
		select {
		case <-b.ctx.Done():
			fmt.Println("Run: Time to return")
			return
		case b.OutChannel <- value:
		default:
			return
		}
	}

}
