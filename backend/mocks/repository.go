package mocks

import (
	"context"
	"reflect"

	"github.com/mujhtech/b0/database/models"
	"go.uber.org/mock/gomock"
)

// MockProjectRepository is a mock of AppRepository interface
type MockProjectRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProjectRepositoryMockRecorder
}

// MockProjectRepositoryMockRecorder is the mock recorder for MockProjectRepository
type MockProjectRepositoryMockRecorder struct {
	mock *MockProjectRepository
}

// NewMockProjectRepository creates a new mock instance
func NewMockProjectRepository(ctrl *gomock.Controller) *MockProjectRepository {
	mock := &MockProjectRepository{ctrl: ctrl}
	mock.recorder = &MockProjectRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProjectRepository) EXPECT() *MockProjectRepositoryMockRecorder {
	return m.recorder
}

// CreateProject mocks base method
func (m *MockProjectRepository) CreateProject(arg0 context.Context, arg1 *models.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockProjectRepositoryMockRecorder) CreateProject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockProjectRepository)(nil).CreateProject), arg0, arg1)
}

// UpdateProject mocks base method
func (m *MockProjectRepository) UpdateProject(arg0 context.Context, arg1 *models.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProject indicates an expected call of UpdateProject.
func (mr *MockProjectRepositoryMockRecorder) UpdateProject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockProjectRepository)(nil).UpdateProject), arg0, arg1)
}

// FindProjectByOwnerID mocks base method
func (m *MockProjectRepository) FindProjectByOwnerID(arg0 context.Context, arg1 string) ([]*models.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProjectByOwnerID", arg0, arg1)
	ret0, _ := ret[0].([]*models.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProjectByOwnerID indicates an expected call of FindProjectByOwnerID.
func (mr *MockProjectRepositoryMockRecorder) FindProjectByOwnerID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProjectByOwnerID", reflect.TypeOf((*MockProjectRepository)(nil).FindProjectByOwnerID), arg0, arg1)
}

// FindProjectByID mocks base method
func (m *MockProjectRepository) FindProjectByID(arg0 context.Context, arg1 string) (*models.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProjectByID", arg0, arg1)
	ret0, _ := ret[0].(*models.Project)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

// FindProjectByID indicates an expected call of FindProjectByID.
func (mr *MockProjectRepositoryMockRecorder) FindProjectByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProjectByID", reflect.TypeOf((*MockProjectRepository)(nil).FindProjectByID), arg0, arg1)
}

// FindProjectBySlug mocks base method
func (m *MockProjectRepository) FindProjectBySlug(arg0 context.Context, arg1 string) (*models.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProjectBySlug", arg0, arg1)
	ret0, _ := ret[0].(*models.Project)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

// FindProjectBySlug indicates an expected call of FindProjectBySlug.
func (mr *MockProjectRepositoryMockRecorder) FindProjectBySlug(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProjectBySlug", reflect.TypeOf((*MockProjectRepository)(nil).FindProjectByID), arg0, arg1)
}

// DeleteProject mocks base method
func (m *MockProjectRepository) DeleteProject(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProject indicates an expected call of DeleteProject.
func (mr *MockProjectRepositoryMockRecorder) DeleteProject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockProjectRepository)(nil).DeleteProject), arg0, arg1)
}

// CountByOwnerID mocks base method
func (m *MockProjectRepository) CountByOwnerID(arg0 context.Context, arg1 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountByOwnerID", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountByOwnerID indicates an expected call of CountByOwnerID.
func (mr *MockProjectRepositoryMockRecorder) CountByOwnerID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountByOwnerID", reflect.TypeOf((*MockProjectRepository)(nil).CountByOwnerID), arg0, arg1)
}

// MockEndpointRepository is a mock of AppRepository interface
type MockEndpointRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEndpointRepositoryMockRecorder
}

// MockEndpointRepositoryMockRecorder is the mock recorder for MockProjectRepository
type MockEndpointRepositoryMockRecorder struct {
	mock *MockEndpointRepository
}

