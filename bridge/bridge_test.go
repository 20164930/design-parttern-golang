package bridge

import "testing"

func Test_bridge(t *testing.T) {
	huaweiPhoneA := &huaweiPhone{&phone{&Red{}}, "nova"}
	huaweiPhoneB := &huaweiPhone{&phone{&Green{}}, "Mate"}
	if huaweiPhoneA.getAttribute() != "nova/Red:#DC143C" {
		t.Error("huaweiPhoneA is incorrect")
	}
	if huaweiPhoneB.getAttribute() != "Mate/Green:#00FF00" {
		t.Error("huaweiPhoneB is incorrect")
	}
}
