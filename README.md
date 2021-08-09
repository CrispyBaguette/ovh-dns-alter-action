# OVH alter DNS record

This action updates an existing DNS record using the OVH API.

## Inputs

## `application-key`

**Required** Application key.

## `application-secret`

**Required** Application secret.

## `consumer-key`

**Required** Consumer key.

## `api-endpoint`

API endpoint. Defaults to `ovh-eu`.

## `dns-zone`

**Required** DNS zone.

## `record-id`

**Required** DNS record id.

## `subdomain`

Subdomain.

## `target`

DNS record target.

## `ttl`

DNS record TTL.

## Example usage

```yaml
uses: CrispyBaguette/ovh-dns-alter-action@v1.0
with:
  application-key: foo
  application-secret: bar
  consumer-key: far
  api-endpoint: ovh-us
  dns-zone: example.com
  record-id: 42
  target: "new_record_value"
```
