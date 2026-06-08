package main

import "testing"

func Test_handlerRegister(t *testing.T) {
	tests := []struct {
		description string // description of this test case
		// descriptiond input parameters for target function.
		s       *state
		cmd     command
		wantErr bool
	}{
		{
			description: "errors without any args",
			s:           &state{},
			cmd:         command{name: "register", args: nil},
			wantErr:     true,
		},
		{
			description: "errors without a name in the args",
			s:           &state{},
			cmd:         command{name: "register", args: []string{""}},
			wantErr:     true,
		},
		// {
		// 	description: "fetches an existing user by name",
		// 	s:           &state{},
		// 	cmd:         command{name: "register", args: []string{""}},
		// 	wantErr:     false,
		// },
		// {
		// 	description: "errors with a non-existing user",
		// 	s:           &state{},
		// 	cmd:         command{name: "register", args: []string{""}},
		// 	wantErr:     true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			gotErr := handlerRegister(tt.s, tt.cmd)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("handlerRegister() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("handlerRegister() succeeded unexpectedly")
			}
		})
	}
}
