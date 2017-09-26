package country

import (
	"database/sql"

	"openthailand/errors"
	"openthailand/models"
)

type DatabaseRepository struct {
	DB *sql.DB
}

func (d *DatabaseRepository) GetProvinceDetail(provinceID string) (models.Province, error) {
	stmt, err := d.DB.Prepare(`select id,name,name_th from province where id = $1`)
	if err != nil {
		return models.Province{}, errors.ServerError("%v", err).Error()
	}
	defer stmt.Close()
	province := models.Province{}
	err = stmt.QueryRow(provinceID).Scan(&province.ID, &province.Name, &province.NameThai)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Province{}, errors.BadRequest("Province not found").Error()
		}
		return models.Province{}, errors.ServerError("%v", err).Error()
	}

	return province, nil
}

func (d *DatabaseRepository) GetProvinces() ([]models.Province, error) {
	stmt, err := d.DB.Prepare(`select id,name,name_th from province`)
	if err != nil {
		return []models.Province{}, errors.ServerError("%v", err).Error()
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return []models.Province{}, errors.ServerError("%v", err).Error()
	}
	defer rows.Close()

	list := []models.Province{}
	for rows.Next() {
		province := models.Province{}
		err = rows.Scan(&province.ID, &province.Name, &province.NameThai)
		if err != nil {
			return []models.Province{}, errors.ServerError("%v", err).Error()
		}

		list = append(list, province)
	}
	return list, nil
}

func (d *DatabaseRepository) GetDistrictsByProvinceID(provinceID string) ([]models.District, error) {
	stmt, err := d.DB.Prepare(`select id,name,name_th from district where province_id=$1`)
	if err != nil {
		return []models.District{}, errors.ServerError("%v", err).Error()
	}
	defer stmt.Close()

	rows, err := stmt.Query(provinceID)
	if err != nil {
		return []models.District{}, errors.ServerError("%v", err).Error()
	}
	defer rows.Close()

	list := []models.District{}
	for rows.Next() {
		district := models.District{}
		err = rows.Scan(&district.ID, &district.Name, &district.NameThai)
		if err != nil {
			return []models.District{}, errors.ServerError("%v", err).Error()
		}

		list = append(list, district)
	}
	return list, nil
}
func (d *DatabaseRepository) GetSubdistrictsByDistrictID(districtID string) ([]models.Subdistrict, error) {
	stmt, err := d.DB.Prepare(`select id,name,name_th,lat,lng from subdistrict where district_id=$1`)
	if err != nil {
		return []models.Subdistrict{}, errors.ServerError("%v", err).Error()
	}
	defer stmt.Close()

	rows, err := stmt.Query(districtID)
	if err != nil {
		return []models.Subdistrict{}, errors.ServerError("%v", err).Error()
	}
	defer rows.Close()

	list := []models.Subdistrict{}
	for rows.Next() {
		subdistrict := models.Subdistrict{}
		err = rows.Scan(&subdistrict.ID, &subdistrict.Name, &subdistrict.NameThai, &subdistrict.Location.Lat, &subdistrict.Location.Lng)
		if err != nil {
			return []models.Subdistrict{}, errors.ServerError("%v", err).Error()
		}

		list = append(list, subdistrict)
	}
	return list, nil
}
