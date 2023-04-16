package test

import "testing"
import "github.com/stretchr/testify/assert"

func TestSquare1(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(81, square(9), "square(9) should be 81")
}

func TestSquare2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(9, square(3), "square(3) should be 9")
}

// Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -coverprofile=C:\_dev\_TEMPA~1\Local\Temp\vscode-goFRSlPZ\go-code-cover test.02

// ok  	test.02	0.049s	coverage: 50.0% of statements
