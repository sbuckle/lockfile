package lockfile

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLockFileExists(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "lock")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	lock := New(tmpfile.Name())
	if err := lock.Lock(); err == nil {
		t.Error("Call to lock file should have failed")
	}
}

func TestLockUnlock(t *testing.T) {
	lock := New("/tmp/test.lock")
	err := lock.Lock()
	if err != nil {
		t.Error("Failed to acquire file lock")
		return
	}
	err = lock.Unlock()
	if err != nil {
		t.Error("Failed to unlock file")
	}
}
