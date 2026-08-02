# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.1.0] - 2026-08-02

First tagged release of garabic — a set of functions for Arabic text processing in Go.

### Added

- `Normalize` — normalize Arabic text for searching and indexing.
- `RemoveHarakat` — strip diacritics (tashkeel) from Arabic text.
- `Shape` — reconstruct Arabic and Persian glyphs (including پ چ ژ گ) so text renders connected and in the right order in images and other non-RTL-aware surfaces.
- `SpellNumber` — spell out numbers in Arabic words.
- `Tashkeel` — add matching diacritics to Arabic text (initial support).
- `Justify` — justify an Arabic sentence to a given width using kashida (tatweel) stretching, with space-padding fallback.
- `ToArabicDigits`, `ToPersianDigits`, `ToEnglishDigits` — convert between English, Arabic and Persian digits.
- `IsArabic`, `IsArabicLetter` — Arabic-script detection helpers.
- Test suites for Arabic and Persian shaping, digits, justification and tashkeel (96% coverage).
- GitHub Actions workflows for build checks and race-enabled tests.

### Fixed

- `Shape` no longer misplaces diacritics on words combining tashkeel with a lam-alef ligature (لا / لأ / لإ / لآ), and handles tatweel in input correctly.
- `Tashkeel` no longer panics when a sentence ends with a preposition (حرف جر).
- `Justify` no longer panics on empty input, single words, or tashkeel with lam-alef; width math is now rune-based and exact.
- `ToPersianDigits` now emits proper Extended Arabic-Indic codepoints (۰–۹, U+06F0–U+06F9) for all ten digits instead of a mix of Arabic and Persian forms; `ToEnglishDigits` accepts both.

### Changed

- Module path is `github.com/abdullahdiaa/garabic`.
- Requires Go 1.25+; `golang.org/x/text` upgraded to v0.40.0.

[v0.1.0]: https://github.com/AbdullahDiaa/garabic/releases/tag/v0.1.0
