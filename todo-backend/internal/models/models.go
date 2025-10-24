package models
type Status string

const (
	Pending Status = "pending"
	InProgress ="inprogress"
	Completed ="completed"
)

type Todo struct {
	Id     int
	Name   string
	Status Status
}
type Profile struct{
	Username string
	Hashed_Password string
}