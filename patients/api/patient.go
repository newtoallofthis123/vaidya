package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/newtoallofthis123/patients/types"
)

func (s *ApiServer) handlePatientCreate(c *gin.Context) {
	var patientRequest types.PatientRequest

	if err := c.ShouldBindBodyWithJSON(&patientRequest); err != nil {
		s.logger.Error(fmt.Sprintln("Error binding request to types.patientRequest:", err))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	patient, err := s.store.NewPatient(patientRequest)
	if err != nil {
		s.logger.Error(fmt.Sprintln("Error creating a new patient:", err))
	}

	c.JSON(201, patient)
}

func (s *ApiServer) handlePatientGet(c *gin.Context) {
	id := c.Param("id")

	patient, err := s.store.GetPatient(id)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error getting patient with id %s:", id))
		c.JSON(404, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(200, patient)
}
