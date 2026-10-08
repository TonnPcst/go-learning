# Lesson 1 — Go and the Go Toolchain

> Phase 1 · Go Fundamentals
> Code: [`lesson01/`](../lesson01/)
> Go version: 1.27.1

---

## 1. What Go is

A compiled, statically typed language created at Google (2009). Design goals:

- **Simplicity over expressiveness.** About 25 keywords. No classes, inheritance, decorators or exceptions.
- **Fast compilation.** Large projects build in seconds.
- **Built-in concurrency.** Goroutines and channels are part of the language.
- **One way to do things.** `gofmt` has no options, so there are no style debates.

**Mindset shift from NestJS:** Go deliberately leaves out "magic". There's no DI container, no decorators and no reflection-heavy framework. You write a bit more code, and in exchange any Go codebase shows exactly what it does.

## 2. Why Go is popular for backend work

- **One static binary.** Deploying is copying a file. No `node_modules`, no runtime to install.
- **Real parallelism.** One process uses every CPU core.
- **Low memory and fast startup.** Good for containers, Kubernetes and serverless.
- **A strong standard library:** `net/http`, `encoding/json`, `database/sql`, `context`, `testing`, `log/slog`.
- **Cloud-native tools are written in Go:** Docker, Kubernetes, Terraform, Prometheus, etcd.

## 3. Go vs Node.js / TypeScript

| Concern | Node.js / TypeScript | Go |
|---|---|---|
| Execution | JIT-compiled JS on V8 | Ahead-of-time compiled machine code |
| Types | Erased at runtime; `any` escape hatch | Enforced at compile time, present at runtime |
| Concurrency | Single-threaded event loop + `async/await` | Goroutines on all cores; blocking code is normal |
| Errors | `throw` / `try/catch` | Errors are return values: `val, err := f()` |
| OOP | Classes, inheritance, decorators | Structs + methods + implicit interfaces |
| DI | NestJS container, `@Injectable()` | Pass dependencies into constructors yourself |
| Packages | npm, `package.json`, `node_modules` | Go modules, `go.mod`, global module cache |
| Format / lint | Prettier + ESLint configs | `gofmt` + `go vet`, built in |
| Testing | Jest / Vitest (third-party) | `go test` + `testing` package (built in) |
| Exports | `export` keyword | **Capitalized = exported**, lowercase = package-private |

**The big one:** Node must never block the event loop, so everything is `async`. In Go you write plain blocking code (`resp, err := http.Get(url)`); the runtime parks that goroutine and runs others. There's no function colouring, so `async` doesn't spread through the codebase.

## 4. Go compiler vs Node.js runtime

```text
Node:  app.ts ──tsc──► app.js ──(ship with node + node_modules)──► V8 runs/JITs at runtime
Go:    *.go   ──go build──► ./app   (native machine code, self-contained)
```

Node 22.18+/23.6+ can also run `.ts` directly by stripping types, but it's still interpreted and JIT-compiled at runtime.

The Go binary **contains** a small runtime: the garbage collector and the goroutine scheduler. Nothing needs to be installed on the target machine.

Cross-compiling for a Linux server from a Mac:

```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o app ./cmd/api
```

`CGO_ENABLED=0` produces a fully static binary with no dependency on the system's C library. Cross-compiling already disables cgo by default, but a native build inside a Docker builder image does not, so set it explicitly.

## 5. Project structure: module → package → file

- **Module:** a versioned unit defined by `go.mod` (roughly a `package.json` project).
- **Package:** **one directory = one package.** Every file in it declares the same `package xyz`.
- **File:** files in the same package see each other's identifiers without importing.

```text
lesson01/
├── go.mod
├── cmd/
│   └── greeter/
│       └── main.go          ← package main: executable entry point
└── internal/
    └── greet/
        ├── greet.go         ← package greet
        └── greet_test.go    ← tests live next to the code
```

Conventions:

