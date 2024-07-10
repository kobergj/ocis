// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	connector "github.com/owncloud/ocis/v2/services/collaboration/pkg/connector"

	fileinfo "github.com/owncloud/ocis/v2/services/collaboration/pkg/connector/fileinfo"

	io "io"

	mock "github.com/stretchr/testify/mock"
)

// FileConnectorService is an autogenerated mock type for the FileConnectorService type
type FileConnectorService struct {
	mock.Mock
}

type FileConnectorService_Expecter struct {
	mock *mock.Mock
}

func (_m *FileConnectorService) EXPECT() *FileConnectorService_Expecter {
	return &FileConnectorService_Expecter{mock: &_m.Mock}
}

// CheckFileInfo provides a mock function with given fields: ctx
func (_m *FileConnectorService) CheckFileInfo(ctx context.Context) (fileinfo.FileInfo, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CheckFileInfo")
	}

	var r0 fileinfo.FileInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (fileinfo.FileInfo, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) fileinfo.FileInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fileinfo.FileInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileConnectorService_CheckFileInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckFileInfo'
type FileConnectorService_CheckFileInfo_Call struct {
	*mock.Call
}

// CheckFileInfo is a helper method to define mock.On call
//   - ctx context.Context
func (_e *FileConnectorService_Expecter) CheckFileInfo(ctx interface{}) *FileConnectorService_CheckFileInfo_Call {
	return &FileConnectorService_CheckFileInfo_Call{Call: _e.mock.On("CheckFileInfo", ctx)}
}

func (_c *FileConnectorService_CheckFileInfo_Call) Run(run func(ctx context.Context)) *FileConnectorService_CheckFileInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *FileConnectorService_CheckFileInfo_Call) Return(_a0 fileinfo.FileInfo, _a1 error) *FileConnectorService_CheckFileInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileConnectorService_CheckFileInfo_Call) RunAndReturn(run func(context.Context) (fileinfo.FileInfo, error)) *FileConnectorService_CheckFileInfo_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteFile provides a mock function with given fields: ctx, lockID
func (_m *FileConnectorService) DeleteFile(ctx context.Context, lockID string) (string, error) {
	ret := _m.Called(ctx, lockID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteFile")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, lockID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, lockID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, lockID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileConnectorService_DeleteFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteFile'
type FileConnectorService_DeleteFile_Call struct {
	*mock.Call
}

// DeleteFile is a helper method to define mock.On call
//   - ctx context.Context
//   - lockID string
func (_e *FileConnectorService_Expecter) DeleteFile(ctx interface{}, lockID interface{}) *FileConnectorService_DeleteFile_Call {
	return &FileConnectorService_DeleteFile_Call{Call: _e.mock.On("DeleteFile", ctx, lockID)}
}

func (_c *FileConnectorService_DeleteFile_Call) Run(run func(ctx context.Context, lockID string)) *FileConnectorService_DeleteFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *FileConnectorService_DeleteFile_Call) Return(_a0 string, _a1 error) *FileConnectorService_DeleteFile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileConnectorService_DeleteFile_Call) RunAndReturn(run func(context.Context, string) (string, error)) *FileConnectorService_DeleteFile_Call {
	_c.Call.Return(run)
	return _c
}

// GetLock provides a mock function with given fields: ctx
func (_m *FileConnectorService) GetLock(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetLock")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileConnectorService_GetLock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLock'
type FileConnectorService_GetLock_Call struct {
	*mock.Call
}

// GetLock is a helper method to define mock.On call
//   - ctx context.Context
func (_e *FileConnectorService_Expecter) GetLock(ctx interface{}) *FileConnectorService_GetLock_Call {
	return &FileConnectorService_GetLock_Call{Call: _e.mock.On("GetLock", ctx)}
}

func (_c *FileConnectorService_GetLock_Call) Run(run func(ctx context.Context)) *FileConnectorService_GetLock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *FileConnectorService_GetLock_Call) Return(_a0 string, _a1 error) *FileConnectorService_GetLock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileConnectorService_GetLock_Call) RunAndReturn(run func(context.Context) (string, error)) *FileConnectorService_GetLock_Call {
	_c.Call.Return(run)
	return _c
}

