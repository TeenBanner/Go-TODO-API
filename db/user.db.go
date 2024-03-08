package db

import "github.com/simple-rest-API/model"

type UserRecord struct {
	UserRecordId int
	Users        map[int]model.User
	UsersTasks   map[int]model.Task
}

/*func StartUserDB(record *TaskRecord) UserRecord {
	RecordInstance := make(map[int]model.User)
	return UserRecord{
		UserRecordId: 0,
		Users:        RecordInstance,
		UsersTasks:   record,
	}
}*/

func (UR *UserRecord) CreateUser(User *model.User) error {
	UR.UserRecordId++

	UR.Users[User.ID] = *User

	return nil
}

func (UR *UserRecord) GetALLUsers() (model.Users, error) {
	var user model.Users

	for _, v := range UR.Users {
		user = append(user, v)
	}

	return user, nil
}

func (UR *UserRecord) UpdateUser(ID int, User *model.User) error {
	if User == nil {
		return model.ErrUserDoNotExist
	}
	if _, ok := UR.Users[ID]; !ok {
		return model.ErrUserDoNotExist
	}

	UR.Users[ID] = *User

	return nil
}

func (UR *UserRecord) DeleteUser(ID int) error {
	if _, ok := UR.Users[ID]; !ok {
		return model.ErrUserDoNotExist
	}

	delete(UR.Users, ID)

	return nil
}

func (UR *UserRecord) GetUserByID(ID int) (model.User, error) {
	user, ok := UR.Users[ID]
	if !ok {
		return model.User{}, model.ErrUserDoNotExist
	}

	return user, nil
}
