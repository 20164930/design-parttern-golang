package builder

import "testing"

func TestBuilder(t *testing.T) {
	director := Director{}
	builderA := GetABuilder()
	builderB := GetBBuilder()
	mealA := director.buyMeal(builderA)
	mealB := director.buyMeal(builderB)
	if mealA.humBurger != 1 || mealA.pizza!=1 || mealA.cola!=1 || mealA.chips != 1 {
		t.Error("mealA is error")
	}
	if mealB.humBurger != 2 || mealB.pizza!=0 || mealB.cola!=1 || mealB.chips != 2 {
		t.Error("mealB is error")
	}
}
