package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Role struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Permission struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RolePermission struct {
	RoleID       uint `gorm:"primary_key"`
	PermissionID uint `gorm:"primary_key"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Status int `json:"status"`
	Data   any `json:"data"`
}

type ResData struct {
	Username  string `json:"userName"`
	GrantId   string `json:"grantId"`
	UserType  string `json:"userType"`
	Category  string `json:"category"`
	Sub       string `json:"sub"`
	ClientId  string `json:"clientId"`
	RealmName string `json:"realmName"`
	JTI       string `json:"atHash"`
	Ext       Ext    `json:"ext"`
	Active    bool   `json:"active"`
}
type Ext struct {
	TenantId string `json:"tenantId"`
}
