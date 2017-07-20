package xdg

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockDefaulter struct {
	mock.Mock
}

func (m *mockDefaulter) defaultDataHome() string {
	args := m.Called()
	return args.String(0)
}
func (m *mockDefaulter) defaultDataDirs() []string {
	args := m.Called()
	return args.Get(0).([]string)
}
func (m *mockDefaulter) defaultConfigHome() string {
	args := m.Called()
	return args.String(0)
}
func (m *mockDefaulter) defaultConfigDirs() []string {
	args := m.Called()
	return args.Get(0).([]string)
}
func (m *mockDefaulter) defaultCacheHome() string {
	args := m.Called()
	return args.String(0)
}

func TestDataHome_WithoutXDG(t *testing.T) {
	assert := assert.New(t)
	expected := "/some/path"
	mockDef := new(mockDefaulter)
	mockDef.On("defaultDataHome").Return(expected)
	setDefaulter(mockDef)
	os.Setenv("XDG_DATA_HOME", "")

	actual := DataHome()
	mockDef.AssertExpectations(t)
	assert.Equal(expected, actual)
}

func TestDataHome_WithXDG(t *testing.T) {
	assert := assert.New(t)
	expected := "/some/path"
	mockDef := new(mockDefaulter)
	mockDef.On("defaultDataHome").Return("/wrong/path")
	setDefaulter(mockDef)
	os.Setenv("XDG_DATA_HOME", expected)

	actual := DataHome()
	mockDef.AssertNotCalled(t, "defaultDataHome")
	assert.Equal(expected, actual)
}

func TestDataHome_Application(t *testing.T) {
	assert := assert.New(t)
	root := "/some/path"
	vendor := "OpenPeeDeeP"
	app := "XDG"
	expected := filepath.Join(root, vendor, app)
	mockDef := new(mockDefaulter)
	appXDG := New(vendor, app)
	mockDef.On("defaultDataHome").Return(root)
	setDefaulter(mockDef)
	os.Setenv("XDG_DATA_HOME", "")

	actual := appXDG.DataHome()
	mockDef.AssertExpectations(t)
	assert.Equal(expected, actual)
}

func TestDataDirs_WithoutXDG(t *testing.T) {
	assert := assert.New(t)
	expected := []string{"/some/path", "/some/other/path"}
	mockDef := new(mockDefaulter)
	mockDef.On("defaultDataDirs").Return(expected)
	setDefaulter(mockDef)
	os.Setenv("XDG_DATA_DIRS", "")

	actual := DataDirs()
	mockDef.AssertExpectations(t)
	assert.Equal(expected, actual)
}

func TestDataDirs_WithXDG(t *testing.T) {
	assert := assert.New(t)
	expected := []string{"/some/path", "/some/other/path"}
	mockDef := new(mockDefaulter)
	mockDef.On("defaultDataDirs").Return([]string{"/wrong/path"})
	setDefaulter(mockDef)
	os.Setenv("XDG_DATA_DIRS", strings.Join(expected, seperator))

	actual := DataDirs()
	mockDef.AssertNotCalled(t, "defaultDataDirs")
	assert.Equal(expected, actual)
}

func TestDataDirs_Application(t *testing.T) {
	assert := assert.New(t)
	root := []string{"/some/path", "/some/other/path"}
	vendor := "OpenPeeDeeP"
	app := "XDG"
	expected := make([]string, len(root))
	for i, r := range root {
		expected[i] = filepath.Join(r, vendor, app)
	}
	mockDef := new(mockDefaulter)
	appXDG := New(vendor, app)
	mockDef.On("defaultDataDirs").Return(root)
	setDefaulter(mockDef)
	os.Setenv("XDG_DATA_DIRS", "")

	actual := appXDG.DataDirs()
	mockDef.AssertExpectations(t)
	assert.Equal(expected, actual)
}

func TestConfigHome_WithoutXDG(t *testing.T) {
	assert := assert.New(t)
	expected := "/some/path"
	mockDef := new(mockDefaulter)
	mockDef.On("defaultConfigHome").Return(expected)
	setDefaulter(mockDef)
	os.Setenv("XDG_CONFIG_HOME", "")

	actual := ConfigHome()
	mockDef.AssertExpectations(t)
	assert.Equal(expected, actual)
}

