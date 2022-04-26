package handler

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	models "github.com/MIHAIL33/Service-Nats-streaming/model"
	"github.com/MIHAIL33/Service-Nats-streaming/pkg/service"
	mock_service "github.com/MIHAIL33/Service-Nats-streaming/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_createModel(t *testing.T) {
	type mockBehavior func(r *mock_service.MockModel, model models.Model)

	var jsonInput string = "{\"order_uid\":\"14\",\"track_number\":\"WBILMTESTTRACK\",\"entry\":\"WBIL\",\"delivery\":{\"name\":\"Test Testov\",\"phone\":\"+9720000000\",\"zip\":\"2639809\",\"city\":\"Kiryat Mozkin\",\"address\":\"Ploshad Mira 15\",\"region\":\"Kraiot\",\"email\":\"test@gmail.com\"},\"payment\":{\"transaction\":\"b563feb7b2b84b6test\",\"request_id\":\"\",\"currency\":\"USD\",\"provider\":\"wbpay\",\"amount\":1817,\"payment_dt\":1637907727,\"bank\":\"alpha\",\"delivery_cost\":1500,\"goods_total\":317,\"custom_fee\":0},\"items\":[{\"chrt_id\":9934930,\"track_number\":\"WBILMTESTTRACK\",\"price\":453,\"rid\":\"ab4219087a764ae0btest\",\"name\":\"Mascaras\",\"sale\":30,\"size\":\"0\",\"total_price\":317,\"nm_id\":2389212,\"brand\":\"Vivienne Sabo\",\"status\":202}],\"locale\":\"en\",\"internal_signature\":\"\",\"customer_id\":\"test\",\"delivery_service\":\"meest\",\"shardkey\":\"9\",\"sm_id\":99,\"date_created\":\"2021-11-26T06:22:19Z\",\"oof_shard\":\"1\"}"
	var modelInput models.Model
	err := json.Unmarshal([]byte(jsonInput), &modelInput)
	if err != nil {
		return
	}

	testTable := []struct {
		name string
		inputBody string
		inputModel models.Model
		mockBehavior mockBehavior
		expectedStatusCode int
		expectedRequestBody string
	} {
		{
			name: "OK",
			inputBody: jsonInput,
			inputModel: modelInput,
			mockBehavior: func(s *mock_service.MockModel, model models.Model)  {
				s.EXPECT().Create(model).Return(&model, nil)
			},
			expectedStatusCode: 200,
			expectedRequestBody: jsonInput,
		},
		{
			name: "invalid data",
			inputBody: "1",
			mockBehavior: func(r *mock_service.MockModel, model models.Model) {},
			expectedStatusCode: 400,
			expectedRequestBody: `{"message":"json: cannot unmarshal number into Go value of type models.Model"}`,
		},
		{
			name: "invalid json",
			inputBody: `{"customer": "cust"}`,
			mockBehavior: func(r *mock_service.MockModel, model models.Model) {},
			expectedStatusCode: 400,
			expectedRequestBody: `{"message":"Key: 'Model.Order_uid' Error:Field validation for 'Order_uid' failed on the 'required' tag\nKey: 'Model.Track_number' Error:Field validation for 'Track_number' failed on the 'required' tag\nKey: 'Model.Entry' Error:Field validation for 'Entry' failed on the 'required' tag\nKey: 'Model.Delivery.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'Model.Delivery.Phone' Error:Field validation for 'Phone' failed on the 'required' tag\nKey: 'Model.Delivery.Zip' Error:Field validation for 'Zip' failed on the 'required' tag\nKey: 'Model.Delivery.City' Error:Field validation for 'City' failed on the 'required' tag\nKey: 'Model.Delivery.Address' Error:Field validation for 'Address' failed on the 'required' tag\nKey: 'Model.Delivery.Region' Error:Field validation for 'Region' failed on the 'required' tag\nKey: 'Model.Delivery.Email' Error:Field validation for 'Email' failed on the 'required' tag\nKey: 'Model.Payment.Transaction' Error:Field validation for 'Transaction' failed on the 'required' tag\nKey: 'Model.Payment.Currency' Error:Field validation for 'Currency' failed on the 'required' tag\nKey: 'Model.Payment.Provider' Error:Field validation for 'Provider' failed on the 'required' tag\nKey: 'Model.Payment.Amount' Error:Field validation for 'Amount' failed on the 'required' tag\nKey: 'Model.Payment.Payment_dt' Error:Field validation for 'Payment_dt' failed on the 'required' tag\nKey: 'Model.Payment.Bank' Error:Field validation for 'Bank' failed on the 'required' tag\nKey: 'Model.Payment.Delivery_cost' Error:Field validation for 'Delivery_cost' failed on the 'required' tag\nKey: 'Model.Payment.Goods_total' Error:Field validation for 'Goods_total' failed on the 'required' tag\nKey: 'Model.Items' Error:Field validation for 'Items' failed on the 'required' tag\nKey: 'Model.Locale' Error:Field validation for 'Locale' failed on the 'required' tag\nKey: 'Model.Customer_id' Error:Field validation for 'Customer_id' failed on the 'required' tag\nKey: 'Model.Delivery_service' Error:Field validation for 'Delivery_service' failed on the 'required' tag\nKey: 'Model.Shardkey' Error:Field validation for 'Shardkey' failed on the 'required' tag\nKey: 'Model.Sm_id' Error:Field validation for 'Sm_id' failed on the 'required' tag\nKey: 'Model.Date_created' Error:Field validation for 'Date_created' failed on the 'required' tag\nKey: 'Model.Oof_shard' Error:Field validation for 'Oof_shard' failed on the 'required' tag"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockModel(c)
			testCase.mockBehavior(repo, testCase.inputModel)

			services := &service.Service{Model:  repo}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/models", handler.createModel)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/models", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}