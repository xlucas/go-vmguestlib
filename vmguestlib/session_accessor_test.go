package vmguestlib

import "testing"

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

func TestGetMemMapped(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetMemMapped()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetMemActive(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetMemActive()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetMemOverhead(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetMemOverhead()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetMemBallooned(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetMemBallooned()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetMemSwapped(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetMemSwapped()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetMemShared(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetMemShared()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetMemSharedSaved(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetMemSharedSaved()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetMemUsed(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetMemUsed()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetHostNumCPUCores(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetHostNumCPUCores()
	if err != nil {
		t.Fatal(err)
	}
}
