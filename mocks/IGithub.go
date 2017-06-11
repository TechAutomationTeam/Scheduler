// Code generated by mockery v1.0.0
package mocks

import github "github.com/google/go-github/github"
import mock "github.com/stretchr/testify/mock"
import models "github.com/Golang-Coach/Scheduler/models"

// IGithub is an autogenerated mock type for the IGithub type
type IGithub struct {
	mock.Mock
}

// GetLastCommitInfo provides a mock function with given fields: owner, repositoryName
func (_m *IGithub) GetLastCommitInfo(owner string, repositoryName string) (*github.RepositoryCommit, error) {
	ret := _m.Called(owner, repositoryName)

	var r0 *github.RepositoryCommit
	if rf, ok := ret.Get(0).(func(string, string) *github.RepositoryCommit); ok {
		r0 = rf(owner, repositoryName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.RepositoryCommit)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(owner, repositoryName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPackageRepoInfo provides a mock function with given fields: owner, repositoryName
func (_m *IGithub) GetPackageRepoInfo(owner string, repositoryName string) (*models.RepositoryInfo, error) {
	ret := _m.Called(owner, repositoryName)

	var r0 *models.RepositoryInfo
	if rf, ok := ret.Get(0).(func(string, string) *models.RepositoryInfo); ok {
		r0 = rf(owner, repositoryName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.RepositoryInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(owner, repositoryName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRateLimitInfo provides a mock function with given fields: owner, repositoryName
func (_m *IGithub) GetRateLimitInfo(owner string, repositoryName string) (*github.RateLimits, error) {
	ret := _m.Called(owner, repositoryName)

	var r0 *github.RateLimits
	if rf, ok := ret.Get(0).(func(string, string) *github.RateLimits); ok {
		r0 = rf(owner, repositoryName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.RateLimits)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(owner, repositoryName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReadMe provides a mock function with given fields: owner, repositoryName
func (_m *IGithub) GetReadMe(owner string, repositoryName string) (string, error) {
	ret := _m.Called(owner, repositoryName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(owner, repositoryName)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(owner, repositoryName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
