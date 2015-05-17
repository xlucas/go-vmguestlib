package vmguestlib

import "testing"

func TestGetCPUReservation(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetCPUReservation(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetCPULimit(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetCPULimit(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetCPUShares(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetCPUShares(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetHostProcessorSpeed(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetHostProcessorSpeed(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetMemReservation(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetMemReservation(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetMemLimit(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetMemLimit(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetMemShares(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetMemShares(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetMemMapped(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetMemMapped(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetMemActive(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetMemActive(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetMemOverhead(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetMemOverhead(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetMemBallooned(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetMemBallooned(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetMemSwapped(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetMemSwapped(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetMemShared(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetMemShared(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetMemSharedSaved(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetMemSharedSaved(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetMemUsed(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetMemUsed(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetHostNumCPUCores(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetHostNumCPUCores(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetTimeElapsed(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetTimeElapsed(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetCPUStolen(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetCPUStolen(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetMemTargetSize(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetMemTargetSize(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetCPUUsed(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetCPUUsed(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetHostCPUUsed(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetHostCPUUsed(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetHostMemSwapped(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetHostMemSwapped(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetHostMemShared(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetHostMemShared(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetHostMemUsed(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetHostMemUsed(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetHostMemPhys(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetHostMemPhys(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetHostMemPhysFree(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetHostMemPhysFree(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetHostMemKernOvhd(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetHostMemKernOvhd(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetHostMemMapped(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetHostMemMapped(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}

func TestGetHostMemUnmapped(t *testing.T) {
	var err error
	var s *Session

	if s, err = NewSession(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetHostMemUnmapped(); err != nil {
		t.Fatal(err)
	}

	s.Close()
}
