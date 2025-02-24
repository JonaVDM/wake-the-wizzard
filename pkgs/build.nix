{
  pkgs,
  version,
  frontend,
  ...
}: let
in
  pkgs.buildGo124Module {
    pname = "wol";
    inherit version;
    src = ./..;

    vendorHash = "sha256-5JVm6JdpcqRLlOyx2uuwbkMtjSySWk3TiEq4dRTwuYI=";

    preBuild = ''
      cp -r ${frontend}/dist frontend
    '';
  }
