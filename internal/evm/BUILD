load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
	name = "evm_lib",
	srcs = ["evm.go"],
	importpath = "root/internal/evm",
	visibility = ["//:__subpackages__"],
)

go_test(
	name = "evm_test",
	srcs = ["evm_test.go"],
	deps = [":evm_lib"],
)