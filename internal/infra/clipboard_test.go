package infra

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCommandRunner struct {
	mock.Mock
}

func (m *MockCommandRunner) Run(name string, args ...string) error {
	arguments := []interface{}{name}
	for _, arg := range args {
		arguments = append(arguments, arg)
	}
	return m.Called(arguments...).Error(0)
}

func (m *MockCommandRunner) LookPath(name string) (string, error) {
	args := m.Called(name)
	return args.String(0), args.Error(1)
}

func (m *MockCommandRunner) SetText(text string) {
	m.Called(text)
}

type MockOSDetector struct {
	mock.Mock
}

func (m *MockOSDetector) GetOS() string {
	return m.Called().String(0)
}

func TestClipboard_Copy_OSX(t *testing.T) {
	// Arrange
	mockRunner := new(MockCommandRunner)
	mockOS := new(MockOSDetector)

	mockOS.On("GetOS").Return("darwin")
	mockRunner.On("SetText", "test text").Return()
	mockRunner.On("Run", "pbcopy").Return(nil)

	clipboard := NewClipboard(WithCommandRunner(mockRunner), WithOSDetector(mockOS))

	// Act
	err := clipboard.Copy("test text")

	// Assert
	assert.NoError(t, err)
	mockOS.AssertExpectations(t)
	mockRunner.AssertExpectations(t)
}

func TestClipboard_Copy_Linux_Xclip(t *testing.T) {
	// Arrange
	mockRunner := new(MockCommandRunner)
	mockOS := new(MockOSDetector)

	mockOS.On("GetOS").Return("linux")
	mockRunner.On("LookPath", "xclip").Return("/usr/bin/xclip", nil)
	mockRunner.On("SetText", "test text").Return()
	mockRunner.On("Run", "xclip", "-selection", "clipboard").Return(nil)

	clipboard := NewClipboard(WithCommandRunner(mockRunner), WithOSDetector(mockOS))

	// Act
	err := clipboard.Copy("test text")

	// Assert
	assert.NoError(t, err)
	mockOS.AssertExpectations(t)
	mockRunner.AssertExpectations(t)
}

func TestClipboard_Copy_Linux_Xsel_Fallback(t *testing.T) {
	// Arrange
	mockRunner := new(MockCommandRunner)
	mockOS := new(MockOSDetector)

	mockOS.On("GetOS").Return("linux")
	mockRunner.On("LookPath", "xclip").Return("", errors.New("not found"))
	mockRunner.On("SetText", "test text").Return()
	mockRunner.On("LookPath", "xsel").Return("/usr/bin/xsel", nil)
	mockRunner.On("Run", "xsel", "--clipboard", "--input").Return(nil)

	clipboard := NewClipboard(WithCommandRunner(mockRunner), WithOSDetector(mockOS))

	// Act
	err := clipboard.Copy("test text")

	// Assert
	assert.NoError(t, err)
	mockOS.AssertExpectations(t)
	mockRunner.AssertExpectations(t)
}

func TestClipboard_Copy_Linux_NoCommandsAvailable(t *testing.T) {
	// Arrange
	mockRunner := new(MockCommandRunner)
	mockOS := new(MockOSDetector)

	mockOS.On("GetOS").Return("linux")
	mockRunner.On("LookPath", "xclip").Return("", errors.New("not found"))
	mockRunner.On("SetText", "test text").Return()
	mockRunner.On("LookPath", "xsel").Return("", errors.New("not found"))

	clipboard := NewClipboard(WithCommandRunner(mockRunner), WithOSDetector(mockOS))

	// Act
	err := clipboard.Copy("test text")

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "neither xclip nor xsel found")
	mockOS.AssertExpectations(t)
	mockRunner.AssertExpectations(t)
}

func TestClipboard_Copy_UnsupportedOS(t *testing.T) {
	// Arrange
	mockRunner := new(MockCommandRunner)
	mockOS := new(MockOSDetector)

	mockOS.On("GetOS").Return("windows")
	mockRunner.On("SetText", "test text").Return()

	clipboard := NewClipboard(WithCommandRunner(mockRunner), WithOSDetector(mockOS))

	// Act
	err := clipboard.Copy("test text")

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "clipboard copy not supported on windows")
	mockOS.AssertExpectations(t)
}

func TestClipboard_Copy_OSX_Error(t *testing.T) {
	// Arrange
	mockRunner := new(MockCommandRunner)
	mockOS := new(MockOSDetector)
	expectedError := errors.New("command failed")

	mockOS.On("GetOS").Return("darwin")
	mockRunner.On("SetText", "test text").Return()
	mockRunner.On("Run", "pbcopy").Return(expectedError)

	clipboard := NewClipboard(WithCommandRunner(mockRunner), WithOSDetector(mockOS))

	// Act
	err := clipboard.Copy("test text")

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockOS.AssertExpectations(t)
	mockRunner.AssertExpectations(t)
}

func TestClipboard_Copy_Linux_Xclip_Error(t *testing.T) {
	// Arrange
	mockRunner := new(MockCommandRunner)
	mockOS := new(MockOSDetector)
	expectedError := errors.New("command failed")

	mockOS.On("GetOS").Return("linux")
	mockRunner.On("SetText", "test text").Return()
	mockRunner.On("LookPath", "xclip").Return("/usr/bin/xclip", nil)
	mockRunner.On("Run", "xclip", "-selection", "clipboard").Return(expectedError)

	clipboard := NewClipboard(WithCommandRunner(mockRunner), WithOSDetector(mockOS))

	// Act
	err := clipboard.Copy("test text")

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockOS.AssertExpectations(t)
	mockRunner.AssertExpectations(t)
}
