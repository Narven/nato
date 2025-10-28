# NATO

Nato is a tiny CLI tool that converts words in to NATO phonetic alphabet

## Usage

`nato <sentence>`

```sh
nato 123
# One Two Three
nato TW11 5BA
# Tango Whiskey One One Five Bravo Alfa
```

## Installation

### Homebrew (macOS)

```sh
brew tap Narven/tap
brew install nato
```

### Manual Installation (macOS)

If you see a Gatekeeper warning about malware:

1. **Temporary solution** (bypass the warning):
   ```sh
   xattr -d com.apple.quarantine /path/to/nato
   ```

2. **Or allow it through System Settings**:
   - Try running it once
   - Go to System Settings → Privacy & Security
   - Click "Open Anyway" next to the nato warning

3. **Proper solution**: The binary is ad-hoc signed, but for full Gatekeeper trust, you'd need an Apple Developer account code signature.

### Linux & Windows

Download the binary from [releases](https://github.com/Narven/nato/releases) and use directly.

## Credits
All the credits go to [EvanHahn](https://codeberg.org/EvanHahn/dotfiles/src/commit/843b9ee13d949d346a4a73ccee2a99351aed285b/home/bin/bin/nato), I just converted it to [Go](https://go.dev/).
