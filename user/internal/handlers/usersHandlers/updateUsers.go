package usershandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user/domain/model"
	"user/internal/store"
	"user/pkg/response"

	"github.com/julienschmidt/httprouter"
)

// update user ...
func UpdateUser(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		req := &model.User{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Eror during JSON request decoding. Request body: %v, Err msg: %w", r.Body, err)
			json.NewEncoder(w).Encode(response.Error{Messsage: fmt.Sprintf("Eror during JSON request decoding. Request body: %v, Err msg: %v", r.Body, err)})
			return
		}

		//id := 
		
		err := s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
		}

		u, err := s.User().FindByID(req.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("Cant find user. Err msg:%v.", err)
			return
		}

		email := req.Email
		if email != "" {
			u.Email = email
		}

		role := req.Role
		if role != "" {
			u.Role = model.Role(role)
		}

		name := req.Name
		if name != "" {
			u.Name = name
		}

		surname := req.Surname
		if surname != "" {
			u.Surname = surname
		}

		middleName := req.MiddleName
		if surname != "" {
			u.MiddleName = middleName
		}

		sex := req.Sex
		if sex != "" {
			u.Sex = model.Sex(sex)
		}

		//TODO
		/*layout := "2006-01-02"
		dateOfBirth := req.DateOfBirth
		if dateOfBirth != "" {
			dateOfBirth, err := time.Parse(layout, r.FormValue("DateOfBirth"))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("DateOfBirth"))
				return
			}
			u.DateOfBirth = dateOfBirth
		}*/

		address := req.Address
		if address != "" {
			u.Address = address
		}

		phone := req.Phone
		if phone != "" {
			u.Phone = phone
		}

		photo := req.Photo
		if photo != "" {
			u.Photo = photo
		}

		err = u.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Data is not valid. Err msg:%v.", err)
			return
		}

		err = s.User().Update(u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.Logger.Errorf("Can't update user. Err msg:%v.", err)
			return
		}

		s.Logger.Info("Update user with id = %d", u.UserID)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response.Info{Messsage: fmt.Sprintf("Update user with id = %d", u.UserID)})

	}
}
