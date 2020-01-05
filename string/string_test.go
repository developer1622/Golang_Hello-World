package string

import (
	"testing"
)

func Test_Reverse(t *testing.T) {
  testcase:={
	  name string, // name of the subtest
	  input string // input to the function
	  expectErr bool // will this method fail for this case?
	  expectedOutput string // what is the expected output
  }{
	  name:"valid string",
	  input:"tata",
	  expectErr:false,
	  expectedOutput:"atat",
  }

