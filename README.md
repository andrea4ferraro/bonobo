# Liquidity Position Tracker

Liquidity Position Tracker is a local Go application for tracking DeFi liquidity positions.

The application stores LP information locally and performs simple calculations without connecting to any blockchain or wallet.

## Features

- Add LP positions
- Calculate current value
- Profit/Loss estimation
- JSON storage
- Report export

## Run

```bash
go run ./cmd
```

Example

```
Pool: ETH / USDC

Current Value: $1820

Profit: +$120
```

## Future Improvements

- CSV export
- Historical snapshots
- Charts
- Multiple portfolios
- Token price API
