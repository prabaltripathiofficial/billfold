# FAQ

**Where is my data?** In `~/.billfold/invoices.json`. It is plain JSON; back it up like any
other file. `BILLFOLD_FILE=/path/to/file.json billfold list` reads a different one.

**How do I export to JSON instead of CSV?** `billfold export --format json`.

**Can I open the CSV in Excel or Google Sheets?** Yes: `billfold export --out invoices.csv`,
then open or import the file.

**Can I use several currencies?** Yes. Each invoice carries its own currency, and `total`
prints one line per currency. billfold never converts between them.

**How do I edit or delete an invoice?** There is no command for it yet. Edit
`~/.billfold/invoices.json` by hand; keep the ids unique.

**How do I upgrade?** `go install github.com/prabaltripathiofficial/billfold/cmd/billfold@latest`,
then `billfold version`.
