package argparser

import (
	"errors"
	"fmt"
)

type TsvRow []string

type TsvHeader struct {
	Columns map[string]int
}

func (th TsvHeader) ColumnValue(columnName string, row TsvRow) (string, error) {
	if val, ok := th.Columns[columnName]; ok {
		return row[val], nil
	}
	return "", errors.New(fmt.Sprintf("missing column %s", columnName))
}
