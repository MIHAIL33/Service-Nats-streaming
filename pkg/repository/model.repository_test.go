package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	models "github.com/MIHAIL33/Service-Nats-streaming/model"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestModelRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewModelRepository(db)

	type args struct {
		model models.Model
	}

	type mockBehavior func(model models.Model)

	var jsonInput string = "{\"order_uid\":\"14\",\"track_number\":\"WBILMTESTTRACK\",\"entry\":\"WBIL\",\"delivery\":{\"name\":\"Test Testov\",\"phone\":\"+9720000000\",\"zip\":\"2639809\",\"city\":\"Kiryat Mozkin\",\"address\":\"Ploshad Mira 15\",\"region\":\"Kraiot\",\"email\":\"test@gmail.com\"},\"payment\":{\"transaction\":\"b563feb7b2b84b6test\",\"request_id\":\"\",\"currency\":\"USD\",\"provider\":\"wbpay\",\"amount\":1817,\"payment_dt\":1637907727,\"bank\":\"alpha\",\"delivery_cost\":1500,\"goods_total\":317,\"custom_fee\":0},\"items\":[{\"chrt_id\":9934930,\"track_number\":\"WBILMTESTTRACK\",\"price\":453,\"rid\":\"ab4219087a764ae0btest\",\"name\":\"Mascaras\",\"sale\":30,\"size\":\"0\",\"total_price\":317,\"nm_id\":2389212,\"brand\":\"Vivienne Sabo\",\"status\":202}],\"locale\":\"en\",\"internal_signature\":\"\",\"customer_id\":\"test\",\"delivery_service\":\"meest\",\"shardkey\":\"9\",\"sm_id\":99,\"date_created\":\"2021-11-26T06:22:19Z\",\"oof_shard\":\"1\"}"
	var modelInput models.Model
	err = json.Unmarshal([]byte(jsonInput), &modelInput)
	if err != nil {
		return
	}

	testTable := []struct {
		name         string
		input        models.Model
		mockBehavior mockBehavior
		output       models.Model
		wantErr      bool
	}{
		{
			name:   "OK",
			input:  modelInput,
			output: modelInput,
			mockBehavior: func(model models.Model) {
				// createModelQuery := fmt.Sprintf(`INSERT INTO %s (model) VALUES(%s) RETURNING *`, modelsTable, jsonInput)
				// mock.ExpectQuery(createModelQuery)
				fmt.Println(model)
				jsonModel, err := json.Marshal(model)
				if err != nil {
				 	return
				}
				fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
				//rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				//mock.ExpectQuery("INSERT INTO users").WithArgs("first_name", "last_name", "username", "password").WillReturnRows(rows)
				//rows := sqlmock.NewRows([]string{"model"})
				//mock.ExpectQuery("INSERT INTO models").WithArgs(jsonModel).WillReturnRows(rows)
				//row := sqlmock.NewRows([]string{"model"}).AddRow("model")
				//mock.ExpectQuery(`INSERT INTO models`).WithArgs(out).WillReturnRows(row)
				mock.ExpectExec(`INSERT INTO models`).WithArgs(jsonModel).WillReturnResult(sqlmock.NewResult(1, 1))
				//fmt.Println(row)
				//mock.ExpectQuery(`INSERT INTO models`).WithArgs("model").WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.input)

			got, err := r.Create(testCase.input)

			fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			fmt.Println(got)
			fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.output, *got)
			}
		})
	}
}
