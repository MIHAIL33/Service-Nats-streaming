package repository

import (
	"encoding/json"
	"errors"
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
				jsonModel, err := json.Marshal(model)
				if err != nil {
				 	return
				}
				rows := sqlmock.NewRows([]string{"model"}).AddRow(jsonModel)
				mock.ExpectQuery(`INSERT INTO models`).WithArgs(jsonModel).WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name:   "Empty object",
			input:  models.Model{},
			output: modelInput,
			mockBehavior: func(model models.Model) {
				jsonModel, err := json.Marshal(model)
				if err != nil {
				 	return
				}
				rows := sqlmock.NewRows([]string{"model"}).AddRow(jsonModel).RowError(0, errors.New("insert error"))
				mock.ExpectQuery(`INSERT INTO models`).WithArgs(jsonModel).WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.input)

			got, err := r.Create(testCase.input)

			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.output, *got)
			}
		})
	}
}

func TestModelRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewModelRepository(db)

	type mockBehavior func()

	var jsonInput string = "{\"order_uid\":\"14\",\"track_number\":\"WBILMTESTTRACK\",\"entry\":\"WBIL\",\"delivery\":{\"name\":\"Test Testov\",\"phone\":\"+9720000000\",\"zip\":\"2639809\",\"city\":\"Kiryat Mozkin\",\"address\":\"Ploshad Mira 15\",\"region\":\"Kraiot\",\"email\":\"test@gmail.com\"},\"payment\":{\"transaction\":\"b563feb7b2b84b6test\",\"request_id\":\"\",\"currency\":\"USD\",\"provider\":\"wbpay\",\"amount\":1817,\"payment_dt\":1637907727,\"bank\":\"alpha\",\"delivery_cost\":1500,\"goods_total\":317,\"custom_fee\":0},\"items\":[{\"chrt_id\":9934930,\"track_number\":\"WBILMTESTTRACK\",\"price\":453,\"rid\":\"ab4219087a764ae0btest\",\"name\":\"Mascaras\",\"sale\":30,\"size\":\"0\",\"total_price\":317,\"nm_id\":2389212,\"brand\":\"Vivienne Sabo\",\"status\":202}],\"locale\":\"en\",\"internal_signature\":\"\",\"customer_id\":\"test\",\"delivery_service\":\"meest\",\"shardkey\":\"9\",\"sm_id\":99,\"date_created\":\"2021-11-26T06:22:19Z\",\"oof_shard\":\"1\"}"
	var modelInput models.Model
	err = json.Unmarshal([]byte(jsonInput), &modelInput)
	if err != nil {
		return
	}
	jsonModel, err := json.Marshal(modelInput)
	if err != nil {
		 return
	}

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		output       []models.Model
		wantErr      bool
	}{
		{
			name:   "OK",
			output: []models.Model{modelInput, modelInput, modelInput},
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"model"}).AddRow(jsonModel).AddRow(jsonModel).AddRow(jsonModel)
				mock.ExpectQuery(`SELECT (.+) FROM models`).WithArgs().WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name:   "No Records",
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"model"})
				mock.ExpectQuery(`SELECT (.+) FROM models`).WithArgs().WillReturnRows(rows)
			},
			wantErr: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()

			got, err := r.GetAll()

			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.output, *got)
			}
		})
	}
}