- **`package main` + `func main()`** marks an executable. Any other package name is a library.
- **`cmd/<name>/`** holds one directory per binary (later: `cmd/api`, `cmd/worker`).
- **`internal/`** is enforced by the compiler: only code under the parent of `internal/` can import it. Here the parent is the module root, so other modules can't import `lesson01/internal/greet`.
- Group packages by **what they do** (`user`, `charger`, `postgres`), not by technical layer (`controllers/`, `services/`, `dtos/`). Don't copy the NestJS layout.
- **Import path = module path + folder path:** `example.com/lesson01` + `/internal/greet`. The **package name** is the `package` line, and it should match the folder's last part (`greet`). In code you write `greet.Hello`.

## 6. `go.mod` and `go.sum`

```bash
go mod init example.com/lesson01
```

```text
module example.com/lesson01   // import prefix for every package in the module

go 1.27.1                     // minimum Go version
```

After `go get github.com/jackc/pgx/v5`:

```go
require github.com/jackc/pgx/v5 v5.7.1
```

| File | Purpose | npm analogy |
|---|---|---|
| `go.mod` | Module path, Go version, **minimum** version of each dependency (no `^` / `~` ranges) | `package.json` |
| `go.sum` | Cryptographic checksums of downloaded modules; detects tampered or changed code | Nothing exact (closest: the integrity hashes in `package-lock.json`) |

- Go uses **Minimal Version Selection**: it builds with the minimum versions listed in `go.mod`, so builds are reproducible without a lockfile.
- **Commit both files.**
- `go mod tidy` adds missing dependencies and removes unused ones.
- Imports always use the full module path (`"example.com/lesson01/internal/greet"`). There are no relative imports.

## 7. Commands

```bash
go run ./cmd/greeter                  # compile to a temp dir and run
go build -o bin/greeter ./cmd/greeter # produce a native binary
go test ./...                         # run tests in every package (./... = recursive)
go test -v ./...                      # verbose: list each test
go test -run 'TestHello' ./...        # run only tests matching a regex
go test -cover ./...                  # show coverage percentage
go vet ./...                          # static checks for suspicious code
gofmt -l .                            # list files that aren't formatted
go mod tidy                           # sync go.mod/go.sum with the imports
go doc ./internal/greet               # show a package's doc comments
go env GOOS GOARCH                    # show the target OS/arch (darwin arm64 on your Mac)
go version                            # show the installed Go version
```

Set your editor to run **gopls / goimports on save**. It adds and removes imports automatically, so you'll rarely see the unused-import error.

## Syntax notes

- `func Add(a, b int) int`: the type comes **after** the name; `a, b int` shares one type.
- **Unused imports and unused local variables are compile errors**, not warnings.
  - A side-effect-only import uses the blank identifier: `import _ "github.com/lib/pq"`.
- No semicolons and no parentheses around `if` conditions. Braces are required, and `{` stays on the same line.
- Tests: file `*_test.go`, function `TestXxx(t *testing.T)`. Use a plain `if` plus `t.Errorf`; there's no `expect()`.
  - `t.Errorf` marks the test failed and keeps going; `t.Fatalf` stops the test immediately.
  - Use `%q` in failure messages so empty strings and whitespace are visible.
- Doc comments start with the identifier name (`// Hello returns ...`) and are prose, not JSDoc `@param` blocks. Cover what it returns and any special cases. A blank `//` line starts a new paragraph; an indented line becomes a code block.

---

## Exercise 1: greeter

**Task:** module `example.com/lesson01`, package `internal/greet` with an exported `Hello(name string) string` (`"Hello, <name>!"`, or `"Hello, stranger!"` for an empty name) and an unexported helper, `cmd/greeter/main.go`, and two tests.

**Solution:** see [`lesson01/internal/greet/greet.go`](../lesson01/internal/greet/greet.go). `Hello` delegates to the unexported `displayName`, which trims whitespace and falls back to `"stranger"`.

