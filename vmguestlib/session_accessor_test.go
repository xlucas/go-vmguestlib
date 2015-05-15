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

func TestGetCPUUsed(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetCPUUsed()
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

func TestGetTimeElapsed(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetTimeElapsed()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetCpuStolen(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetCpuStolen()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetMemTargetSize(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetMemTargetSize()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetHostCPUUsed(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetHostCPUUsed()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetHostMemSwapped(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetHostMemSwapped()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetHostMemShared(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetHostMemShared()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetHostMemUsed(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetHostMemUsed()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetHostMemPhys(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetHostMemPhys()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetHostMemPhysFree(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetHostMemPhysFree()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetHostMemKernOvhd(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetHostMemKernOvhd()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetHostMemMapped(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetHostMemMapped()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetHostMemUnmapped(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.GetHostMemUnmapped()
	if err != nil {
		t.Fatal(err)
	}
}
