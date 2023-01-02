load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "bazel_skylib",
    sha256 = "74d544d96f4a5bb630d465ca8bbcfe231e3594e5aae57e1edbf17a6eb3ca2506",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.3.0/bazel-skylib-1.3.0.tar.gz",
        "https://github.com/bazelbuild/bazel-skylib/releases/download/1.3.0/bazel-skylib-1.3.0.tar.gz",
    ],
)

load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")

bazel_skylib_workspace()

http_archive(
    name = "aspect_rules_js",
    sha256 = "c4a5766a45dff25b2eb1789d7a2decfda81b281fc88350d24687620c37fefb25",
    strip_prefix = "rules_js-1.14.0",
    url = "https://github.com/aspect-build/rules_js/archive/refs/tags/v1.14.0.tar.gz",
)

http_archive(
    name = "rules_nodejs",
    sha256 = "08337d4fffc78f7fe648a93be12ea2fc4e8eb9795a4e6aa48595b66b34555626",
    urls = ["https://github.com/bazelbuild/rules_nodejs/releases/download/5.8.0/rules_nodejs-core-5.8.0.tar.gz"],
)

http_archive(
    name = "aspect_rules_ts",
    sha256 = "6406905c5f7c5ca6dedcca5dacbffbf32bb2a5deb77f50da73e7195b2b7e8cbc",
    strip_prefix = "rules_ts-1.0.5",
    url = "https://github.com/aspect-build/rules_ts/archive/refs/tags/v1.0.5.tar.gz",
)

http_archive(
    name = "aspect_rules_jest",
    sha256 = "300e2a75285bb47560d6bf743c87e38c3c0a139e0b988a6fff8719c155dc3780",
    strip_prefix = "rules_jest-0.14.2",
    url = "https://github.com/aspect-build/rules_jest/archive/refs/tags/v0.14.2.tar.gz",
)

# Node toolchain setup ==========================
load("@rules_nodejs//nodejs:repositories.bzl", "DEFAULT_NODE_VERSION", "nodejs_register_toolchains")

nodejs_register_toolchains(
    name = "nodejs",
    node_version = DEFAULT_NODE_VERSION,
)

# rules_js setup ================================
load("@aspect_rules_js//js:repositories.bzl", "rules_js_dependencies")

rules_js_dependencies()

# rules_js npm setup ============================
load("@aspect_rules_js//npm:npm_import.bzl", "npm_translate_lock")

npm_translate_lock(
    name = "npm",
    pnpm_lock = "//:pnpm-lock.yaml",
    npmrc = "//:.npmrc",
    verify_node_modules_ignored = "//:.bazelignore",
)

# rules_ts npm setup ============================
load("@npm//:repositories.bzl", "npm_repositories")

npm_repositories()

load("@aspect_rules_ts//ts:repositories.bzl", "rules_ts_dependencies", LATEST_TS_VERSION = "LATEST_VERSION")

rules_ts_dependencies(ts_version = LATEST_TS_VERSION)

# rules_jest setup ==============================
load("@aspect_rules_jest//jest:dependencies.bzl", "rules_jest_dependencies")

rules_jest_dependencies()

load("@aspect_rules_jest//jest:repositories.bzl", "jest_repositories")

jest_repositories(
    name = "jest",
    jest_version = "v28.1.0",
)

load("@jest//:npm_repositories.bzl", jest_npm_repositories = "npm_repositories")

jest_npm_repositories()
