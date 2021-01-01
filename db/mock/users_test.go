// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chinmobi/gin-mvc/model"
)

var usersDB *UsersDB

func setUp() *UsersDB {
	if usersDB == nil {
		usersDB = newUsersDB()
	}
	return usersDB
}

func tearDown() {
	if usersDB != nil {
		usersDB.clear()
	}
}


func newUser(nickname, email string) *model.UserEntity {
	user := &model.UserEntity{
		Nickname: nickname,
		Email: email,
	}
	return user
}

var testUsers = []struct {
	nickname    string
	email       string
	wantID      uint32
	created     bool
}{
	{"nicknameA", "", USER_ID_INIT_VALUE + 1, true},
	{"nicknameB", "", 0, false},
	{"", "emailA", USER_ID_INIT_VALUE + 2, true},
	{"", "emailB", 0, false},
	{"nicknameC", "emailC", USER_ID_INIT_VALUE + 3, true},
	{"nicknameD", "emailD", 0, false},
}

func setUpTestUsers(db *UsersDB) *UsersDB {
	for _, user := range testUsers {
		if user.created {
			db.CreateUser(newUser(user.nickname, user.email))
		}
	}
	return db
}

func getUserByID(users []model.UserEntity, uid uint32) (*model.UserEntity, bool) {
	for i, cnt := 0, len(users); i < cnt; i++ {
		if users[i].ID == uid {
			return &users[i], true
		}
	}
	return nil, false
}

// User's test cases

func TestCreateUser(t *testing.T) {
	var tests = []struct {
		nickname    string
		email       string
		wantID      uint32
		errOccurred bool
	}{
		{"nicknameA", "", USER_ID_INIT_VALUE + 1, false},
		{"nicknameA", "", 0, true},
		{"", "emailA", USER_ID_INIT_VALUE + 2, false},
		{"", "emailA", 0, true},
		{"nicknameB", "emailB", USER_ID_INIT_VALUE + 3, false},
		{"nicknameB", "emailB", 0, true},
		{"", "", 0, true},
	}

	assert := assert.New(t)

	db := setUp()
	defer tearDown()

	for _, test := range tests {
		user, err := db.CreateUser(newUser(test.nickname, test.email))
		if err != nil && !test.errOccurred {
			t.Fatalf("err: %v\n", err)
		}

		assert.Equal(test.nickname, user.Nickname)
		assert.Equal(test.email, user.Email)
		assert.Equal(test.wantID, user.ID)
	}
}

func TestFindAllUsersEmpty(t *testing.T) {
	assert := assert.New(t)

	db := setUp()
	defer tearDown()

	results, err := db.FindAllUsers()
	if err != nil {
			t.Fatalf("err: %v\n", err)
	}

	assert.Empty(results)
}

func TestFindAllUsers(t *testing.T) {
	assert := assert.New(t)

	db := setUpTestUsers(setUp())
	defer tearDown()

	results, err := db.FindAllUsers()
	if err != nil {
			t.Fatalf("err: %v\n", err)
	}

	assert.NotEmpty(results)

	count := 0
	for _, test := range testUsers {
		if !test.created {
			continue
		}

		user, found := getUserByID(results, test.wantID)
		if !found {
			t.Errorf("User not found by: %d\n", test.wantID)
		}

		count++

		assert.Equal(test.nickname, user.Nickname)
		assert.Equal(test.email, user.Email)
		assert.Equal(test.wantID, user.ID)
	}

	assert.Equal(len(results), count)
}

func TestFindUserByID(t *testing.T) {
	assert := assert.New(t)

	db := setUpTestUsers(setUp())
	defer tearDown()

	for _, test := range testUsers {
		if !test.created {
			continue
		}

		user, err := db.FindUserByID(test.wantID)
		if err != nil {
			t.Errorf("err: %v\n", err)
			continue
		}

		assert.Equal(test.nickname, user.Nickname)
		assert.Equal(test.email, user.Email)
		assert.Equal(test.wantID, user.ID)
	}
}

