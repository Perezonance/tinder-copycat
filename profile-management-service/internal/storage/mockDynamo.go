package storage

//MockDynamo emulates a key value store db with an Articles table
type MockDynamo struct {
	c Config
	// ArticlesTable map[int]models.Profile
}

//NewMockDynamo creates a new MockDynamo DB with a blank Articles table
func NewMockDynamo() *MockDynamo {
	return &MockDynamo{}
}
