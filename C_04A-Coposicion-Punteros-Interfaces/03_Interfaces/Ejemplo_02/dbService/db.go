package dbService

import "fmt"

type Mysql struct {
	User     string
	Password string
	Host     string
	Port     uint16
	DbName   string
}

// es una forma de saber si nos falta por implementar alguno de los metodos de la interface
var _ DbConnection = &Mysql{}

// Mysql - definimos un constructor (por buenas practicas)
func MysqlDB(user, password, host, dbname string, port uint16) *Mysql {
	return &Mysql{
		User:     user,
		Password: password,
		Host:     host,
		Port:     port,
		DbName:   dbname,
	}
}

func (m *Mysql) Conect() {
	fmt.Printf("\nConnecting to mysql://%v:%v@%v:%v/%v",
		m.User,
		m.Password,
		m.Host,
		m.Port,
		m.DbName,
	)
}

func (m Mysql) Disconect() {
	fmt.Println("\nDesconectandose de la DB")
}
