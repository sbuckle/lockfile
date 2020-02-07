package lockfile

import "os"

type Lockfile string

func New(path string) Lockfile {
	return Lockfile(path)
}

func (l Lockfile) Lock() error {
	f, err := os.OpenFile(string(l), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0000)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func (l Lockfile) Unlock() error {
	return os.Remove(string(l))
}
