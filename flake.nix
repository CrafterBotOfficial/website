{
    description = "My personal website";
    inputs = {
        nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05";
        flake-utils.url = "github:numtide/flake-utils";
    };

    outputs = { self, nixpkgs, flake-utils, }:
        flake-utils.lib.eachSystem [ "x86_64-linux" "aarch64-linux" ] (system:
            let
                pkgs = import nixpkgs {
                    inherit system;
                    config.allowUnfree = true;
                };
                website = pkgs.buildGoModule {
                    pname = "website";
                    version = "0.1.9";
                    src = ./.;
                    vendorHash = "sha256-rcCF98dA8RR8vPeS0ivNgVfyaOKMnsC3XbhkdcSfb3w=";
                };
                dockerImage = pkgs.dockerTools.buildImage {
                    name = "website";
                    tag = "latest";
                    contents = [ website ];
                    copyToRoot = [ ./config.json ];
                    config = {
                        Cmd = [ "/bin/website" ];
                        ExposedPorts = { "8080/tcp" = {}; };
                        Env = [
                            "PORT=8080"
                        ];
                    };
                };
            in
            {
                packages.default = website;
                packages.dockerImage = dockerImage;
                devShells.default = pkgs.mkShell {
                    buildInputs = [ 
                    ];
                };
            });
}
