package main

import (
	"sorm"
	"testing"
)

type User struct {
	Name string `sorm:"PRIMARY KEY"`
	Age  int
}

var e, _ = sorm.NewEngine("sqlite3", "simple.db")

func TestSession_CreateTable(t *testing.T) {
	s := e.NewSession().Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("failed to create table User")
	}
}
