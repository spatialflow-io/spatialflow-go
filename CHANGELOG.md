# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.0.0] - 2026-10-08

The Go, Python and Node SDKs ship this release together.

**Breaking changes.** The module path is now `github.com/spatialflow-io/spatialflow-go/v2`, so every import changes. The generated client no longer has the operations removed from the API since 1.1.0, which returned 404: the simulation calls (`/simulations`, 15 operations), the route tester (`/route-tester/test`, 2), the test helpers (`/test/*`, 4) and the `GET` email verification calls under `/auth/verify-email` (2).

### Added

- `VerifyWorkflowSignature` verifies a workflow Webhook action delivery: `X-SpatialFlow-Signature` (`sha256=<hex>` HMAC-SHA256 over `<timestamp>.<raw body>`) and `X-SpatialFlow-Timestamp`, with a default 5 minute tolerance in both directions. A body that isn't JSON is returned as a string. `VerifyWebhookSignature` is unchanged.

### Fixed

- **Webhook signature verification**: `VerifyWebhookSignature` now matches the platform's `X-SF-Signature: sha256=<hex>` (HMAC-SHA256 over the raw body), replacing the Stripe-style `t=,v1=` scheme that rejected every real webhook.
- **Webhook payload mapping**: `WebhookEvent` maps the backend `event` and `timestamp` keys via a custom `UnmarshalJSON`, while still accepting the legacy `type` / `created_at` shape, so `event.Type` is populated.
- The `tolerance` parameter is retained but deprecated and ignored; `DefaultTimestampTolerance` is kept as a deprecated constant.

## [1.1.0] - 2026-04-05

### Added

- Generated API client in `spatialflow/_generated/`, produced by OpenAPI Generator when the module is synced to the public repository
- `ErrConflict` sentinel error, returned for HTTP 409 responses
- Tag-triggered release workflow in the public repository (`go build` and `go test` on Go 1.21 to 1.23)

### Changed

- `SDKVersion` and the `User-Agent` header report `1.1.0` (they reported `0.1.0` before)

## [0.1.0] - 2024-12-04

### Added

- Initial alpha release
- Client configuration with API key and JWT token support
- Custom HTTP transport with auth header injection
- Automatic retry with exponential backoff and jitter
- Rate limit handling with `Retry-After` header support
- Typed errors with sentinel error support (`ErrAuthentication`, `ErrNotFound`, etc.)
- Generic pagination helper (`Pager[T]`)
- Webhook signature verification (HMAC-SHA256)
- File upload helper with presigned URL workflow
- Job polling with timeout and progress callbacks
- Unit tests for all core helpers

### Known Limitations

- Generated API client not yet included (run `generate.sh go` to generate)
- Integration tests require live API (env-gated)
- No streaming support yet

[2.0.0]: https://github.com/spatialflow-io/spatialflow-go/compare/v1.1.0...v2.0.0
[1.1.0]: https://github.com/spatialflow-io/spatialflow-go/compare/v0.1.0...v1.1.0
[0.1.0]: https://github.com/spatialflow-io/spatialflow-go/releases/tag/v0.1.0
