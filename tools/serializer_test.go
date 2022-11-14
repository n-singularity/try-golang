package tools_test

import (
	"executive-monitoring/config"
	"executive-monitoring/models/aep"
	monitoring "executive-monitoring/models/monitoring"
	swaggerstructs "executive-monitoring/swagger_structs/aep"
	"executive-monitoring/tools"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"gotest.tools/assert"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func TestSerialize(t *testing.T) {
	re := regexp.MustCompile(`^(.*` + "executive-monitoring" + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		t.Error(err.Error())
		return
	}

	config.ConnectDBTest()
	db := config.DBPostgres.Connect
	db.AutoMigrate(&monitoring.ExmonAgent{})
	db.AutoMigrate(&aep.AepEvent{})
	db.AutoMigrate(&aep.AepParticipant{})
	db.AutoMigrate(&aep.AepTypes{})

	name := "jalan-jalan"
	active := true
	timeNow := time.Now()
	id := 1

	agentCode := "21"
	aepEventTypes := aep.AepTypes{
		ID:        id,
		Name:      &name,
		Active:    &active,
		CreatedAt: &timeNow,
		UpdatedAt: &timeNow,
	}

	timeNow = time.Date(2021, 03, 01, 13, 33, 30, 0, time.UTC)

	aepEvent := aep.AepEvent{
		ID:          2,
		AepTypeID:   id,
		Active:      true,
		StartTime:   timeNow,
		EndTime:     timeNow,
		Title:       "Sequis Recruiter",
		Location:    "Jakarta",
		Description: "Agent baru",
		AgentCode:   "32",
	}

	aepParticipant := aep.AepParticipant{
		ID:         1,
		AgentCode:  agentCode,
		AepEventID: aepEvent.ID,
		Active:     true,
		StatusedAt: &timeNow,
	}

	participantName := "RaidenX"
	var i int = 21
	participantCode := int64(i)

	exmonAgent := monitoring.ExmonAgent{
		AgentName: participantName,
		AgentCode: &participantCode,
	}

	creatorName := "Rusli"
	var z int = 32
	creatorCode := int64(z)

	exmonAgent2 := monitoring.ExmonAgent{
		AgentName: creatorName,
		AgentCode: &creatorCode,
	}

	db.Create(&aepEventTypes)
	db.Create(&aepEvent)
	db.Create(&aepParticipant)
	db.Create(&exmonAgent)
	db.Create(&exmonAgent2)

	serializer := TestJSONAEPEventDetailSerialize{}

	var aepEventData aep.AepEvent
	db.Preload("AepParticipants.Agent").
		Preload("Agent").
		First(&aepEventData)

	got := tools.Serialize(aepEventData, &serializer, serializer).(TestJSONAEPEventDetailSerialize)

	assert.DeepEqual(t, got.Location, aepEvent.Location)
	assert.DeepEqual(t, got.Description, aepEvent.Description)
	assert.DeepEqual(t, got.Title, aepEvent.Title)
	assert.DeepEqual(t, got.CreatorAgentCode, aepEvent.AgentCode)
	assert.DeepEqual(t, got.CreatorName, creatorName)
	assert.DeepEqual(t, got.StartAt, timeNow)
	assert.DeepEqual(t, got.EndAt, timeNow)

	AgentParticipants := got.AgentParticipants.([]aep.AgentParticipantSerializer)
	assert.DeepEqual(t, AgentParticipants[0].AgentCode, strconv.Itoa(int(participantCode)))
	assert.DeepEqual(t, AgentParticipants[0].AgentName, participantName)

	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE;", db.NewScope(&aep.AepEvent{}).TableName()))
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE;", db.NewScope(&aep.AepParticipant{}).TableName()))
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE;", db.NewScope(&aep.AepTypes{}).TableName()))
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE;", db.NewScope(&monitoring.ExmonAgent{}).TableName()))
}

type TestJSONAEPEventDetailSerialize struct {
	TestJSONAEPEventDetail
	AgentParticipants interface{} `json:"agent_participants"`
}

type TestJSONAEPEventDetail struct {
	ID                int64                            `json:"id" example:1`
	AepTypeID         int                              `json:"aep_type_id" example:1`
	Title             string                           `json:"title" example:"asuransi"`
	Location          string                           `json:"location" example:"jakarta"`
	Description       string                           `json:"description" example:"ayoo masukk!"`
	StartAt           time.Time                        `json:"start_at"  example:"2021-02-02T07:03:15Z"`
	EndAt             time.Time                        `json:"end_at" example:"2021-02-02T07:09:15Z"`
	CreatorName       string                           `json:"creator_name" example:"Budi"`
	CreatorAgentCode  string                           `json:"creator_agent_code" example:"123"`
	AgentParticipants swaggerstructs.AgentParticipants `json:"agent_participants"` //keperluan swagger
}
