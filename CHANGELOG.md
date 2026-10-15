# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.0] - 2026-10-15

### Security

- **Built and tested with Go 1.26.9 (was Go 1.21).** `go.mod` now declares
  `go 1.26.0` and CI runs on 1.26.9, which carries the standard-library security
  fixes Go 1.21 stopped receiving in August 2024 (including the 2026-10-09
  `net/http` batch). A new `govulncheck` job scans the module on every PR, on
  push to `main`, and weekly. Dependencies are already at their latest versions.
- **Every GitHub Action in CI is pinned to a full commit SHA** (with its version
  in a trailing comment), so a retagged or compromised action can't change what
  CI runs. `test.yml` moves off the floating `checkout@v4` / `setup-go@v5` /
  `codecov-action@v4` tags onto the same releases spawn uses (checkout v7.0.1,
  setup-go v7.0.0, codecov-action v7.1.1). Dependabot now bumps the pins weekly,
  and `workflows_test.go` fails on any unpinned action or a missing Dependabot
  entry (#4).

### Added

- **The changelog policy is now enforced in CI rather than by habit.** A PR that changes
  non-test Go source without touching `CHANGELOG.md` fails, and `changelog_test.go`
  checks `[Unreleased]` for duplicate group headings, unknown group names, entries
  outside a group, and releases missing a compare link.
  This policy has been suite-wide for a while but only `spawn` enforced it — where it
  immediately earned its keep, catching a duplicate `### Fixed` **four times in one
  session** and a PR that had merged with no entry at all (found only at the next
  release, against an empty `[Unreleased]`, with the entries reconstructed from the diff
  at tag time).
  `scripts/changelog-consolidate.py` (and `make changelog-fix` where there's a Makefile)
  merges duplicate groups mechanically, because two PRs each adding their own
  `### Fixed` is a routine conflict that merges cleanly for git and badly for the format
  — not a mistake worth hand-fixing each time.

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

[Unreleased]: https://github.com/spore-host/spore-host-plugin-globus/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/spore-host/spore-host-plugin-globus/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/spore-host/spore-host-plugin-globus/releases/tag/v0.1.0
