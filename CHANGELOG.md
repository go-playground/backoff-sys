# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.3.0] - 2023-05-08
### Changed
- re-export ErrMaxAttemptsReached from github.com/go-playground/pkg to allow for more reusable coe between packages.

## [1.2.0] - 2023-05-07
### Added
- Added MaxAttempts(..) to configuration and ErrorMaxAttemptsReached to be returned by the Sleep(...) function when configured. 

[Unreleased]: https://github.com/go-playground/backoff-sys/compare/v1.3.0...HEAD
[1.3.0]: https://github.com/go-playground/backoff-sys/compare/v1.3.0...v1.2.0
[1.2.0]: https://github.com/go-playground/backoff-sys/commit/v1.2.0