func TestModelRepository_GetById(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewModelRepository(db)

	type mockBehavior func(id string)

	var jsonInput string = "{\"order_uid\":\"14\",\"track_number\":\"WBILMTESTTRACK\",\"entry\":\"WBIL\",\"delivery\":{\"name\":\"Test Testov\",\"phone\":\"+9720000000\",\"zip\":\"2639809\",\"city\":\"Kiryat Mozkin\",\"address\":\"Ploshad Mira 15\",\"region\":\"Kraiot\",\"email\":\"test@gmail.com\"},\"payment\":{\"transaction\":\"b563feb7b2b84b6test\",\"request_id\":\"\",\"currency\":\"USD\",\"provider\":\"wbpay\",\"amount\":1817,\"payment_dt\":1637907727,\"bank\":\"alpha\",\"delivery_cost\":1500,\"goods_total\":317,\"custom_fee\":0},\"items\":[{\"chrt_id\":9934930,\"track_number\":\"WBILMTESTTRACK\",\"price\":453,\"rid\":\"ab4219087a764ae0btest\",\"name\":\"Mascaras\",\"sale\":30,\"size\":\"0\",\"total_price\":317,\"nm_id\":2389212,\"brand\":\"Vivienne Sabo\",\"status\":202}],\"locale\":\"en\",\"internal_signature\":\"\",\"customer_id\":\"test\",\"delivery_service\":\"meest\",\"shardkey\":\"9\",\"sm_id\":99,\"date_created\":\"2021-11-26T06:22:19Z\",\"oof_shard\":\"1\"}"
	var modelInput models.Model
	err = json.Unmarshal([]byte(jsonInput), &modelInput)
	if err != nil {
		return
	}
	jsonModel, err := json.Marshal(modelInput)
	if err != nil {
		 return
	}

	testTable := []struct {
		name         string
		input		 string
		mockBehavior mockBehavior
		output       models.Model
		wantErr      bool
	}{
		{
			name:   "OK",
			input: 	"1",
			output: modelInput,
			mockBehavior: func(id string) {
				rows := sqlmock.NewRows([]string{"model"}).AddRow(jsonModel)
				mock.ExpectQuery(`SELECT (.+) FROM models WHERE (.+)`).WithArgs(id).WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name:   "Not Found",
			input: 	"2",
			output: modelInput,
			mockBehavior: func(id string) {
				rows := sqlmock.NewRows([]string{"model"})
				mock.ExpectQuery(`SELECT (.+) FROM models WHERE (.+)`).WithArgs(id).WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.input)

			got, err := r.GetById(testCase.input)

			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.output, *got)
			}
		})
	}
}

func TestModelRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewModelRepository(db)

	type mockBehavior func(id string)

	var jsonInput string = "{\"order_uid\":\"14\",\"track_number\":\"WBILMTESTTRACK\",\"entry\":\"WBIL\",\"delivery\":{\"name\":\"Test Testov\",\"phone\":\"+9720000000\",\"zip\":\"2639809\",\"city\":\"Kiryat Mozkin\",\"address\":\"Ploshad Mira 15\",\"region\":\"Kraiot\",\"email\":\"test@gmail.com\"},\"payment\":{\"transaction\":\"b563feb7b2b84b6test\",\"request_id\":\"\",\"currency\":\"USD\",\"provider\":\"wbpay\",\"amount\":1817,\"payment_dt\":1637907727,\"bank\":\"alpha\",\"delivery_cost\":1500,\"goods_total\":317,\"custom_fee\":0},\"items\":[{\"chrt_id\":9934930,\"track_number\":\"WBILMTESTTRACK\",\"price\":453,\"rid\":\"ab4219087a764ae0btest\",\"name\":\"Mascaras\",\"sale\":30,\"size\":\"0\",\"total_price\":317,\"nm_id\":2389212,\"brand\":\"Vivienne Sabo\",\"status\":202}],\"locale\":\"en\",\"internal_signature\":\"\",\"customer_id\":\"test\",\"delivery_service\":\"meest\",\"shardkey\":\"9\",\"sm_id\":99,\"date_created\":\"2021-11-26T06:22:19Z\",\"oof_shard\":\"1\"}"
	var modelInput models.Model
	err = json.Unmarshal([]byte(jsonInput), &modelInput)
	if err != nil {
		return
	}
	jsonModel, err := json.Marshal(modelInput)
	if err != nil {
		 return
	}

	testTable := []struct {
		name         string
		input		 string
		mockBehavior mockBehavior
		output       models.Model
		wantErr      bool
	}{
		{
			name:   "OK",
			input: 	"1",
			output: modelInput,
			mockBehavior: func(id string) {
				rows := sqlmock.NewRows([]string{"model"}).AddRow(jsonModel)
				mock.ExpectQuery(`DELETE FROM models WHERE (.+)`).WithArgs(id).WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name:   "Not Found",
			input: 	"2",
			output: modelInput,
			mockBehavior: func(id string) {
				rows := sqlmock.NewRows([]string{"model"})
				mock.ExpectQuery(`DELETE FROM models WHERE (.+)`).WithArgs(id).WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.input)

			got, err := r.Delete(testCase.input)

			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.output, *got)
			}
		})
	}
}