func TestFindUserByNickname(t *testing.T) {
	assert := assert.New(t)

	db := setUpTestUsers(setUp())
	defer tearDown()

	for _, test := range testUsers {
		user, err := db.FindUserByNickname(test.nickname)
		if err != nil {
			t.Errorf("err: %v\n", err)
			continue
		}

		if !test.created || test.nickname == "" {
			assert.Nil(user)
			continue
		}

		if user == nil {
			assert.NotNil(user, test.nickname)
			continue
		}

		assert.Equal(test.nickname, user.Nickname)
		assert.Equal(test.email, user.Email)
		assert.Equal(test.wantID, user.ID)
	}
}

func TestFindUserByEmail(t *testing.T) {
	assert := assert.New(t)

	db := setUpTestUsers(setUp())
	defer tearDown()

	for _, test := range testUsers {
		user, err := db.FindUserByEmail(test.email)
		if err != nil {
			t.Errorf("err: %v\n", err)
			continue
		}

		if !test.created || test.email == "" {
			assert.Nil(user)
			continue
		}

		if user == nil {
			assert.NotNil(user, test.email)
			continue
		}

		assert.Equal(test.nickname, user.Nickname)
		assert.Equal(test.email, user.Email)
		assert.Equal(test.wantID, user.ID)
	}
}

func setUpUpdates() []*model.UserEntity {
	updates := make([]*model.UserEntity, 0, 3)

	for _, test := range testUsers {
		if !test.created {
			continue
		}

		nickname := test.nickname
		nickname += "updated"

		email := test.email
		if email != "" {
			email += "updated"
		}

		u := newUser(nickname, email)

		u.ID = test.wantID

		updates = append(updates, u)
	}

	return updates
}

func TestUpdateUser(t *testing.T) {
	updates := setUpUpdates()

	assert := assert.New(t)

	db := setUpTestUsers(setUp())
	defer tearDown()

	for _, u := range updates {
		_, err := db.UpdateUser(u.ID, u)
		if err != nil {
			t.Errorf("err: %v\n", err)
			continue
		}
	}

	for _, test := range testUsers {
		if !test.created {
			continue
		}

		nickname := test.nickname
		nickname += "updated"

		email := test.email
		if email != "" {
			email += "updated"
		}

		// Find by old values

		if test.nickname != "" {
			user, err := db.FindUserByNickname(test.nickname)
			if err != nil {
				t.Errorf("err: %v\n", err)
				continue
			}

			assert.Nil(user)
		}

		if test.email != "" {
			user, err := db.FindUserByEmail(test.email)
			if err != nil {
				t.Errorf("err: %v\n", err)
				continue
			}

			assert.Nil(user)
		}

		// Find by new values

		if nickname != "" {
			user, err := db.FindUserByNickname(nickname)
			if err != nil {
				t.Errorf("err: %v\n", err)
				continue
			}

			assert.NotNil(user)

			assert.Equal(nickname, user.Nickname)
			assert.Equal(email, user.Email)
			assert.Equal(test.wantID, user.ID)
		}

		if email != "" {
			user, err := db.FindUserByEmail(email)
			if err != nil {
				t.Errorf("err: %v\n", err)
				continue
			}

			assert.NotNil(user)

			assert.Equal(nickname, user.Nickname)
			assert.Equal(email, user.Email)
			assert.Equal(test.wantID, user.ID)
		}
	}
}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	db := setUpTestUsers(setUp())
	defer tearDown()

	for _, test := range testUsers {
		if !test.created {
			continue
		}

		if _, err := db.DeleteUser(test.wantID); err != nil {
			t.Errorf("err: %v\n", err)
			continue
		}

		if test.nickname != "" {
			user, err := db.FindUserByNickname(test.nickname)
			if err != nil {
				t.Errorf("err: %v\n", err)
				continue
			}

			assert.Nil(user)
		}

		if test.email != "" {
			user, err := db.FindUserByEmail(test.email)
			if err != nil {
				t.Errorf("err: %v\n", err)
				continue
			}

			assert.Nil(user)
		}
	}
}
