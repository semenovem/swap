package tinkoffapi

// SandboxRegisterRequest Регистрация клиента в sandbox
// https://tinkoffcreditsystems.github.io/invest-openapi/swagger-ui/#/sandbox/post_sandbox_register
type SandboxRegisterRequest struct {
  BrokerAccountType string `json:"brokerAccountType"`
}

type SandboxRegisterReply struct {
  TrackingId string `json:"trackingId"`
  Status     string `json:"status"`
  Payload    struct {
    BrokerAccountType string `json:"brokerAccountType"`
    BrokerAccountId   string `json:"brokerAccountId"`
  } `json:"payload"`
}

// SandboxCurrenciesBalanceRequest --------------------------------------
// Выставление баланса по валютным позициям
// https://tinkoffcreditsystems.github.io/invest-openapi/swagger-ui/#/sandbox/post_sandbox_currencies_balance
type SandboxCurrenciesBalanceRequest struct {
  Currency string `json:"currency"`
  Balance  int    `json:"balance"`
}

type SandboxCurrenciesBalanceReply struct {
  TrackingId string `json:"trackingId"`
  Payload    struct {
  } `json:"payload"`
  Status string `json:"status"`
}

// SandboxPositionsBalanceRequest --------------------------------------
// Выставление баланса по инструментным позициям
// https://tinkoffcreditsystems.github.io/invest-openapi/swagger-ui/#/sandbox/post_sandbox_positions_balance
type SandboxPositionsBalanceRequest struct {
  Figi    string `json:"figi"`
  Balance int    `json:"balance"`
}

type SandboxPositionsBalanceReply struct {
  TrackingId string `json:"trackingId"`
  Payload    struct {
  } `json:"payload"`
  Status string `json:"status"`
}
