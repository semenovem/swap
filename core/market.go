package core

type MarketStocksRequest struct{}

type MarketStocksReply struct {
  Total       int
  Instruments []*Instrument
}

type Instrument struct {
  Figi              string
  Ticker            string
  Isin              string
  MinPriceIncrement float32
  Lot               int
  MinQuantity       int
  Currency          string
  Name              string
  Type              string
}
