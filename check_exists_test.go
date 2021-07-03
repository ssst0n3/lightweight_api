package lightweight_api

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_api/test/model"
	"github.com/ssst0n3/lightweight_api/test/test_data"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func WrapperTestResourceCheckResourceExists(t *testing.T, executor func(id uint) bool) {
	t.Run("not exists", func(t *testing.T) {
		assert.NoError(t, test_data.InitEmptyChallenge(DB))
		exists := executor(1)
		assert.Equal(t, false, exists)
	})
	t.Run("exists", func(t *testing.T) {
		assert.NoError(t, test_data.InitEmptyChallenge(DB))
		DB.Create(&test_data.Challenge1)
		exists := executor(test_data.Challenge1.ID)
		assert.True(t, true, exists)
		assert.False(t, executor(2))
	})
}

func WrapperTestResourceCheckResourceExistsExceptSelf(t *testing.T, executor func(id uint) bool) {
	t.Run("not exists", func(t *testing.T) {
		assert.NoError(t, test_data.InitEmptyChallenge(DB))
		exists := executor(1)
		assert.Equal(t, false, exists)
	})
	t.Run("not exists 2", func(t *testing.T) {
		assert.NoError(t, test_data.InitEmptyChallenge(DB))
		DB.Create(&test_data.Challenge1)
		exists := executor(test_data.Challenge1.ID)
		assert.Equal(t, false, exists)
	})
	t.Run("exists", func(t *testing.T) {
		assert.NoError(t, test_data.InitEmptyChallenge(DB))
		DB.Create(&test_data.Challenge1)
		DB.Create(&test_data.ChallengeSameName)
		exists := executor(test_data.Challenge1.ID)
		assert.Equal(t, true, exists)
	})
}

func TestResource_CheckResourceExistsByIdAutoParseParam(t *testing.T) {
	WrapperTestResourceCheckResourceExists(t, func(i uint) bool {
		c := &gin.Context{Params: []gin.Param{
			{Key: "id", Value: strconv.Itoa(int(i))},
		}}
		exists, id, err := challenge.CheckResourceExistsByIdAutoParseParam(c)
		assert.NoError(t, err)
		assert.Equal(t, int64(i), id)
		return exists
	})
}

func TestResource_CheckResourceExistsById(t *testing.T) {
	WrapperTestResourceCheckResourceExists(t, func(id uint) bool {
		exists, err := challenge.CheckResourceExistsById(id)
		assert.NoError(t, err)
		return exists
	})
}

func TestResource_CheckResourceExistsByGuid(t *testing.T) {
	WrapperTestResourceCheckResourceExists(t, func(i uint) bool {
		var name string
		if i == 1 {
			name = test_data.Challenge1.Name
		} else if i == 2 {
			name = "not_exists"
		}
		return challenge.CheckResourceExistsByGuid(challenge.GuidFieldJsonTag, name)
	})
}

func TestResource_CheckResourceExistsByModelPtrWithGuid(t *testing.T) {
	WrapperTestResourceCheckResourceExists(t, func(i uint) bool {
		var c model.Challenge
		if i == 1 {
			c = test_data.Challenge1
		} else if i == 2 {
			c = test_data.Challenge2
		}
		return challenge.CheckResourceExistsByModelPtrWithGuid(&c, challenge.GuidFieldJsonTag)
	})
}

func TestResource_CheckResourceExistsExceptSelfByGuid(t *testing.T) {
	WrapperTestResourceCheckResourceExistsExceptSelf(t, func(id uint) bool {
		return challenge.CheckResourceExistsExceptSelfByGuid(challenge.GuidFieldJsonTag, test_data.Challenge1.Name, id)
	})
}

func TestResource_CheckResourceExistsExceptSelfByModelPtrWithGuid(t *testing.T) {
	WrapperTestResourceCheckResourceExistsExceptSelf(t, func(id uint) bool {
		return challenge.CheckResourceExistsExceptSelfByModelPtrWithGuid(&test_data.Challenge1, challenge.GuidFieldJsonTag, id)
	})
}
