package storage

type (
	storageKey struct {
		Logged string
		Email  string
	}
)

var StorageKey = storageKey{
	Logged: "logged",
	Email:  "email",
}
