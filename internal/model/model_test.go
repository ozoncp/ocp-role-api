package model

import "testing"

func TestServiceRolesString(t *testing.T) {
	cases := []struct {
		in   ServiceRoles
		want string
	}{
		{ServiceRoles{"A", []string{}}, "ServiceRoles{Service: A, Operations: []}"},
		{ServiceRoles{"A", []string{"op1"}}, "ServiceRoles{Service: A, Operations: [op1]}"},
		{ServiceRoles{"A", []string{"op1", "op2"}}, "ServiceRoles{Service: A, Operations: [op1, op2]}"},
	}

	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("expected %s, got %s", c.want, got)
		}
	}
}
