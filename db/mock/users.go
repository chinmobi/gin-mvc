// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mock

// NOTE: The user dao is just for demo, could to be removed for real project.

import (
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/chinmobi/gin-mvc/errors"
	"github.com/chinmobi/gin-mvc/model"
)

const USER_ID_INIT_VALUE = 10000

type UsersDB struct {
	rw            sync.RWMutex
	entities      map[uint32]model.UserEntity
	nicknameIndex map[string]uint32
	emailIndex    map[string]uint32
	uids          uint32
}

func NewUsersDB() *UsersDB {
	users := &UsersDB{
		entities: make(map[uint32]model.UserEntity),
		nicknameIndex: make(map[string]uint32),
		emailIndex: make(map[string]uint32),
		uids: USER_ID_INIT_VALUE,
	}

	return users
}

func notFoundByUidError(uid uint32) error {
	return errors.NewErrNotFoundBy("User", "ID", strconv.FormatUint(uint64(uid), 10))
}

func notFoundByNicknameError(nickname string) error {
	return errors.NewErrNotFoundBy("User", "Nickname", nickname)
}

func notFoundByEmailError(email string) error {
	return errors.NewErrNotFoundBy("User", "Email", email)
}

func nicknameAlreadyExistsError(nickname string) error {
	return errors.NewErrAlreadyExistsFor("User", "Nickname", nickname)
}

func emailAlreadyExistsError(email string) error {
	return errors.NewErrAlreadyExistsFor("User", "Email", email)
}

func (users *UsersDB) CreateUser(u *model.UserEntity) (*model.UserEntity, error) {
	nickname := u.Nickname
	email := u.Email

	if nickname == "" && email == "" {
		return u, errors.NewErrLackOfParameter("Nickname", "Email")
	}

	users.rw.Lock()
	defer users.rw.Unlock()

	if email != "" {
		if _, exists := users.emailIndex[email]; exists {
			return u, emailAlreadyExistsError(email)
		}
	}

	if nickname != "" {
		if _, exists := users.nicknameIndex[nickname]; exists {
			return u, nicknameAlreadyExistsError(nickname)
		}
	}

	uid := users.generateID()
	u.ID = uid

	if email != "" {
		users.emailIndex[email] = uid
	}

	if nickname != "" {
		users.nicknameIndex[nickname] = uid
	}

	users.entities[uid] = *u

	return u, nil
}

func (users *UsersDB) FindAllUsers() ([]model.UserEntity, error) {
	users.rw.RLock()
	defer users.rw.RUnlock()

	results := make([]model.UserEntity, 0, len(users.entities))

	for _, e := range users.entities {
		results = append(results, e)
	}

	return results, nil
}

func (users *UsersDB) doFindUserByID(uid uint32) (*model.UserEntity, error) {
	if e, exists := users.entities[uid]; exists {
		return &e, nil
	}

	return nil, notFoundByUidError(uid)
}

func (users *UsersDB) FindUserByID(uid uint32) (*model.UserEntity, error) {
	users.rw.RLock()
	defer users.rw.RUnlock()

	return users.doFindUserByID(uid)
}

func (users *UsersDB) FindUserByNickname(nickname string) (*model.UserEntity, error) {
	users.rw.RLock()
	defer users.rw.RUnlock()

	if uid, exists := users.nicknameIndex[nickname]; exists {
		return users.doFindUserByID(uid)
	}

	return nil, nil
}

func (users *UsersDB) FindUserByEmail(email string) (*model.UserEntity, error) {
	users.rw.RLock()
	defer users.rw.RUnlock()

	if uid, exists := users.emailIndex[email]; exists {
		return users.doFindUserByID(uid)
	}

	return nil, nil
}

func (users *UsersDB) UpdateUser(uid uint32, u *model.UserEntity) (*model.UserEntity, error) {
	users.rw.Lock()
	defer users.rw.Unlock()

	e, exists := users.entities[uid]
	if !exists {
		return u, notFoundByUidError(uid)
	}

	modifies := 0

	nickname := u.Nickname
	if nickname != "" && nickname != e.Nickname {
		if _, exists := users.nicknameIndex[nickname]; exists {
			return u, nicknameAlreadyExistsError(nickname)
		}
		modifies++
	}

	email := u.Email
	if email != "" && email != e.Email {
		if _, exists := users.emailIndex[email]; exists {
			return u, emailAlreadyExistsError(email)
		}
		modifies++
	}

	users.replaceUserNickname(&e, nickname)
	users.replaceUserEmail(&e, email)

	if u.PasswordHash != "" && u.PasswordHash != e.PasswordHash {
		e.PasswordHash = u.PasswordHash
		modifies++
	}

	if modifies > 0 {
		users.entities[uid] = e
	}

	return &e, nil
}

func (users *UsersDB) DeleteUser(uid uint32) (uint64, error) {
	users.rw.Lock()
	defer users.rw.Unlock()

	e, exists := users.entities[uid]
	if !exists {
		return 0, notFoundByUidError(uid)
	}

	nickname := e.Nickname
	if nickname != "" {
		delete(users.nicknameIndex, nickname)
	}

	email := e.Email
	if email != "" {
		delete(users.emailIndex, email)
	}

	delete(users.entities, e.ID)

	return 1, nil
}

func (users *UsersDB) generateID() uint32 {
	val := atomic.AddUint32(&users.uids, 1)
	return val
}

func (users *UsersDB) replaceUserNickname(u *model.UserEntity, nickname string) {
	if nickname != "" && nickname != u.Nickname {
		if u.Nickname != "" {
			delete(users.nicknameIndex, u.Nickname)
		}

		u.Nickname = nickname

		users.nicknameIndex[nickname] = u.ID
	}
}

func (users *UsersDB) replaceUserEmail(u *model.UserEntity, email string) {
	if email != "" && email != u.Email {
		if u.Email != "" {
			delete(users.emailIndex, u.Email)
		}

		u.Email = email

		users.emailIndex[email] = u.ID
	}
}

func (users *UsersDB) clear() {
	users.nicknameIndex = make(map[string]uint32)
	users.emailIndex = make(map[string]uint32)
	users.entities = make(map[uint32]model.UserEntity)

	atomic.StoreUint32(&users.uids, USER_ID_INIT_VALUE)
}
