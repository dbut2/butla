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

## Redirection

For a given request to {DOMAIN}/{PATH}, the first item in this list found in the config will be returned with a temporary redirect (307):

1. {DOMAIN}/{PATH}
2. default/{PATH}
3. {DOMAIN}/default
4. default/default
