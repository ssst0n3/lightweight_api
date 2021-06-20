package lightweight_api

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_api/test/test_data"
	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, true, exists)
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
	c := &gin.Context{Params: []gin.Param{
		{Key: "id", Value: "1"},
	}}
	WrapperTestResourceCheckResourceExists(t, func(i uint) bool {
		exists, id, err := challenge.CheckResourceExistsByIdAutoParseParam(c)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), id)
		return exists
	})
}

func TestResource_CheckResourceExistsById(t *testing.T) {
	id := uint(1)
	WrapperTestResourceCheckResourceExists(t, func(i uint) bool {
		return challenge.CheckResourceExistsById(id)
	})
}

func TestResource_CheckResourceExistsByGuid(t *testing.T) {
	WrapperTestResourceCheckResourceExists(t, func(i uint) bool {
		return challenge.CheckResourceExistsByGuid(challenge.GuidFieldJsonTag, test_data.Challenge1.Name)
	})
}

func TestResource_CheckResourceExistsByModelPtrWithGuid(t *testing.T) {
	WrapperTestResourceCheckResourceExists(t, func(i uint) bool {
		return challenge.CheckResourceExistsByModelPtrWithGuid(&test_data.Challenge1, challenge.GuidFieldJsonTag)
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
