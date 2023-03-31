package tests

//type TestSuiteEnv struct {
//	suite.Suite
//	DB *gorm.DB
//}
//
//// SetupSuite Tests are run before they start
//func (suite *TestSuiteEnv) SetupSuite() {
//	//initiate repository
//	r := repositories.NewRepository(suite.DB)
//
//	suite.db = database.GetDB()
//}
//
//// TearDownTest Running after each test
//func (suite *TestSuiteEnv) TearDownTest() {
//	suite.DB.ClearTable()
//}
//
//// TearDownSuite Running after all tests are completed
//func (suite *TestSuiteEnv) TearDownSuite() {
//	suite.db.Close()
//}
//
//// TestSuite This gets run automatically by `go test` so we call `suite.Run` inside it
//func TestSuite(t *testing.T) {
//	// This is what actually runs our suite
//	suite.Run(t, new(TestSuiteEnv))
//}
