package api

type entry struct {
  typ  string
  path string
}

type descApi struct {
  *sandbox
  *portfolio
  *market
}

// https://tinkoffcreditsystems.github.io/invest-openapi/swagger-ui/#/sandbox
type sandbox struct {
  register         *entry // Регистрация клиента в sandbox
  currencyBalance  *entry // Выставление баланса по валютным позициям
  positionsBalance *entry // Выставление баланса по инструментным позициям
  remove           *entry // Удаление счета
  clear            *entry // Удаление всех позиций
}

// https://tinkoffcreditsystems.github.io/invest-openapi/swagger-ui/#/portfolio
type portfolio struct {
  portfolio  *entry // Получение портфеля клиента
  currencies *entry // Получение валютных активов клиента
}

// https://tinkoffcreditsystems.github.io/invest-openapi/swagger-ui/#/market
type market struct {
  stocks     *entry // Получение списка акций
  bonds      *entry // Получение списка облигаций
  etfs       *entry // Получение списка ETF
  currencies *entry // Получение списка валютных пар
  orderbook  *entry // Получение стакана по FIGI
  candles    *entry // Получение исторических свечей по FIGI
  byFigi     *entry // Получение инструмента по FIGI
  byTicker   *entry // Получение инструмента по тикеру
}

// Err500 Ошибка сервиса
type Err500 struct {
  TrackingId string `json:"trackingId"`
  Status     string `json:"status"`
  Payload    struct {
    Message string `json:"message"`
    Code    string `json:"code"`
  } `json:"payload"`
}
