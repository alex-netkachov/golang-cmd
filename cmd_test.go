package cmd

import "testing"

// Test_Parse checks correctness of the command line
// parsing by running test cases.
func Test_Parse(t *testing.T) {
	test(t, ``, []string{""})
	test(t, `a`, []string{"a"})
	test(t, `a `, []string{"a"})
	test(t, `a  `, []string{"a"})
	test(t, `ab`, []string{"ab"})
	test(t, `ab `, []string{"ab"})
	test(t, `ab  `, []string{"ab"})
	test(t, `a c`, []string{"a", "c"})
	test(t, `a c `, []string{"a", "c"})
	test(t, `a cd`, []string{"a", "cd"})
	test(t, `a cd `, []string{"a", "cd"})
	test(t, `""`, []string{""})
	test(t, `""""`, []string{""})
	test(t, `" "`, []string{" "})
	test(t, `"a"`, []string{"a"})
	test(t, `"a"""`, []string{"a"})
	test(t, `a""`, []string{"a"})
	test(t, `a""b`, []string{"ab"})
}

func test(t *testing.T, test string, expected []string) {
	cmd, args := Parse(test)
	actual := append([]string{cmd}, args...)
	if len(actual) != len(expected) {
		t.Errorf("%#v => %#v != %#v", test, actual, expected)
	}
	for i := 0; i < len(actual) && i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("%#v => %#v != %#v", test, actual, expected)
		}
	}
}
