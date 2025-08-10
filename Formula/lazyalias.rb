class Lazyalias < Formula
  desc "Interactive CLI to manage frequent commands with YAML config and clipboard support"
  homepage "https://github.com/sergiorivas/lazyalias"
  version "NEW_VERSION"

  on_macos do
    on_arm do
      url "https://github.com/sergiorivas/lazyalias/releases/download/vNEW_VERSION/lazyalias-darwin-arm64"
      sha256 "SHA256_DARWIN_ARM64"
      def install
        bin.install "lazyalias-darwin-arm64" => "lazyalias"
      end
    end
    on_intel do
      url "https://github.com/sergiorivas/lazyalias/releases/download/vNEW_VERSION/lazyalias-darwin-amd64"
      sha256 "SHA256_DARWIN_AMD64"
      def install
        bin.install "lazyalias-darwin-amd64" => "lazyalias"
      end
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/sergiorivas/lazyalias/releases/download/vNEW_VERSION/lazyalias-linux-arm64"
      sha256 "SHA256_LINUX_ARM64"
      def install
        bin.install "lazyalias-linux-arm64" => "lazyalias"
      end
    end
    on_intel do
      url "https://github.com/sergiorivas/lazyalias/releases/download/vNEW_VERSION/lazyalias-linux-amd64"
      sha256 "SHA256_LINUX_AMD64"
      def install
        bin.install "lazyalias-linux-amd64" => "lazyalias"
      end
    end
  end

  test do
    assert_match "Welcome", shell_output("#{bin}/lazyalias 2>&1", 1)
  end
end
