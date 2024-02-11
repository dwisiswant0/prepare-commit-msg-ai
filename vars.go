package main

import _ "embed"
import "path/filepath"

//go:embed scripts/prepare-commit-msg
var prepareCommitMsgByte []byte

var hookPath = filepath.Join(".git", "hooks", "prepare-commit-msg")
