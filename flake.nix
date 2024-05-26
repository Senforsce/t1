{
  description = "thunderf1sh";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs?ref=nixos-unstable";
    gitignore = {
      url = "github:hercules-ci/gitignore.nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    xc = {
      url = "github:joerdav/xc";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = { self, nixpkgs, gitignore, xc }:
    let
      allSystems = [
        "x86_64-linux" # 64-bit Intel/AMD Linux
        "aarch64-linux" # 64-bit ARM Linux
        "x86_64-darwin" # 64-bit Intel macOS
        "aarch64-darwin" # 64-bit ARM macOS
      ];
      forAllSystems = f: nixpkgs.lib.genAttrs allSystems (system: f {
        inherit system;
        pkgs = import nixpkgs { inherit system; };
      });
    in
    {
      packages = forAllSystems ({ pkgs, ... }: rec {
        default = thunderfish;

        thunderfish = pkgs.buildGo121Module {
          name = "thunderfish";
          src = gitignore.lib.gitignoreSource ./.;
          subPackages = [ "cmd/t1" ];
          vendorHash = "sha256-4tHofTnSNI/MBmrGdGsLNoXjxUC0+Gwp3PzzUwfUkQU=";
          CGO_ENABLED = 0;
          flags = [
            "-trimpath"
          ];
          ldflags = [
            "-s"
            "-w"
            "-extldflags -static"
          ];
        };

        thunderf1sh-docs = pkgs.buildNpmPackage {
          name = "thunderf1sh-docs";
          src = gitignore.lib.gitignoreSource ./docs;
          npmDepsHash = "sha256-i6clvSyHtQEGl2C/wcCXonl1W/Kxq7WPTYH46AhUvDM=";
          installPhase = ''
            mkdir -p $out/share
            cp -r build/ $out/share/docs
          '';
        };
      });

      # `nix develop` provides a shell containing development tools.
      devShell = forAllSystems ({ system, pkgs }:
        pkgs.mkShell {
          buildInputs = with pkgs; [
            (golangci-lint.override { buildGoModule = buildGo121Module; })
            go_1_21
            goreleaser
            nodejs
            xc.packages.${system}.xc
          ];
        });

      # This flake outputs an overlay that can be used to add t1 and
      # thunderf1sh-docs to nixpkgs as per https://senforsce.com/thunderf1sh/quick-start/installation/#nix
      #
      # Example usage:
      #
      # nixpkgs.overlays = [
      #   inputs.thunderf1sh.overlays.default
      # ];
      overlays.default = final: prev: {
        thunderf1sh = self.packages.${final.stdenv.system}.thunderf1sh;
        thunderf1sh-docs = self.packages.${final.stdenv.system}.thunderf1sh-docs;
      };
    };
}

