package helper

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/takeshi22/cozyLiv/infrastructure"
	"github.com/takeshi22/cozyLiv/model/entity"
)

func ReadCsv() []entity.Prefecture {
	filePath, _ := filepath.Abs("static/toshi.csv")
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	r := csv.NewReader(bytes.NewBuffer(data))

	var entities []entity.Prefecture

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		elias := strings.Split(record[2], " ")[0]

		entity := entity.Prefecture{
			ID:    uuid.New(),
			Name:  record[1],
			Elias: elias,
		}
		entities = append(entities, entity)
	}

	prefectures := []entity.Prefecture{}

	infrastructure.Db.Find(&prefectures)
	fmt.Println(prefectures)

	return entities
}
