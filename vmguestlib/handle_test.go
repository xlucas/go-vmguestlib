package vmguestlib

import "testing"

func TestOpenAndCloseHandle(t *testing.T) {
	h, err := NewHandle()
	if err != nil {
		t.Fatal(err)
	}
	err = h.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateInfo(t *testing.T) {
	h, err := NewHandle()
	if err != nil {
		t.Fatal(err)
	}
	err = h.UpdateInfo()
	if err != nil {
		t.Fatal(err)
	}
	err = h.Close()
	if err != nil {
		t.Fatal(err)
	}
}
