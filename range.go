package tddbc

import (
	"errors"
	"fmt"
)

type Range struct {
	lower int64
	upper int64
}

func Create(lower int64, upper int64) (*Range, error) {
	if lower > upper {
		return nil, errors.New("エラー：上端点より下端点が大きい閉区間を作ることはできない")
	}
	return &Range{
		lower: lower,
		upper: upper,
	}, nil
}

func (r *Range) ToString() string {
	return fmt.Sprintf("[%d,%d]", r.lower, r.upper)
}

func (r *Range) IncludeNum(n int64) (bool, error) {
	if n >= r.lower && n <= r.upper {
		return true, nil
	}
	return false, nil
}

func (r *Range) EqualRange(cmp *Range) (bool, error) {
	if r.lower == cmp.lower && r.upper == cmp.upper {
		return true, nil
	}
	return false, nil
}

func (r *Range) IncludeRange(cmp *Range) (bool, error) {
	if r.lower <= cmp.lower && cmp.upper <= r.upper {
		return true, nil
	}
	return false, nil
}
