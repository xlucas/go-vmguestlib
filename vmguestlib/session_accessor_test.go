package vmguestlib

import "testing"

func TestGetCPULimit(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetCPULimit()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetCPUReservation(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetCPUReservation()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetCPUShares(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetCPUShares()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetElaspedTime(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetElaspedTime()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetHostProcessorSpeed(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetHostProcessorSpeed()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetMemReservation(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetMemReservation()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetMemLimit(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetMemLimit()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetMemShares(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetMemShares()
	if err != nil {
		t.Fatal(err)
	}
}
