package main

type EnglishWord struct {
	Name        string   `json:"name"`
	Translation string   `json:"translation"`
	Noun        string   `json:"noun"`
	Adjective   string   `json:"adjective"`
	Verb        string   `json:"verb"`
	Adverb      string   `json:"adverb"`
	Examples    []string `json:"examples"`
}
