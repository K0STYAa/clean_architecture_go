// Code generated by mockery v2.11.0. DO NOT EDIT.

package mockStorage

import (
	context "forgoproject/pkg/type/context"
	contact "forgoproject/services/contact/internal/domain/contact"

	mock "github.com/stretchr/testify/mock"

	testing "testing"

	uuid "github.com/google/uuid"
)

// ContactInGroup is an autogenerated mock type for the ContactInGroup type
type ContactInGroup struct {
	mock.Mock
}

// AddContactsToGroup provides a mock function with given fields: ctx, groupID, contactIDs
func (_m *ContactInGroup) AddContactsToGroup(ctx context.Context, groupID uuid.UUID, contactIDs ...uuid.UUID) error {
	_va := make([]interface{}, len(contactIDs))
	for _i := range contactIDs {
		_va[_i] = contactIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, ...uuid.UUID) error); ok {
		r0 = rf(ctx, groupID, contactIDs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateContactIntoGroup provides a mock function with given fields: ctx, groupID, contacts
func (_m *ContactInGroup) CreateContactIntoGroup(ctx context.Context, groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	_va := make([]interface{}, len(contacts))
	for _i := range contacts {
		_va[_i] = contacts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*contact.Contact
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, ...*contact.Contact) []*contact.Contact); ok {
		r0 = rf(ctx, groupID, contacts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*contact.Contact)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, ...*contact.Contact) error); ok {
		r1 = rf(ctx, groupID, contacts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteContactFromGroup provides a mock function with given fields: ctx, groupID, contactID
func (_m *ContactInGroup) DeleteContactFromGroup(ctx context.Context, groupID uuid.UUID, contactID uuid.UUID) error {
	ret := _m.Called(ctx, groupID, contactID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) error); ok {
		r0 = rf(ctx, groupID, contactID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewContactInGroup creates a new instance of ContactInGroup. It also registers a cleanup function to assert the mocks expectations.
func NewContactInGroup(t testing.TB) *ContactInGroup {
	mock := &ContactInGroup{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}