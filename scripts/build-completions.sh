#!/bin/bash

set -euo pipefail

rm -rf completions
mkdir completions

go build .

./clive completion bash > completions/clive.bash
./clive completion zsh  > completions/clive.zsh
./clive completion fish > completions/clive.fish
