# butla

butla is a simple URL shortening service written in Go, built in Docker. It is made for speed and nothing else.

- ❌ No UI
- ❌ No GUI
- ❌ No CLI
- ❌ No API
- ❌ No volume mounts
- ❌ No container networking
- ❌ No database
- ✅ Simple in-memory configuration
- ✅ Stupid simple
- ✅ Stupid fast

## Deployment

1. Create a `config.yaml` file (see [`config.example.yaml`](config.example.yaml))

2. Run with Docker:
   ```bash
   docker run -p 8080:8080 -v $(pwd)/config.yaml:/app/config.yaml ghcr.io/dbut2/butla
   ```

3. Or use Docker Compose:
   ```yaml
   services:
     butla:
       image: ghcr.io/dbut2/butla
       ports:
         - "8080:8080"
       volumes:
         - ./config.yaml:/app/config.yaml
   ```

## Redirection

For a given request to {DOMAIN}/{PATH}, the first item in this list found in the config will be returned with a temporary redirect (307):

1. {DOMAIN}/{PATH}
2. default/{PATH}
3. {DOMAIN}/default
4. default/default
