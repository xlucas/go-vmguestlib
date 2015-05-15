package vmguestlib

import "testing"

func TestOpenAndCloseSession(t *testing.T) {
	h, err := NewHandle()
	if err != nil {
		t.Fatal(err)
	}
	s, err := Open(h)
	if err != nil {
		t.Fatal(err)
	}
	err = s.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewSessionAndClose(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	err = s.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestRefreshSession(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.Refresh()
	if err != nil {
		t.Fatal(err)
	}
	err = s.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestRefreshInfoSession(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.RefreshInfo()
	if err != nil {
		t.Fatal(err)
	}
	err = s.Close()
	if err != nil {
		t.Fatal(err)
	}
}
