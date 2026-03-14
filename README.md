# fitter-happier

_Fitter, happier, more productive, A pig in a cage on antibiotics._

A collection of small, standalone CLI utilities written in pure Go. Each tool is designed to solve a specific problem without external dependencies, making them easy to build and use across different platforms.

## Tools

### `treeview`

Displays a directory tree structure, excluding common development folders.

**Usage:**
```bash
treeview [path]
```

- `path` — directory to display (leave empty for current directory)
- Excludes clutter folders like `node_modules`, `vendor`, `.git`, etc.

---

### `genkey`

Generates a cryptographically secure random key, encoded as base64. Equivalent to `openssl rand -base64`.

**Usage:**
```bash
genkey [bytes]
```

- `bytes` — number of random bytes to generate (default: `32`)
- Output is a base64-encoded string printed to stdout

**Examples:**
```bash
genkey        # 32 bytes default → 44-char base64 string
genkey 64     # 64 bytes → 88-char base64 string
```

---

## Building

Each tool is a self-contained binary with no external dependencies. Build any of them with:

```bash
go build -o <name> ./<name>
```

Or install directly to your Go bin:

```bash
go install ./<name>
```
  
## Contact

Feel free to reach out to discuss collaboration opportunities or to say hello!

- [**My Email**](mailto:matheus.felipe.19rt@gmail.com)
- [**My LinkedIn Profile**](https://www.linkedin.com/in/matheus-mortari-19rt)
- [**My GitHub Profile**](https://github.com/matimortari)

## License

This project is licensed under the [**MIT License**](./LICENSE).