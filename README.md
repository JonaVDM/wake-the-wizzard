# Wake the wizard

A simple Wake on Lan server.

## Running via docker

To run the app within a container you need to run it on host mode.

```bash
docker volume create wol-storage
docker run --rm -d --network host --volume wol-storage:/wol-server ghcr.io/jonavdm/wake-the-wizzard
```

After this, the website will be available at http://[ip]:3080.
