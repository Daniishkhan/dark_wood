package main

type DarkWoodMap struct {
	Name        string
	Description string
	Locations   []Location
}

type Location struct {
	Name        string
	Description string
}
