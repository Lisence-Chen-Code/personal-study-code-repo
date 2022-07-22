package test

import (
	"time"
)

type Error struct {
}

func (e *Error) Error() string {
	return "this callback need to retry."
}

type Retryable struct {
	MaxRetryTimes int
	Delay         time.Duration
	Callback      func() error
}

func (r *Retryable) Run() error {
	var err error
	for i := 0; i < r.MaxRetryTimes; i++ {
		err = r.Callback()
		if err == nil {
			break
		}
		if r.Delay > 0 {
			time.Sleep(r.Delay)
		}
	}
	return err
}

func Do(callback func() error, maxRetryTime int, delays ...time.Duration) error {
	delay := time.Duration(0)
	if len(delays) > 0 {
		delay = delays[0]
	}
	retryAble := &Retryable{
		MaxRetryTimes: maxRetryTime,
		Delay:         delay,
		Callback:      callback,
	}
	return retryAble.Run()
}
