package storage

type (
	storageKey struct {
		Logged   string
		Email    string
		Name     string
		Verified string
	}
)

var StorageKey = storageKey{
	Logged:   "logged",
	Email:    "email",
	Name:     "name",
	Verified: "verified",
}
