// Copyright 2020 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// Code generated by mockery v2.1.0. DO NOT EDIT.

package mocks

import (
	context "context"

	jwt "github.com/mendersoftware/useradm/jwt"
	mock "github.com/stretchr/testify/mock"

	model "github.com/mendersoftware/useradm/model"

	oid "github.com/mendersoftware/go-lib-micro/mongo/oid"
)

// DataStore is an autogenerated mock type for the DataStore type
type DataStore struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, u
func (_m *DataStore) CreateUser(ctx context.Context, u *model.User) error {
	ret := _m.Called(ctx, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) error); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteToken provides a mock function with given fields: ctx, tokenID
func (_m *DataStore) DeleteToken(ctx context.Context, tokenID oid.ObjectID) error {
	ret := _m.Called(ctx, tokenID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, oid.ObjectID) error); ok {
		r0 = rf(ctx, tokenID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTokens provides a mock function with given fields: ctx
func (_m *DataStore) DeleteTokens(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTokensByUserId provides a mock function with given fields: ctx, userId
func (_m *DataStore) DeleteTokensByUserId(ctx context.Context, userId string) error {
	ret := _m.Called(ctx, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *DataStore) DeleteUser(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSettings provides a mock function with given fields: ctx
func (_m *DataStore) GetSettings(ctx context.Context) (map[string]interface{}, error) {
	ret := _m.Called(ctx)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(context.Context) map[string]interface{}); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokenById provides a mock function with given fields: ctx, id
func (_m *DataStore) GetTokenById(ctx context.Context, id oid.ObjectID) (*jwt.Token, error) {
	ret := _m.Called(ctx, id)

	var r0 *jwt.Token
	if rf, ok := ret.Get(0).(func(context.Context, oid.ObjectID) *jwt.Token); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, oid.ObjectID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *DataStore) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	ret := _m.Called(ctx, email)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserById provides a mock function with given fields: ctx, id
func (_m *DataStore) GetUserById(ctx context.Context, id string) (*model.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: ctx, fltr
func (_m *DataStore) GetUsers(ctx context.Context, fltr model.UserFilter) ([]model.User, error) {
	ret := _m.Called(ctx, fltr)

	var r0 []model.User
	if rf, ok := ret.Get(0).(func(context.Context, model.UserFilter) []model.User); ok {
		r0 = rf(ctx, fltr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.UserFilter) error); ok {
		r1 = rf(ctx, fltr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ping provides a mock function with given fields: ctx
func (_m *DataStore) Ping(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveSettings provides a mock function with given fields: ctx, s
func (_m *DataStore) SaveSettings(ctx context.Context, s map[string]interface{}) error {
	ret := _m.Called(ctx, s)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) error); ok {
		r0 = rf(ctx, s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveToken provides a mock function with given fields: ctx, token
func (_m *DataStore) SaveToken(ctx context.Context, token *jwt.Token) error {
	ret := _m.Called(ctx, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *jwt.Token) error); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: ctx, id, u
func (_m *DataStore) UpdateUser(ctx context.Context, id string, u *model.UserUpdate) error {
	ret := _m.Called(ctx, id, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *model.UserUpdate) error); ok {
		r0 = rf(ctx, id, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
