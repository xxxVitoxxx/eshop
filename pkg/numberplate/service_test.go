package numberplate_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xxxVitoxxx/eshop/pkg/numberplate"
	"github.com/xxxVitoxxx/eshop/storage/fake"
)

func TestPutCondition(t *testing.T) {
	fakeRepo := fake.NewConditionRepo()
	s := numberplate.NewService(fakeRepo)

	// mock condition
	fakeRepo.HowMany = 20
	fakeRepo.HowLong = 60
	fakeRepo.Remind = 5

	t.Run("PutCondition success", func(t *testing.T) {
		c := numberplate.PutCondition{
			HowMany: 15,
			HowLong: 40,
			Remind:  10,
		}
		assert.NoError(t, s.PutCondition(c))
		assert.Equal(t, c.HowMany, fakeRepo.HowMany)
		assert.Equal(t, c.HowLong, fakeRepo.HowLong)
		assert.Equal(t, c.Remind, fakeRepo.Remind)
	})
}
