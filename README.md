# Table of Contents

- [Overview](#overview)
- [Installing](#installing)
- [Usage](#usage)

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

# Usage
After installing, the command can be used immediately by using `kt`. The
main command and all sub-commands have help flags available to help get started.
```
Usage:
  kt [command]

Available Commands:
  all         top all pods across all namespaces
  help        Help about any command
  node        top all pods in a node
  pod         top a specific pod
  ```