func TestConfigHome_WithXDG(t *testing.T) {
	assert := assert.New(t)
	expected := "/some/path"
	mockDef := new(mockDefaulter)
	mockDef.On("defaultConfigHome").Return("/wrong/path")
	setDefaulter(mockDef)
	os.Setenv("XDG_CONFIG_HOME", expected)

	actual := ConfigHome()
	mockDef.AssertNotCalled(t, "defaultConfigHome")
	assert.Equal(expected, actual)
}

func TestConfigHome_Application(t *testing.T) {
	assert := assert.New(t)
	root := "/some/path"
	vendor := "OpenPeeDeeP"
	app := "XDG"
	expected := filepath.Join(root, vendor, app)
	mockDef := new(mockDefaulter)
	appXDG := New(vendor, app)
	mockDef.On("defaultConfigHome").Return(root)
	setDefaulter(mockDef)
	os.Setenv("XDG_CONFIG_HOME", "")

	actual := appXDG.ConfigHome()
	mockDef.AssertExpectations(t)
	assert.Equal(expected, actual)
}

func TestConfigDirs_WithoutXDG(t *testing.T) {
	assert := assert.New(t)
	expected := []string{"/some/path", "/some/other/path"}
	mockDef := new(mockDefaulter)
	mockDef.On("defaultConfigDirs").Return(expected)
	setDefaulter(mockDef)
	os.Setenv("XDG_CONFIG_DIRS", "")

	actual := ConfigDirs()
	mockDef.AssertExpectations(t)
	assert.Equal(expected, actual)
}

func TestConfigDirs_WithXDG(t *testing.T) {
	assert := assert.New(t)
	expected := []string{"/some/path", "/some/other/path"}
	mockDef := new(mockDefaulter)
	mockDef.On("defaultConfigDirs").Return([]string{"/wrong/path"})
	setDefaulter(mockDef)
	os.Setenv("XDG_CONFIG_DIRS", strings.Join(expected, seperator))

	actual := ConfigDirs()
	mockDef.AssertNotCalled(t, "defaultConfigDirs")
	assert.Equal(expected, actual)
}

func TestConfigDirs_Application(t *testing.T) {
	assert := assert.New(t)
	root := []string{"/some/path", "/some/other/path"}
	vendor := "OpenPeeDeeP"
	app := "XDG"
	expected := make([]string, len(root))
	for i, r := range root {
		expected[i] = filepath.Join(r, vendor, app)
	}
	mockDef := new(mockDefaulter)
	appXDG := New(vendor, app)
	mockDef.On("defaultConfigDirs").Return(root)
	setDefaulter(mockDef)
	os.Setenv("XDG_CONFIG_DIRS", "")

	actual := appXDG.ConfigDirs()
	mockDef.AssertExpectations(t)
	assert.Equal(expected, actual)
}

func TestCacheHome_WithoutXDG(t *testing.T) {
	assert := assert.New(t)
	expected := "/some/path"
	mockDef := new(mockDefaulter)
	mockDef.On("defaultCacheHome").Return(expected)
	setDefaulter(mockDef)
	os.Setenv("XDG_CACHE_HOME", "")

	actual := CacheHome()
	mockDef.AssertExpectations(t)
	assert.Equal(expected, actual)
}

func TestCacheHome_WithXDG(t *testing.T) {
	assert := assert.New(t)
	expected := "/some/path"
	mockDef := new(mockDefaulter)
	mockDef.On("defaultCacheHome").Return("/wrong/path")
	setDefaulter(mockDef)
	os.Setenv("XDG_CACHE_HOME", expected)

	actual := CacheHome()
	mockDef.AssertNotCalled(t, "defaultCacheHome")
	assert.Equal(expected, actual)
}

func TestCacheHome_Application(t *testing.T) {
	assert := assert.New(t)
	root := "/some/path"
	vendor := "OpenPeeDeeP"
	app := "XDG"
	expected := filepath.Join(root, vendor, app)
	mockDef := new(mockDefaulter)
	appXDG := New(vendor, app)
	mockDef.On("defaultCacheHome").Return(root)
	setDefaulter(mockDef)
	os.Setenv("XDG_CACHE_HOME", "")

	actual := appXDG.CacheHome()
	mockDef.AssertExpectations(t)
	assert.Equal(expected, actual)
}
