package syncs

import "testing"

func TestIsSingleInstance(t *testing.T) {
	ok, err := IsSingleInstance("CE8FoDlQ")
	if err != nil {
		t.Error(err)
		return
	}
	if !ok {
		t.Errorf("IsSingleInstance should return true")
		return
	}
	ok, err = IsSingleInstance("CE8FoDlQ")
	if err != nil {
		t.Error(err)
		return
	}
	if ok {
		t.Errorf("IsSingleInstance should return false")
		return
	}
}
