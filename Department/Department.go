package Department

type Department struct {
	DepartmentName string
	CutOffMark     int32
	Fees           float64
}

var biology = Department{"biology", 200, 40000}
var engineering = Department{"engineeering", 250, 50000}
var medicine = Department{"medicine", 290, 80000}
var fineArts = Department{"fine arts", 180, 25000}

var Depats = []Department{biology, engineering, medicine, fineArts}

func GetDepartments() []Department {
	departments := []Department{biology, engineering, medicine, fineArts}
	return departments
}

func GetDepartmentFees(depName string) float64 {
	var fees float64
	for _, val := range Depats {
		if val.DepartmentName == depName {
			fees = val.Fees
		}
	}
	return fees
}
