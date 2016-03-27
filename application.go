package main

type Application struct {
	Beacons     Beacons `json:"beacons"`
	Types       Types   `json:"types"`
	IconURL     string  `json:"icon_url"`
	Active      bool    `json:"active"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	PublicKey   string  `json:"public_key"`
	Id          string  `json:"id"`
}

type Applications []Application
