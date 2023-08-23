package erp


func (m *ModelUser) TableName() string {
	return "erp_user"
}
func (m *ModelRole) TableName() string {
	return "erp_role"
}
func (m *ModelMenu) TableName() string {
	return "erp_menu"
}
func (m *ModelUserRole) TableName() string {
	return "erp_user_role"
}
func (m *ModelRoleMenu) TableName() string {
	return "erp_role_menu"
}
func (m *ModelAccount) TableName() string {
	return "erp_account"
}
func (m *ModelUserAccount) TableName() string {
	return "erp_user_account"
}
func (m *ModelExpense) TableName() string {
	return "erp_expense"
}
func (m *ModelIncome) TableName() string {
	return "erp_income"
}
func (m *ModelPlan) TableName() string {
	return "erp_plan"
}