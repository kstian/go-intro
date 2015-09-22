package main
import (
	"time"
	"sort"
	"fmt"
	"github.com/ktarou/gointro/model"
)

func main() {
	const DATE_LAYOUT =  "2006-02-01"
	date := func(date string) time.Time {
		dt,_ := time.Parse(DATE_LAYOUT, date)
		return dt
	}
	print := fmt.Println
	patient := []model.Patient{
		{1, "Alan", "Duron",date("1990-01-11")},
		{2, "Dadang", "Sihombing", date("1998-01-11")},
		{3, "Bobby", "Edward", date("1950-01-11")},
	}
	print("BEFORE SORT", patient)
	sort.Sort(model.ByFirstName(patient))
	print("AFTER SORT BY FIRSTNAME", patient)
	sort.Sort(model.ByLastName(patient))
	print("AFTER SORT BY LASTNAME", patient)
	sort.Sort(model.ByDob(patient))
	print("AFTER SORT BY DOB", patient)
}