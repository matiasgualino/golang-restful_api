package main

type Type struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Attributes Attributes `json:"attributes"`
}

type Types []Type
