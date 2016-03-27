package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TypesIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(types); err != nil {
		panic(err)
	}
}

func TypeShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var type_id int
	var err error
	if type_id, err = strconv.Atoi(vars["type_id"]); err != nil {
		panic(err)
	}
	type_obj := RepoFindType(type_id)
	if type_obj.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(type_obj); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

func TypeCreate(w http.ResponseWriter, r *http.Request) {
	var type_obj Type
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &type_obj); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateType(type_obj)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

// --------------------------------AttributeTypes--------------------------------
func AttributeTypesIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(attributeTypes); err != nil {
		panic(err)
	}
}

// --------------------------------BeaconTypes--------------------------------
func BeaconTypesShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var beacon_id = vars["beacon_id"]
	var beacon = RepoFindBeacon(beacon_id)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(beacon.Types); err != nil {
		panic(err)
	}
	return

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

func BeaconTypesCreate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var beacon_id = vars["beacon_id"]
	var beacon = RepoFindBeacon(beacon_id)

	var types Types

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &types); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	RepoDestroyBeacon(beacon_id)
	beacon.Types = types

	beacon = RepoCreateBeacon(beacon)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(beacon.Types); err != nil {
		panic(err)
	}
}

// --------------------------------Application--------------------------------
func ApplicationsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(applications); err != nil {
		panic(err)
	}
}

func ApplicationShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var application_id = vars["application_id"]
	var application = RepoFindApplication(application_id)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(application); err != nil {
		panic(err)
	}
	return

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

func ApplicationBeaconsShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var application_id = vars["application_id"]
	var application = RepoFindApplication(application_id)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(application.Beacons); err != nil {
		panic(err)
	}
	return

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

func ApplicationTypesShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var application_id = vars["application_id"]
	var application = RepoFindApplication(application_id)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(application.Types); err != nil {
		panic(err)
	}
	return

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

func ApplicationBeaconsCreate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var application_id = vars["application_id"]
	var application = RepoFindApplication(application_id)

	var beacons Beacons

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &beacons); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	RepoDestroyApplication(application_id)
	application.Beacons = beacons

	application = RepoCreateApplication(application)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(application.Beacons); err != nil {
		panic(err)
	}
}

func ApplicationTypesCreate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var application_id = vars["application_id"]
	var application = RepoFindApplication(application_id)

	var types Types

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &types); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	RepoDestroyApplication(application_id)
	application.Types = types

	application = RepoCreateApplication(application)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(application.Types); err != nil {
		panic(err)
	}
}
