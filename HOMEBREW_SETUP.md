# Homebrew Tap Setup Guide

This guide will help you set up a Homebrew tap so users can install `nato` with `brew install nato`.

## Important: Tap Naming

The `homebrew-tap` repository can contain formulas for **multiple projects**. You don't need a separate tap for each tool!

- ✅ `homebrew-tap` - One tap for all your tools (recommended)
- ❌ `homebrew-nato-tap` - Unnecessary, unless you want project-specific taps

If you already have a `homebrew-tap` repository, GoReleaser will just add the `nato.rb` formula alongside your other formulas. If you don't have one yet, create it once and reuse it for all future projects.

## Steps to Set Up Homebrew Tap

### 1. Create the Homebrew Tap Repository (If you don't have one)

Create a new GitHub repository called `homebrew-tap`:

```bash
# On GitHub, create a new repository named: homebrew-tap
# Make sure it's public and owned by the Narven organization/user
```

Or via command line:
```bash
# Create local repo
mkdir homebrew-tap
cd homebrew-tap
git init

# Create a README
echo "# Homebrew Tap for Narven Tools" > README.md
git add README.md
git commit -m "Initial commit"

# Create on GitHub (requires GitHub CLI)
gh repo create Narven/homebrew-tap --public --source=. --remote=origin

# Push
git push -u origin main
```

### 2. Update GoReleaser Token Permissions

Make sure your `GITHUB_TOKEN` has permissions to push to the `homebrew-tap` repository:
- Go to GitHub Settings → Developer settings → Personal access tokens
- Ensure your token has `repo` scope with write access

### 3. Create Your First Release

When you create a git tag and run GoReleaser, it will automatically:
1. Build the binaries
2. Create a GitHub release
3. Generate and push the Homebrew formula to `homebrew-tap`

```bash
# Tag your release
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0

# Run GoReleaser (will auto-update the tap)
export GITHUB_TOKEN=your_token_here
goreleaser release --clean
```

### 4. Verify the Formula

After running GoReleaser, check that the formula was created:
- Go to `https://github.com/Narven/homebrew-tap`
- You should see a `Formula/nato.rb` file

### 5. Test the Installation

Users can now install your tool:

```bash
brew tap Narven/tap
brew install nato
```

Or they can install directly:
```bash
brew install Narven/tap/nato
```

## Troubleshooting

### GoReleaser fails to update the tap
- Check that `homebrew-tap` repository exists
- Verify your `GITHUB_TOKEN` has write access to that repo
- Make sure the repository name matches exactly: `homebrew-tap`

### Formula not updating
- Check GoReleaser logs for errors
- Verify the tap repository is public
- Ensure the formula file is in the correct location (`Formula/nato.rb`)

## Alternative: Submit to Homebrew Core

If you want `brew install nato` without the tap, you can submit to Homebrew's core repository:

1. Read the [Homebrew contribution guide](https://docs.brew.sh/Adding-Software-to-Homebrew)
2. Create a pull request to [homebrew-core](https://github.com/Homebrew/homebrew-core)
3. Requirements:
   - Your project must have stable releases (at least a few)
   - Binary downloads must work
   - Project must be actively maintained
   - Must have documentation

The personal tap approach is faster and easier for getting started!
