package test_data

import (
	"github.com/ssst0n3/lightweight_api/test/model"
	"gorm.io/gorm"
)

var (
	Challenge1 = model.Challenge{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "challenge1",
	}
	ChallengeSameName = model.Challenge{
		Model: gorm.Model{
			ID: 2,
		},
		Name: Challenge1.Name,
	}
	Challenge1Update = model.Challenge{
		Model: gorm.Model{
			ID: 2,
		},
		Name:   Challenge1.Name,
		Score:  20,
		Solved: true,
	}
	Challenge2 = model.Challenge{
		Model: gorm.Model{
			ID: 2,
		},
		Name: "challenge2",
	}
	Challenges = []model.Challenge{
		Challenge1,
	}
)

func InitEmptyChallenge(DB *gorm.DB) (err error) {
	err = DB.Migrator().DropTable(&model.Challenge{})
	if err != nil {
		return
	}
	err = DB.AutoMigrate(&model.Challenge{})
	return
}
