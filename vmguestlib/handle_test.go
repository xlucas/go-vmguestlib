package vmguestlib

import "testing"

func TestOpenAndCloseHandle(t *testing.T) {
	if h, err := NewHandle(); err != nil {
		t.Fatal(err)
	} else {
		h.Close()
	}
}

func TestUpdateInfo(t *testing.T) {
	if h, err := NewHandle(); err != nil {
		t.Fatal(err)
	} else if err := h.UpdateInfo(); err != nil {
		t.Fatal(err)
	} else {
		h.Close()
	}
}
