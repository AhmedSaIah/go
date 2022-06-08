package main

import "testing"

func TestOK(t *testing.T) {
	type args struct {
		b BankAccount
	}
	tests := map[string]struct {
		name string
		args args
		Err  bool
	}{
		"Case it was not OK": {
			args: args{BankAccount{
				Name:   "Ahmed",
				ID:     0,
				Salary: 3000,
				BankID: 2,
			}},
			Err: true,
		},
		"Case it was OK": {
			args: args{BankAccount{
				Name:   "Ahmed",
				ID:     444,
				Salary: 3000,
				BankID: 1,
			}},
			Err: false,
		},
		"Case it was kinda okay": {
			args: args{BankAccount{
				Name:   "",
				ID:     11,
				Salary: 0,
				BankID: 1,
			}},
			Err: false,
		},
	}
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			if err := OK(tt.args.b); (err != nil) != tt.Err {
				t.Errorf("OK() error = %v, wantErr %v", err, tt.Err)
			}
		})
	}
}
