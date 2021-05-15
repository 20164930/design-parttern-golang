package factory

import "testing"

func TestFactory(t *testing.T) {
	factory1 := FactoryA{}
	factory2 := FactoryB{}
	product1 := factory1.createProduct()
	product2 := factory2.createProduct()
	if product1.show() != "productA" {
		t.Error("produce productA error")
	}

	if product2.show() != "productB" {
		t.Error("produce productB error")
	}
}
