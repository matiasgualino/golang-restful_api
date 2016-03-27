package main

import "fmt"

var currentId int

var types Types
var attributeTypes AttributeTypes
var beacons Beacons

// Give us some seed data
func init() {

	att := Attributes{
		Attribute{Name: "factor", Type: "Double"},
	}

	mesaAtt := Attributes{
		Attribute{Name: "number", Type: "Integer"},
	}

	RepoCreateType(Type{Name: "Red", Attributes: att})
	RepoCreateType(Type{Name: "Blue", Attributes: att})
	RepoCreateType(Type{Name: "Mesa", Attributes: mesaAtt})

	RepoCreateAttributeType(AttributeType{Name: "Entero", Id: "Integer"})
	RepoCreateAttributeType(AttributeType{Name: "Decimal", Id: "Double"})
	RepoCreateAttributeType(AttributeType{Name: "Cadena", Id: "String"})
	RepoCreateAttributeType(AttributeType{Name: "Booleano", Id: "Boolean"})

	multipleTypes := Types{
		Type{Name: "Red", Attributes: att},
		Type{Name: "Mesa", Attributes: mesaAtt},
	}

	types := Types{
		Type{Name: "Blue", Attributes: att},
	}

	RepoCreateBeacon(Beacon{Id: "1-1-1-1", PublicKey: "1", Minor: "1", Major: "1", UUID: "1", Types: multipleTypes})
	RepoCreateBeacon(Beacon{Id: "2-2-2-2", PublicKey: "2", Minor: "2", Major: "2", UUID: "2", Types: types})

}

// ------------------------------------------------------------------------------------

func RepoFindType(id int) Type {
	for _, t := range types {
		if t.Id == id {
			return t
		}
	}
	// return empty Type if not found
	return Type{}
}

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
	return fmt.Errorf("Could not find type with id of %d to delete", id)
}

// ------------------------------------------------------------------------------------

func RepoCreateAttributeType(at AttributeType) AttributeType {
	attributeTypes = append(attributeTypes, at)
	return at
}

// ------------------------------------------------------------------------------------

func RepoFindBeacon(id string) Beacon {
	for _, b := range beacons {
		if b.Id == id {
			return b
		}
	}
	return Beacon{}
}

func RepoCreateBeacon(b Beacon) Beacon {
	beacons = append(beacons, b)
	return b
}

func RepoDestroyBeacon(id string) error {
	for i, b := range beacons {
		if b.Id == id {
			beacons = append(beacons[:i], beacons[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find beacon with id of %d to delete", id)
}
