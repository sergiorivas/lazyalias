#!/usr/bin/env bash
set -euo pipefail

# Simple, interactive release script using gum
# - Bumps version (patch|minor|major|custom)
# - Builds cross-platform binaries (Makefile)
# - Generates SHA256SUMS
# - Updates Homebrew Formula (macOS+Linux) and Scoop manifest (Windows)
# - Commits, tags, pushes
# - Publishes GitHub Release with assets via gh

REPO_OWNER="sergiorivas"
REPO_NAME="lazyalias"
FORMULA_PATH="Formula/lazyalias.rb"
SCOOP_PATH="scoop/lazyalias.json"

require() {
  command -v "$1" >/dev/null 2>&1 || {
    echo "Missing dependency: $1"
    exit 1
  }
}

require gum
require git
require go
require gh
require make

section() { gum style --foreground 212 --bold "$1"; }
ok()      { gum style --foreground 42  "$1"; }
warn()    { gum style --foreground 178 "$1"; }
err()     { gum style --foreground 196 "$1"; }

# 0) Safety checks
section "Pre-flight checks"
if [[ -n "$(git status --porcelain)" ]]; then
  err "Working tree is not clean. Commit or stash your changes first."
  git status --porcelain
  exit 1
fi
git fetch --tags >/dev/null 2>&1 || true

# 1) Choose version bump
CURRENT_VERSION="$(tr -d ' \n' < VERSION)"
section "Current version: ${CURRENT_VERSION}"

CHOICE="$(gum choose --header="Select version bump" patch minor major custom)"
if [[ "$CHOICE" == "custom" ]]; then
  NEW_VERSION="$(gum input --placeholder "e.g. 0.1.12" --prompt "Enter new version: ")"
  [[ "$NEW_VERSION" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]] || {
    err "Invalid semantic version: $NEW_VERSION"
    exit 1
  }
else
  IFS='.' read -r MA MI PA <<<"$CURRENT_VERSION"
  case "$CHOICE" in
    patch) PA=$((PA+1));;
    minor) MI=$((MI+1)); PA=0;;
    major) MA=$((MA+1)); MI=0; PA=0;;
  esac
  NEW_VERSION="${MA}.${MI}.${PA}"
fi
NEW_TAG="v${NEW_VERSION}"

gum style --border normal --padding "1 2" --margin "1 0" \
  "Release summary" \
  "From: ${CURRENT_VERSION}" \
  "To:   ${NEW_VERSION} (${NEW_TAG})"

gum confirm "Proceed with release?" || { warn "Aborted."; exit 0; }

# 2) Update VERSION
section "Updating VERSION"
printf "%s\n" "${NEW_VERSION}" > VERSION
ok "VERSION updated to ${NEW_VERSION}"

# 3) Build all binaries
section "Building cross-platform binaries"
rm -f bin/lazyalias-* || true
gum spin --title "make build" -- make build
ok "Build complete"

# 4) Generate SHA256SUMS
section "Generating SHA256SUMS"
gum spin --title "make sha256" -- make sha256
[[ -f SHA256SUMS ]] || { err "SHA256SUMS not generated"; exit 1; }

# Helper to extract checksum for a specific artifact
checksum() { awk -v f="bin/$1" '$2==f {print $1}' SHA256SUMS; }

SHA_DARWIN_AMD64="$(checksum 'lazyalias-darwin-amd64')"
SHA_DARWIN_ARM64="$(checksum 'lazyalias-darwin-arm64')"
SHA_LINUX_AMD64="$(checksum 'lazyalias-linux-amd64')"
SHA_LINUX_ARM64="$(checksum 'lazyalias-linux-arm64')"
SHA_WIN_AMD64="$(checksum 'lazyalias-windows-amd64.exe')"
SHA_WIN_ARM64="$(checksum 'lazyalias-windows-arm64.exe')"

for v in SHA_DARWIN_AMD64 SHA_DARWIN_ARM64 SHA_LINUX_AMD64 SHA_LINUX_ARM64 SHA_WIN_AMD64 SHA_WIN_ARM64; do
  [[ -n "${!v}" ]] || { err "Missing checksum: $v"; exit 1; }
done
ok "Checksums ready"

# 5) Update Homebrew Formula (macOS + Linux using prebuilt binaries)
section "Updating Homebrew Formula (${FORMULA_PATH})"
mkdir -p "$(dirname "${FORMULA_PATH}")"
cat > "${FORMULA_PATH}" <<EOF
class Lazyalias < Formula
  desc "Interactive CLI to manage frequent commands with YAML config and clipboard support"
  homepage "https://github.com/${REPO_OWNER}/${REPO_NAME}"
  version "${NEW_VERSION}"

  on_macos do
    on_arm do
      url "https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${NEW_TAG}/lazyalias-darwin-arm64"
      sha256 "${SHA_DARWIN_ARM64}"
      def install
        bin.install "lazyalias-darwin-arm64" => "lazyalias"
      end
    end
    on_intel do
      url "https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${NEW_TAG}/lazyalias-darwin-amd64"
      sha256 "${SHA_DARWIN_AMD64}"
      def install
        bin.install "lazyalias-darwin-amd64" => "lazyalias"
      end
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${NEW_TAG}/lazyalias-linux-arm64"
      sha256 "${SHA_LINUX_ARM64}"
      def install
        bin.install "lazyalias-linux-arm64" => "lazyalias"
      end
    end
    on_intel do
      url "https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${NEW_TAG}/lazyalias-linux-amd64"
      sha256 "${SHA_LINUX_AMD64}"
      def install
        bin.install "lazyalias-linux-amd64" => "lazyalias"
      end
    end
  end

  test do
    assert_match "Welcome", shell_output("\#{bin}/lazyalias 2>&1", 1)
  end
