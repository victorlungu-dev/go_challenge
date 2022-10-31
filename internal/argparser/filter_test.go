package argparser

import "testing"

func TestNewStringFilter(t *testing.T) {
	expName := "testName"
	sf := NewStringFilter(expName, "testValue")
	if sf.Name != "testName" {
		t.Errorf("Expected name %s got %s", expName, sf.Name)
	}
}

func TestStringFilter_IsEqual(t *testing.T) {
	expValue := "testVale"
	sf := NewStringFilter("testName", expValue)
	if !sf.IsEqual(expValue) {
		t.Errorf("Expected %s == %s", sf.Value, expValue)
	}
}

func TestStringFilters_Filter(t *testing.T) {
	sf := NewStringFilter("test1", "v1")
	sf2 := NewStringFilter("test2", "v2")
	h := &TsvHeader{Columns: map[string]int{
		"test1": 0,
		"test2": 1,
	}}
	row := TsvRow{"v1", "v2"}
	sfs := StringFilters{sf, sf2}
	if !sfs.Filter(h, row) {
		t.Errorf("Expected test1 == v1 and test2 == v2")
	}

	h = &TsvHeader{Columns: map[string]int{
		"test3": 0,
		"test2": 1,
	}}
	row = TsvRow{"v1", "v2"}
	sfs = StringFilters{sf, sf2}
	if sfs.Filter(h, row) {
		t.Errorf("Expected false filter presence of not known column")
	}

	h = &TsvHeader{Columns: map[string]int{
		"test1": 0,
		"test2": 1,
	}}
	row = TsvRow{"v3", "v2"}
	sfs = StringFilters{sf, sf2}
	if sfs.Filter(h, row) {
		t.Errorf("Expected test1 != v1 and test2 == v2")
	}

}
