package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"swagger-gin/module/utils"

	"github.com/gin-gonic/gin"
)

func connect() (*sql.DB, error) {
	connStr := "user=postgres password=postgresql123 dbname=tesapi sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}

type Items struct {
	Id          int         `json:"id" example:"1"`
	Itemname    string      `json:"itemname" example:"Book"`
	Price       int         `json:"price" example:"5000"`
	Createddate interface{} `json:"createddate" `
}

// PingExample godoc
// @Security token
// @Summary get item
// @Schemes
// @Description get item
// @Tags item
// @Accept json
// @Produce json
// @Success 200 {array} Items
// @Router /item/get-item [get]
func GetItem(g *gin.Context) {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from items ORDER BY id DESC")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Items

	for rows.Next() {
		var each = Items{}
		var err = rows.Scan(&each.Id, &each.Itemname, &each.Price, &each.Createddate)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	// for _, each := range result {
	// 	fmt.Println(each)
	// }

	var respone utils.Respone
	respone.Result = true
	respone.Message = ""
	respone.Data = result
	g.JSON(http.StatusOK, respone)
}

// PingExample godoc
// @Security token
// @Summary insert item
// @Schemes
// @Description insert item
// @Tags item
// @Param request body Items true "Payload Body [RAW]"
// @Accept json
// @Produce json
// @Success 200 {array} utils.Respone
// @Router /item/insert-item [post]
func InsertItem(g *gin.Context) {
	var requestBody Items

	if err := g.BindJSON(&requestBody); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(requestBody)

	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO items (itemname, price, createddate)
		VALUES ($1, $2, $3)
		RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, requestBody.Itemname, requestBody.Price, requestBody.Createddate).Scan(&id)
	if err != nil {
		panic(err)
	}

	type new_id struct {
		Id int `json:"id" example:"1"`
	}

	var return_id new_id
	return_id.Id = id

	var respone utils.Respone
	respone.Result = true
	respone.Message = "Penyimpanan Berhasil"
	respone.Data = return_id
	g.JSON(http.StatusOK, respone)
}

// PingExample godoc
// @Summary update item
// @Security token
// @Schemes
// @Description update item
// @Tags item
// @Param request body Items true "Payload Body [RAW]"
// @Accept json
// @Produce json
// @Success 200 {array} utils.Respone
// @Router /item/update-item [put]
func UpdateItem(g *gin.Context) {
	var requestBody Items

	if err := g.BindJSON(&requestBody); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(requestBody)

	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	sqlStatement := `
		UPDATE items set itemname=$1, price=$2, createddate=$3
		WHERE id=$4
		RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, requestBody.Itemname, requestBody.Price, requestBody.Createddate, requestBody.Id).Scan(&id)
	if err != nil {
		panic(err)
	}

	type new_id struct {
		Id int `json:"id" example:"1"`
	}

	var return_id new_id
	return_id.Id = id

	var respone utils.Respone
	respone.Result = true
	respone.Message = "Update Berhasil"
	respone.Data = return_id
	g.JSON(http.StatusOK, respone)
}

// PingExample godoc
// @Summary delete item
// @Security token
// @Schemes
// @Description delete item
// @Tags item
// @Param id   path int true "id"
// @Accept json
// @Produce json
// @Success 200 {array} utils.Respone
// @Router /item/delete-item/{id} [delete]
func DeleteItem(g *gin.Context) {
	id := g.Param("id")
	fmt.Println(id)

	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	sqlStatement := `
		DELETE FROM items
		WHERE id=$1`

	_, err = db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}

	type new_id struct {
		Id int `json:"id" example:"1"`
	}

	var return_id new_id
	intVar, err := strconv.Atoi(id)
	return_id.Id = intVar

	var respone utils.Respone
	respone.Result = true
	respone.Message = "Delete Berhasil"
	respone.Data = return_id
	g.JSON(http.StatusOK, respone)
}
