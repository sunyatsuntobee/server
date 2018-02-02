package models

import "github.com/sunyatsuntobee/server/logger"

// PhotoLive Model
type PhotoLive struct {
	ID                    int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	ExpectMembers         int    `xorm:"expect_members INT NOTNULL"`
	AdProgress            string `xorm:"ad_progress VARCHAR(100) NOTNULL"`
	ActivityStageID       int    `xorm:"activity_stage_id INT"`
	ManagerID             int    `xorm:"manager_id INT"`
	PhotographerManagerID int    `xorm:"photographer_manager_id INT"`
}

type PhotoLiveDataAccessObject struct{}

var PhotoLiveDAO *PhotoLiveDataAccessObject

func (*PhotoLiveDataAccessObject) TableName() string {
	return "photo_lives"
}

func (*PhotoLiveDataAccessObject) InsertOne(photoLive *PhotoLive) {
	_, err := orm.Table(PhotoLiveDAO.TableName()).InsertOne(photoLive)
	logger.LogIfError(err)
}

func (*PhotoLiveDataAccessObject) FindAll() []PhotoLive {
	l := make([]PhotoLive, 0)
	err := orm.Table(PhotoLiveDAO.TableName()).Find(&l)
	logger.LogIfError(err)
	return l
}

func (*PhotoLiveDataAccessObject) UpdateByID(id int, photoLive *PhotoLive) {
	_, err := orm.Table(PhotoLiveDAO.TableName()).ID(id).Update(photoLive)
	logger.LogIfError(err)
}

func (*PhotoLiveDataAccessObject) DeleteByID(id int) {
	var photoLive PhotoLive
	_, err := orm.Table(PhotoLiveDAO.TableName()).ID(id).Unscoped().
		Delete(&photoLive)
	logger.LogIfError(err)
}

func (*PhotoLiveDataAccessObject) FindFullAll() []PhotoLiveFull {
	l := make([]PhotoLiveFull, 0)
	err := orm.Table(PhotoLiveDAO.TableName()).
		Join("INNER", ActivityStageDAO.TableName(),
			"photo_lives.activity_stage_id=activity_stages.id").
		Join("INNER", ActivityDAO.TableName(),
			"activity_stages.activity_id=activities.id").
		Join("INNER", OrganizationDAO.TableName(),
			"activities.organization_id=organizations.id").
		Join("INNER", UserDAO.TableName(),
			"photo_lives.manager_id=users.id").
		Join("INNER", UserDAO.TableName(),
			"photo_lives.photographer_manager_id=users.id").
		Join("INNER", "photo_live_supervisor_relationships",
			"photo_lives.id=photo_live_supervisor_relationships.photo_live_id").
		Join("INNER", UserDAO.TableName(),
			"photo_live_supervisor_relationships.supervisor_id=users.id").
		Find(&l)
	logger.LogIfError(err)
	return l
}

func (*PhotoLiveDataAccessObject) FindFullByID(id int) []PhotoLiveFull {
	l := make([]PhotoLiveFull, 0)
	err := orm.Table(PhotoLiveDAO.TableName()).
		Where("photo_lives.id=?", id).
		Join("INNER", ActivityStageDAO.TableName(),
			"photo_lives.activity_stage_id=activity_stages.id").
		Join("INNER", ActivityDAO.TableName(),
			"activity_stages.activity_id=activities.id").
		Join("INNER", OrganizationDAO.TableName(),
			"activities.organization_id=organizations.id").
		Join("INNER", UserDAO.TableName(),
			"photo_lives.manager_id=users.id").
		Join("INNER", UserDAO.TableName(),
			"photo_lives.photographer_manager_id=users.id").
		Join("INNER", "photo_live_supervisor_relationships",
			"photo_lives.id=photo_live_supervisor_relationships.photo_live_id").
		Join("INNER", UserDAO.TableName(),
			"photo_live_supervisor_relationships.supervisor_id=users.id").
		Find(&l)
	logger.LogIfError(err)
	return l
}
