package stub

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type GetPetNamesStub struct { //liststart1
	Entities
}

func (ps GetPetNamesStub) GetPets(userID string) ([]Pet, error) {
	switch userID {
	case "1":
		return []Pet{{Name: "Bubbles"}}, nil
	case "2":
		return []Pet{{Name: "Stampy"}, {Name: "Snowball II"}}, nil
	default:
		return nil, fmt.Errorf("不正なID: %s", userID)
	}
} //listend1

func TestLogicGetPetNames(t *testing.T) { //liststart2
	data := []struct {
		name     string
		userID   string
		petNames []string
	}{
		{"case1", "1", []string{"Bubbles"}},
		{"case2", "2", []string{"Stampy", "Snowball II"}},
		{"case3", "3", nil},
	}
	l := Logic{GetPetNamesStub{}}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			petNames, err := l.GetPetNames(d.userID)
			if err != nil {
				t.Error(err)
			}
			if diff := cmp.Diff(d.petNames, petNames); diff != "" {
				t.Error(diff)
			}
		})
	}
} //listend2

type EntitiesStub struct { //liststart3
	getUser     func(id string) (User, error)
	getPets     func(userID string) ([]Pet, error)
	getChildren func(userID string) ([]Person, error)
	getFriends  func(userID string) ([]Person, error)
	saveUser    func(user User) error
} //listend3

//liststart4
func (es EntitiesStub) GetUser(id string) (User, error) {
	return es.getUser(id)
}

func (es EntitiesStub) GetPets(userID string) ([]Pet, error) {
	return es.getPets(userID)
}
//listend4

func (es EntitiesStub) GetChildren(userID string) ([]Person, error) {
	return es.getChildren(userID)
}

func (es EntitiesStub) GetFriends(userID string) ([]Person, error) {
	return es.getFriends(userID)
}

func (es EntitiesStub) SaveUser(user User) error {
	return es.saveUser(user)
}

func TestLogicGetPetNames2(t *testing.T) { //liststart5
	data := []struct {
		name     string
		getPets  func(userID string) ([]Pet, error)
		userID   string
		petNames []string
		errMsg   string
	}{
		{"case1", func(userID string) ([]Pet, error) {
			return []Pet{{Name: "Bubbles"}}, nil
		}, "1", []string{"Bubbles"}, ""},
		{"case2", func(userID string) ([]Pet, error) {
			return []Pet{{Name: "Stampy"}, {Name: "Snowball II"}}, nil
		}, "2", []string{"Stampy", "Snowball II"}, ""},
		{"case3", func(userID string) ([]Pet, error) {
			return nil, errors.New("不正なID: 3")
		}, "3", nil, "不正なID: 3"},
	}
	l := Logic{}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			l.Entities = EntitiesStub{getPets: d.getPets}
			petNames, err := l.GetPetNames(d.userID)
			if diff := cmp.Diff(petNames, d.petNames); diff != "" {
				t.Error(diff)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("予想されたエラーメッセージ `%s`, 得られたエラーメッセージ `%s`", d.errMsg, errMsg)
			}
		})
	}
} //listend5
