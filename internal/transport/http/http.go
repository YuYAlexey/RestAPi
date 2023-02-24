package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/adYushinW/RestAPi/internal/app"
	"github.com/adYushinW/RestAPi/internal/log"
)

func Service(app *app.App, logger log.Logger) error {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {

		var (
			respBody interface{}
			status   int
			err      error
		)

		switch r.Method {

		case http.MethodGet:

			if r.URL.Query().Get("id") != "" {

				id, errr := strconv.Atoi(r.URL.Query().Get("id"))

				if errors.Is(errr, strconv.ErrSyntax) {
					logger.Error(r, http.StatusBadRequest, errr)
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("Wrong Request"))
					return
				}

				respBody, err = app.GetOnlyOne(id)
				status = http.StatusOK

			} else {
				state := r.URL.Query().Get("state")
				date1 := r.URL.Query().Get("date1")
				date2 := r.URL.Query().Get("date2")
				sort := r.URL.Query().Get("sort")
				limit := r.URL.Query().Get("limit")

				respBody, err = app.GetTodo(state, date1, date2, sort, limit)
				status = http.StatusOK
			}

		case http.MethodPost:

			state, errr := strconv.ParseBool(r.URL.Query().Get("state"))
			date := r.URL.Query().Get("date")
			name := r.URL.Query().Get("name")

			if errors.Is(errr, strconv.ErrSyntax) {
				logger.Error(r, http.StatusBadRequest, errr)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			respBody, err = app.AddNew(state, date, name)
			status = http.StatusOK

		case http.MethodPut:

			id, _ := strconv.Atoi(r.URL.Query().Get("id"))
			state, _ := strconv.ParseBool(r.URL.Query().Get("state"))

			if errors.Is(err, strconv.ErrSyntax) {
				logger.Error(r, http.StatusBadRequest, err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			respBody, err = app.ChangeStatus(id, state)
			status = http.StatusOK

		case http.MethodDelete:
			var id int
			id, err = strconv.Atoi(r.URL.Query().Get("id"))

			if errors.Is(err, strconv.ErrSyntax) {
				logger.Error(r, http.StatusBadRequest, err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			respBody, err = app.Delete(id)
			status = http.StatusOK

		default:
			logger.Info(r, http.StatusBadRequest, nil)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err != nil {
			logger.Error(r, http.StatusInternalServerError, err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		response(w, status, respBody)
	})

	return http.ListenAndServe(":8080", nil)
}

func response(w http.ResponseWriter, status int, body interface{}) {
	response, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
