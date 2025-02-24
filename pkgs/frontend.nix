{
  pkgs,
  version,
}:
with pkgs; let
  pname = "wol-frontend";
  src = ../frontend;
in
  stdenv.mkDerivation (finalAttrs: {
    inherit src version pname;

    pnpmDeps = pnpm_10.fetchDeps {
      inherit
        pname
        version
        src
        ;
      hash = "sha256-H18wD/6VJ1bjkNrrS8GjE0ci0qJtUxQ32P7JBNHXCr8=";
    };

    nativeBuildInputs = [
      nodejs_22
      pnpm_10.configHook
    ];

    buildInputs = [
      nodejs_22
    ];

    buildPhase = ''
      runHook preBuild
      pnpm build
      runHook postBuild
    '';

    installPhase = ''
      runHook preInstall

      mkdir -p $out/
      cp -r dist/ $out

      runHook postInstall
    '';
  })
