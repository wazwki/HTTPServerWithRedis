package handlers

import (
	"encoding/json"
	"net/http"
	"rediscache/internal/db"
	"rediscache/internal/storage"
	"rediscache/internal/types"
	"strconv"
)

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	row := []types.User{}
	rows, err := db.DB.Query(`SELECT id, firstname, lastname FROM users`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer rows.Close()
	for rows.Next() {
		data := types.User{}
		err = rows.Scan(&data.ID, &data.Firstname, &data.Lastname)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !storage.CheckKey(data.ID) {
			storage.AddKey(data)
		}
		row = append(row, data)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(row)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var data types.User
	id := r.PathValue("id")
	idd, _ := strconv.Atoi(id)

	if !storage.CheckKey(idd) {
		rows, err := db.DB.Query(`SELECT id, firstname, lastname FROM users WHERE id=$1`, id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer rows.Close()

		data = types.User{}
		err = rows.Scan(&data.ID, &data.Firstname, &data.Lastname)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		storage.AddKey(data)
	} else {
		data = storage.GetKey(idd)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user types.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, err := db.DB.Exec(`INSERT INTO users(firstname, lastname) VALUES ($1, $2)`, user.Firstname, user.Lastname)
	id, _ := data.LastInsertId()
	user.ID = int(id)
	storage.AddKey(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func PutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := r.PathValue("id")

	var user types.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec(`UPDATE users SET firstname=$1, lastname=$2 WHERE id=$3`, user.Firstname, user.Lastname, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idd, _ := strconv.Atoi(id)
	if !storage.CheckKey(idd) {
		storage.AddKey(user)
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := r.PathValue("id")
	_, err := db.DB.Query(`DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idd, _ := strconv.Atoi(id)
	if storage.CheckKey(idd) {
		storage.DeleteKey(idd)
	}
	w.WriteHeader(http.StatusOK)
}
