package session

import "testing"

type User struct {
	Name string `sorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	s := NewSession().Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("failed to create table User")
	}
}

func TestSession_Model(t *testing.T) {
	s := NewSession().Model(&User{})
	table := s.RefTable()
	s.Model(&Session{})
	if table.Name != "User" || s.RefTable().Name != "Session" {
		t.Fatal("failed to change model")
	}
}
