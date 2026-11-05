# billfold

Invoices for freelancers, from the terminal. One binary, one JSON file, no account.

```sh
go install github.com/prabaltripathiofficial/billfold/cmd/billfold@latest

billfold add --client "Globex" --amount 1250.00 --currency USD
billfold list
billfold paid 1
billfold total --month 2026-10
billfold export --format csv --out invoices.csv
```

- [Usage](docs/usage.md) — every command and flag
- [FAQ](docs/faq.md) — answers to the questions people ask most
- [Changelog](CHANGELOG.md)

## Community

Questions and bug reports go to `#help` on the billfold Discord. Maintainers work
cases in Slack and post the outcome back in your thread.

## Contributing

```sh
go vet ./... && go test ./...
```

MIT licensed.
