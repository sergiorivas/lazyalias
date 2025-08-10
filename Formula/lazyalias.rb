class Lazyalias < Formula
  desc "Interactive CLI to manage frequent commands with YAML config and clipboard support"
  homepage "https://github.com/sergiorivas/lazyalias"
  version "v0.1.14"

  on_macos do
    on_arm do
      url "https://github.com/sergiorivas/lazyalias/releases/download/vv0.1.14/lazyalias-darwin-arm64"
      sha256 "ad03694a1d19f223e0c13fdbdba9e65395c17795ad89e0a43c630487f3a3169d"
      def install
        bin.install "lazyalias-darwin-arm64" => "lazyalias"
      end
    end
    on_intel do
      url "https://github.com/sergiorivas/lazyalias/releases/download/vv0.1.14/lazyalias-darwin-amd64"
      sha256 "501240adafe1872406495fe9489df8ab95eac6694c2269cb055ad1ee9820170d"
      def install
        bin.install "lazyalias-darwin-amd64" => "lazyalias"
      end
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/sergiorivas/lazyalias/releases/download/vv0.1.14/lazyalias-linux-arm64"
      sha256 "9e121b64579fd4bd6034d887f945d94b6072dcbacdafee3f64c22a5f50e320f1"
      def install
        bin.install "lazyalias-linux-arm64" => "lazyalias"
      end
    end
    on_intel do
      url "https://github.com/sergiorivas/lazyalias/releases/download/vv0.1.14/lazyalias-linux-amd64"
      sha256 "dce986239eebc2f34c6f858b2aacef8a7c5f822c9813d277761c6228c29644b0"
      def install
        bin.install "lazyalias-linux-amd64" => "lazyalias"
      end
    end
  end

  test do
    assert_match "Welcome", shell_output("\#{bin}/lazyalias 2>&1", 1)
  end
end
