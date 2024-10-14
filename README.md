# butla

butla is a simple URL shortening service written in Go, built in Docker. It is made for speed and nothing else.

- ❌ No UI
- ❌ No GUI
- ❌ No CLI
- ❌ No API
- ❌ No volume mounts
- ❌ No container networking
- ❌ No database
- ✅ Simple embedded in-memory configuration
- ✅ Stupid simple
- ✅ Stupid fast

## Usage

As the config is embedded you will need to build your own image, you can do this by cloning the repo, updating the config and docker building.

```shell
git clone https://github.com/dbut2/butla.git
cd butla
docker build -t butla .
```

There is an example compose.yaml included with some labels for Traefik.

## Redirection

For a given request to {DOMAIN}/{PATH}, the first item in this list found in the config will be returned with a temporary redirect (307):

1. {DOMAIN}/{PATH}
2. default/{PATH}
3. {DOMAIN}/default
4. default/default
