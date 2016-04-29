package main

import (
	"time"

	pb "github.com/kelseyhightower/ticker/quotepb"
)

func Quote(symbol string) pb.StockQuote {
	// Simulate remote database overhead
	time.Sleep(10 * time.Millisecond)

	sq := pb.StockQuote{
		Name:              "Company X",
		Symbol:            symbol,
		LastPrice:         524.49,
		Change:            15.6,
		ChangePercent:     3.06549549018453,
		Timestamp:         time.Now().Format(time.RFC3339),
		MarketCap:         476497591530,
		Volume:            397562,
		Change_YTD:        532.1729,
		ChangePercent_YTD: -1.44368493773359,
		High:              52499,
		Low:               519.175,
		Open:              519.175,
	}

	return sq
}