// Lock provides a mock function with given fields: ctx, lockID, oldLockID
func (_m *FileConnectorService) Lock(ctx context.Context, lockID string, oldLockID string) (string, error) {
	ret := _m.Called(ctx, lockID, oldLockID)

	if len(ret) == 0 {
		panic("no return value specified for Lock")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (string, error)); ok {
		return rf(ctx, lockID, oldLockID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, lockID, oldLockID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, lockID, oldLockID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileConnectorService_Lock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Lock'
type FileConnectorService_Lock_Call struct {
	*mock.Call
}

// Lock is a helper method to define mock.On call
//   - ctx context.Context
//   - lockID string
//   - oldLockID string
func (_e *FileConnectorService_Expecter) Lock(ctx interface{}, lockID interface{}, oldLockID interface{}) *FileConnectorService_Lock_Call {
	return &FileConnectorService_Lock_Call{Call: _e.mock.On("Lock", ctx, lockID, oldLockID)}
}

func (_c *FileConnectorService_Lock_Call) Run(run func(ctx context.Context, lockID string, oldLockID string)) *FileConnectorService_Lock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *FileConnectorService_Lock_Call) Return(_a0 string, _a1 error) *FileConnectorService_Lock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileConnectorService_Lock_Call) RunAndReturn(run func(context.Context, string, string) (string, error)) *FileConnectorService_Lock_Call {
	_c.Call.Return(run)
	return _c
}

// PutRelativeFileRelative provides a mock function with given fields: ctx, ccs, stream, streamLength, target
func (_m *FileConnectorService) PutRelativeFileRelative(ctx context.Context, ccs connector.ContentConnectorService, stream io.Reader, streamLength int64, target string) (*connector.PutRelativeResponse, *connector.PutRelativeHeaders, error) {
	ret := _m.Called(ctx, ccs, stream, streamLength, target)

	if len(ret) == 0 {
		panic("no return value specified for PutRelativeFileRelative")
	}

	var r0 *connector.PutRelativeResponse
	var r1 *connector.PutRelativeHeaders
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, connector.ContentConnectorService, io.Reader, int64, string) (*connector.PutRelativeResponse, *connector.PutRelativeHeaders, error)); ok {
		return rf(ctx, ccs, stream, streamLength, target)
	}
	if rf, ok := ret.Get(0).(func(context.Context, connector.ContentConnectorService, io.Reader, int64, string) *connector.PutRelativeResponse); ok {
		r0 = rf(ctx, ccs, stream, streamLength, target)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connector.PutRelativeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, connector.ContentConnectorService, io.Reader, int64, string) *connector.PutRelativeHeaders); ok {
		r1 = rf(ctx, ccs, stream, streamLength, target)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*connector.PutRelativeHeaders)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, connector.ContentConnectorService, io.Reader, int64, string) error); ok {
		r2 = rf(ctx, ccs, stream, streamLength, target)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FileConnectorService_PutRelativeFileRelative_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutRelativeFileRelative'
type FileConnectorService_PutRelativeFileRelative_Call struct {
	*mock.Call
}

// PutRelativeFileRelative is a helper method to define mock.On call
//   - ctx context.Context
//   - ccs connector.ContentConnectorService
//   - stream io.Reader
//   - streamLength int64
//   - target string
func (_e *FileConnectorService_Expecter) PutRelativeFileRelative(ctx interface{}, ccs interface{}, stream interface{}, streamLength interface{}, target interface{}) *FileConnectorService_PutRelativeFileRelative_Call {
	return &FileConnectorService_PutRelativeFileRelative_Call{Call: _e.mock.On("PutRelativeFileRelative", ctx, ccs, stream, streamLength, target)}
}

func (_c *FileConnectorService_PutRelativeFileRelative_Call) Run(run func(ctx context.Context, ccs connector.ContentConnectorService, stream io.Reader, streamLength int64, target string)) *FileConnectorService_PutRelativeFileRelative_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(connector.ContentConnectorService), args[2].(io.Reader), args[3].(int64), args[4].(string))
	})
	return _c
}

