package model

import "testing"

func TestUserID(t *testing.T) {
	u := &User{
		Name:     "vincent",
		Password: "pwd",
		ID:       "abc",
	}
	if u.GetID() != "abc" {
		t.Errorf("get id fails %v: expected: %v", u.GetID(), "abc")
	}
}

func TestUserName(t *testing.T) {
	u := &User{
		Name:     "vincent",
		Password: "pwd",
		ID:       "abc",
	}
	if u.GetName() != "vincent" {
		t.Errorf("get GetName fails %v: expected: %v", u.GetName(), "vincent")
	}
}
