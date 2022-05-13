// package store implements the storing of data in database
package store

import (
	"assignment/models"
	"database/sql"
	"errors"
)
// Datastore 
type Datastore struct {
	db *sql.DB
}

func NewDatastore(db *sql.DB) *Datastore {
	return &Datastore{db}
}

func (d *Datastore) AdmitPatient(patient *models.Patient)  error {
	_, err := d.db.Exec("INSERT INTO patient (name, age,email,phone,covid) VALUES (?, ?, ?,?,?)", patient.Name, patient.Age,patient.Email,patient.Phone, patient.Covid)
	if err != nil {
		return errors.New("error inserting patient")
	}
	return nil
}

func (d *Datastore) GetPatient(id int) (*models.Patient, error) {
	patient := &models.Patient{}
	err := d.db.QueryRow("SELECT name, age,email,phone, covid FROM patient WHERE id = ?", id).Scan(&patient.Name, &patient.Age,&patient.Email,&patient.Phone, &patient.Covid)
	if err != nil {
		return  &models.Patient{}, errors.New("error retrieving patient")
	}
	return patient, nil
}

func (d *Datastore) UpdatePatient(id int, patient *models.Patient) error {
	_, err := d.db.Exec("UPDATE patient SET name = ?, age = ?, email = ? , phone = ? ,covid = ? WHERE id = ?", patient.Name, patient.Age, patient.Email,patient.Phone,patient.Covid, id)
	if err != nil {
		return errors.New("error updating patient status")
	}
	return nil
}

func (d *Datastore) RemovePatient(id int) error {
	_, err := d.db.Exec("DELETE FROM patient WHERE id = ?", id)
	if err != nil {
		return errors.New("error removing patient")
	}
	return nil
}

func (d *Datastore) GetAllPatients() ([]*models.Patient, error) {
	rows, err := d.db.Query("SELECT id, name, age,email,phone, covid FROM patient")
	if err != nil {
		return []*models.Patient{}, errors.New("error retrieving patients")
	}
	defer rows.Close()
	patients := []*models.Patient{}
	for rows.Next() {
		patient := &models.Patient{}
		err := rows.Scan(&patient.Id, &patient.Name, &patient.Age,&patient.Email,&patient.Phone, &patient.Covid)
		if err != nil {
			return []*models.Patient{}, errors.New("error retrieving patients")
		}
		patients = append(patients, patient)
	}
	return patients, nil
}