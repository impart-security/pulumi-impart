# Changelog

## [0.10.3] - 2025-09-15

### Changed

- Updated dependencies

## [0.10.2] - 2025-08-20

### Changed

- Updated dependencies

## [0.10.1] - 2025-06-25

### Changed

- Updated dependencies

## [0.10.0] - 2025-05-19

### Added

- Rule resource which allows to configure both scfrip and recipe rule types.
- RuleClientIdentifier resource to manage client identifiers.
- RuleClientIdentifierStorage resource to manage client identifier storages.
- CoreRuleConfig resource to configure core rules.

### Changed

- Deprecated RuleScript and RuleScriptDependencies. Please migrate to Rule and RuleDependencies resources.

## [0.9.1] - 2025-04-22

### Changed

- Updated dependencies

## [0.9.0] - 2025-03-21

### Add

- External Link resource

## [0.8.1] - 2025-01-09

### Changed

- Updated dependencies

## [0.8.0] - 2024-11-11

### Added

- Label resource
- TagMetadata resource
- labels attribute to rule scripts, rule test cases, lists, monitors, tags
- description attribute to lists and to test cases messages and assertions

### Changed

- List implementation to use PUT and PATCH items endpoint

## [0.7.0] - 2024-09-23

### Added

- Added RuleTestCase resource

## [0.6.0] - 2024-06-17

### Added

- List functionality attribute

## [0.5.0] - 2024-06-11

### Added

- List resource
- ApiBinding disabled attribute

## [0.4.0] - 2024-04-10

### Added

- Monitor resource
- Notification template resource
- Connector resource
- Rule script dependencies resource

## [0.3.0] - 2024-02-14

### Changed

- Add json support for the log binding configuration
- Add pattern_type field
- Rename grok_pattern to pattern

## [0.2.0] - 2024-01-16

### Changed

- Fixed source hash validation
- Renamed impart_binding to impart_api_binding
- Added advanced optios to impart_api_binding
- Added impart_log_binding

## [0.1.10] - 2023-12-30

### Changed

- Updated deps

## [0.1.9] - 2023-12-29

### Changed

- Updated deps

## [0.1.8] - 2023-08-18

### Changed

- Updated deps
- Updated docs

## [0.1.7] - 2023-06-20

### Changed

- Updated docs layout
- Fixed terraform provider bridge

## [0.1.6] - 2023-06-07

### Changed

- Undo display name and title change

## [0.1.5] - 2023-06-05

### Changed

- Changed display name and title

## [0.1.4] - 2023-06-05

### Fixed

- Fixed plugin download url

## [0.1.3] - 2023-06-03

### Fixed

- Fixed registry docs

## [0.1.2] - 2023-06-02

### Added

- Changelog

## [0.1.1] - 2023-05-30

### Added

- Registry docs

## [0.1.0] - 2023-05-25

### Added

- Initial release
