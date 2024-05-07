# Home

Home acts as a reverse proxy to direct all traffic between projects from a single domain

Currently points to 2 services:
- {DOMAIN}/shorten points to https://shortener-prod-hqkniphctq-km.a.run.app/shorten
- else, {DOMAIN}/* points to https://dylanbutler-dev.web.app


- https://dylanbutler.dev
- https://dylanbutler.dev/shorten
