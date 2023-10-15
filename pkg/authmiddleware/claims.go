package authmiddleware

import (
	"crm-system/pkg/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

const (
	AccessTokenTTL  = time.Hour * 8
	RefreshTokenTTL = time.Hour * 24 * 7
)

type BaseClaims struct {
	jwt.StandardClaims
	ID   uuid.UUID `json:"id"`
	Role model.UserRole
}

// Claims represent set of fields stored in JWT payload.
type AccessClaims struct {
	BaseClaims
	AccessUUID string `json:"access_uuid"`
}

// RefreshClaims represent JSON Web Token Claims for refresh token.
type RefreshClaims struct {
	BaseClaims
	RefreshUUID string `json:"refresh_uuid"`
}

// Tokens represent user tokens
type Tokens struct {
	Access  string `json:"accessToken"`
	Refresh string `json:"refreshToken"`
}

// NewClaims method for create BaseClaims with StandartClaims and user data
func NewClaims(idClaims uuid.UUID, role model.UserRole, ttl time.Duration) BaseClaims {
	return BaseClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			Id:        uuid.NewV4().String(),
			IssuedAt:  time.Now().Unix(),
		},
		ID:   idClaims,
		Role: role,
	}
}

// Generate Access and Refresh Claims
func GenerateClaims(idClaims uuid.UUID, role model.UserRole) (*AccessClaims, *RefreshClaims) {
	access := AccessClaims{
		BaseClaims: NewClaims(idClaims, role, AccessTokenTTL),
	}

	refresh := RefreshClaims{
		BaseClaims: NewClaims(idClaims, role, RefreshTokenTTL),
	}

	access.AccessUUID = refresh.Id
	refresh.RefreshUUID = access.Id

	return &access, &refresh
}
