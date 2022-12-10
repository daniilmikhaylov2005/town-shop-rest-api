package repository

import (
	"github.com/daniilmikhaylov2005/town-shop-rest-api/models"
)

func GetAllGoods(category string) ([]models.Good, error) {
	db := getConnection()
	defer db.Close()

	query := `SELECT * FROM goods WHERE category=$1`
	rows, err := db.Query(query, category)

	if err != nil {
		return []models.Good{}, err
	}

	var goods []models.Good

	defer rows.Close()
	for rows.Next() {
		var good models.Good

		if err := rows.Scan(&good.ID, &good.Name, &good.Description, &good.Image, &good.Category); err != nil {
			return []models.Good{}, err
		}

		goods = append(goods, good)
	}

	return goods, nil
}

func GetGoodByIdAndCategory(category string, id int) (models.Good, error) {
	db := getConnection()
	defer db.Close()

	var good models.Good

	query := `SELECT * FROM goods WHERE category=$1 AND id=$2`
	row := db.QueryRow(query, category, id)

	if err := row.Scan(&good.ID, &good.Name, &good.Description, &good.Image, &good.Category); err != nil {
		return models.Good{}, err
	}

	return good, nil
}

func GetGoodById(id int) (models.Good, error) {
	db := getConnection()
	defer db.Close()

	var good models.Good

	query := `SELECT * FROM goods WHERE id=$1`
	row := db.QueryRow(query, id)

	if err := row.Scan(&good.ID, &good.Name, &good.Description, &good.Image, &good.Category); err != nil {
		return models.Good{}, err
	}

	return good, nil
}

func InsertGood(good models.Good) (int, error) {
	db := getConnection()
	defer db.Close()

	query := `INSERT INTO goods (name, description, category, image) VALUES ($1, $2, $3, $4) RETURNING id`
	row := db.QueryRow(query, good.Name, good.Description, good.Category, good.Image)

	var id int

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func UpdateGood(good models.Good, goodId int) (int, error) {
	db := getConnection()
	defer db.Close()

	query := `UPDATE goods SET name=$1, description=$2, category=$3 WHERE id=$4 RETURNING id`
	row := db.QueryRow(query, good.Name, good.Description, good.Category, goodId)

	var id int

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func DeleteGood(id int) (int, error) {
	db := getConnection()
	defer db.Close()

	query := `DELETE FROM goods WHERE id=$1 RETURNING id`
	row := db.QueryRow(query, id)

	var deletedId int
	if err := row.Scan(&deletedId); err != nil {
		return 0, err
	}

	return deletedId, nil
}
