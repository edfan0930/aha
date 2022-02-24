package storage

type (
	storageKey struct {
		Logged string
		Email  string
		Name   string
	}
)

var StorageKey = storageKey{
	Logged: "logged",
	Email:  "email",
	Name:   "name",
}
