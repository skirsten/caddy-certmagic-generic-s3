# Caddy Certmagic Generic S3

Implements <https://github.com/thomersch/certmagic-generic-s3> as a Caddy storage plugin. With this plugin you can use S3 as certificate storage without any additional Redis for locking.

## Building

```sh
go get -u github.com/caddyserver/xcaddy/cmd/xcaddy

xcaddy build --output ./caddy --with github.com/skirsten/caddy-certmagic-generic-s3
```

## Config

```json
{
  "storage": {
    "module": "generic_s3",
    "endpoint": "s3.amazonaws.com",
    "bucket": "my-bucket",
    "access_key_id": "...",
    "secret_access_key": "...",
    "prefix": "caddy",
    "encryption_key": "32 character key here"
  }
}
```
