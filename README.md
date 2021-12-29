# manufacturer-parser

Service written on Golang to get actual MAC OUI list from [here](https://linuxnet.ca/ieee/oui/nmap-mac-prefixes) and store to MongoDB. 
Needed for dealt with actual list in Wimark One.

## Services and APIs

Needed mongodb via DBURL env var to pass before execute.

## How to build

To build a binary just use:

```bash

make

```

To build docker image use provided `Dockerfile`:

```bash
docker build  .
# or
docker-compose build 
```

## How to use

See docker-compose.yml

## Copyright

Wimark Systems, 2021