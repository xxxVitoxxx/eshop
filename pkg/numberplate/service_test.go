package numberplate_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xxxVitoxxx/eshop/pkg/numberplate"
	"github.com/xxxVitoxxx/eshop/storage/fake"
)

func TestPutCondition(t *testing.T) {
	fakeNPRepo := fake.NewNumberPlateRepo()
	s := numberplate.NewService(fakeNPRepo)

	// mock condition
	fakeNPRepo.HowMany = 20
	fakeNPRepo.HowLong = 60
	fakeNPRepo.Remind = 5

	t.Run("PutCondition success", func(t *testing.T) {
		c := numberplate.PutCondition{
			HowMany: 15,
			HowLong: 40,
			Remind:  10,
		}
		assert.NoError(t, s.PutCondition(c))
		assert.Equal(t, c.HowMany, fakeNPRepo.HowMany)
		assert.Equal(t, c.HowLong, fakeNPRepo.HowLong)
		assert.Equal(t, c.Remind, fakeNPRepo.Remind)
	})
}
