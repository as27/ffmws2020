package user

type User struct {
	ID   int
	Name string
}

type UserStorage struct {
	Users []User
}

// Add fügt einen neuen User zum UserStorage hinzu, wenn dieser
// noch nicht vorhanden ist. Bei Erfolg wird true zurück gegeben.
// Sollte es bereits einen User mit dieser ID geben, so darf der
// User nicht hinzugefügt werden. In dem Fall ist der Rückgabewert
// false
func (us *UserStorage) Add(u User) bool {

	return true
}

// GetUserByID sucht im Storage den User mit der angegebenen ID.
// Ist der User vorhanden so wird dieser zurück gegeben. Der zweite
// Parameter ist dabei true.
// Sollte die ID nicht vorhanden sein, so wird User{} und false
// zurück gegeben.
func (us *UserStorage) GetUserByID(id int) (User, bool) {

	return User{}, true
}
