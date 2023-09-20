package profile

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/shirrashko/BuildingAServer-step2/config"
	"github.com/shirrashko/BuildingAServer-step2/pkg/api/profile/model"
	profileBL "github.com/shirrashko/BuildingAServer-step2/pkg/bl/profile"
	"github.com/shirrashko/BuildingAServer-step2/pkg/db"
	profileRep "github.com/shirrashko/BuildingAServer-step2/pkg/repository/profile"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type APIIntegrationTestSuite struct {
	suite.Suite
	Recorder *httptest.ResponseRecorder
	Context  *gin.Context
	Handler  Handler
	DB       *sql.DB
}

func (s *APIIntegrationTestSuite) SetupSuite() {
	// Create a mock configuration for testing
	mockConfig := config.DBConfig{
		Host:         "localhost",
		Port:         6432,
		User:         "srashkovits",
		Password:     "password",
		DatabaseName: "postgres",
	}

	// Call the NewDBClient function with the mock configuration
	dbClient, err := db.NewDBClient(mockConfig)
	if err != nil {
		s.T().Fatal("Failed to initialize DB client:", err)
	}

	s.DB = dbClient
}

func (s *APIIntegrationTestSuite) SetupTest() {
	s.Recorder = httptest.NewRecorder()
	s.Context, _ = gin.CreateTestContext(s.Recorder)

	// Assuming your NewHandler function and NewService are correct
	rep := profileRep.NewRepository(s.DB)
	service := profileBL.NewService(&rep)
	s.Handler = NewHandler(&service)

	// Assuming you have SQL file to setup schema
	// s.executeSqlFile("./schema/create_profile_table.sql", "create-profile-table")
}

func (s *APIIntegrationTestSuite) TearDownTest() {
	// Cleanup after each test
	// s.executeSqlFile("./schema/delete_profile_table.sql", "delete-profile-table")
}

func (s *APIIntegrationTestSuite) TestCreateProfile() {
	// Arrange
	newProfile := model.CreateProfileRequest{
		Profile: model.BaseUserProfile{
			Username:      "newUsername",
			FullName:      "newFullName",
			Bio:           "newBio",
			ProfilePicURL: "http://example.com/new-pic.png",
		},
	}

	marshaledProfile, err := json.Marshal(newProfile)
	if err != nil {
		panic("couldn't marshal profile")
	}

	req, _ := http.NewRequest(http.MethodPost, "/profile/users", bytes.NewBuffer(marshaledProfile))
	req.Header.Set("Content-Type", "application/json")
	s.Context.Request = req

	// Act
	s.Handler.createProfile(s.Context)

	// Assert Response
	assert.Equal(s.T(), http.StatusCreated, s.Recorder.Code)

	// Assert Insertion (This will depend on how your service returns the new ID)
	var returnedID int // This type should match whatever your service returns
	if json.Unmarshal(s.Recorder.Body.Bytes(), &returnedID) != nil {
		panic("Couldn't unmarshal returned object")
	}
	// Use returnedID to fetch the created profile from DB and compare
}

func TestProfileAPITestSuite(t *testing.T) {
	suite.Run(t, new(APIIntegrationTestSuite))
}
