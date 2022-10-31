package argparser

type StringFilter struct {
	Name  string
	Value string
}

func (sf StringFilter) IsEqual(s string) bool {
	return sf.Value == s
}

type StringFilters []StringFilter

func (sfs StringFilters) Filter(h *TsvHeader, row TsvRow) bool {
	for _, f := range sfs {
		v, err := h.ColumnValue(f.Name, row)
		if err != nil {
			return false
		}
		if !f.IsEqual(v) {
			return false
		}
	}
	return true
}

func NewStringFilter(name string, value string) StringFilter {
	return StringFilter{
		Name:  name,
		Value: value,
	}
}
