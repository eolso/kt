# Table of Contents

- [Overview](#overview)
- [Installing](#installing)

# Overview
kt is a small utility that extends the functionality of `kubectl top`.

kt adds the following functionality:
* top all of the pods on a node
* top a pod regardless of current namespace
* top all pods across all namespaces
* sort output by name, cpu usage, or memory usage

# Installing
Use `go get` to install the latest version of the library. This command will
install the `kt` executable along with the library and its dependencies:

    go get -u github.com/ericolsonnv/kt

