class Lazyalias < Formula
  desc "Interactive CLI to manage frequent commands with YAML config and clipboard support"
  homepage "https://github.com/sergiorivas/lazyalias"
  version "0.1.16"

  on_macos do
    on_arm do
      url "https://github.com/sergiorivas/lazyalias/releases/download/v0.1.16/lazyalias-darwin-arm64"
      sha256 "4fdf0e6ff370517b2d57f861dd6e76dc02c12a99fa7d6249c3b5872e10ccdd2f"
      def install
        bin.install "lazyalias-darwin-arm64" => "lazyalias"
      end
    end
    on_intel do
      url "https://github.com/sergiorivas/lazyalias/releases/download/v0.1.16/lazyalias-darwin-amd64"
      sha256 "937c19f5bf2647e72d30b1733d6b040e97ef4ef07d43d0e3d8bd1e9a2272b19a"
      def install
        bin.install "lazyalias-darwin-amd64" => "lazyalias"
      end
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/sergiorivas/lazyalias/releases/download/v0.1.16/lazyalias-linux-arm64"
      sha256 "e339a6ce9cedef17274c62c2cb34c901431114a531f6f34ac7d07042fd44c9c1"
      def install
        bin.install "lazyalias-linux-arm64" => "lazyalias"
      end
    end
    on_intel do
      url "https://github.com/sergiorivas/lazyalias/releases/download/v0.1.16/lazyalias-linux-amd64"
      sha256 "490e7336b0a02ee8b7a35c9646ce9c191e50d2659e430b5ca6a6295b92f82075"
      def install
        bin.install "lazyalias-linux-amd64" => "lazyalias"
      end
    end
  end

  test do
    assert_match "Welcome", shell_output("\#{bin}/lazyalias 2>&1", 1)
  end
end
