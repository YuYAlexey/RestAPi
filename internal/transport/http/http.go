package http

import (
	"encoding/json"
	"net/http"

	"github.com/adYushinW/RestAPi/internal/app"
)

func Service(app *app.App) error {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		todos, err := app.GetAllTodo()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		response, err := json.Marshal(todos)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})

	http.HandleFunc("/todo/sort", func(w http.ResponseWriter, r *http.Request) {
		sort := r.URL.Query().Get("key1")

		todos, err := app.GetAllSort(sort)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		response, err := json.Marshal(todos)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})

	http.HandleFunc("/todo/Undone", func(w http.ResponseWriter, r *http.Request) {

		todos, err := app.GetAllSortUndone()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		response, err := json.Marshal(todos)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})

	http.HandleFunc("/state", func(w http.ResponseWriter, r *http.Request) {
		sort := r.URL.Query().Get("sort")

		todos, err := app.GetAllState(sort)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		response, err := json.Marshal(todos)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})

	http.HandleFunc("/one", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		todos, err := app.GetOnlyOne(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		response, err := json.Marshal(todos)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})

	http.HandleFunc("/todo/row", func(w http.ResponseWriter, r *http.Request) {

		state := r.URL.Query().Get("state")
		date1 := r.URL.Query().Get("date1")
		date2 := r.URL.Query().Get("date2")
		limit := r.URL.Query().Get("limit")

		todos, err := app.GetSomeRow(state, date1, date2, limit)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		response, err := json.Marshal(todos)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})

	return http.ListenAndServe(":8080", nil)
}
