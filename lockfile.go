package lockfile

import (
	"errors"
	"os"
	"time"
)

var (
	ErrTimeout  = errors.New("Failed to acquire in specified time interval")
	ErrNotExist = errors.New("Lockfile does not exist")
)

type Lockfile struct {
	path     string
	retries  int
	interval int // sleep period (seconds)
}

type Option func(*Lockfile)

func SetMaxRetries(i int) Option {
	return func(f *Lockfile) {
		f.retries = i
	}
}

func SetInterval(i int) Option {
	return func(f *Lockfile) {
		f.interval = i
	}
}

func New(path string, opts ...Option) Lockfile {
	lf := Lockfile{path: path, retries: 1}
	for _, opt := range opts {
		opt(&lf)
	}
	return lf
}

func (l Lockfile) Lock() error {
	i := 1
	for {
		if i > l.retries {
			return ErrTimeout
		}
		f, err := os.OpenFile(l.path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0000)
		if err == nil {
			defer f.Close()
			break
		}
		if !os.IsExist(err) {
			return err
		}
		if l.interval > 0 {
			time.Sleep(time.Duration(l.interval) * time.Second)
		}
		i += 1
	}
	return nil
}

func (l Lockfile) Unlock() error {
	if _, err := os.Stat(l.path); os.IsNotExist(err) {
		return ErrNotExist
	}
	return os.Remove(l.path)
}
