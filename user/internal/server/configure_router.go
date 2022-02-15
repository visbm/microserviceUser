package server

import (
	pethandlers "user/internal/handlers/pethandlers"
	usershandlers "user/internal/handlers/usersHandlers"
	"user/internal/store"
)

// ConfigureRouter ...
func (s *Server) ConfigureRouter() {

	s.Router.Handle("GET", "/pets", pethandlers.AllPetsHandler(store.New(s.Config)))
	s.Router.Handle("GET", "/pets/id", pethandlers.GetPetByID(store.New(s.Config)))
	s.Router.Handle("POST", "/pets/delete", pethandlers.DeletePets(store.New(s.Config)))

	s.Router.Handle("GET", "/users", usershandlers.AllUsersHandler(store.New(s.Config)))
	s.Router.Handle("GET", "/users/id/", usershandlers.GetUserByID(store.New(s.Config)))
	s.Router.Handle("POST", "/users/delete", usershandlers.DeleteUser(store.New(s.Config)))
	s.Router.Handle("POST", "/users/update", usershandlers.UpdateUser(store.New(s.Config)))
	s.Router.Handle("POST", "/users/new", usershandlers.NewUser(store.New(s.Config)))

	/*s.Router.Handle("POST", "/login", handlers.LoginHandle(store.New(s.Config)))
	s.Router.Handle("POST", "/logout", handlers.LogoutHandle(store.New(s.Config)))
	s.Router.Handle("POST", "/refresh", handlers.RefreshHandle(store.New(s.Config)))
	s.Router.Handle("POST", "/registration", handlers.RegistrationHandle(store.New(s.Config)))*/
}
