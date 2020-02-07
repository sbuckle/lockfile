package lockfile

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLockRemoved(t *testing.T) {
	lock := New("/tmp/test.lock")
	if err := lock.Unlock(); err != ErrNotExist {
		t.Error("Expected to receive ErrNotExist error")
	}
}

func TestLockOption(t *testing.T) {
	retries := 5
	lock := New("/tmp/test.lock", SetMaxRetries(retries))
	if lock.retries != retries {
		t.Errorf("Expected %d got %d", retries, lock.retries)
	}
}

func TestFailToAcquireLock(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "lock")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	lock := New(tmpfile.Name(), SetMaxRetries(2), SetInterval(2))
	err = lock.Lock()
	if err != nil {
		if err != ErrTimeout {
			t.Error("Expected ErrTimeout error")
		}
	}
}

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
