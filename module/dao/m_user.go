package dao

import (
	"fmt"
	"html"
	"strings"
	"swagger-gin/module/config"

	"golang.org/x/crypto/bcrypt"
)

type UserGetModel struct {
	Id_user  string `json:"id_user"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}

type M_User struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
}

func UserGet() (bool, string, interface{}) {
	db, err_conn := config.ConnectDataBase()
	if err_conn != nil {
		fmt.Println(err_conn.Error())
		return false, "Error Connection", ""
	}
	defer db.Close()

	rows, err := db.Query("select id_user,username,fullname from m_user ORDER BY id_user DESC")
	if err != nil {
		fmt.Println(err.Error())
		return false, "Error Sql", ""
	}
	defer rows.Close()
	// fmt.Println(rows)

	var result []UserGetModel

	for rows.Next() {
		var each = UserGetModel{}
		var err = rows.Scan(&each.Id_user, &each.Username, &each.Fullname)

		if err != nil {
			fmt.Println(err.Error())
			return false, "Error Sql", ""
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return false, "Error Sql", ""
	}
	// fmt.Println(result)

	return true, "", result
}

func UserGetById(id int) (bool, string, interface{}) {
	db, err_conn := config.ConnectDataBase()
	if err_conn != nil {
		fmt.Println(err_conn.Error())
		return false, "Error Connection", ""
	}
	defer db.Close()

	var result = UserGetModel{}

	err_sql := db.
		QueryRow("select id_user,username,fullname from m_user WHERE id_user=$1", id).
		Scan(&result.Id_user, &result.Username, &result.Fullname)
	if err_sql != nil {
		return false, "Error Sql", ""
	}

	return true, "", result
}

func SaveUser(u M_User) (bool, string, int) {
	db, err_conn := config.ConnectDataBase()
	if err_conn != nil {
		fmt.Println(err_conn.Error())
		return false, "Error Connection", 0
	}
	defer db.Close()

	//turn password into hash
	hashedPassword, err_pass := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err_pass != nil {
		return false, "Error Hasher", 0
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	// fmt.Println(u.Password)
	// fmt.Println(u.Username)

	var err error
	sqlStatement := `
		INSERT INTO m_user (username, fullname, password)
		VALUES ($1, $2, $3)
		RETURNING id_user`
	id_user := 0
	err = db.QueryRow(sqlStatement, u.Username, u.Fullname, u.Password).Scan(&id_user)
	if err != nil {
		panic(err)
	}

	return true, "Penyimpanan Berhasil", id_user
}

func UpdateUser(u UserGetModel) (bool, string, int) {
	db, err_conn := config.ConnectDataBase()
	if err_conn != nil {
		fmt.Println(err_conn.Error())
		return false, "Error Connection", 0
	}
	defer db.Close()

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	// fmt.Println(u.Username)

	var err error
	sqlStatement := `
		UPDATE m_user SET
		username =$1,
		fullname =$2
		WHERE id_user =$3
		RETURNING id_user`
	id_user := 0
	err = db.QueryRow(sqlStatement, u.Username, u.Fullname, u.Id_user).Scan(&id_user)
	if err != nil {
		panic(err)
	}

	return true, "Proses Update Berhasil", id_user
}

func UserDeleteById(id int) (bool, string, interface{}) {
	db, err_conn := config.ConnectDataBase()
	if err_conn != nil {
		fmt.Println(err_conn.Error())
		return false, "Error Connection", ""
	}
	defer db.Close()

	sqlStatement := `
		DELETE FROM m_user
		WHERE id_user=$1`

	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}

	return true, "Proses Delete berhasil", ""
}
