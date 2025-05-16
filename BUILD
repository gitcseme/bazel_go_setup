load("@rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "client",
    srcs = ["cmd/driver/main.go"],
    visibility = ["//visibility:public"],
    deps = ["//internal/evm:evm_lib"]
)
