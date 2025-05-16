load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_go//proto:def.bzl", "go_proto_library")

proto_library(
	name = "evm_proto",
	srcs = ["evm.proto"],
	visibility = ["//visibility:public"]
)

go_proto_library(
	name = "evm_go_proto",
	protos = [":evm_proto"],
	importpath = "root/proto",
	visibility = ["//visibility:public"],
)
