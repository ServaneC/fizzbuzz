package fizzbuzz

type Params struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

func (p Params) IsValid() bool {
	return p.Int1 > 0 &&
		p.Int2 > 0 &&
		p.Limit > 0 &&
		p.Str1 != "" &&
		p.Str2 != ""
}
