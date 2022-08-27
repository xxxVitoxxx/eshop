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
