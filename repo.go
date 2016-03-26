package main

import "fmt"

var currentId int

var types Types

// Give us some seed data
func init() {

	att := Attributes{
		Attribute{Name: "factor", Type: "Double"},
	}

	RepoCreateType(Type{Name: "Red", Attributes: att})
	RepoCreateType(Type{Name: "Blue", Attributes: att})
}

func RepoFindType(id int) Type {
	for _, t := range types {
		if t.Id == id {
			return t
		}
	}
	// return empty Type if not found
	return Type{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateType(t Type) Type {
	currentId += 1
	t.Id = currentId
	types = append(types, t)
	return t
}

func RepoDestroyType(id int) error {
	for i, t := range types {
		if t.Id == id {
			types = append(types[:i], types[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
