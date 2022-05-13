package service

import (
	"assignment/models"
	"errors"
)

type PatientService struct {
	patientStore HospitalService
}

func NewService(patientStore HospitalService) *PatientService {
	return &PatientService{patientStore}
}

func (p *PatientService) AddService(patient *models.Patient) error {
	if patient.Covid != true && patient.Covid != false {
		return errors.New("Invalid covid status")
	}
	if patient.Age < 0 {
		return errors.New("Invalid age")
	}
	if patient.Name == "" {
		return errors.New("Invalid name")
	}
	if patient.Email == "" {
		return errors.New("Invalid email")
	}
	if patient.Phone == "" {
		return errors.New("Invalid phone")
	}


	err := p.patientStore.AdmitPatient(patient)
	if err != nil {
		return errors.New("error inserting patient")
	}
	return nil
}

func (p *PatientService) GetService(id int) (*models.Patient, error) {
	if id < 0 {
		return &models.Patient{}, errors.New("Invalid id")
	}

	res, err :=  p.patientStore.GetPatient(id)
	if err != nil {
		return &models.Patient{}, errors.New("error retrieving patient")
	}
	return res, nil
}

func (p *PatientService) UpdateService(id int, patient *models.Patient) error {
	if id < 0 {
		return errors.New("Invalid id")
	}
	if patient.Covid != true && patient.Covid != false {
		return errors.New("Invalid covid status")
	}
	if patient.Age < 0 {
		return errors.New("Invalid age")
	}
	if patient.Name == "" {
		return errors.New("Invalid name")
	}
	if patient.Email == "" {
		return errors.New("Invalid email")
	}
	if patient.Phone == "" {
		return errors.New("Invalid phone")
	}

	err :=  p.patientStore.UpdatePatient(id, patient)
	if err != nil {
		return errors.New("error updating patient status")
	}
	return nil
}

func (p *PatientService) RemoveService(id int) error {
	if id < 0 {
		return errors.New("invalid id")
	}
	err :=  p.patientStore.RemovePatient(id)
	if err != nil {
		return errors.New("error removing patient")
	}
	return nil
}

func (p *PatientService) GetAllService() ([]*models.Patient, error) {
	 res,err :=p.patientStore.GetAllPatients()
	if err != nil {
		return []*models.Patient{}, errors.New("error retrieving all patients")
	}
	return res, nil
}