end
EOF
ok "Formula written"

# 6) Update Scoop manifest (Windows)
section "Updating Scoop manifest (${SCOOP_PATH})"
mkdir -p "$(dirname "${SCOOP_PATH}")"
cat > "${SCOOP_PATH}" <<EOF
{
  "version": "${NEW_VERSION}",
  "description": "Interactive CLI to manage frequent commands with YAML config and clipboard support",
  "homepage": "https://github.com/${REPO_OWNER}/${REPO_NAME}",
  "license": "MIT",
  "architecture": {
    "64bit": {
      "url": "https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${NEW_TAG}/lazyalias-windows-amd64.exe",
      "hash": "${SHA_WIN_AMD64}",
      "bin": "lazyalias-windows-amd64.exe"
    },
    "arm64": {
      "url": "https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${NEW_TAG}/lazyalias-windows-arm64.exe",
      "hash": "${SHA_WIN_ARM64}",
      "bin": "lazyalias-windows-arm64.exe"
    }
  }
}
EOF
ok "Scoop manifest written"

# 7) Release notes (auto or manual)
section "Release notes"
NOTES_MODE="$(gum choose --header="How to prepare release notes?" "Auto (from git log)" "Write now")"
NOTES_FILE="$(mktemp)"
if [[ "$NOTES_MODE" == "Write now" ]]; then
  gum write --placeholder "Write release notes here..." > "${NOTES_FILE}"
else
  PREV_TAG="$(git describe --tags --abbrev=0 2>/dev/null || echo "")"
  {
    echo "lazyalias ${NEW_TAG}"
    echo
    if [[ -n "${PREV_TAG}" ]]; then
      git log --pretty=format:"- %s (%h)" "${PREV_TAG}..HEAD"
    else
      echo "- Initial release"
    fi
  } > "${NOTES_FILE}"
fi
ok "Notes prepared: ${NOTES_FILE}"

# 8) Commit changes (without committing binaries)
section "Commit changes"
git add VERSION SHA256SUMS "${FORMULA_PATH}" "${SCOOP_PATH}"
git commit -m "release: ${NEW_TAG}"
ok "Commit created"

# 9) Tag
section "Create tag ${NEW_TAG}"
git tag -a "${NEW_TAG}" -m "lazyalias ${NEW_TAG}"
ok "Tag created"

# 10) Push and create GitHub Release
section "Push and publish"
DRY=false
gum confirm "Dry run (no push, no GitHub release)?" && DRY=true || true

if [[ "$DRY" == "true" ]]; then
  warn "[dry-run] Skipping push and release creation"
  echo "Would run: git push && git push --tags"
  echo "Would create GitHub release with assets:"
else
  gum spin --title "git push" -- git push
  gum spin --title "git push --tags" -- git push --tags
fi

ASSETS=(
  "bin/lazyalias-darwin-amd64"
  "bin/lazyalias-darwin-arm64"
  "bin/lazyalias-linux-amd64"
  "bin/lazyalias-linux-arm64"
  "bin/lazyalias-windows-amd64.exe"
  "bin/lazyalias-windows-arm64.exe"
  "SHA256SUMS"
)

if [[ "$DRY" == "true" ]]; then
  echo "gh release create ${NEW_TAG} ${ASSETS[*]} --title \"lazyalias ${NEW_TAG}\" --notes-file ${NOTES_FILE}"
else
  gum spin --title "gh release create ${NEW_TAG}" -- \
    gh release create "${NEW_TAG}" "${ASSETS[@]}" \
      --title "lazyalias ${NEW_TAG}" \
      --notes-file "${NOTES_FILE}"
  ok "GitHub Release published"
fi

section "Download URLs"
echo "  macOS amd64: https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${NEW_TAG}/lazyalias-darwin-amd64"
echo "  macOS arm64: https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${NEW_TAG}/lazyalias-darwin-arm64"
echo "  Linux amd64: https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${NEW_TAG}/lazyalias-linux-amd64"
echo "  Linux arm64: https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${NEW_TAG}/lazyalias-linux-arm64"
echo "  Windows amd64: https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${NEW_TAG}/lazyalias-windows-amd64.exe"
echo "  Windows arm64: https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${NEW_TAG}/lazyalias-windows-arm64.exe"

ok "All done!"
