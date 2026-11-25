# fitter-happier

_Fitter, happier, more productive, A pig in a cage on antibiotics._

A collection of small, standalone CLI utilities written in pure Go. Each tool is designed to solve a specific problem without external dependencies, making them easy to build and use across different platforms.

## Build and Install

Each tool is its own Go module. From the repo root:

```bash
go -C <name> build -o bin/<name> .
go -C <name> install .
```

`go install` puts the binary in `$(go env GOPATH)/bin` (usually `~/go/bin`) — keep that on your `PATH`.

## Tools

### `treeview`

Displays a directory tree, excluding common development folders by default.

**Usage:**

```bash
treeview [path] [flags]
```

- `path` — directory to display (default: current directory)
- `-exclude` — comma-separated names to exclude (adds to defaults)
- `-no-default-exclude` — disable default exclusions

**Examples:**

```bash
treeview
treeview . -exclude "coverage,tmp"
treeview . -no-default-exclude -exclude ".git,node_modules"
```

---

### `genkey`

Generates cryptographically secure random bytes as Base64. Equivalent to `openssl rand -base64`.

**Usage:**

```bash
genkey [bytes]
```

- `bytes` — number of random bytes (default: `32`)

**Examples:**

```bash
genkey      # 32 bytes → 44-char base64 string
genkey 64   # 64 bytes → 88-char base64 string
```

---

### `timer`

Stopwatch (no args) or countdown (with a duration). Updates in place; Ctrl+C stops.

**Usage:**

```bash
timer [duration]
```

- no args — stopwatch until Ctrl+C
- `duration` — Go duration (`10s`, `5m`, `1h30m`) or bare seconds (`90`)

**Examples:**

```bash
timer        # stopwatch
timer 25s    # countdown 25 seconds
timer 5m     # countdown 5 minutes
timer 90     # countdown 90 seconds
```

---

### `filesize`

Prints the size of files or directories (directories are summed recursively). Human-readable by default; does not follow symlinks. A single path prints only the size; multiple paths print `path` then size.

**Usage:**

```bash
filesize [path...] [flags]
```

- `path` — file or directory (default: current directory)
- `-bytes` — print size in bytes instead of human-readable units

## Contact

Feel free to reach out to discuss collaboration opportunities or to say hello!

- [**My Email**](mailto:matheus.felipe.19rt@gmail.com)
- [**My LinkedIn Profile**](https://www.linkedin.com/in/matheus-mortari-19rt)
- [**My GitHub Profile**](https://github.com/matimortari)

## License

This project is licensed under the [**MIT License**](./LICENSE).
