package vmguestlib

import "testing"

func TestOpenAndCloseSession(t *testing.T) {
	var err error
	var h *Handle
	var s *Session

	if h, err = NewHandle(); err != nil {
		t.Fatal(err)
	}
	if s, err = Open(h); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestNewSessionAndClose(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestRefreshSession(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.Refresh(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestRefreshInfoSession(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.RefreshInfo(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}
