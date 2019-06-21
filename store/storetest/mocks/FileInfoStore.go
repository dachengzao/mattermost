// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/mattermost/mattermost-server/model"
import store "github.com/mattermost/mattermost-server/store"

// FileInfoStore is an autogenerated mock type for the FileInfoStore type
type FileInfoStore struct {
	mock.Mock
}

// AttachToPost provides a mock function with given fields: fileId, postId, creatorId
func (_m *FileInfoStore) AttachToPost(fileId string, postId string, creatorId string) store.StoreChannel {
	ret := _m.Called(fileId, postId, creatorId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string, string) store.StoreChannel); ok {
		r0 = rf(fileId, postId, creatorId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// ClearCaches provides a mock function with given fields:
func (_m *FileInfoStore) ClearCaches() {
	_m.Called()
}

// DeleteForPost provides a mock function with given fields: postId
func (_m *FileInfoStore) DeleteForPost(postId string) store.StoreChannel {
	ret := _m.Called(postId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(postId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *FileInfoStore) Get(id string) store.StoreChannel {
	ret := _m.Called(id)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetByPath provides a mock function with given fields: path
func (_m *FileInfoStore) GetByPath(path string) store.StoreChannel {
	ret := _m.Called(path)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetForPost provides a mock function with given fields: postId, readFromMaster, allowFromCache
func (_m *FileInfoStore) GetForPost(postId string, readFromMaster bool, allowFromCache bool) store.StoreChannel {
	ret := _m.Called(postId, readFromMaster, allowFromCache)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, bool, bool) store.StoreChannel); ok {
		r0 = rf(postId, readFromMaster, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetForUser provides a mock function with given fields: userId
func (_m *FileInfoStore) GetForUser(userId string) store.StoreChannel {
	ret := _m.Called(userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// InvalidateFileInfosForPostCache provides a mock function with given fields: postId
func (_m *FileInfoStore) InvalidateFileInfosForPostCache(postId string) {
	_m.Called(postId)
}

// PermanentDelete provides a mock function with given fields: fileId
func (_m *FileInfoStore) PermanentDelete(fileId string) store.StoreChannel {
	ret := _m.Called(fileId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(fileId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PermanentDeleteBatch provides a mock function with given fields: endTime, limit
func (_m *FileInfoStore) PermanentDeleteBatch(endTime int64, limit int64) store.StoreChannel {
	ret := _m.Called(endTime, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(int64, int64) store.StoreChannel); ok {
		r0 = rf(endTime, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PermanentDeleteByUser provides a mock function with given fields: userId
func (_m *FileInfoStore) PermanentDeleteByUser(userId string) store.StoreChannel {
	ret := _m.Called(userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Save provides a mock function with given fields: info
func (_m *FileInfoStore) Save(info *model.FileInfo) store.StoreChannel {
	ret := _m.Called(info)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.FileInfo) store.StoreChannel); ok {
		r0 = rf(info)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}