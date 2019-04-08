package models

import (
	"errors"
)

type Agenda struct {
	UID     uint32 `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

var (
	ErrAgendaNotFound = errors.New("Directorio no encontrado")
)

func NewAgenda(agenda Agenda) (bool, error) {
	con := Connect()
	defer con.Close()
	tx, err := con.Begin()

	if err != nil {
		return false, err
	}
	sql := "insert into agenda(name, address, phone) values($1, $2,$3)"
	{
		stmt, err := tx.Prepare(sql)
		if err != nil {
			tx.Rollback()
			return false, err
		}
		defer stmt.Close()
		_, err = stmt.Exec(agenda.Name, agenda.Address, agenda.Phone)
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}
	return true, tx.Commit()
}

func GetAgenda() ([]Agenda, error) {
	con := Connect()
	defer con.Close()

	sql := "select *from agenda"
	rs, err := con.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	var directorio []Agenda
	for rs.Next() {
		var agenda Agenda
		err := rs.Scan(&agenda.UID, &agenda.Name, &agenda.Address, &agenda.Phone)
		if err != nil {
			return nil, err
		}
		directorio = append(directorio, agenda)
	}

	return directorio, nil
}

func GetAgendaId(id uint32) (Agenda, error) {
	con := Connect()
	defer con.Close()
	sql := "select *from agenda where id=$1"
	rs, err := con.Query(sql, id)
	if err != nil {
		return Agenda{}, err
	}
	defer rs.Close()
	var agenda Agenda
	for rs.Next() {
		err := rs.Scan(&agenda.UID, &agenda.Name, &agenda.Address, &agenda.Phone)
		if err != nil {
			return Agenda{}, err
		}
	}
	if agenda.UID == 0 {
		return Agenda{}, ErrAgendaNotFound
	}
	return agenda, nil
}

func UpdateAgenda(agenda Agenda) (int64, error) {
	con := Connect()
	defer con.Close()
	sql := "update agenda set name=$1, address=$2, phone=$3 where id=$4"
	stmt, err := con.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	rs, err := stmt.Exec(agenda.Name, agenda.Address, agenda.Phone, agenda.UID)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}

func DeleteAgenda(uid uint32) (int64, error) {
	con := Connect()
	defer con.Close()
	sql := "delete from agenda where id=$1"
	stmt, err := con.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	rs, err := stmt.Exec(uid)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()

}
