package main

import "github.com/AleHts29/CTD-BackEnd_III-Golang/C_04A-Coposicion-Punteros-Interfaces/03_Interfaces/Ejemplo_02/dbService"

func main() {
	//var db dbService.DbConnection
	//db := dbService.MysqlDB("admin", "123qwe", "localhost", "dbtest", 3000)

	db := dbService.NewSqlitedb("/var/www/users.sqlite")
	db.Conect()
	db.Disconect()
}
