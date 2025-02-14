{pkgs}:
with pkgs; let
  version = "v0.1";
  pnpmHash = "sha256-H18wD/6VJ1bjkNrrS8GjE0ci0qJtUxQ32P7JBNHXCr8=";
  goHash = "sha256-5JVm6JdpcqRLlOyx2uuwbkMtjSySWk3TiEq4dRTwuYI=";

  frontend = stdenv.mkDerivation (finalAttrs: {
    pname = "wol-frontend";
    inherit version;

    src = ./frontend;
    pnpmDeps = pnpm.fetchDeps {
      inherit
        (finalAttrs)
        pname
        version
        src
        ;
      hash = pnpmHash;
    };

    nativeBuildInputs = [
      nodejs
      pnpm.configHook
    ];

    buildInputs = [
      nodejs
    ];

    buildPhase = ''
      runHook preBuild

      pnpm build

      runHook postBuild
    '';

    installPhase = ''
      runHook preInstall

      mkdir -p $out
      cp -r dist $out/dist

      runHook postInstall
    '';
  });
in
  buildGo124Module {
    name = "wol";
    inherit version;
    src = ./.;
    vendorHash = goHash;

    preBuild = ''
      cp -r ${frontend}/dist frontend/dist
    '';
  }
