package replayer

import (
	"sync"
	"time"
)

type Player[T any] struct {
	buffer      chan *triggerEvent[T]
	HandleError func(error)
	sync.WaitGroup
}

func (p *Player[T]) Play(playFunc func(T) error) {
	for event := range p.buffer {
		p.Add(1)
		go func() {
			timer := time.NewTimer(event.delay)
			defer timer.Stop()
			<-timer.C
			if err := playFunc(event.data); err != nil {
				p.HandleError(err)
			}
			p.Done()
		}()
	}
	p.Wait()
}
