class Lazyalias < Formula
  desc "Interactive CLI to manage frequent commands with YAML config and clipboard support"
  homepage "https://github.com/sergiorivas/lazyalias"
  version "0.1.18"

  on_macos do
    on_arm do
      url "https://github.com/sergiorivas/lazyalias/releases/download/v0.1.18/lazyalias-darwin-arm64"
      sha256 "e75537351329b43274865dca313fc9a5298347fee8e1b8a664f6be97e04b3163"
      def install
        bin.install "lazyalias-darwin-arm64" => "lazyalias"
      end
    end
    on_intel do
      url "https://github.com/sergiorivas/lazyalias/releases/download/v0.1.18/lazyalias-darwin-amd64"
      sha256 "ad8ae6ad10788ed033623e27fe976b9e27adbfe6e6c85101d86e53d0ebdb9946"
      def install
        bin.install "lazyalias-darwin-amd64" => "lazyalias"
      end
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/sergiorivas/lazyalias/releases/download/v0.1.18/lazyalias-linux-arm64"
      sha256 "1c3ea2ebc32bbb684c25c7a26bfdfc7be2f0e213e7e43ec2cdbbf5dc40437564"
      def install
        bin.install "lazyalias-linux-arm64" => "lazyalias"
      end
    end
    on_intel do
      url "https://github.com/sergiorivas/lazyalias/releases/download/v0.1.18/lazyalias-linux-amd64"
      sha256 "76482cb14d8c4ca5bfa560663c76c8a0027bb7ee85660674fc4cb2744cf75eab"
      def install
        bin.install "lazyalias-linux-amd64" => "lazyalias"
      end
    end
  end

  test do
    assert_match "Welcome", shell_output("\#{bin}/lazyalias 2>&1", 1)
  end
end
