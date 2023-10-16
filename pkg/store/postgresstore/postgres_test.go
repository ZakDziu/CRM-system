package postgresstore_test

import (
	"crm-system/pkg/config"
	"crm-system/pkg/logger"
	"crm-system/pkg/model"
	"crm-system/pkg/store/postgresstore"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)

type StoreSuite struct {
	suite.Suite
	store *postgresstore.PostgresStore

	AuthUserFixture *postgresstore.FixtureAuthUser
	UserFixture     *postgresstore.FixtureUser
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(StoreSuite))
}

func (s *StoreSuite) SetupSuite() {

	cf, err := config.New()
	if err != nil {
		logger.Fatalf("Can't read config file: %s", err.Error())
	}
	conf := cf.DBPostgresConfig

	db, err := postgresstore.NewPostgresStore(&conf)
	s.Nil(err)
	s.store = db

	s.AuthUserFixture = postgresstore.NewFixtureAuthUser()
	s.UserFixture = postgresstore.NewFixtureUser()

	s.cleanDB()
}

func (s *StoreSuite) cleanDB() {
	s.store.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.User{})
	s.store.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.AuthUser{})

}

func (s *StoreSuite) TearDownTest() {
	s.cleanDB()
}
