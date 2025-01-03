package models

type Request struct {
	Authorization string            `json:"authorization"`
	Ip            string            `json:"ip"`
	UserAgent     string            `json:"userAgent"`
	Params        map[string]string `json:"params"`
	Body          interface{}       `json:"body"`
	Headers       map[string]string `json:"headers"`
	Files         interface{}       `json:"files"`
}
