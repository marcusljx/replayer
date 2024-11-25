package replayer

import (
	"errors"
	"time"
)

type Configuration[T any] struct {
	Source       Source[T]
	GetTimestamp func(int, T) time.Time
	HandleError  func(error)
	BufferSize   int // defaults to 1
	Speed        float64

	timer time.Timer
}

func (c *Configuration[T]) Compile() (*Player[T], error) {
	if c.Source == nil {
		return nil, errors.New("source is nil")
	}
	if !c.Source.hasMore() {
		return nil, errors.New("source is empty")
	}

	initialT := c.Source.next()
	initialTimestamp := c.GetTimestamp(0, initialT)
	buffer := make(chan *triggerEvent[T], c.BufferSize)
	buffer <- &triggerEvent[T]{
		delay: 0,
		data:  initialT,
	}

	// feed source
	go func() {
		i := 1
		for c.Source.hasMore() {
			nextData := c.Source.next()

			// block if buffer is full
			evt := &triggerEvent[T]{
				delay: time.Duration(float64(c.GetTimestamp(i, nextData).Sub(initialTimestamp)) / c.Speed),
				data:  nextData,
			}
			buffer <- evt
			i++
		}

		close(buffer)
	}()

	return &Player[T]{
		buffer:      buffer,
		HandleError: c.HandleError,
	}, nil
}

type triggerEvent[T any] struct {
	delay time.Duration
	data  T
}
