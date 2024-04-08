package main

type person struct {
	alive    bool
	children []string
}

type ThroneInheritance struct {
	root   string
	people map[string]person
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{root: kingName, people: map[string]person{kingName: {alive: true}}}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	this.people[childName] = person{alive: true}
	this.people[parentName] = person{alive: true, children: append(this.people[parentName].children, childName)}
}

func (this *ThroneInheritance) Death(name string) {
	this.people[name] = person{alive: false, children: this.people[name].children}
}

func (this *ThroneInheritance) GetInheritanceOrder() (ans []string) {
	var dfs func(name string)
	dfs = func(name string) {
		if this.people[name].alive {
			ans = append(ans, name)
		}
		for _, p := range this.people[name].children {
			dfs(p)
		}
	}
	dfs(this.root)
	return
}
