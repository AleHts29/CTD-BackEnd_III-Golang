package dbService

import "fmt"

type sqlitedb struct {
	file string
}

func NewSqlitedb(file string) *sqlitedb {
	return &sqlitedb{file: file}
}

func (s sqlitedb) Conect() {
	//TODO implement me
	fmt.Println("\nconectandoce a sqlite..")
}

func (s sqlitedb) Disconect() {
	//TODO implement me
	fmt.Println("\ndesconectandoce de sqlite..")
}
