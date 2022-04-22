// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	domain "github.com/odpf/stencil/server/domain"

	mock "github.com/stretchr/testify/mock"
)

// GetSchemaData is an autogenerated mock type for the getSchemaData type
type GetSchemaData struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1, _a2
func (_m *GetSchemaData) Execute(_a0 http.ResponseWriter, _a1 *http.Request, _a2 map[string]string) (*domain.Metadata, []byte, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *domain.Metadata
	if rf, ok := ret.Get(0).(func(http.ResponseWriter, *http.Request, map[string]string) *domain.Metadata); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Metadata)
		}
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(http.ResponseWriter, *http.Request, map[string]string) []byte); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(http.ResponseWriter, *http.Request, map[string]string) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
