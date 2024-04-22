package user


import "gorm.io/gorm"
	

// ini interface
type RepositoryUser interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
}



// kontract db nya
type repositoryUser struct{
	db *gorm.DB
}

// ini menginstialisasikan menjadi db yg baru
func NewRepositoryUser(db *gorm.DB) *repositoryUser {
	return &repositoryUser{db: db}
}




func (r *repositoryUser) Save(user User) (User, error) {
     err := r.db.Create(&user).Error

	 if err != nil {
		return user, err
	 }

	 return user, nil
}


func (r *repositoryUser) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
