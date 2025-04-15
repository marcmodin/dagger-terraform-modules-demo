{
  description = "A flake that adds the latest version of the Dagger Cli via github/dagger/nix";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    dagger.url = "github:dagger/nix";
  };

  outputs = { self, nixpkgs, flake-utils, dagger, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {
        packages.default = dagger.packages.${system}.dagger;

        devShells.default = pkgs.mkShell {
          packages = [ dagger.packages.${system}.dagger ];
        };
      }
    );
}