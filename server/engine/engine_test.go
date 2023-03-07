package engine

import (
	"server/azure"
	"server/config"
	"server/controllers"
	"server/model"
	"server/persistence"
	"testing"
)

var testCount = 3

func TestGetLatestExecution(t *testing.T) {
	testStep := model.Step{}
	testStep.ID = 5 //Arbitrary stepID
	testdb := persistence.NewPersistentDB(config.GetEnvironment().DB_PATH)
	testDeploymentsClient := azure.NewDeploymentsClient(nil)
	//Create new engine to get latest execution
	engine := NewEngine(controllers.NewExitController().Context(), testdb, testDeploymentsClient)
	latestExecution := engine.GetLatestExecution(testStep)
	//Check if returned value is invalid
	if latestExecution.Status == "Failed" {
		t.Error("Incorrect output")
	}

	//Check if latestExecution returns count <= 0
	if latestExecution.database.Instance.Model(&model.Execution{}).Where("step_id = ?", testStep.ID).Count() <= 0 {
		t.Error("Error: Count <= 0")
	}

	// engine, err := &Engine{
	// 	database: &persistence.Database{
	// 		Instance: &gorm.DB{}, //Mock gorm.DB call?
	// 	},
	// }
	// if(engine.database.Instance == nil) {
	// 	t.Fatal(err)
	// }

	// //Expected values
	// stepID := 2
	// expectedExecution := model.Execution {
	// 	StepID: int64(stepID),
	// }
	// //Count <= 0
	// if() {

	// }
	// //Count > 0

}
