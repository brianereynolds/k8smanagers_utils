package k8smanagers_utils

import "testing"
import "context"

func Test_compareVersions(t *testing.T) {

	result, err := CompareVersions("1.24.0", "1.24.1")
	expected := -1
	if (err != nil) || (result != expected) {
		t.Errorf("compareVersions(1.24.0, 1.24.1) = %d; want %d", result, expected)
	}

	result, err = CompareVersions("1.24.1", "1.24.2")
	expected = -1
	if (err != nil) || (result != expected) {
		t.Errorf("compareVersions(1.24.1, 1.24.2) = %d; want %d", result, expected)
	}

	result, err = CompareVersions("1.24.1", "1.24.0")
	expected = 1
	if (err != nil) || (result != expected) {
		t.Errorf("compareVersions(1.24.1, 1.24.0) = %d; want %d", result, expected)
	}

	result, err = CompareVersions("1.24.4", "1.24")
	if err.Error() != "version numbers must be of the format V.v.n" {
		t.Errorf("compareVersions(1.24.4, 1.24) should return an error")
	}

	result, err = CompareVersions("1.24.0", "1.24.0")
	expected = 0
	if (err != nil) || (result != expected) {
		t.Errorf("compareVersions(1.24.0, 1.24.0) = %d; want %d", result, expected)
	}
}

func Test_GetManagedClusterClient(t *testing.T) {
	ctx := context.Background()

	_, err := GetManagedClusterClient(ctx, "SUBID")

	if err != nil {
		print(err.Error())
	}
}

func Test_GetAgentPoolClient(t *testing.T) {
	ctx := context.Background()

	_, err := GetAgentPoolClient(ctx, "SUBID")

	if err != nil {
		print(err.Error())
	}
}

func Test_GetClientSet(t *testing.T) {
	ctx := context.Background()

	_, err := GetClientSet(ctx, nil)

	if err != nil {
		print(err.Error())
	}
}

func Test_IsLowercaseAndNumbers(t *testing.T) {
	result := IsLowercaseAndNumbers(context.Background(), "abcabc")
	expected := true
	if result != expected {
		t.Errorf("Test failed ")
	}

	result = IsLowercaseAndNumbers(context.Background(), "Abcabc")
	expected = false
	if result != expected {
		t.Errorf("Test failed ")
	}

	result = IsLowercaseAndNumbers(context.Background(), "abcabc349A")
	expected = false
	if result != expected {
		t.Errorf("Test failed ")
	}

	result = IsLowercaseAndNumbers(context.Background(), "abc-1da")
	expected = false
	if result != expected {
		t.Errorf("Test failed ")
	}

	result = IsLowercaseAndNumbers(context.Background(), "abc_1da")
	expected = false
	if result != expected {
		t.Errorf("Test failed ")
	}

	result = IsLowercaseAndNumbers(context.Background(), "abc+1da")
	expected = false
	if result != expected {
		t.Errorf("Test failed ")
	}
}

func Test_StartsWithANumber(t *testing.T) {
	result := StartsWithANumber(context.Background(), "1abc")
	expected := true
	if result != expected {
		t.Errorf("Test failed ")
	}

	result = StartsWithANumber(context.Background(), "")
	expected = false
	if result != expected {
		t.Errorf("Test failed ")
	}

	result = StartsWithANumber(context.Background(), "a111a")
	expected = false
	if result != expected {
		t.Errorf("Test failed ")
	}
}
