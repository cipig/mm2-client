package http

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"mm2_client/config"
	"mm2_client/constants"
	"mm2_client/helpers"
	"mm2_client/services"
	"os"
	"strconv"
)

type MM2GenericRequest struct {
	Method   string `json:"method"`
	Userpass string `json:"userpass"`
}

var gRuntimeUserpass = ""

const GMM2Endpoint = "http://127.0.0.1:7783"

type GenericEnableAnswer struct {
	Coin                  string `json:"coin"`
	Address               string `json:"address"`
	Balance               string `json:"balance"`
	RequiredConfirmations int    `json:"required_confirmations"`
	RequiresNotarization  bool   `json:"requires_notarization"`
	UnspendableBalance    string `json:"unspendable_balance"`
	Result                string `json:"result"`
}

func (answer *GenericEnableAnswer) ToTable() {
	symbol := config.RetrieveUSDSymbolIfSupported(answer.Coin)
	val := "0"
	if len(symbol) > 0 {
		if v, ok := services.BinancePriceRegistry.Load(symbol); ok {
			val = helpers.BigFloatMultiply(answer.Balance, v.(string), 8)
		}
	}

	data := [][]string{
		{answer.Coin, answer.Address, answer.Balance, val, strconv.Itoa(answer.RequiredConfirmations), strconv.FormatBool(answer.RequiresNotarization), answer.UnspendableBalance, answer.Result},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Coin", "Address", "Balance", "Balance (USD)", "Confirmations", "Notarization", "Unspendable", "Status"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}

func ToTable(answers []GenericEnableAnswer) {
	var data [][]string

	for _, answer := range answers {
		symbol := config.RetrieveUSDSymbolIfSupported(answer.Coin)
		val := "0"
		if len(symbol) > 0 {
			if v, ok := services.BinancePriceRegistry.Load(symbol); ok {
				val = helpers.BigFloatMultiply(answer.Balance, v.(string), 8)
			}
		}

		//helpers.BigFloatMultiply(answer.Balance, answers)
		cur := []string{answer.Coin, answer.Address, answer.Balance, val, strconv.Itoa(answer.RequiredConfirmations),
			strconv.FormatBool(answer.RequiresNotarization), answer.UnspendableBalance, answer.Result}
		data = append(data, cur)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Coin", "Address", "Balance", "Balance (USD)", "Confirmations", "Notarization", "Unspendable", "Status"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}

func NewGenericRequest(method string) *MM2GenericRequest {
	if gRuntimeUserpass == "" {
		gRuntimeUserpass = config.NewMM2ConfigFromFile(constants.GMM2ConfPath).RPCPassword
	}
	return &MM2GenericRequest{Method: method, Userpass: gRuntimeUserpass}
}

func (req MM2GenericRequest) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}
