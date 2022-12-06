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

func GetGoodById(category string, id int) (models.Good, error) {
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
