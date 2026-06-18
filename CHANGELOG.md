# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed
- Module path and all repo links moved from `github.com/scttfrdmn/...` to
  `github.com/spore-host/...` to match the repo's home in the spore-host org.

## [0.1.0] - 2026-03-22

### Added
- `globus-connect-personal/plugin.yaml`: install Globus Connect Personal as a data transfer endpoint
- Local provision step uses `globus endpoint create --personal` to generate a one-time setup key
- Setup key is pushed to the remote instance via SSH tunnel
- Endpoint ID captured in `outputs.endpoint_id` for future deprovision support
- `pgrep` health check (1-minute interval)

[Unreleased]: https://github.com/spore-host/spore-host-plugin-globus/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/spore-host/spore-host-plugin-globus/releases/tag/v0.1.0
