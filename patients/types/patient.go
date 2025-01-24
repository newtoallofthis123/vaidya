package types

type PatientRequest struct {
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Age         uint32 `json:"age,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Address     string `json:"address,omitempty"`
	Identity    string `json:"identity,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Email       string `json:"email,omitempty"`
	Description string `json:"description,omitempty"`
	Recurring   bool   `json:"recurring,omitempty"`
}

type Patient struct {
	PatientID   string `json:"patient_id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Age         uint32 `json:"age,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Address     string `json:"address,omitempty"`
	Identity    string `json:"identity,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Email       string `json:"email,omitempty"`
	Description string `json:"description,omitempty"`
	Recurring   bool   `json:"recurring,omitempty"`
}

type PatientRecord struct {
	Patient     Patient  `json:"patient,omitempty"`
	SessionID   string   `json:"session_id,omitempty"`
	Problems    []string `json:"problems,omitempty"`
	Medicines   []string `json:"medicines,omitempty"`
	Conditions  []string `json:"conditions,omitempty"`
	Diagnosis   string   `json:"diagnosis,omitempty"`
	NextSession string   `json:"next_session,omitempty"`
}
