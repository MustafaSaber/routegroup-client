// +build tools

package main

import (
        _ "k8s.io/code-generator/cmd/client-gen"
        _ "k8s.io/code-generator/cmd/deepcopy-gen"
)