func (_c *FileConnectorService_PutRelativeFileRelative_Call) Return(_a0 *connector.PutRelativeResponse, _a1 *connector.PutRelativeHeaders, _a2 error) *FileConnectorService_PutRelativeFileRelative_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *FileConnectorService_PutRelativeFileRelative_Call) RunAndReturn(run func(context.Context, connector.ContentConnectorService, io.Reader, int64, string) (*connector.PutRelativeResponse, *connector.PutRelativeHeaders, error)) *FileConnectorService_PutRelativeFileRelative_Call {
	_c.Call.Return(run)
	return _c
}

// PutRelativeFileSuggested provides a mock function with given fields: ctx, ccs, stream, streamLength, target
func (_m *FileConnectorService) PutRelativeFileSuggested(ctx context.Context, ccs connector.ContentConnectorService, stream io.Reader, streamLength int64, target string) (*connector.PutRelativeResponse, error) {
	ret := _m.Called(ctx, ccs, stream, streamLength, target)

	if len(ret) == 0 {
		panic("no return value specified for PutRelativeFileSuggested")
	}

	var r0 *connector.PutRelativeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, connector.ContentConnectorService, io.Reader, int64, string) (*connector.PutRelativeResponse, error)); ok {
		return rf(ctx, ccs, stream, streamLength, target)
	}
	if rf, ok := ret.Get(0).(func(context.Context, connector.ContentConnectorService, io.Reader, int64, string) *connector.PutRelativeResponse); ok {
		r0 = rf(ctx, ccs, stream, streamLength, target)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connector.PutRelativeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, connector.ContentConnectorService, io.Reader, int64, string) error); ok {
		r1 = rf(ctx, ccs, stream, streamLength, target)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileConnectorService_PutRelativeFileSuggested_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutRelativeFileSuggested'
type FileConnectorService_PutRelativeFileSuggested_Call struct {
	*mock.Call
}

// PutRelativeFileSuggested is a helper method to define mock.On call
//   - ctx context.Context
//   - ccs connector.ContentConnectorService
//   - stream io.Reader
//   - streamLength int64
//   - target string
func (_e *FileConnectorService_Expecter) PutRelativeFileSuggested(ctx interface{}, ccs interface{}, stream interface{}, streamLength interface{}, target interface{}) *FileConnectorService_PutRelativeFileSuggested_Call {
	return &FileConnectorService_PutRelativeFileSuggested_Call{Call: _e.mock.On("PutRelativeFileSuggested", ctx, ccs, stream, streamLength, target)}
}

func (_c *FileConnectorService_PutRelativeFileSuggested_Call) Run(run func(ctx context.Context, ccs connector.ContentConnectorService, stream io.Reader, streamLength int64, target string)) *FileConnectorService_PutRelativeFileSuggested_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(connector.ContentConnectorService), args[2].(io.Reader), args[3].(int64), args[4].(string))
	})
	return _c
}

func (_c *FileConnectorService_PutRelativeFileSuggested_Call) Return(_a0 *connector.PutRelativeResponse, _a1 error) *FileConnectorService_PutRelativeFileSuggested_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileConnectorService_PutRelativeFileSuggested_Call) RunAndReturn(run func(context.Context, connector.ContentConnectorService, io.Reader, int64, string) (*connector.PutRelativeResponse, error)) *FileConnectorService_PutRelativeFileSuggested_Call {
	_c.Call.Return(run)
	return _c
}

