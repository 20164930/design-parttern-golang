package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {

	instance1 := &singleton{}
	instance2 := &singleton{}
	for i:=0;i <100;i++ {
		go createInstance(instance1)
		go createInstance(instance2)
		if instance1 != instance2 {
			t.Errorf("Expect instance to equal, but not equal.\n")
		}
	}

}
func createInstance(single *singleton){
	single = GetInstance()
}
