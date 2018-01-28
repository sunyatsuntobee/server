package models

type Activities struct {
	ID	int		`xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	Name    string	`xorm:"name VARCHAR(45) NOTNULL"`
	Description		string	`xorm:"description VARCHAR(45) NOTNULL"`
	Category    string	`xorm:"category VARCHAR(45) NOTNULL"`
	PosterUrl	string `xorm:"poster_url VARCHAR(45)"`
	Logo	[]int    `xorm:"logo BLOB"`
}



type Organization_contact_relatonships struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	OrganizationID int `xorm:"organization_id INT NOTNULL INDEX(organization_id_idx)"`
	ContactID int `xorm:"contact_id INT NOTNULL INDEX(contact_id_idx)"`
}

type Organzation_departments struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	Name string `xorm:"name VARCHAR(45) NOTNULL"`
	OrganizationID int `xorm:"organization_id INt NOTNULL INDEX(organization_id_idx)"`
}

type Organzation_login_logs struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	LoginTime time.Time `xorm:"login_time NOTNULL"`
	LoginLocation string `xorm:"login_location NOTNULL"`
	LoginDevice string `xorm:"login_device NOTNULL"`
	OrganizationID int `xorm:"organization_id NOTNULL INDEX(organization_id_idx)"`
}

type Organzations struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	Name string `xorm:"name VARCHAR(45) NOTNULL"`
	Phone string `xorm:"phone VARCHAR(45) NOTNULL UNIQUE"`
	Password string `xorm:"password VARCHAR(45) NOTNULL"`
	Collage string `xorm:"collage VARCHAR(45) NOTNULL"`
	Logo []int `xorm:"logo BLOB"`
	Description string `xorm:"description VARCHAR(45)"`
}

type Photo_comments struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	Title string `xorm:"title VARCHAR(45) NOTNULL"`
	Content string `xorm:"content VARCHAR(45) NOTNULL"`
	UserID int `xorm:"user_id INT NOTNULL INDEX(user_id_idx)"`
	PhotoID int `xorm:"photo_id INT NOTNULL INDEX(photo_id_idx)"`
}

type Photo_live_supervisor_relationships struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	PhotoLiveID int `xorm:"photo_live_id INT NOTNULL INDEX(photo_live_id_idx)"`
	PhotoSuperVisorRelationshipsCol int `xorm:"photo_supervisor_relationshipscol INT NOTNULL INDEX(supervisor_id_idx)"`
}

type Photo_lives struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	ExpectMembers int `xorm:"expect_members INT NOTNULL"`
	AdProgress string `xorm:"ad_progress VARCHAR(45) NOTNULL"`
	ActivityStageID int `xorm:"activity_stage_id INT"`
	ManagerID int `xorm:"manager_id INT"`
	PhotographerManagerID int `xorm:"photographer_manager_id INT"`
}

type Photo_tags struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	Tag string `xorm:"tag VARCHAR(45) NOTNULL"`
	PhotoID int `xorm:"photo_id INT NOTNULL UNIQUE INDEX(photo_id_UNIQUE)"`
}

type User_activity_relationships struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	UserID int `xorm:"user_id INT NOTNULL INDEX(user_id_idx)"`
	ActivityID int `xorm:"activity_id INT NOTNULL(activity_id_idx)"`
}

type User_login_logs struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	LoginTime time.Time `xorm:"login_time DATETIME NOTNULL"`
	LoginLocation string `xorm:"login_location VARCHAR(45) NOTNULL"`
	LoginDevice string `xorm:"login_device VARCHAR(45) NOTNULL"`
	UserID int `xorm:"user_id INT NOTNULL INDEX(user_id_idx)"`
}

type User_organization_relationships struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	UserID int `xorm:"user_id NOTNULL INDEX(user_id_idx)"`
	OrganizationID int `xorm:"organization_id INT NOTNULL INDEX(organization_id_idx)"`
}

type User_photo_relationships struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	UserID int `xorm:"user_id INT NOTNULL INDEX(user_id_idx)"`
	LikedPhotoID int `xorm:"liked_photo_id INT NOTNULL INDEX(liked_photo_id_idx)"`
}

type User_user_relationships struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	UserID int `xorm:"user_id INT NOTNULL INDEX(user_id_idx)"`
	LikedUserID int `xorm:"liked_user_id INT NOTNULL INDEX(liked_user_id_idx)"`
}