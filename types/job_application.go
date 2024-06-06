package types

type ApplicationStatus string

const (
	Applied   ApplicationStatus = "applied"
	Interview ApplicationStatus = "interview"
	Offered   ApplicationStatus = "offered"
	Rejected  ApplicationStatus = "rejected"
)
