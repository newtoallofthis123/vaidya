package db

import (
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/newtoallofthis123/patients/types"
	"github.com/newtoallofthis123/ranhash"
)

func (d *Store) NewPatient(pr types.PatientRequest) (string, error) {
	patientId := ranhash.GenerateRandomString(12)

	_, err := d.pq.Insert("patients").Columns("id", "name", "age", "gender", "address", "identity", "phone", "conditions", "problems", "description").
		Values(patientId, pr.Name, pr.Age, pr.Gender, pr.Address, pr.Identity, pr.Phone, pr.Conditions, pr.Problems, pr.Description).
		RunWith(d.db).Exec()

	return patientId, err
}

func (d *Store) GetPatient(patientId string) (types.Patient, error) {
	row := d.pq.Select("*").From("patients").Where(squirrel.Eq{"id": patientId}).RunWith(d.db).QueryRow()

	var pr types.Patient
	var createdAt time.Time

	err := row.Scan(
		&pr.PatientID,
		&pr.Name,
		&pr.Age,
		&pr.Gender,
		&pr.Address,
		&pr.Identity,
		&pr.Phone,
		&pr.Conditions,
		&pr.Problems,
		&pr.Description,
		&pr.Medicines,
		&pr.Diagnosis,
		&pr.NextSession,
		&createdAt,
	)
	if err != nil {
		return types.Patient{}, err
	}

	return pr, nil
}

func (d *Store) EditPatient(patient types.Patient) error {
	_, err := d.pq.Update("patients").
		Set("name", patient.Name).
		Set("age", patient.Age).
		Set("gender", patient.Gender).
		Set("address", patient.Address).
		Set("identity", patient.Identity).
		Set("phone", patient.Phone).
		Set("conditions", patient.Conditions).
		Set("problems", patient.Problems).
		Set("description", patient.Description).
		Set("diagnosis", patient.Diagnosis).
		Set("medicines", patient.Medicines).
		Set("next_session", patient.NextSession).
		Where(squirrel.Eq{"id": patient.PatientID}).
		RunWith(d.db).Exec()

	return err
}
