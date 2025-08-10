class Lazyalias < Formula
  desc "Interactive CLI to manage frequent commands with YAML config and clipboard support"
  homepage "https://github.com/sergiorivas/lazyalias"
  version "0.1.17"

  on_macos do
    on_arm do
      url "https://github.com/sergiorivas/lazyalias/releases/download/v0.1.17/lazyalias-darwin-arm64"
      sha256 "4c8233ce371ece90e410c0d1a38f55a5022f497c592060e1c9ad5d64c7095dea"
      def install
        bin.install "lazyalias-darwin-arm64" => "lazyalias"
      end
    end
    on_intel do
      url "https://github.com/sergiorivas/lazyalias/releases/download/v0.1.17/lazyalias-darwin-amd64"
      sha256 "6d924b880fe72dc5891c47a7d8f24a028b0362e0d321c68433e50e9d8dcf67c4"
      def install
        bin.install "lazyalias-darwin-amd64" => "lazyalias"
      end
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/sergiorivas/lazyalias/releases/download/v0.1.17/lazyalias-linux-arm64"
      sha256 "31aa4c9fb838bd746e8b933a059e9761f5429da3aa4de139852502a225ee477b"
      def install
        bin.install "lazyalias-linux-arm64" => "lazyalias"
      end
    end
    on_intel do
      url "https://github.com/sergiorivas/lazyalias/releases/download/v0.1.17/lazyalias-linux-amd64"
      sha256 "6d4e894af12e666a27b678ebce061dd885c3833d6a1bae76225b75d9edc07076"
      def install
        bin.install "lazyalias-linux-amd64" => "lazyalias"
      end
    end
  end

  test do
    assert_match "Welcome", shell_output("\#{bin}/lazyalias 2>&1", 1)
  end
end
