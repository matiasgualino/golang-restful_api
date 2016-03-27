package main

type Beacon struct {
	Types     Types  `json:"types"`
	Major     string `json:"major"`
	Minor     string `json:"minor"`
	UUID      string `json:"uuid"`
	PublicKey string `json:"public_key"`
	Id        string `json:"id"`
}

type Beacons []Beacon
