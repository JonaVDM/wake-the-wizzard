FROM node:18-alpine AS frontend-builder
WORKDIR /app
RUN npm i -g pnpm
COPY frontend/package.json .
COPY frontend/pnpm-lock.yaml .
RUN pnpm install --frozen-lockfile
COPY frontend .
RUN npm run build

FROM golang:1.19.4-alpine AS go-builder
WORKDIR /app
COPY . .
COPY --from=frontend-builder /app/dist /app/frontend/dist
RUN go build .

# TODO: find out why scratch didn't work
FROM alpine
COPY --from=go-builder /app/wake-the-wizzard /wol
ENV GIN_MODE=release
EXPOSE 3080
CMD ["/wol", "serve"]
