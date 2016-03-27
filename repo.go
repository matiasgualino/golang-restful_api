package main

import "fmt"

var currentId int

var types Types
var attributeTypes AttributeTypes
var beacons Beacons
var applications Applications

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

	redType := Type{Name: "Red", Attributes: att}
	mesaType := Type{Name: "Mesa", Attributes: mesaAtt}
	blueType := Type{Name: "Blue", Attributes: att}

	mixedTypes := Types{
		redType,
		mesaType,
		blueType,
	}

	colorTypes := Types{
		blueType,
		redType,
	}

	redBeacon := Beacon{Id: "1-1-1-1", PublicKey: "1", Minor: "1", Major: "1", UUID: "1", Types: Types{redType}}
	blueBeacon := Beacon{Id: "1-2-2-2", PublicKey: "1", Minor: "2", Major: "2", UUID: "2", Types: Types{blueType}}

	colorBeacons := Beacons{
		redBeacon,
		blueBeacon,
	}

	mixedBeacon := Beacon{Id: "2-1-1-1", PublicKey: "2", Minor: "1", Major: "1", UUID: "1", Types: Types{redType, mesaType}}
	blueMixedBeacon := Beacon{Id: "2-2-2-2", PublicKey: "2", Minor: "2", Major: "2", UUID: "2", Types: Types{blueType}}

	mixedTypesBeacons := Beacons{
		mixedBeacon,
		blueMixedBeacon,
	}

	colorApplication := Application{Id: "123456", PublicKey: "1", Description: "Aplicacion paleta de colores",
		Name: "Color Palette", Active: true, Types: colorTypes, Beacons: colorBeacons}

	mixedTypesApplication := Application{Id: "654321", PublicKey: "2", Description: "Aplicacion con multiples tipos",
		Name: "Mixed Types", Active: true, Types: mixedTypes, Beacons: mixedTypesBeacons}

	RepoCreateBeacon(redBeacon)
	RepoCreateBeacon(blueBeacon)
	RepoCreateBeacon(mixedBeacon)
	RepoCreateBeacon(blueMixedBeacon)

	RepoCreateApplication(colorApplication)
	RepoCreateApplication(mixedTypesApplication)

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

// ------------------------------------------------------------------------------------

func RepoFindApplication(id string) Application {
	for _, a := range applications {
		if a.Id == id {
			return a
		}
	}
	return Application{}
}

func RepoCreateApplication(a Application) Application {
	applications = append(applications, a)
	return a
}

func RepoDestroyApplication(id string) error {
	for i, a := range applications {
		if a.Id == id {
			applications = append(applications[:i], applications[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find application with id of %d to delete", id)
}
