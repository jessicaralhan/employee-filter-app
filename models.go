package main

type Employee struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Country string `json:"country"`
}
