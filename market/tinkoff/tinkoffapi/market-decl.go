package tinkoffapi

import "time"

// MarketStocksRequest Получение списка акций
// https://tinkoffcreditsystems.github.io/invest-openapi/swagger-ui/#/market/get_market_stocks
type MarketStocksRequest struct{}

type MarketStocksReply struct {
  TrackingId string `json:"trackingId"`
  Status     string `json:"status"`
  Payload    struct {
    Total       int `json:"total"`
    Instruments []struct {
      Figi              string  `json:"figi"`
      Ticker            string  `json:"ticker"`
      Isin              string  `json:"isin"`
      MinPriceIncrement float32 `json:"minPriceIncrement"`
      Lot               int     `json:"lot"`
      MinQuantity       int     `json:"minQuantity"`
      Currency          string  `json:"currency"`
      Name              string  `json:"name"`
      Type              string  `json:"type"`
    } `json:"instruments"`
  } `json:"payload"`
}

// MarketBondsRequest Получение списка облигаций
// https://tinkoffcreditsystems.github.io/invest-openapi/swagger-ui/#/market/get_market_bonds
type MarketBondsRequest struct{}

type MarketBondsReply struct {
  TrackingId string `json:"trackingId"`
  Status     string `json:"status"`
  Payload    struct {
    Total       int `json:"total"`
    Instruments []struct {
      Figi              string `json:"figi"`
      Ticker            string `json:"ticker"`
      Isin              string `json:"isin"`
      MinPriceIncrement int    `json:"minPriceIncrement"`
      Lot               int    `json:"lot"`
      MinQuantity       int    `json:"minQuantity"`
      Currency          string `json:"currency"`
      Name              string `json:"name"`
      Type              string `json:"type"`
    } `json:"instruments"`
  } `json:"payload"`
}

// MarketEtfsRequest Получение списка ETF
// https://tinkoffcreditsystems.github.io/invest-openapi/swagger-ui/#/market/get_market_etfs
type MarketEtfsRequest struct{}

type MarketEtfsReply struct {
  TrackingId string `json:"trackingId"`
  Status     string `json:"status"`
  Payload    struct {
    Total       int `json:"total"`
    Instruments []struct {
      Figi              string `json:"figi"`
      Ticker            string `json:"ticker"`
      Isin              string `json:"isin"`
      MinPriceIncrement int    `json:"minPriceIncrement"`
      Lot               int    `json:"lot"`
      MinQuantity       int    `json:"minQuantity"`
      Currency          string `json:"currency"`
      Name              string `json:"name"`
      Type              string `json:"type"`
    } `json:"instruments"`
  } `json:"payload"`
}

// MarketCurrenciesRequest Получение списка валютных пар
// https://tinkoffcreditsystems.github.io/invest-openapi/swagger-ui/#/market/get_market_currencies
type MarketCurrenciesRequest struct{}

type MarketCurrenciesReply struct {
  TrackingId string `json:"trackingId"`
  Status     string `json:"status"`
  Payload    struct {
    Total       int `json:"total"`
    Instruments []struct {
      Figi              string `json:"figi"`
      Ticker            string `json:"ticker"`
      Isin              string `json:"isin"`
      MinPriceIncrement int    `json:"minPriceIncrement"`
      Lot               int    `json:"lot"`
      MinQuantity       int    `json:"minQuantity"`
      Currency          string `json:"currency"`
      Name              string `json:"name"`
      Type              string `json:"type"`
    } `json:"instruments"`
  } `json:"payload"`
}

// MarketOrderbookRequest Получение стакана по FIGI
// https://tinkoffcreditsystems.github.io/invest-openapi/swagger-ui/#/market/get_market_orderbook
type MarketOrderbookRequest struct {
  figi  string // FIGI
  depth int    // Глубина стакана [1..20]
}

type MarketOrderbookReply struct {
  TrackingId string `json:"trackingId"`
  Status     string `json:"status"`
  Payload    struct {
    Figi  string `json:"figi"`
    Depth int    `json:"depth"`
    Bids  []struct {
      Price    int `json:"price"`
      Quantity int `json:"quantity"`
    } `json:"bids"`
    Asks []struct {
      Price    int `json:"price"`
      Quantity int `json:"quantity"`
    } `json:"asks"`
    TradeStatus       string `json:"tradeStatus"`
    MinPriceIncrement int    `json:"minPriceIncrement"`
    FaceValue         int    `json:"faceValue"`
    LastPrice         int    `json:"lastPrice"`
    ClosePrice        int    `json:"closePrice"`
    LimitUp           int    `json:"limitUp"`
    LimitDown         int    `json:"limitDown"`
  } `json:"payload"`
}

// MarketСandlesRequest Получение исторических свечей по FIGI
// https://tinkoffcreditsystems.github.io/invest-openapi/swagger-ui/#/market/get_market_candles
type MarketСandlesRequest struct {
  figi     string // FIGI
  from     string // Начало временного промежутка 2019-08-19T18:38:33+03:00
  to       string // Конец временного промежутка 2019-08-19T18:38:33+03:00
  interval string // 1min, 2min, 3min, 5min, 10min, 15min, 30min, hour, day, week, month
}

type MarketСandlesReply struct {
  TrackingId string `json:"trackingId"`
  Status     string `json:"status"`
  Payload    struct {
    Figi     string `json:"figi"`
    Interval string `json:"interval"`
    Candles  []struct {
      Figi     string    `json:"figi"`
      Interval string    `json:"interval"`
      O        int       `json:"o"`
      C        int       `json:"c"`
      H        int       `json:"h"`
      L        int       `json:"l"`
      V        int       `json:"v"`
      Time     time.Time `json:"time"`
    } `json:"candles"`
  } `json:"payload"`
}

// MarketByFigiRequest Получение инструмента по FIGI
// https://tinkoffcreditsystems.github.io/invest-openapi/swagger-ui/#/market/get_market_search_by_figi
type MarketByFigiRequest struct {
  figi string // FIGI
}

type MarketByFigiReply struct {
  TrackingId string `json:"trackingId"`
  Status     string `json:"status"`
  Payload    struct {
    Figi              string `json:"figi"`
    Ticker            string `json:"ticker"`
    Isin              string `json:"isin"`
    MinPriceIncrement int    `json:"minPriceIncrement"`
    Lot               int    `json:"lot"`
    Currency          string `json:"currency"`
    Name              string `json:"name"`
    Type              string `json:"type"`
  } `json:"payload"`
}

// MarketByTickerRequest Получение инструмента по тикеру
// https://tinkoffcreditsystems.github.io/invest-openapi/swagger-ui/#/market/get_market_search_by_ticker
type MarketByTickerRequest struct {
  ticker string // Тикер инструмента
}

type MarketByTickerReply struct {
  TrackingId string `json:"trackingId"`
  Status     string `json:"status"`
  Payload    struct {
    Total       int `json:"total"`
    Instruments []struct {
      Figi              string `json:"figi"`
      Ticker            string `json:"ticker"`
      Isin              string `json:"isin"`
      MinPriceIncrement int    `json:"minPriceIncrement"`
      Lot               int    `json:"lot"`
      MinQuantity       int    `json:"minQuantity"`
      Currency          string `json:"currency"`
      Name              string `json:"name"`
      Type              string `json:"type"`
    } `json:"instruments"`
  } `json:"payload"`
}
