package simplefactory

import (
	"testing"
)

func TestFactoryMethod(t *testing.T) {

	assert := []string{"A", "B", "C"}

	factory := &Factory{}
	products := []Product{
		factory.Create(assert[0]),
		factory.Create(assert[1]),
		factory.Create(assert[2]),
	}

	for i, product := range products {
		if product.getQuality() != i + 1 {
			t.Errorf("create product error : %s quality is %d", assert[i], product.getQuality())
		}
	}

}
