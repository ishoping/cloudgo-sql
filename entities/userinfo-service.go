package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

var userInfoInsertStmt = "INSERT userinfo SET username=?,departname=?,created=?"
var userInfoQueryAll = "SELECT * FROM userinfo"

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {

	session := engine.NewSession()
	defer session.Close()

	// add Begin() before any action
	err := session.Begin()

	_, err = session.Exec(userInfoInsertStmt, u.UserName, u.DepartName, u.CreateAt)
	if err != nil {
		session.Rollback()
		return err
	}

	// add Commit() after all actions
	err = session.Commit()
	if err != nil {
		return err
	}

	return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {

	results, err := engine.QueryString(userInfoQueryAll)
	checkErr(err)

	ulist := make([]UserInfo, 0, 0)
	for _, result := range results {
		u := UserInfo{}
		u.FillStruct(conv(result))
		ulist = append(ulist, u)
	}
	return ulist
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	var result = make(map[string]string)
	has, err := engine.Where("uid = ?", id).Get(&result)
	checkErr(err)
	if has {
		u := UserInfo{}
		u.FillStruct(conv(result))
		return &u
	}
	return &UserInfo{}
}
