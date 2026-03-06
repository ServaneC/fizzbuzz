package fizzbuzz

import "errors"

type Params struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

func (p Params) IsValid() error {
	if p.Int1 <= 0 {
		return errors.New("int1 must be > 0")
	}
	if p.Int2 <= 0 {
		return errors.New("int2 must be > 0")
	}
	if p.Limit <= 0 {
		return errors.New("limit must be > 0")
	}
	if p.Str1 == "" {
		return errors.New("str1 cannot be empty")
	}
	if p.Str2 == "" {
		return errors.New("str2 cannot be empty")
	}
	return nil
}
