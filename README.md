# monomail-sync

A web interface for imapsync. Sync emails between IMAP servers with a clean UI.

## Quick Start

```bash
go run .
```

Access at http://localhost:8000

Admin login: `admin` / `admin`

## Configuration (Optional)

Create `config/config.yml` to override defaults:

```yaml
language: en
port: 8000
databaseInfo:
  adminName: admin
  adminPass: admin
  databasePath: ./db.db
sourceAndDestination:
  SourceServer: imap.example.com
  SourceMail: "@example.com"
  DestinationServer: imap.example.com
  DestinationMail: "@example.com"
email:
  SMTPHost: smtp.example.com
  SMTPPort: 587
  SMTPFrom: example
  SMTPUser: example@example.com
  SMTPPassword: password
```

Or run with custom config:
```bash
go run . -config config/config.yml
```

## Features

- Single user/bulk email sync
- Real-time progress tracking
- Dashboard with sync statistics
- IMAP account validation
- Health check endpoint (`/health`)
- Log rotation

## Endpoints

- `/` - Main sync interface
- `/admin` - Admin panel with settings
- `/login` - Admin login
- `/health` - Health check for monitoring

## Tech Stack

- Go + Gin
- SQLite
- HTMX + Alpine.js
- Tailwind CSS (daisyUI)

## License

GPL-3.0
