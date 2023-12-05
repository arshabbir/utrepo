package dao

import (
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("DAOMAIN : Started")
	m.Run()
}

func TestA(t *testing.T) {
	log.Println("TestA")
}
