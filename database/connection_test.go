package database

import (
	"testing"
)

func TestCreateInsert_1(t *testing.T) {
	columns := []string{
		"id",
		"nome",
		"descricao",
	}

	got := CreateInsert(1, "pessoa", columns)
	want := "insert into pessoa (id, nome, descricao) values($1,$2,$3)"

	if got != want {
		t.Errorf("CreateInsert de 1 linha\ngot: %v\nwant: %v\n", got, want)
	}
}

func TestCreateInsert_3(t *testing.T) {
	columns := []string{
		"id",
		"nome",
		"descricao",
	}

	got := CreateInsert(3, "pessoa", columns)
	want := "insert into pessoa (id, nome, descricao) values($1,$2,$3),($4,$5,$6),($7,$8,$9)"

	if got != want {
		t.Errorf("CreateInsert de 1 linha\ngot: %v\nwant: %v\n", got, want)
	}
}
