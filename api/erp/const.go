package erp

var ExpenseCategoryMap = map[string]uint32{
	"其他支出":      uint32(ExpenseCategory_ExpenseCategoryOther),
	"采购支出":      uint32(ExpenseCategory_ExpenseCategoryPurchase),
	"人力支出":      uint32(ExpenseCategory_ExpenseCategoryHR),
	"办公杂费":      uint32(ExpenseCategory_ExpenseCategoryOfficeOther),
	"借款支出":      uint32(ExpenseCategory_ExpenseCategoryBorrow),
	"外发加工":      uint32(ExpenseCategory_ExpenseCategoryProcess),
	"采购_办公用品":   uint32(ExpenseCategory_ExpenseCategoryOffice),
	"采购_食堂伙食":   uint32(ExpenseCategory_ExpenseCategoryCanteen),
	"人力支出_员工工资": uint32(ExpenseCategory_ExpenseCategorySalary),
	"人力支出_临时工":  uint32(ExpenseCategory_ExpenseCategoryPartTime),
	"其他支出_备用金":  uint32(ExpenseCategory_ExpenseCategoryStandby),
	"头程支出":      uint32(ExpenseCategory_ExpenseCategoryHead),
	"推广_测评":     uint32(ExpenseCategory_ExpenseCategoryEvaluation),
	"采购_产品采购":   uint32(ExpenseCategory_ExpenseCategoryProductPurchase),
	"运费-国内运费":   uint32(ExpenseCategory_ExpenseCategoryDeliveryCN),
}

var IncomeCategoryMap = map[string]uint32{
	"其他收入":     uint32(IncomeCategory_IncomeCategoryOther),
	"其他收入_备用金": uint32(IncomeCategory_IncomeCategoryStandby),
	"销售收入":     uint32(IncomeCategory_IncomeCategorySale),
}
