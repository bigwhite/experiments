#!/usr/bin/env node
const { spawnSync } = require("child_process");
const fs = require("fs");
const path = require("path");

const OS = { darwin: "darwin", linux: "linux", win32: "windows" }[process.platform];
const ARCH = { x64: "amd64", arm64: "arm64" }[process.arch];

if (!OS || !ARCH) {
  console.error(`不支持的平台：${process.platform}-${process.arch}`);
  process.exit(1);
}

const bin = path.join(
  __dirname, "bin", `hello-${OS}-${ARCH}${OS === "windows" ? ".exe" : ""}`
);

try { fs.chmodSync(bin, 0o755); } catch (_) {} // 保证有执行权限

const r = spawnSync(bin, process.argv.slice(2), { stdio: "inherit" });
if (r.error) {
  console.error(r.error.message);
  process.exit(1);
}
process.exit(r.status === null ? 1 : r.status);

