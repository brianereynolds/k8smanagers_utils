package k8smanagers_utils

import "testing"

func Test_compareVersions(t *testing.T) {

	result, err := compareVersions("1.24.0", "1.24.1")
	expected := -1
	if (err != nil) || (result != expected) {
		t.Errorf("compareVersions(1.24.0, 1.24.1) = %d; want %d", result, expected)
	}

	result, err = compareVersions("1.24.1", "1.24.2")
	expected = -1
	if (err != nil) || (result != expected) {
		t.Errorf("compareVersions(1.24.1, 1.24.2) = %d; want %d", result, expected)
	}

	result, err = compareVersions("1.24.1", "1.24.0")
	expected = 1
	if (err != nil) || (result != expected) {
		t.Errorf("compareVersions(1.24.1, 1.24.0) = %d; want %d", result, expected)
	}

	result, err = compareVersions("1.24.4", "1.24")
	if err.Error() != "version numbers must be of the format V.v.n" {
		t.Errorf("compareVersions(1.24.4, 1.24) should return an error")
	}

	result, err = compareVersions("1.24.0", "1.24.0")
	expected = 0
	if (err != nil) || (result != expected) {
		t.Errorf("compareVersions(1.24.0, 1.24.0) = %d; want %d", result, expected)
	}
}
