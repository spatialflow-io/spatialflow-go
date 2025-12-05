# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

[Unreleased]: https://github.com/spatialflow-io/spatialflow-go/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/spatialflow-io/spatialflow-go/releases/tag/v0.1.0
