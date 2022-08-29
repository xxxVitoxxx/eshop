package mysqldb

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xxxVitoxxx/eshop/pkg/numberplate"
	"github.com/xxxVitoxxx/eshop/storage/conn"
)

func TestPutConditionByStoreName(t *testing.T) {
	db := conn.ConnectToMySQL()
	repo := NewConditionRepo(db)
	db.AutoMigrate(&Condition{})
	defer db.Migrator().DropTable(&Condition{})

	// mock
	mock := Condition{
		StoreName: "eshop",
		HowMany:   20,
		HowLong:   60,
		Remind:    5,
	}
	db.Create(mock)

	t.Run("PutConditionByStoreName success", func(t *testing.T) {
		condition := numberplate.PutCondition{
			HowMany: 10,
			HowLong: 45,
			Remind:  5,
		}
		assert.NoError(t, repo.PutConditionByStoreName(mock.StoreName, condition))

		findDB := Condition{}
		db.Find(&findDB)
		assert.Equal(t, findDB.HowMany, condition.HowMany)
		assert.Equal(t, findDB.HowLong, condition.HowLong)
		assert.Equal(t, findDB.Remind, condition.Remind)
	})
}

func TestCreateCondition(t *testing.T) {
	db := conn.ConnectToMySQL()
	repo := NewConditionRepo(db)
	db.AutoMigrate(&Condition{})
	defer db.Migrator().DropTable(&Condition{})

	t.Run("CreateCondition success", func(t *testing.T) {
		condition := numberplate.Condition{
			StoreName: "eshop",
			HowMany:   15,
			HowLong:   50,
			Remind:    10,
		}
		err := repo.CreateCondition(condition)
		assert.NoError(t, err)

		var findDB []Condition
		db.Find(&findDB)
		assert.Equal(t, condition.StoreName, findDB[0].StoreName)
		assert.Equal(t, condition.HowMany, findDB[0].HowMany)
		assert.Equal(t, condition.HowLong, findDB[0].HowLong)
		assert.Equal(t, condition.Remind, findDB[0].Remind)
	})
}

func TestDeleteConditionByStoreName(t *testing.T) {
	db := conn.ConnectToMySQL()
	repo := NewConditionRepo(db)
	db.AutoMigrate(&Condition{})
	defer db.Migrator().DropTable(&Condition{})

	// mock
	mock := []Condition{
		{
			StoreName: "eshop",
			HowMany:   20,
			HowLong:   60,
			Remind:    5,
		},
		{
			StoreName: "two circle",
			HowMany:   15,
			HowLong:   50,
			Remind:    10,
		},
	}
	db.Create(mock)

	t.Run("DeleteConditionByStoreName success", func(t *testing.T) {
		assert.NoError(t, repo.DeleteConditionByStoreName("two circle"))

		findDB := []Condition{}
		db.Find(&findDB)
		assert.Equal(
			t,
			[]Condition{{
				StoreName: "eshop",
				HowMany:   20,
				HowLong:   60,
				Remind:    5,
			}},
			findDB)
	})
}
