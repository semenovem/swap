package tinkoffapi

// PortfolioRequest --------------------------------------
// Получение портфеля клиента
// https://tinkoffcreditsystems.github.io/invest-openapi/swagger-ui/#/portfolio/get_portfolio
type PortfolioRequest struct {
  BrokerAccountType string `json:"brokerAccountType"`
}

type PortfolioReply struct {
  TrackingId string `json:"trackingId"`
  Status     string `json:"status"`
  Payload    struct {
    Positions []struct {
      Figi           string `json:"figi"`
      Ticker         string `json:"ticker"`
      Isin           string `json:"isin"`
      InstrumentType string `json:"instrumentType"`
      Balance        int    `json:"balance"`
      Blocked        int    `json:"blocked"`
      ExpectedYield  struct {
        Currency string `json:"currency"`
        Value    int    `json:"value"`
      } `json:"expectedYield"`
      Lots                 int `json:"lots"`
      AveragePositionPrice struct {
        Currency string `json:"currency"`
        Value    int    `json:"value"`
      } `json:"averagePositionPrice"`
      AveragePositionPriceNoNkd struct {
        Currency string `json:"currency"`
        Value    int    `json:"value"`
      } `json:"averagePositionPriceNoNkd"`
      Name string `json:"name"`
    } `json:"positions"`
  } `json:"payload"`
}

// PortfolioCurrenciesRequest
// Получение валютных активов клиента
// https://tinkoffcreditsystems.github.io/invest-openapi/swagger-ui/#/portfolio/get_portfolio_currencies
type PortfolioCurrenciesRequest struct {
  BrokerAccountType string `json:"brokerAccountType"`
}

type PortfolioCurrenciesReply struct {
  TrackingId string `json:"trackingId"`
  Status     string `json:"status"`
  Payload    struct {
    Currencies []struct {
      Currency string `json:"currency"`
      Balance  int    `json:"balance"`
      Blocked  int    `json:"blocked"`
    } `json:"currencies"`
  } `json:"payload"`
}
