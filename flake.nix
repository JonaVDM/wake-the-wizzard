{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    nixpkgs,
    flake-utils,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {inherit system;};
    in {
      packages.default = import ./build.nix {inherit pkgs;};

      devShells.default = pkgs.mkShell {
        packages = with pkgs; [
          go
          cobra-cli
          pnpm
          nodejs
        ];
      };
    });
}
