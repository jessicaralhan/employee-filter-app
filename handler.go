package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	country := r.URL.Query().Get("country")

	query := "SELECT id, name, status, country FROM employees WHERE 1=1"
	var args []interface{}

	if status != "" {
		query += " AND status = ?"
		args = append(args, status)
	}
	if country != "" {
		query += " AND country = ?"
		args = append(args, country)
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Status, &emp.Country); err != nil {
			log.Println("Row scan error:", err)
			continue
		}
		employees = append(employees, emp)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
