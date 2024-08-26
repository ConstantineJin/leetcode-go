package main

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	mp := make(map[int]*Employee)
	for _, employee := range employees {
		mp[employee.Id] = employee
	}
	var dfs func(id int) int
	dfs = func(id int) int {
		res := mp[id].Importance
		for _, subordinate := range mp[id].Subordinates {
			res += dfs(subordinate)
		}
		return res
	}
	return dfs(id)
}
