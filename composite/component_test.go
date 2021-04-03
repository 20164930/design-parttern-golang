package composite

import "testing"

func TestComponent(t *testing.T) {
	// 组织A是最高级别，组织B是组织A的下级组织,A有成员101，B有成员100
	groupLevelB := group{[]Component{&employee{"100"}}}
	groupLevelA := group{[]Component{&employee{"101"},&groupLevelB}}
	groupLevelA.printComponent()
}
