package stub

type User struct{}
type Pet struct {
	Name string
}
type Person struct{}

type Entities interface { //liststart1
	GetUser(id string) (User, error)
	GetPets(userID string) ([]Pet, error)
	GetChildren(userID string) ([]Person, error)
	GetFriends(userID string) ([]Person, error)
	SaveUser(user User) error
} //listend1

type Logic struct { //liststart2
	Entities Entities
} //listend2

//liststart3
func (l Logic) GetPetNames(userId string) ([]string, error) { 
	pets, err := l.Entities.GetPets(userId)
	if err != nil {
		return nil, err
	}
	out := make([]string, 0, len(pets))
	for _, p := range pets {
		out = append(out, p.Name)
	}
	return out, nil
} //listend3
