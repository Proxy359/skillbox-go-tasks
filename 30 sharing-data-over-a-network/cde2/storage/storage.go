package storage

import (
	"errors"
	"log"
	"strconv"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
	ID      int    `json:"id"`
}

type Storage struct {
	Store map[int]*User
}

func (s *Storage) Save(u *User) (int, error) {
	if u == nil {
		return 0, errors.New("надостаточно данных для создания пользователя")
	}
	i := len(s.Store) + 1
	u.ID = i
	s.Store[i] = u
	return u.ID, nil
}

type MakeFriendsRequest struct {
	SourceId int `json:"source_id"`
	TargetId int `json:"target_id"`
}

func (s *Storage) Add(u MakeFriendsRequest) (string, error) {
	sourceUser := s.Store[u.SourceId]
	if sourceUser == nil {
		return "", errors.New("первый пользователь запроса не найлен")
	}

	targetUser := s.Store[u.TargetId]
	if targetUser == nil {
		return "", errors.New("второй пользователь запроса не найлен")
	}

	sourceUser.Friends = append(sourceUser.Friends, targetUser.ID)
	targetUser.Friends = append(targetUser.Friends, sourceUser.ID)

	ans := string(sourceUser.Name + "и" + targetUser.Name + "теперь друзья")
	return ans, nil
}

type DelieteUserRequest struct {
	TargetId int `json:"target_id"`
}

func (s *Storage) DelieteU(u DelieteUserRequest) (string, error) {
	user := s.Store[u.TargetId]
	if user == nil {
		return "", errors.New("пользователь запроса не найлен")
	}

	for i := range s.Store {
		a := Find(s.Store[i].Friends, u.TargetId)
		if a != len(s.Store[i].Friends) {

			copy(s.Store[i].Friends[a:], s.Store[i].Friends[a+1:])
			s.Store[i].Friends[len(s.Store[i].Friends)-1] = 0
			s.Store[i].Friends = s.Store[i].Friends[:len(s.Store[i].Friends)-1]
		}
	}
	ans := string("Имя удалённого пользователя" + user.Name)
	delete(s.Store, u.TargetId)
	return ans, nil
}

func Find(a []int, x int) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

func (s *Storage) GetU(u string) ([]*User, error) {
	abc, err := strconv.Atoi(u)
	if err != nil {
		return nil, errors.New("id пользователя введен неорректно")
	}

	user := s.Store[abc]
	if user == nil {
		return nil, errors.New("пользователь запроса не найлен")
	}

	a := []*User{}
	aa := len(user.Friends)
	bb := user.Friends
	for i := 0; i < aa; i++ {
		a = append(a, s.Store[bb[i]])
	}
	return a, nil
}

type NewAgeRequest struct {
	NewAge int `json:"new age"`
}

func (s *Storage) NewAgeU(u NewAgeRequest, idU string) (string, error) {
	abc, err := strconv.Atoi(idU)
	if err != nil {
		return "", errors.New("id пользователя введен неорректно")
	}

	user := s.Store[abc]
	if user == nil {
		return "", errors.New("пользователь запроса не найлен")
	}

	user.Age = u.NewAge
	a := "возраст пользователя успешно обновлён"
	return a, nil
}

func (s *Storage) AllUsers() []*User {
	friendsList := []*User{}
	for _, val := range s.Store {
		friendsList = append(friendsList, val)
	}
	log.Println(friendsList)
	return friendsList
}
