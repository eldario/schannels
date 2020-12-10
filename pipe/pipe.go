package pipe

import (
	"context"
	"fmt"
)

// msgFunc type of format function
type msgFunc func(str string) string

// pipe main struct
type pipe struct {
	ctx        context.Context
	inChannel  chan string
	outChannel chan string
	format     msgFunc
}

// New pipe constructor
func New(ctx context.Context, inChannel chan string, outChannel chan string, format msgFunc) *pipe {
	return &pipe{
		ctx:        ctx,
		inChannel:  inChannel,
		outChannel: outChannel,
		format:     format,
	}
}

// FillValues insert given list of words in channel
func (p *pipe) FillValues(words []string) {
	go func() {
		defer close(p.inChannel)

		for _, word := range words {
			p.inChannel <- word
		}
	}()
}

// Run re-sender method
func (p *pipe) Run() {
	for {
		select {
		case <-p.ctx.Done():
			fmt.Println("Run: Time to return")
			return
		case value, ok := <-p.inChannel:
			if !ok {
				return
			}
			p.outChannel <- p.format(value)
		}
	}
}