// NewMockEndpointRepository creates a new mock instance
func NewMockEndpointRepository(ctrl *gomock.Controller) *MockEndpointRepository {
	mock := &MockEndpointRepository{ctrl: ctrl}
	mock.recorder = &MockEndpointRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEndpointRepository) EXPECT() *MockEndpointRepositoryMockRecorder {
	return m.recorder
}

// CreateEndpoint mocks base method
func (m *MockEndpointRepository) CreateEndpoint(arg0 context.Context, arg1 *models.Endpoint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEndpoint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEndpoint indicates an expected call of CreateEndpoint.
func (mr *MockEndpointRepositoryMockRecorder) CreateEndpoint(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEndpoint", reflect.TypeOf((*MockEndpointRepository)(nil).CreateEndpoint), arg0, arg1)
}

// UpdateEndpoint mocks base method
func (m *MockEndpointRepository) UpdateEndpoint(arg0 context.Context, id string, arg1 *models.Endpoint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEndpoint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEndpoint indicates an expected call of UpdateEndpoint.
func (mr *MockEndpointRepositoryMockRecorder) UpdateEndpoint(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEndpoint", reflect.TypeOf((*MockEndpointRepository)(nil).UpdateEndpoint), arg0, arg1, arg2)
}

// FindEndpointByID mocks base method
func (m *MockEndpointRepository) FindEndpointByID(arg0 context.Context, arg1 string) (*models.Endpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEndpointByID", arg0, arg1)
	ret0, _ := ret[0].(*models.Endpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEndpointByID indicates an expected call of FindEndpointByID.
func (mr *MockEndpointRepositoryMockRecorder) FindEndpointByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEndpointByID", reflect.TypeOf((*MockEndpointRepository)(nil).FindEndpointByID), arg0, arg1)
}

// FindEndpointByProjectID mocks base method
func (m *MockEndpointRepository) FindEndpointByProjectID(arg0 context.Context, arg1 string) ([]*models.Endpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEndpointByProjectID", arg0, arg1)
	ret0, _ := ret[0].([]*models.Endpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEndpointByProjectID indicates an expected call of FindEndpointByProjectID.
func (mr *MockEndpointRepositoryMockRecorder) FindEndpointByProjectID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEndpointByProjectID", reflect.TypeOf((*MockEndpointRepository)(nil).FindEndpointByProjectID), arg0, arg1)
}

// FindEndpointByOwnerID mocks base method
func (m *MockEndpointRepository) FindEndpointByOwnerID(arg0 context.Context, arg1 string) ([]*models.Endpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEndpointByOwnerID", arg0, arg1)
	ret0, _ := ret[0].([]*models.Endpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEndpointByOwnerID indicates an expected call of FindEndpointByOwnerID.
func (mr *MockEndpointRepositoryMockRecorder) FindEndpointByOwnerID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEndpointByOwnerID", reflect.TypeOf((*MockEndpointRepository)(nil).FindEndpointByOwnerID), arg0, arg1)
}

// DeleteEndpoint mocks base method
func (m *MockEndpointRepository) DeleteEndpoint(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEndpoint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEndpoint indicates an expected call of DeleteEndpoint.
func (mr *MockEndpointRepositoryMockRecorder) DeleteEndpoint(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEndpoint", reflect.TypeOf((*MockEndpointRepository)(nil).DeleteEndpoint), arg0, arg1)
}

// MockUserRepository is a mock of AppRepository interface
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method
func (m *MockUserRepository) CreateUser(arg0 context.Context, arg1 *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), arg0, arg1)
}

// UpdateUser mocks base method
func (m *MockUserRepository) UpdateUser(arg0 context.Context, arg1 *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserRepositoryMockRecorder) UpdateUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserRepository)(nil).UpdateUser), arg0, arg1)
}

// FindUserByEmail mocks base method
func (m *MockUserRepository) FindUserByEmail(arg0 context.Context, arg1 any) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockUserRepositoryMockRecorder) FindUserByEmail(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindUserByEmail), arg0, arg1)
}

// FindUserByID mocks base method
func (m *MockUserRepository) FindUserByID(arg0 context.Context, arg1 any) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByID", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByID indicates an expected call of FindUserByID.
func (mr *MockUserRepositoryMockRecorder) FindUserByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockUserRepository)(nil).FindUserByID), arg0, arg1)
}
