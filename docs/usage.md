# Usage

Invoices are stored in `~/.billfold/invoices.json`. Set `BILLFOLD_FILE` to use another file,
for example one per client or per year.

## add

```sh
billfold add --client NAME --amount 1250.00 [--currency USD] [--note TEXT]
```

Amounts take at most two decimals. The currency is a free-form code; billfold never converts
between currencies. The issue date is the moment you run the command.

## list

Prints every invoice with its id, client, total, issue date and whether it is paid.

## paid

```sh
billfold paid 3
```

Marks invoice 3 as paid.

## total

```sh
billfold total --month 2026-10
```

Sums the invoices issued in that month, one line per currency. Without `--month` it uses the
current month.

## export

```sh
billfold export --format csv --out invoices.csv
billfold export --format json > invoices.json
```

`csv` (the default) writes one row per invoice with the columns
`id,client,amount,currency,issued,paid,note`. `json` writes the same invoices as a JSON array.
Without `--out` the export goes to standard output.
