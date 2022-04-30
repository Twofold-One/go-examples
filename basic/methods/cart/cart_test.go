package cart

import (
	"os/user"
	"testing"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/Twofold-One/go-examples/basic/methods/product"
	"github.com/stretchr/testify/assert"
)

func TestTotalPrice(t *testing.T) {
	items := []Item{
		{
			Product: product.Product{
				ID: "p-1234",
				Name: "Product1",
				Price: money.New(1000, "USD"),
			},
			Quantity: 2,
		},
		{
			Product: product.Product{
				ID: "p-12345",
				Name: "Product2",
				Price: money.New(2000, "USD"),
			},
			Quantity: 1,
		},
	}
	
	c := Cart{
		ID: "1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		User: user.User{},
		Items: items,
		CurrencyCode: "USD",
	}

	actual, err := c.TotalPrice()
	assert.NoError(t, err)
	assert.Equal(t, money.New(4000, "USD"), actual)
}

func TestLock(t *testing.T) {
	c := Cart{
		ID: "1234",
	}
	err := c.Lock()
	assert.NoError(t, err)
	assert.True(t, c.isLocked)
	assert.True(t, c.lockedAt.Unix() > 0)
}

func TestLockAlreadyLocked(t *testing.T) {
	c := Cart{
		ID: "1234",
		isLocked: true,
	}
	err := c.Lock()
	assert.Error(t, err)
}