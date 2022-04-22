# Release Notes for Varly

## 0.4.0 - 2022-04-22
### Added
- Added move coverage for the Spanish language translations
- Added a way for you to refresh the **Layer Setup** screen when the layers folder changes externally

### Fixed
- Fixed issue #4 where the **Layer Setup** screen would not update when changes were made in Finder/Explorer
- Fixed issue #3 where the **Artist** and **Base URI** were not used in metadata generation

### Updated
- Updated the number of workers that can run concurrently during generation to the number of physical CPUs X 3
- Updated the overall design of the sidebar and the icons it uses for its links
- Updated the macOS installer to use a **Universal** binary that works for Inter and M1 macs
- Updated a few small parts of the UI to simplify appearance and use of whitespace

## 0.3.0 - 2022-04-18
### Added
- Added new binary for intel based macs and the first M1 mac binary
- Added support for marking assets as duplicates when DNA is a duplicate

### Updated
- Updated the Start screen to show document create, open, recent
- Updated the Start screen to update anytime it's visited
- Updated several small UI elements and workflows

### Fixed
- Fixed issue where Start screen would get stuck when loading layers
- Fixed issue where generation could be easily interrupted
- Fixed issue where confetti would sometimes not fill the screen
- Fixed issue where non-unique DNAs would not generate an asset


## 0.2.0 - 2022-04-16
### Added
- Added multi-language support
- Added language switching via the language menu
- Added partial support for Spanish
- Added screen caching with the ability to refresh as needed

### Updated
- Updated Run screen with better error handling

## 0.1.0 - 2022-04-15
- Initial (early alpha) release
