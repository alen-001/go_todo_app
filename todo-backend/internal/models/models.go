package models
type Status string
const (
	Pending Status = "pending"
	InProgress Status ="inprogress"
	Completed Status ="completed"
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