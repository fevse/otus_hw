package internalhttp

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)


func (s *Server) index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello, user!\n"))
		if err != nil {
			s.Logger.Error(err.Error())
		}
	}
}

func (s *Server) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		estr := EventStr{}
		data, err := io.ReadAll(r.Body)
		if err != nil {
			s.App.Logger.Error(err.Error())
			return 
		}
		err = json.Unmarshal(data, &estr)
		if err != nil {
			s.App.Logger.Error(err.Error())
			return 
		}
		event, err := estr.ParseEvent()
		if err != nil {
			s.App.Logger.Error(err.Error())
			return 
		}
		err = s.App.Storage.CreateEvent(&event)
		if err != nil {
			s.App.Logger.Error(err.Error())
		}
		w.Write([]byte(event.Title))
	}
}

func (s *Server) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		estr := EventStr{}
		data, err := io.ReadAll(r.Body)
		if err != nil {
			s.App.Logger.Error(err.Error())
			return 
		}
		err = json.Unmarshal(data, &estr)
		if err != nil {
			s.App.Logger.Error(err.Error())
			return 
		}
		event, err := estr.ParseEvent()
		if err != nil {
			s.App.Logger.Error(err.Error())
			return 
		}
		err = s.App.Storage.UpdateEvent(event.ID, &event)
		if err != nil {
			s.App.Logger.Error(err.Error())
		}
		w.Write([]byte(event.Title))
	}
}

func (s *Server) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jid := ID{}
		data, err := io.ReadAll(r.Body)
		if err != nil {
			s.App.Logger.Error(err.Error())
			return 
		}
		err = json.Unmarshal(data, &jid)
		if err != nil {
			s.App.Logger.Error(err.Error())
			return 
		}
		id, err := strconv.Atoi(string(jid.ID))
		if err != nil {
			s.App.Logger.Error(err.Error())
			return
		}
		err = s.App.Storage.DeleteEvent(int64(id))
		if err != nil {
			s.App.Logger.Error(err.Error())
		}
	}
}

func (s *Server) Show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events := s.App.Storage.List()
		w.Write([]byte("List: \n"))
		for _, v := range events {
			_, err := w.Write([]byte("#" + v.Title + "\n"))
			if err != nil {
				s.App.Logger.Error(err.Error())
			}
		}
	}
}

func (s *Server) ShowEventDay() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jdate := Date{}
		data, err := io.ReadAll(r.Body)
		if err != nil {
			s.App.Logger.Error(err.Error())
			return 
		}
		err = json.Unmarshal(data, &jdate)
		if err != nil {
			s.App.Logger.Error(err.Error())
			return 
		}
		date, err := time.Parse(time.DateOnly, jdate.Date)
		if err != nil {
			s.App.Logger.Error(err.Error())
			return
		}
		events, err := s.App.Storage.ListOfEventsDay(date)
		if err != nil {
			s.App.Logger.Error(err.Error())
			return
		}
		for _, v := range events {
			_, err := w.Write([]byte("#" + v.Title + "\n"))
			if err != nil {
				s.App.Logger.Error(err.Error())
			}
		}
	}
}

func (s *Server) ShowEventWeek() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jdate := Date{}
		data, err := io.ReadAll(r.Body)
		if err != nil {
			s.App.Logger.Error(err.Error())
			return 
		}
		err = json.Unmarshal(data, &jdate)
		if err != nil {
			s.App.Logger.Error(err.Error())
			return 
		}
		date, err := time.Parse(time.DateOnly, jdate.Date)
		if err != nil {
			s.App.Logger.Error(err.Error())
			return
		}
		events, err := s.App.Storage.ListOfEventsWeek(date)
		if err != nil {
			s.App.Logger.Error(err.Error())
			return
		}
		for _, v := range events {
			_, err := w.Write([]byte("#" + v.Title + "\n"))
			if err != nil {
				s.App.Logger.Error(err.Error())
			}
		}
	}
}

func (s *Server) ShowEventMonth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jdate := Date{}
		data, err := io.ReadAll(r.Body)
		if err != nil {
			s.App.Logger.Error(err.Error())
			return 
		}
		err = json.Unmarshal(data, &jdate)
		if err != nil {
			s.App.Logger.Error(err.Error())
			return 
		}
		date, err := time.Parse(time.DateOnly, jdate.Date)
		if err != nil {
			s.App.Logger.Error(err.Error())
			return
		}
		events, err := s.App.Storage.ListOfEventsMonth(date)
		if err != nil {
			s.App.Logger.Error(err.Error())
			return
		}
		for _, v := range events {
			_, err := w.Write([]byte("#" + v.Title + "\n"))
			if err != nil {
				s.App.Logger.Error(err.Error())
			}
		}
	}
}
