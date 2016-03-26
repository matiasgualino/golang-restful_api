package main

type Attribute struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Attributes []Attribute
