package numberplate_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xxxVitoxxx/eshop/pkg/numberplate"
	"github.com/xxxVitoxxx/eshop/pkg/storage/fake"
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

	t.Run("CreateCondition success", func(t *testing.T) {
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

func TestDeleteConditionByStoreName(t *testing.T) {
	fakeRepo := fake.NewConditionRepo()
	s := numberplate.NewService(fakeRepo)

	// mock condition
	mockCondition := []numberplate.Condition{
		{
			StoreName: "eshop",
			HowMany:   20,
			HowLong:   50,
			Remind:    5,
		},
		{
			StoreName: "two circle",
			HowMany:   20,
			HowLong:   60,
			Remind:    10,
		},
	}
	fakeRepo.Condition = append(fakeRepo.Condition, mockCondition...)

	t.Run("DeleteConditionByStoreName success", func(t *testing.T) {
		assert.NoError(t, s.DeleteConditionByStoreName("eshop"))
		assert.Len(t, fakeRepo.Condition, 1)
		assert.Equal(t, mockCondition[1], fakeRepo.Condition[0])
	})
}
