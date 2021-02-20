package handlers

import "net/http"

//Получить весь справочник
func (h *Handlers) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.srv.GetAllUsers()
	if err != nil {
		apiErrorEncode(w, err)
		return
	}

	apiResponseEncoder(w, users)
}
