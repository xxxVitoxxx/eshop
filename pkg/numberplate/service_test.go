package numberplate_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xxxVitoxxx/eshop/pkg/numberplate"
	"github.com/xxxVitoxxx/eshop/storage/fake"
)

func TestPutConditionByStoreName(t *testing.T) {
	fakeRepo := fake.NewConditionRepo()
	s := numberplate.NewService(fakeRepo)

	// mock condition
	fakeRepo.Condition = append(fakeRepo.Condition, []numberplate.Condition{
		{
			StoreName: "eshop",
			HowMany:   15,
			HowLong:   45,
			Remind:    5,
		},
		{
			StoreName: "two circle",
			HowMany:   10,
			HowLong:   50,
			Remind:    10,
		},
	}...)

	t.Run("PutConditionByStoreName success", func(t *testing.T) {
		c := numberplate.PutCondition{
			HowMany: 20,
			HowLong: 40,
			Remind:  10,
		}
		assert.NoError(t, s.PutConditionByStoreName("eshop", c))
		assert.Equal(t, c.HowMany, fakeRepo.Condition[0].HowMany)
		assert.Equal(t, c.HowLong, fakeRepo.Condition[0].HowLong)
		assert.Equal(t, c.Remind, fakeRepo.Condition[0].Remind)
	})
}

func TestCreateCondition(t *testing.T) {
	fakeRepo := fake.NewConditionRepo()
	s := numberplate.NewService(fakeRepo)

	t.Run("CreateCondition", func(t *testing.T) {
		condition := numberplate.Condition{
			StoreName: "eshop",
			HowMany:   10,
			HowLong:   40,
			Remind:    10,
		}
		assert.NoError(t, s.CreateCondition(condition))
		assert.Equal(t, condition, fakeRepo.Condition[0])
	})
}
