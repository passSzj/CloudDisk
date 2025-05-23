package model

import (
	"fmt"

	"github.com/google/uuid"
	"go-cloud-disk/conf"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Uuid                 string `gorm:"primarykey"`
	UserName             string
	PasswordDigest       string
	NickName             string
	Status               string
	Avatar               string `gorm:"size:1000"`
	UserFileStoreID      string
	UserMainFileFolderID string
}

const (
	// PasswordCount password encryption difficulty
	PasswordCount = 12
	// super admin
	StatusSuperAdmin = "super_admin"
	// admin User
	StatusAdmin = "common_admin"
	// active User
	StatusActiveUser = "active"
	// inactive User
	StatusInactiveUser = "inactive"
	// Suspend User
	StatusSuspendUser = "suspend"
)

// SetPassword encrypt user password to save data
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCount)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword check user password
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

// CreateUser create user in database, and bind a userstore for user
func (user *User) CreateUser() error {
	user.Uuid = uuid.New().String()
	fileStoreId, err := CreateFileStore(user.Uuid)
	if err != nil {
		return fmt.Errorf("create file Store error %v", err)
	}
	mainFileFolderId, err := CreateBaseFileFolder(user.Uuid, fileStoreId)
	if err != nil {
		return fmt.Errorf("create base fileFolder error %v", err)
	}

	user.UserFileStoreID = fileStoreId
	user.UserMainFileFolderID = mainFileFolderId
	if err := DB.Create(user).Error; err != nil {
		return fmt.Errorf("create User error %v", err)
	}

	return nil
}

func createSuperAdmin() error {
	admin := User{
		UserName: conf.AdminUserName,
		NickName: conf.AdminUserName,
		Status:   StatusSuperAdmin,
	}

	if err := admin.SetPassword(conf.AdminPassword); err != nil {
		return fmt.Errorf("set password err when set superadmin %v", err)
	}
	if err := admin.CreateUser(); err != nil {
		return fmt.Errorf("create user err when set superadmin %v", err)
	}

	return nil
}
