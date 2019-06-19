#!/usr/bin/env python3
import os
import sys


args = sys.argv


if "clean" in args:
    os.remove("/home/brice/dev/go-ws/src/byaigo/byaigo")

if "test" in args:
    os.system("go test ~/dev/go-ws/src/byaigo/evaluator/")
    os.system("go test ~/dev/go-ws/src/byaigo/parser/")
    os.system("go test ~/dev/go-ws/src/byaigo/ast/")
    os.system("go test ~/dev/go-ws/src/byaigo/lexer/")
