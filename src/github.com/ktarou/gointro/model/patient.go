package model
import "time"

type Patient struct {
	Id int
	FirstName string
	LastName string
	Dob time.Time
}

type ByFirstName []Patient

func (this ByFirstName) Len() int  {
	return len(this)
}

func (this ByFirstName) Less(i, j int) bool {
	return this[i].FirstName < this[j].FirstName
}

func (this ByFirstName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type ByLastName []Patient

func (this ByLastName) Len() int  {
	return len(this)
}

func (this ByLastName) Less(i, j int) bool {
	return this[i].LastName < this[j].LastName
}

func (this ByLastName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type ByDob []Patient

func (this ByDob) Len() int  {
	return len(this)
}

func (this ByDob) Less(i, j int) bool {
	return this[i].Dob.Before(this[j].Dob)
}

func (this ByDob) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}