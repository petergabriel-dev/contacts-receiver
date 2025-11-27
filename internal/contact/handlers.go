package contact

import (
	"contacts/internal/json"
	"log"
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListContacts(w http.ResponseWriter, r *http.Request) {
	err := h.service.ListContacts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	contacts := struct {
		Contact []string `json:"contact"`
	}{
		Contact: []string{"Contact 1", "Contact 2", "Contact 3"}}

	json.Write(w, http.StatusOK, contacts)
}