**Result** (original run with the first two tests; two whitespace tests were added after review, see Review notes)

```text
$ go run ./cmd/greeter
Hello, Ton!
Hello, stranger!

$ go test -v ./...
?       example.com/lesson01/cmd/greeter        [no test files]
=== RUN   TestHelloWithName
--- PASS: TestHelloWithName (0.00s)
=== RUN   TestHelloEmptyName
--- PASS: TestHelloEmptyName (0.00s)
PASS
ok      example.com/lesson01/internal/greet
```

### Questions and answers

**A. What happens if `main.go` calls the unexported helper?**
It doesn't compile: `name displayName not exported by package greet`. A lowercase identifier is visible only inside its own package, and `main` is a different package. Go has no `export` or `private` keyword: the first letter of the name decides, and the compiler enforces it.

**B. How do you build for a Linux amd64 Kubernetes node? Does the node need Go?**
`GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o greeter ./cmd/greeter`. The node doesn't need Go. The binary is self-contained and includes the runtime, and `CGO_ENABLED=0` removes the dependency on the system's C library.

**C. What's the difference between `go.mod` and `go.sum`? Which one is `package-lock.json`?**
`go.mod` declares the module path, the Go version and the minimum version of each dependency. `go.sum` stores checksums of downloaded modules so changed code is detected. Neither is a true `package-lock.json`. `go.sum` is the closest, as the generated file you commit, but Go doesn't need a lockfile: Minimal Version Selection always uses the minimums in `go.mod`, so builds are already reproducible. Commit both.

**D. What happens with an unused `fmt` import? How does that differ from TypeScript?**
It's a compile error (`"fmt" imported and not used`), and the program won't build. In TypeScript an unused import is only an editor or linter warning (an error only with `noUnusedLocals`), and the code still compiles and runs.

### Common mistakes I hit

| Error | Cause | Fix |
|---|---|---|
| `found packages greeter (greet.go) and main (main.go)` | Two package names in one folder | One folder = one package |
| `import cycle not allowed` | `main.go` imported its own folder | Never import your own package; files in the same package see each other directly |
| Editor added an alias `greeter "…/internal/greet"` | `package greeter` inside folder `greet/` | Package name should match the folder's last part |

### Review notes

- ✅ Correct layout (`cmd/` + `internal/`). Doc comments start with the identifier name and are written in English and Lao (ENG/LAO paragraphs, separated by a blank `//` line).
- ✅ `%q` in test failure messages.
- ✅ Reassigning the parameter (`name = strings.TrimSpace(name)`) is idiomatic: parameters are local variables, so this doesn't affect the caller.
- ✅ The extra `TrimSpace` behaviour (`"   "` becomes stranger, `"  Ton  "` becomes `"Ton"`) is now tested (`TestHelloWhitespaceName`, `TestHelloTrimsName`). Every branch you write needs a test.
- ✅ Removed the `Parameters:` / `Returns` blocks (a JSDoc habit). Go doc comments are prose.
- ⚠️ The four tests each repeat the same `got` / `want` / `if` block. Go's fix is the **table-driven test** (Exercise 1.2).

---

## Exercise 1.2: table-driven tests + CLI args *(in progress)*

1. Replace the four tests with one table-driven `TestHello` using `t.Run` subtests. Cover `"Ton"`, `""`, `"   "` and `"  Ton  "`.
2. `go run ./cmd/greeter Alice` greets Alice; with no argument it greets the stranger. Use `os.Args` and check the length before indexing.
3. Run:
   - `go test -v ./...`
   - `go test -run 'TestHello/<empty-case-name>' -v ./internal/greet`
   - `go test -cover ./...`
   - `GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/greeter-linux ./cmd/greeter && file bin/greeter-linux`
   - Add `bin/` to `.gitignore` first, so build output never gets committed.
4. What is `os.Args[0]`, and what happens if you access `os.Args[1]` without checking the length? How does that compare with `process.argv[2]` in Node?
