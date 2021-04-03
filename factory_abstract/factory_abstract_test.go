package factory_abstract

import "testing"

func TestFactoryAbstract(t *testing.T) {
	factoryCreator := FactoryCreator{}
	fruitFactory := factoryCreator.ProducerFactory("fruit")
	vegetableFactory := factoryCreator.ProducerFactory("vegetable")

	apple := fruitFactory.GetFruitProduct("apple")
	banana := fruitFactory.GetFruitProduct("banana")
	tomato := vegetableFactory.GetVegetableProduct("tomato")
	potato := vegetableFactory.GetVegetableProduct("potato")

	if apple.GetFruitName() != "Apple" {
		t.Error("produce apple error")
	}
	if banana.GetFruitName() != "Banana" {
		t.Error("produce banana error")
	}
	if tomato.GetVegetableQuantity() != 1 {
		t.Error("produce tomato error")
	}
	if potato.GetVegetableQuantity() != 2 {
		t.Error("produce potato error")
	}
}
