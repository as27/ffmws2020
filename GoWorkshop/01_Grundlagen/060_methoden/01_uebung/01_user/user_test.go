package user

import (
	"reflect"
	"testing"
)

func TestUserStorage_Add(t *testing.T) {
	user1 := User{1, "user1"}
	user2 := User{2, "user2"}
	user3 := User{3, "user3"}
	type args struct {
		u User
	}
	tests := []struct {
		name   string
		us     *UserStorage
		args   args
		want   bool
		wantUs UserStorage
	}{
		{
			"hinzufuegen",
			&UserStorage{[]User{user1, user2}},
			args{user3},
			true,
			UserStorage{[]User{user1, user2, user3}},
		},
		{
			"Eintrag vorhanden",
			&UserStorage{[]User{user1, user2}},
			args{User{1, "eins"}},
			false,
			UserStorage{[]User{user1, user2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := tt.us
			got := us.Add(tt.args.u)
			if got != tt.want {
				t.Errorf("UserStorage.Add() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.wantUs, *us) {
				t.Errorf("UserStorage.Add() us:%v, want: %v", *us, tt.wantUs)
			}
		})
	}
}

func TestUserStorage_GetUserByID(t *testing.T) {
	user1 := User{1, "user1"}
	user2 := User{2, "user2"}
	user3 := User{3, "user3"}
	type args struct {
		id int
	}
	tests := []struct {
		name  string
		us    *UserStorage
		args  args
		want  User
		want1 bool
	}{
		{
			"Eintrag vorhanden",
			&UserStorage{[]User{user1, user2, user3}},
			args{1},
			user1,
			true,
		},
		{
			"Eintrag vorhanden",
			&UserStorage{[]User{user1, user2, user3}},
			args{3},
			user3,
			true,
		},
		{
			"Eintrag nicht vorhanden",
			&UserStorage{[]User{user1, user2, user3}},
			args{4},
			User{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.us.GetUserByID(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserStorage.GetUserByID() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("UserStorage.GetUserByID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
