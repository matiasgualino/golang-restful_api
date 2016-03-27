package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TypesIndex",
		"GET",
		"/types",
		TypesIndex,
	},
	Route{
		"TypeCreate",
		"POST",
		"/types",
		TypeCreate,
	},
	Route{
		"TypeShow",
		"GET",
		"/types/{type_id}",
		TypeShow,
	},
	Route{
		"AtributeTypesIndex",
		"GET",
		"/attribute_types",
		AttributeTypesIndex,
	},
	Route{
		"BeaconTypesShow",
		"GET",
		"/beacon/{beacon_id}/types",
		BeaconTypesShow,
	},
	Route{
		"BeaconTypesCreate",
		"POST",
		"/beacon/{beacon_id}/types",
		BeaconTypesCreate,
	},
}
