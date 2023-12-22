package csv_test

import (
	"testing"

	"github.com/lutzcle/viseca-exporter/internal/csv"
	"github.com/lutzcle/viseca-exporter/pkg/viseca"
	"github.com/stretchr/testify/assert"
)

var inputTransaction = viseca.Transaction{
	TransactionID:    "AUTH8c919db2-1c23-43f1-8862-61c31336d9b6",
	CardID:           "0000000AAAAA0000",
	MaskedCardNumber: "XXXXXXXXXXXX0000",
	CardName:         "Mastercard",
	Date:             "2021-10-20T17:05:44",
	ValutaDate:       "2021-10-19T14:11:05",
	ShowTimestamp:    true,
	Amount:           50.55,
	Currency:         "CHF",
	OriginalAmount:   51.55,
	OriginalCurrency: "USD",
	MerchantName:     "Aldi Suisse 00",
	PrettyName:       "ALDI",
	MerchantPlace:    "",
	IsOnline:         false,
	PFMCategory: viseca.PFMCategory{
		ID:                  "cv_groceries",
		Name:                "Lebensmittel",
		LightColor:          "#E2FDD3",
		MediumColor:         "#A5D58B",
		Color:               "#51A127",
		ImageURL:            "https://api.one.viseca.ch/v1/media/categories/icon_with_background/ic_cat_tile_groceries_v2.png",
		TransparentImageURL: "https://api.one.viseca.ch/v1/media/categories/icon_without_background/ic_cat_tile_groceries_v2.png",
	},
	StateType: "authorized",
	Details:   "Aldi Suisse 00",
	Type:      "merchant",
	IsBilled:  false,
	Links: viseca.TransactionLinks{
		Transactiondetails: "/v1/card/0000000AAAAA0000/transaction/AUTH8c919db2-1c23-43f1-8862-61c31336d9b6",
	},
}

func TestTransactionString(t *testing.T) {
	expected := `"AUTH8c919db2-1c23-43f1-8862-61c31336d9b6","0000000AAAAA0000","XXXXXXXXXXXX0000","2021-10-20T17:05:44","2021-10-19T14:11:05","ALDI","50.55","CHF","51.55","USD","false","authorized","Aldi Suisse 00","cv_groceries","Lebensmittel"`

	assert.Equal(t, expected, csv.TransactionString(inputTransaction))
}

func TestTransactionsString(t *testing.T) {
	inputTransactions := []viseca.Transaction{inputTransaction}
	expected :=
		`"TransactionID","CardID","MaskedCardNumber","Date","ValutaDate","Merchant","Amount","Currency","OriginalAmount","OriginalCurrency","IsOnline","StateType","Details","PFMCategoryID","PFMCategoryName"` +
			"\n" +
			`"AUTH8c919db2-1c23-43f1-8862-61c31336d9b6","0000000AAAAA0000","XXXXXXXXXXXX0000","2021-10-20T17:05:44","2021-10-19T14:11:05","ALDI","50.55","CHF","51.55","USD","false","authorized","Aldi Suisse 00","cv_groceries","Lebensmittel"` +
			"\n"

	assert.Equal(t, expected, csv.TransactionsString(inputTransactions))
}
