package user


// Repository defines the interface for user data storage.
type Repository interface {
	Create(u *User) error
	FindByUsername(username string) (*User, error)
}
