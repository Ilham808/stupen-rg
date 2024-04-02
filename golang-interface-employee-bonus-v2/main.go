package main

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name         string
	BaseSalary   int
	WorkingMonth int
}

type Senior struct {
	Name            string
	BaseSalary      int
	WorkingMonth    int
	PerformanceRate float64
}

type Manager struct {
	Name             string
	BaseSalary       int
	WorkingMonth     int
	PerformanceRate  float64
	BonusManagerRate float64
}

func (m Manager) GetBonus() float64 {
	prorata := float64(m.WorkingMonth) / 12
	if prorata > 1 {
		prorata = 1
	}
	return 2*float64(m.BaseSalary)*prorata + (m.PerformanceRate * float64(m.BaseSalary)) + (m.BonusManagerRate * float64(m.BaseSalary))
}

func (s Senior) GetBonus() float64 {
	prorata := float64(s.WorkingMonth) / 12
	if prorata > 1 {
		prorata = 1
	}
	return 2*float64(s.BaseSalary)*prorata + (s.PerformanceRate * float64(s.BaseSalary))
}

func (j Junior) GetBonus() float64 {
	prorata := float64(j.WorkingMonth) / 12
	if prorata > 1 {
		prorata = 1
	}
	return float64(j.BaseSalary) * prorata
}

func EmployeeBonus(employee Employee) float64 {
	return employee.GetBonus()
}

func TotalEmployeeBonus(employees []Employee) float64 {
	totalBonus := 0.0
	for _, emp := range employees {
		totalBonus += emp.GetBonus()
	}
	return totalBonus
}
