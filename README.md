# payjp-grpc-go

based on: https://github.com/po3rin/vue-golang-payment-app

## Tech-stacks

- Go
- gRPC
- frourio
  - Next.js
  - Prisma
- TypeScript
- PAYJP

## Requirements

- node, yarn
- docker-compose
- protoc

## Getting started

Start gRPC server & DB

```bash
% source env.sh
% echo "PAYJP_SECRET_KEY=[YOUR KEY]" > secrets.env
% build
% up
```

Start frourio application

```bash
% cd frontend
% yarn
% export PAYJP_PUBLIC_KEY=[YOUR KEY]
% yarn dev
```