// RefreshLock provides a mock function with given fields: ctx, lockID
func (_m *FileConnectorService) RefreshLock(ctx context.Context, lockID string) (string, error) {
	ret := _m.Called(ctx, lockID)

	if len(ret) == 0 {
		panic("no return value specified for RefreshLock")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, lockID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, lockID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, lockID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileConnectorService_RefreshLock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RefreshLock'
type FileConnectorService_RefreshLock_Call struct {
	*mock.Call
}

// RefreshLock is a helper method to define mock.On call
//   - ctx context.Context
//   - lockID string
func (_e *FileConnectorService_Expecter) RefreshLock(ctx interface{}, lockID interface{}) *FileConnectorService_RefreshLock_Call {
	return &FileConnectorService_RefreshLock_Call{Call: _e.mock.On("RefreshLock", ctx, lockID)}
}

func (_c *FileConnectorService_RefreshLock_Call) Run(run func(ctx context.Context, lockID string)) *FileConnectorService_RefreshLock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *FileConnectorService_RefreshLock_Call) Return(_a0 string, _a1 error) *FileConnectorService_RefreshLock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileConnectorService_RefreshLock_Call) RunAndReturn(run func(context.Context, string) (string, error)) *FileConnectorService_RefreshLock_Call {
	_c.Call.Return(run)
	return _c
}

// RenameFile provides a mock function with given fields: ctx, lockID, target
func (_m *FileConnectorService) RenameFile(ctx context.Context, lockID string, target string) (string, error) {
	ret := _m.Called(ctx, lockID, target)

	if len(ret) == 0 {
		panic("no return value specified for RenameFile")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (string, error)); ok {
		return rf(ctx, lockID, target)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, lockID, target)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, lockID, target)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileConnectorService_RenameFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RenameFile'
type FileConnectorService_RenameFile_Call struct {
	*mock.Call
}

// RenameFile is a helper method to define mock.On call
//   - ctx context.Context
//   - lockID string
//   - target string
func (_e *FileConnectorService_Expecter) RenameFile(ctx interface{}, lockID interface{}, target interface{}) *FileConnectorService_RenameFile_Call {
	return &FileConnectorService_RenameFile_Call{Call: _e.mock.On("RenameFile", ctx, lockID, target)}
}

func (_c *FileConnectorService_RenameFile_Call) Run(run func(ctx context.Context, lockID string, target string)) *FileConnectorService_RenameFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *FileConnectorService_RenameFile_Call) Return(_a0 string, _a1 error) *FileConnectorService_RenameFile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileConnectorService_RenameFile_Call) RunAndReturn(run func(context.Context, string, string) (string, error)) *FileConnectorService_RenameFile_Call {
	_c.Call.Return(run)
	return _c
}

// UnLock provides a mock function with given fields: ctx, lockID
func (_m *FileConnectorService) UnLock(ctx context.Context, lockID string) (string, error) {
	ret := _m.Called(ctx, lockID)

	if len(ret) == 0 {
		panic("no return value specified for UnLock")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, lockID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, lockID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, lockID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileConnectorService_UnLock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnLock'
type FileConnectorService_UnLock_Call struct {
	*mock.Call
}

// UnLock is a helper method to define mock.On call
//   - ctx context.Context
//   - lockID string
func (_e *FileConnectorService_Expecter) UnLock(ctx interface{}, lockID interface{}) *FileConnectorService_UnLock_Call {
	return &FileConnectorService_UnLock_Call{Call: _e.mock.On("UnLock", ctx, lockID)}
}

func (_c *FileConnectorService_UnLock_Call) Run(run func(ctx context.Context, lockID string)) *FileConnectorService_UnLock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *FileConnectorService_UnLock_Call) Return(_a0 string, _a1 error) *FileConnectorService_UnLock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileConnectorService_UnLock_Call) RunAndReturn(run func(context.Context, string) (string, error)) *FileConnectorService_UnLock_Call {
	_c.Call.Return(run)
	return _c
}

// NewFileConnectorService creates a new instance of FileConnectorService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFileConnectorService(t interface {
	mock.TestingT
	Cleanup(func())
}) *FileConnectorService {
	mock := &FileConnectorService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
