class Lazyalias < Formula
  desc "Interactive CLI to manage frequent commands with YAML config and clipboard support"
  homepage "https://github.com/sergiorivas/lazyalias"
  version "v0.1.15"

  on_macos do
    on_arm do
      url "https://github.com/sergiorivas/lazyalias/releases/download/vv0.1.15/lazyalias-darwin-arm64"
      sha256 "1a401bc09084ef281f93acfc83822ab623d9faa82dfa95d15c2e253942a9e5fa"
      def install
        bin.install "lazyalias-darwin-arm64" => "lazyalias"
      end
    end
    on_intel do
      url "https://github.com/sergiorivas/lazyalias/releases/download/vv0.1.15/lazyalias-darwin-amd64"
      sha256 "79732c3fe51a92c751e02e356f2938ad20efe310a9addadeb89d975af5947ce9"
      def install
        bin.install "lazyalias-darwin-amd64" => "lazyalias"
      end
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/sergiorivas/lazyalias/releases/download/vv0.1.15/lazyalias-linux-arm64"
      sha256 "f634ab479813a9432496864c71c099cb2bb964c8e898b0bd0a4b9956fc2372da"
      def install
        bin.install "lazyalias-linux-arm64" => "lazyalias"
      end
    end
    on_intel do
      url "https://github.com/sergiorivas/lazyalias/releases/download/vv0.1.15/lazyalias-linux-amd64"
      sha256 "a2689b6dfe2667b09b9d9d11b89196300cf183d9aec5a64f9f8f3a6bf9925b57"
      def install
        bin.install "lazyalias-linux-amd64" => "lazyalias"
      end
    end
  end

  test do
    assert_match "Welcome", shell_output("\#{bin}/lazyalias 2>&1", 1)
  end
end
