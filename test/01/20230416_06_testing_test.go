package test

import "testing"

func TestSquare1(t *testing.T) {
	rst := square(9)
	if rst != 81 {
		t.Errorf("Square(9) should be 81 but returns %d\n", rst)
	}
}

func TestSquare2(t *testing.T) {
	rst := square(3)
	if rst != 9 {
		t.Errorf("Square(3) should be 9 but returns %d\n", rst)
	}
}

// Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -coverprofile=C:\_dev\_TEMPA~1\Local\Temp\vscode-goFRSlPZ\go-code-cover test/01

// ok  	test/01	0.053s	coverage: 50.0% of statements
