package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/adYushinW/RestAPi/internal/app"
	"github.com/adYushinW/RestAPi/internal/log"
	"github.com/adYushinW/RestAPi/internal/model"
)

func Service(app *app.App) error {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {

		var todos []*model.Todo
		var err error

		switch {

		case r.Method == http.MethodGet:

			if r.URL.Query().Get("id") != "" {

				id, errr := strconv.Atoi(r.URL.Query().Get("id"))

				if errors.Is(errr, strconv.ErrSyntax) {
					log.Error(r, "http", http.StatusBadRequest, errr)
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("Wrong Request"))
					return
				}

				todos, err = app.GetOnlyOne(id)

			} else {
				state := r.URL.Query().Get("state")
				date1 := r.URL.Query().Get("date1")
				date2 := r.URL.Query().Get("date2")
				sort := r.URL.Query().Get("sort")
				limit := r.URL.Query().Get("limit")

				todos, err = app.GetTodo(state, date1, date2, sort, limit)
			}

		case r.Method == http.MethodPost:

			state, errr := strconv.ParseBool(r.URL.Query().Get("state"))
			date := r.URL.Query().Get("date")
			name := r.URL.Query().Get("name")

			if errors.Is(errr, strconv.ErrSyntax) {
				log.Error(r, "http", http.StatusBadRequest, errr)
				w.WriteHeader(http.StatusBadRequest)
			}

			todos, err = app.AddNew(state, date, name)

		case r.Method == http.MethodPut:

			id, _ := strconv.Atoi(r.URL.Query().Get("id"))
			state, _ := strconv.ParseBool(r.URL.Query().Get("state"))

			if errors.Is(err, strconv.ErrSyntax) {
				log.Error(r, "http", http.StatusBadRequest, err)
				w.WriteHeader(http.StatusBadRequest)
			}

			todos, err = app.ChangeStatus(id, state)

		}

		if err != nil {
			log.Error(r, "http", http.StatusInternalServerError, err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		response, err := json.Marshal(todos)
		if err != nil {
			log.Error(r, "http", http.StatusInternalServerError, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		log.Info(r, "http", http.StatusOK, nil)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})

	http.HandleFunc("/todo/delete", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodDelete {
			log.Info(r, "http", http.StatusBadRequest, nil)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(r.URL.Query().Get("id"))

		if errors.Is(err, strconv.ErrSyntax) {
			log.Error(r, "http", http.StatusBadRequest, err)
			w.WriteHeader(http.StatusBadRequest)
		}

		todos, _ := app.Delete(id)

		response, err := json.Marshal(todos)

		if err != nil {
			log.Error(r, "http", http.StatusInternalServerError, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		log.Info(r, "http", http.StatusOK, nil)
		w.WriteHeader(http.StatusOK)
		w.Write(response)

	})

	return http.ListenAndServe(":8080", nil)
}
