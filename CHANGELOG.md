# Changelog

See also [Keep a Changelog](https://keepachangelog.com/en/1.1.0/)
and [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed
- Build without symbol table and debug information.

## [0.1.0] - 2024-04-21

### Added

- Functions:
  - `file.DocumentFromJSON`
  - `file.IsCommandName`
- Package `tool`.
- Tool `ema`.

### Changed

- Package `file` was improved:
  - Type `Content` implements `io.WriterTo` to write EMA form of content.
  - Type `Text` is based on `string` instead of `[]string`.

## [0.0.2] - 2024-04-12

### Added

- `LICENSE` file.

## [0.0.1] - 2024-04-12

### Added

- First package `file`.

[Unreleased]: https://github.com/dark-pink/ema/compare/v0.1.0...main
[0.1.0]: https://github.com/dark-pink/ema/compare/v0.0.2...v0.1.0
[0.0.2]: https://github.com/dark-pink/ema/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/dark-pink/ema/tag/v0.0.1
