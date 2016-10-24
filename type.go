package ginsa

type Ref struct {
	Ref    string `json:"ref"`
	URL    string `json:"url"`
	Object Object `json:"object"`
}

type Object struct {
	Type string `json:"type"`
	SHA  string `json:"sha"`
	URL  string `json:"url"`
}

type Bank struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Kana string `json:"kana"`
	Hira string `json:"hira"`
	Roma string `json:"roma"`
}

type Branch struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Kana string `json:"kana"`
	Hira string `json:"hira"`
	Roma string `json:"roma"`
}
