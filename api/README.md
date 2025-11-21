Net63 VPN API

Lightweight API to index and serve OpenVPN configuration files.

This service can either fetch a plain-text list of VPN config URLs from a remote URL or index a local `./flies` directory and serve files directly. It caches results locally to reduce network usage.

Features

- JSON API returning a list of config URLs and metadata
- Plain-text endpoint that returns one URL per line (suitable for simple clients)
- Optional local indexing of `./flies` directory
- Local file serving endpoint for `.ovpn` files (or any files placed in `./flies`)
- Simple cache stored in `~/.cache/net63-vpn/vpn_list.json`

Quick start

1. Install requirements (prefer a virtualenv):

```bash
python -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
```

2. Place any `.ovpn` files you want to serve in the `flies/` directory next to `app.py`:

```bash
mkdir -p flies
cp /path/to/example.ovpn flies/
```

3. Run the app:

```bash
export INDEX_LOCAL=true      # optional: prefer local ./flies
# or set FILES_BASE_URL=https://cdn.example.com/vpns
python app.py
```

Endpoints

- GET /api/health — health check
- GET /api/vpn/list — JSON list of URLs, cached
- GET /api/vpn/plainlist — plain text list (one URL per line)
- POST /api/vpn/reindex — force reindex / fetch and refresh cache
- GET /api/local-files/<filename> — serve files from `./flies` (download)
- GET /api/vpn/download/<filename> — download an item from the JSON list (remote fetch + cache)
- GET /api/vpn/cache — info about cached `.ovpn` files
- DELETE /api/vpn/cache/<filename> — remove a cached file
- DELETE /api/vpn/cache — clear cache

Configuration

Set these environment variables to control behavior:

- VPN_LIST_URL — remote plain-text list URL (one URL per line). Default: https://example.com/vpn-list.txt
- INDEX_LOCAL — when "true", index `./flies` instead of fetching remote list. Default: false
- FILES_BASE_URL — if set, local file entries will be returned as absolute URLs using this base. Default: (empty — returns `/api/local-files/<name>`)

Security notes

- Serving files from `./flies` is convenient for local hosting, but add auth or IP filtering if exposing to the public internet.
- Be careful with `FILES_BASE_URL` and hosting credentials.

Example: Return plain text list

```bash
curl http://localhost:5000/api/vpn/plainlist
```

This will print a newline-separated list of URLs that clients can download.

Next steps

- Add authentication for protected downloads
- Add optional gzip compression for plain list
- Add monitoring/metrics

---
Net63 VPN — a small file-indexing API for OpenVPN configs.
# net63
