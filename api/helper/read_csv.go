package helper

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/takeshi22/cozyLiv/infrastructure"
	"github.com/takeshi22/cozyLiv/model/entity"
)

func ReadCsv() []entity.City {
	filePath, _ := filepath.Abs("static/city.csv")
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	r := csv.NewReader(bytes.NewBuffer(data))

	var entities []entity.City

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		prefecture := entity.Prefecture{}
		infrastructure.Db.Where("code = ?", record[2]).Take(&prefecture)

		entity := entity.City{
			ID:           uuid.New(),
			Name:         record[1],
			PrefectureID: prefecture.ID,
		}

		entities = append(entities, entity)
	}

	return entities
}
