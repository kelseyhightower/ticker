package quote

import "time"

type Quotes struct {
	StockQuotes []*StockQuote
}

type StockQuote struct {
	Name             string
	Symbol           string
	LastPrice        float64
	Change           float64
	ChangePercent    float64
	Timestamp        string
	MarketCap        float64
	Volume           int
	ChangeYTD        float64
	ChangePercentYTD float64
	High             float64
	Low              float64
	Open             float64
}

func Quote(symbol string) StockQuote {
	// Simulate remote database overhead
	time.Sleep(10 * time.Millisecond)

	sq := StockQuote{
		Name:             "Company X",
		Symbol:           symbol,
		LastPrice:        524.49,
		Change:           15.6,
		ChangePercent:    3.06549549018453,
		Timestamp:        time.Now().Format(time.RFC3339),
		MarketCap:        476497591530,
		Volume:           397562,
		ChangeYTD:        532.1729,
		ChangePercentYTD: -1.44368493773359,
		High:             52499,
		Low:              519.175,
		Open:             519.175,
	}

	return sq
}
