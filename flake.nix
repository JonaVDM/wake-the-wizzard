{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    nixpkgs,
    flake-utils,
    self,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      lastModifiedDate = self.lastModifiedDate or self.lastModified or "19700101";
      version = builtins.substring 0 8 lastModifiedDate;
      pkgs = import nixpkgs {inherit system;};
    in {
      packages = rec {
        frontend = import ./pkgs/frontend.nix {inherit pkgs version;};
        wol = import ./pkgs/build.nix {inherit pkgs version frontend;};
      };
      packages.default = pkgs.buildGo124Module {
        pname = "wol";
        inherit version;
        src = ./.;

        vendorHash = "sha256-5JVm6JdpcqRLlOyx2uuwbkMtjSySWk3TiEq4dRTwuYI=";
      };

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
