package models

const (
    DoctorRole       = "doctor"
    ReceptionistRole = "receptionist"
)

var ValidRoles = map[string]bool{
    DoctorRole:       true,
    ReceptionistRole: true,
}
