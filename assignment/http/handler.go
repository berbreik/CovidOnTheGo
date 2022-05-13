package http

import (
	"assignment/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PatientHandler struct {
	patientService HospitalHandler
}

func NewHandler(patientService HospitalHandler) *PatientHandler {
	return &PatientHandler{patientService}
}

func (h *PatientHandler) AddHandler(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient
	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		http.Error(w, "Error parsing patient", http.StatusBadRequest)
		return
	}
	err = h.patientService.AddService(&patient)
	if err != nil {
		http.Error(w, "Error adding patient", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *PatientHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Error parsing id", http.StatusBadRequest)
		return
	}

	patient, err := h.patientService.GetService(id)
	if err != nil {
		http.Error(w, "Error getting patient", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(patient)

}

func (h *PatientHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Error parsing id", http.StatusBadRequest)
		return
	}

	var patient models.Patient
	err = json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		http.Error(w, "Error parsing patient", http.StatusBadRequest)
		return
	}

	err = h.patientService.UpdateService(id, &patient)
	if err != nil {
		http.Error(w, "Error updating patient", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (h *PatientHandler) RemoveHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Error parsing id", http.StatusBadRequest)
		return
	}

	err = h.patientService.RemoveService(id)
	if err != nil {
		http.Error(w, "Error removing patient", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *PatientHandler) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	patients, err := h.patientService.GetAllService()
	if err != nil {
		http.Error(w, "Error getting patients", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(patients)
}
