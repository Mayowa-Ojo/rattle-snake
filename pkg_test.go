package snake

import (
	"strings"
	"testing"
)

func TestToSnakeCase(t *testing.T) {
	res := ToSnakeCase("ModeratorsCanInvite")
	contains := strings.Contains(res, "_")
	expectedStr := "moderators_can_invite"

	if res != expectedStr {
		t.Errorf("expected result to be %s, got %s", expectedStr, res)
	}

	if !contains {
		t.Error("expected formatted string to contain '_'")
	}

	if string(res[10]) != "_" {
		t.Errorf("expected '_' at index 10, got %s", string(res[10]))
	}

	if string(res[10]) != "_" {
		t.Errorf("expected '_' at index 10, got %s", string(res[13]))
	}
}

func TestToKebabCase(t *testing.T) {
	res := ToKebabCase("ModeratorsCanInvite")
	contains := strings.Contains(res, "-")
	expectedStr := "moderators-can-invite"

	if res != expectedStr {
		t.Errorf("expected result to be %s, got %s", expectedStr, res)
	}

	if !contains {
		t.Error("expected formatted string to contain '-'")
	}

	if string(res[10]) != "-" {
		t.Errorf("expected '-' at index 10, got %s", string(res[10]))
	}

	if string(res[10]) != "-" {
		t.Errorf("expected '-' at index 10, got %s", string(res[13]))
	}
}
