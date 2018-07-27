package models

import "github.com/sunyatsuntobee/server/logger"

// DepartmentAddressList Model
type DepartmentAddressList struct {
	ID             int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	Name           string `xorm:"name VARCHAR(45) NOTNULL" json:"name"`
	//OrganizationID int    `xorm:"organization_id INT NOTNULL INDEX(organization_id_idx)" json:"organization_id"`
	DepartmentID   int    `xorm:"department_id INT NOTNULL" json:"department_id"`
}

//  DepartmentAddressListDataAccessObject provides database access for Model
//  DepartmentAddressList
type DepartmentAddressListDataAccessObject struct{}

// DepartmentAddressListDAO instance of DepartmentAddressListDataAccessObject
var DepartmentAddressListDAO *DepartmentAddressListDataAccessObject

// NewDepartmentAddressList creates a new Department AddressList
func NewDepartmentAddressList(name string,
	departmentID int) *DepartmentAddressList {
	return &DepartmentAddressList{
		Name:           name,
		DepartmentID:   departmentID,
	}
}

// TableName returns table name
func (*DepartmentAddressListDataAccessObject) TableName() string {
	return "organization_departments_address_list"
}

// DeleteByOID deletes all organization departments according to an
// organization ID
func (*DepartmentAddressListDataAccessObject) DeleteByDepartmentID(departmentID int) {
	var buf DepartmentAddressList
	_, err := orm.Table(DepartmentAddressListDAO.TableName()).
		Where("organization_id=?", departmentID).
		Unscoped().Delete(&buf)
	logger.LogIfError(err)
}

// DeleteByUserName deletes a department addresslist according to a username
func (*DepartmentAddressListDataAccessObject) DeleteByUserName(userName string) {
	var buf DepartmentAddressList
	_, err := orm.Table(DepartmentAddressListDAO.TableName()).
		Where("name=?", userName).
		Unscoped().Delete(&buf)
	logger.LogIfError(err)
}

// InsertOne inserts a new Department AddressList
func (*DepartmentAddressListDataAccessObject) InsertOne(
	addressList *DepartmentAddressList) {
	_, err := orm.Table(DepartmentAddressListDAO.TableName()).
		InsertOne(addressList)
	logger.LogIfError(err)
}

// FindAll finds all addresslists
func (*DepartmentAddressListDataAccessObject) FindAll() []DepartmentAddressList {
	addressLists := make([]DepartmentAddressList, 0)
	err := orm.Table(DepartmentAddressListDAO.TableName()).Find(&addressLists)
	logger.LogIfError(err)
	return addressLists
}

// UpdateOne updates a addresslist
func (*DepartmentAddressListDataAccessObject) UpdateOne(addressList *DepartmentAddressList) {
	_, err := orm.Table(DepartmentAddressListDAO.TableName()).ID(addressList.ID).Update(addressList)
	logger.LogIfError(err)
}

// FindByID finds an addresslist by ID
func (*DepartmentAddressListDataAccessObject) FindByID(id int) (DepartmentAddressList, bool) {
	var addressList DepartmentAddressList
	has, err := orm.Table(DepartmentAddressListDAO.TableName()).Where("id=?", id).Get(&addressList)
	logger.LogIfError(err)
	return addressList, has
}

// FindByName finds a user by name
func (*DepartmentAddressListDataAccessObject) FindByName(name string) (DepartmentAddressList, bool) {
	var addressList DepartmentAddressList
	has, err := orm.Table(DepartmentAddressListDAO.TableName()).Where("name=?", name).Get(&addressList)
	logger.LogIfError(err)
	return addressList, has
}
