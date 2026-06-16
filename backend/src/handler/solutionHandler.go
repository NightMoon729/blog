package handler

import (
	"blog/backend/src/Service"
	"encoding/json"
	"net/http"
	"strconv"
)

type SolutionHandler struct {
	svc Service.SolutionService
}

func NewSolutionHandler(svc Service.SolutionService) *SolutionHandler {
	return &SolutionHandler{svc: svc}
}

func (h *SolutionHandler) GetSolutions(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	n, _ := strconv.Atoi(r.URL.Query().Get("size"))
	// 2. 调 service
	solutions, err := h.svc.GetSolutions(page, n)
	if err != nil {
		http.Error(w, "internal error", 500)
		return
	}
	// 3. 返回 JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(solutions)
}
