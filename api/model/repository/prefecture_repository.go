package repository

import "github.com/takeshi22/cozyLiv/model/entity"

type PrefectureRepository interface {
	CrateMany([]entity.Prefecture) ([]*entity.Prefecture, error)
}
