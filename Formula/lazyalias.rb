class Lazyalias < Formula
  desc "LazyAlias: interactive CLI to manage and run frequently used project commands"
  homepage "https://github.com/sergiorivas/lazyalias"
  url "https://github.com/sergiorivas/lazyalias/archive/refs/tags/v0.1.13.tar.gz"
  sha256 "d5558cd419c8d46bdc958064cb97f963d1ea793866414c025906ec15033512ed"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", bin/"lazyalias", "./cmd/lazyalias"
  end

  def caveats
    <<~EOS
      LazyAlias builds from source via Homebrew (macOS/Linux).
      Windows users: download the .exe from GitHub Releases.
    EOS
  end
end
