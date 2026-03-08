class Switchboard < Formula
  desc "Low-latency encrypted tmux session router"
  homepage "https://github.com/arcaven/switchboard"
  version "VERSION_PLACEHOLDER"
  license "MIT"

  if Hardware::CPU.arm?
    url "https://github.com/arcaven/switchboard/releases/download/TAG_PLACEHOLDER/switchboard-darwin-arm64"
    sha256 "SHA256_ARM64_PLACEHOLDER"
  else
    url "https://github.com/arcaven/switchboard/releases/download/TAG_PLACEHOLDER/switchboard-darwin-amd64"
    sha256 "SHA256_AMD64_PLACEHOLDER"
  end

  def install
    binary_name = Hardware::CPU.arm? ? "switchboard-darwin-arm64" : "switchboard-darwin-amd64"
    bin.install binary_name => "switchboard"
  end

  test do
    assert_match "switchboard", shell_output("#{bin}/switchboard --version 2>&1")
  end
end
