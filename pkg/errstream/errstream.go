package errstream

import "sync"

type ErrStream struct {
	wg      *sync.WaitGroup
	ErrChan chan error
}

func New(bufferSize int) *ErrStream {
	wg := new(sync.WaitGroup)
	errCH := make(chan error, bufferSize)
	return &ErrStream{
		wg:      wg,
		ErrChan: errCH,
	}
}

func (e *ErrStream) Go(fn func() error) {
	e.wg.Add(1)
	go func() {
		defer e.wg.Done()
		if err := fn(); err != nil {
			e.ErrChan <- err
		}
	}()
}
