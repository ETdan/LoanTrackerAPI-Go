package domain

type SystemLog struct {
	Login_Successful_Attempts string `json:"login_Successful_attempts" bson:"login_Successful_attempts"`
	Login_Failed_Attempts     string `json:"login_Failed_attempts" bson:"login_Failed_attempts"`
	Approved_loans            string `json:"approved_loans" bson:"approved_loans"`
	Rejected_loans            string `json:"rejected_loans" bson:"rejected_loans"`
	Pending_loans             string `json:"pending_loans" bson:"pending_loans"`
	Password_Reset_Completion string `json:"Password_Reset_Completion" bson:"Password_Reset_Completion"`
	Password_Reset_Request    string `json:"Password_Reset_Request" bson:"Password_Reset_Request"`
}

type SystemLogRepository interface {
	GetSystemLog() (SystemLog, error)
	UpdateSystemLog(systemLog SystemLog) error
}
