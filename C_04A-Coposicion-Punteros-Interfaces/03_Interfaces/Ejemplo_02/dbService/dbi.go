package dbService

type DbConnection interface {
	Conect()
	Disconect()
	//CreateTable(name string)
}
