package api

import "net/http"

// Описание tinkoff api
var declEntry = &descApi{
  // -----------------------------
  sandbox: &sandbox{
    register: &entry{
      typ:  http.MethodPost,
      path: "/sandbox/register",
    },
    currencyBalance: &entry{
      typ:  http.MethodPost,
      path: "/sandbox/currencies/balance",
    },
    positionsBalance: &entry{
      typ:  http.MethodPost,
      path: "/sandbox/positions/balance",
    },
    remove: &entry{
      typ:  http.MethodPost,
      path: "/sandbox/remove",
    },
    clear: &entry{
      typ:  http.MethodPost,
      path: "/sandbox/clear",
    },
  },

  // -----------------------------
  portfolio: &portfolio{
    portfolio: &entry{
      typ:  http.MethodGet,
      path: "/portfolio",
    },
    currencies: &entry{
      typ:  http.MethodGet,
      path: "/portfolio/currencies",
    },
  },

  // -----------------------------
  market: &market{
    stocks: &entry{
      typ:  http.MethodGet,
      path: "/market/stocks",
    },
    bonds: &entry{
      typ:  http.MethodGet,
      path: "/market/bonds",
    },
    etfs: &entry{
      typ:  http.MethodGet,
      path: "/market/etfs",
    },
    currencies: &entry{
      typ:  http.MethodGet,
      path: "/market/currencies",
    },
    orderbook: &entry{
      typ:  http.MethodGet,
      path: "/market/orderbook",
    },
    candles: &entry{
      typ:  http.MethodGet,
      path: "/market/candles",
    },
    byFigi: &entry{
      typ:  http.MethodGet,
      path: "/market/search/by-figi",
    },
    byTicker: &entry{
      typ:  http.MethodGet,
      path: "/market/search/by-ticker",
    },
  },
}
