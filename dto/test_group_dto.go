package dto

type CreateTestGroupRequest struct {
	TestGroupName string `json:"test_group_name" binding:"required,min=3,max=100"`
}

type TestGroupResponse struct {
	TestGroupID   string `json:"test_group_id"`
	TestGroupName string `json:"test_group_name"`
}
// doubt -- testgroup response what ? cleared
//how a we are gonna get test_group by userid if we won't get user id in request