package usecase

import "sygap-cmdb/entities"

type ICiClass interface {
	GetCiClassList() ([]entities.Ci_Class, error)
}

func (m *model) GetCiClassList() ([]entities.Ci_Class, error) {
	var CiClassList []entities.Ci_Class
	err := m.mysql.Find(&CiClassList).Error
	if err != nil {
		return CiClassList, err
	}
	return CiClassList, nil
}
