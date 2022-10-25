{ pkgs ? import <nixpkgs> {} }:
  pkgs.stdenv.mkDerivation {
  name = "go-shell";

  buildInputs = with pkgs; [ 
    go
    golangci-lint
  ];

  shellHook = ''
    export GOPATH=`pwd`
    export PATH=$GOPATH/bin:$PATH
  '';
}
