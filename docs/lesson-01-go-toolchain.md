# ບົດທີ 1 — Go ແລະ Go Toolchain

> Phase 1 · ພື້ນຖານ Go
> ໂຄດ: [`lesson01/`](../lesson01/)
> ເວີຊັນ Go: 1.27.1

---

## 1. Go ແມ່ນຫຍັງ

Go ເປັນພາສາແບບ compiled ແລະ statically typed ທີ່ Google ສ້າງຂຶ້ນ (ປີ 2009). ເປົ້າໝາຍການອອກແບບ:

- **ລຽບງ່າຍ ສຳຄັນກວ່າຂຽນສັ້ນ.** ມີ keyword ພຽງ 25 ຄຳ. ບໍ່ມີ class, inheritance, decorator ຫຼື exception.
- **Compile ໄວ.** ໂປຣເຈັກໃຫຍ່ build ໄດ້ພາຍໃນບໍ່ເທົ່າໃດວິນາທີ.
- **Concurrency ມີມາໃນພາສາເລີຍ.** Goroutine ແລະ channel ເປັນສ່ວນໜຶ່ງຂອງພາສາ.
- **ມີວິທີດຽວໃນການເຮັດ.** `gofmt` ບໍ່ມີ option ໃຫ້ຕັ້ງຄ່າ, ຈຶ່ງບໍ່ຕ້ອງຖຽງກັນເລື່ອງ style.

**ປ່ຽນແນວຄິດຈາກ NestJS:** Go ຕັ້ງໃຈບໍ່ມີ "magic". ບໍ່ມີ DI container, ບໍ່ມີ decorator ແລະ ບໍ່ມີ framework ທີ່ໃຊ້ reflection ຫຼາຍ. ຕ້ອງຂຽນໂຄດຫຼາຍຂຶ້ນໜ້ອຍໜຶ່ງ, ແຕ່ແລກມາກັບໂຄດ Go ທີ່ອ່ານແລ້ວເຫັນເລີຍວ່າມັນເຮັດຫຍັງ.

## 2. ເປັນຫຍັງ Go ຈຶ່ງນິຍົມໃຊ້ເຮັດ backend

- **Static binary ໄຟລ໌ດຽວ.** Deploy ແມ່ນແຄ່ສຳເນົາໄຟລ໌. ບໍ່ມີ `node_modules`, ບໍ່ຕ້ອງຕິດຕັ້ງ runtime.
- **Parallelism ແທ້.** Process ດຽວໃຊ້ໄດ້ທຸກ core ຂອງ CPU.
- **ໃຊ້ໜ່ວຍຄວາມຈຳໜ້ອຍ ແລະ ເປີດໄວ.** ເໝາະກັບ container, Kubernetes ແລະ serverless.
- **Standard library ແຂງແຮງ:** `net/http`, `encoding/json`, `database/sql`, `context`, `testing`, `log/slog`.
- **ເຄື່ອງມື cloud-native ຂຽນດ້ວຍ Go:** Docker, Kubernetes, Terraform, Prometheus, etcd.

## 3. Go ທຽບກັບ Node.js / TypeScript

| ຫົວຂໍ້ | Node.js / TypeScript | Go |
|---|---|---|
| ການ run | JS ທີ່ JIT-compile ເທິງ V8 | Compile ເປັນ machine code ລ່ວງໜ້າ (ahead-of-time) |
| Type | ຖືກລຶບຖິ້ມຕອນ runtime; ມີ `any` ໃຫ້ຫຼົບ | ກວດຕອນ compile ແລະ ຍັງມີຢູ່ຕອນ runtime |
| Concurrency | Event loop thread ດຽວ + `async/await` | Goroutine ເທິງທຸກ core; ຂຽນ blocking code ເປັນເລື່ອງປົກກະຕິ |
| Error | `throw` / `try/catch` | Error ເປັນຄ່າທີ່ return: `val, err := f()` |
| OOP | Class, inheritance, decorator | Struct + method + interface ແບບ implicit |
| DI | NestJS container, `@Injectable()` | ສົ່ງ dependency ເຂົ້າ constructor ເອງ |
| Package | npm, `package.json`, `node_modules` | Go modules, `go.mod`, module cache ລວມ |
| Format / lint | ຕັ້ງຄ່າ Prettier + ESLint | `gofmt` + `go vet` ມີມາໃນຕົວ |
| Testing | Jest / Vitest (third-party) | `go test` + package `testing` (ມີມາໃນຕົວ) |
| Export | Keyword `export` | **ຕົວພິມໃຫຍ່ = export**, ຕົວພິມນ້ອຍ = ໃຊ້ໄດ້ພາຍໃນ package ເທົ່ານັ້ນ |

**ຄວາມແຕກຕ່າງໃຫຍ່ທີ່ສຸດ:** ໃນ Node ຫ້າມ block event loop, ທຸກຢ່າງຈຶ່ງຕ້ອງເປັນ `async`. ໃນ Go ຂຽນ blocking code ທຳມະດາໄດ້ເລີຍ (`resp, err := http.Get(url)`); runtime ຈະພັກ goroutine ນັ້ນໄວ້ ແລ້ວໄປ run goroutine ອື່ນ. ບໍ່ມີບັນຫາ function colouring, ສະນັ້ນ `async` ຈະບໍ່ແຜ່ລາມໄປທົ່ວໂຄດ.

## 4. Go compiler ທຽບກັບ Node.js runtime

```text
Node:  app.ts ──tsc──► app.js ──(ship with node + node_modules)──► V8 runs/JITs at runtime
Go:    *.go   ──go build──► ./app   (native machine code, self-contained)
```

Node 22.18+/23.6+ ສາມາດ run ໄຟລ໌ `.ts` ໄດ້ໂດຍກົງດ້ວຍການລຶບ type ອອກ, ແຕ່ກໍຍັງ interpret ແລະ JIT-compile ຕອນ runtime ຄືເກົ່າ.

Binary ຂອງ Go **ມີ** runtime ນ້ອຍໆຢູ່ໃນຕົວ: garbage collector ແລະ goroutine scheduler. ເຄື່ອງປາຍທາງບໍ່ຕ້ອງຕິດຕັ້ງຫຍັງເລີຍ.

Cross-compile ໃຫ້ server Linux ຈາກ Mac:

```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o app ./cmd/api
```

`CGO_ENABLED=0` ເຮັດໃຫ້ໄດ້ binary ແບບ static ທັງໝົດ, ບໍ່ເພິ່ງ C library ຂອງລະບົບ. ຕອນ cross-compile, cgo ຖືກປິດໂດຍ default ຢູ່ແລ້ວ, ແຕ່ຕອນ build ປົກກະຕິໃນ Docker builder image ມັນບໍ່ຖືກປິດ, ສະນັ້ນໃຫ້ກຳນົດເອງໃຫ້ຊັດເຈນ.

## 5. ໂຄງສ້າງໂປຣເຈັກ: module → package → file

- **Module:** ໜ່ວຍທີ່ມີເວີຊັນ, ກຳນົດໂດຍ `go.mod` (ຄ້າຍກັບໂປຣເຈັກທີ່ມີ `package.json`).
- **Package:** **ໜຶ່ງໂຟນເດີ = ໜຶ່ງ package.** ທຸກໄຟລ໌ໃນໂຟນເດີຕ້ອງປະກາດ `package xyz` ຊື່ດຽວກັນ.
- **File:** ໄຟລ໌ໃນ package ດຽວກັນເຫັນຊື່ຂອງກັນ ແລະ ກັນໄດ້ໂດຍບໍ່ຕ້ອງ import.

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

ທຳນຽມ:

- **`package main` + `func main()`** ໝາຍເຖິງໂປຣແກຣມທີ່ run ໄດ້ (executable). Package ຊື່ອື່ນແມ່ນ library.
- **`cmd/<name>/`** ໜຶ່ງໂຟນເດີຕໍ່ໜຶ່ງ binary (ຕໍ່ໄປ: `cmd/api`, `cmd/worker`).
- **`internal/`** compiler ບັງຄັບໃຊ້ແທ້: ມີແຕ່ໂຄດທີ່ຢູ່ພາຍໃຕ້ໂຟນເດີແມ່ຂອງ `internal/` ເທົ່ານັ້ນທີ່ import ໄດ້. ໃນນີ້ໂຟນເດີແມ່ຄື root ຂອງ module, ສະນັ້ນ module ອື່ນ import `lesson01/internal/greet` ບໍ່ໄດ້.
- ຈັດກຸ່ມ package ຕາມ **ໜ້າທີ່ຂອງມັນ** (`user`, `charger`, `postgres`), ບໍ່ແມ່ນຕາມ layer ທາງເທັກນິກ (`controllers/`, `services/`, `dtos/`). ຢ່າກັອບປີ້ layout ຂອງ NestJS.
- **Import path = module path + folder path:** `example.com/lesson01` + `/internal/greet`. **ຊື່ package** ແມ່ນຊື່ໃນແຖວ `package`, ແລະ ຄວນກົງກັບສ່ວນສຸດທ້າຍຂອງຊື່ໂຟນເດີ (`greet`). ໃນໂຄດຈະຂຽນເປັນ `greet.Hello`.

## 6. `go.mod` ແລະ `go.sum`

```bash
go mod init example.com/lesson01
```

```text
module example.com/lesson01   // import prefix for every package in the module

go 1.27.1                     // minimum Go version
```

ຫຼັງຈາກ `go get github.com/jackc/pgx/v5`:

```go
require github.com/jackc/pgx/v5 v5.7.1
```

| ໄຟລ໌ | ໜ້າທີ່ | ທຽບກັບ npm |
|---|---|---|
| `go.mod` | Module path, ເວີຊັນ Go, **ເວີຊັນຕ່ຳສຸດ** ຂອງແຕ່ລະ dependency (ບໍ່ມີ range ແບບ `^` / `~`) | `package.json` |
| `go.sum` | Checksum ຂອງ module ທີ່ດາວໂຫຼດມາ; ກວດຈັບໂຄດທີ່ຖືກແກ້ໄຂ ຫຼື ປ່ຽນແປງ | ບໍ່ມີອັນທີ່ກົງແທ້ (ໃກ້ທີ່ສຸດ: integrity hash ໃນ `package-lock.json`) |

- Go ໃຊ້ **Minimal Version Selection**: build ດ້ວຍເວີຊັນຕ່ຳສຸດທີ່ຂຽນໄວ້ໃນ `go.mod` ສະເໝີ, ສະນັ້ນ build ຊ້ຳໄດ້ຜົນຄືເກົ່າໂດຍບໍ່ຕ້ອງມີ lockfile.
- **Commit ທັງສອງໄຟລ໌.**
- `go mod tidy` ເພີ່ມ dependency ທີ່ຂາດ ແລະ ລຶບອັນທີ່ບໍ່ໃຊ້ອອກ.
- ການ import ໃຊ້ module path ເຕັມສະເໝີ (`"example.com/lesson01/internal/greet"`). ບໍ່ມີ relative import.

## 7. ຄຳສັ່ງ

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

ຕັ້ງ editor ໃຫ້ run **gopls / goimports ຕອນ save**. ມັນຈະເພີ່ມ ແລະ ລຶບ import ໃຫ້ອັດຕະໂນມັດ, ຈຶ່ງບໍ່ຄ່ອຍເຈີ error ເລື່ອງ import ທີ່ບໍ່ໄດ້ໃຊ້.

## ບັນທຶກ syntax

- `func Add(a, b int) int`: type ຢູ່ **ຫຼັງ** ຊື່; `a, b int` ໃຊ້ type ດຽວກັນ.
- **Import ທີ່ບໍ່ໄດ້ໃຊ້ ແລະ local variable ທີ່ບໍ່ໄດ້ໃຊ້ ເປັນ compile error**, ບໍ່ແມ່ນ warning.
  - Import ທີ່ຕ້ອງການພຽງ side effect ໃຫ້ໃຊ້ blank identifier: `import _ "github.com/lib/pq"`.
- ບໍ່ມີ semicolon ແລະ ບໍ່ມີວົງເລັບອ້ອມເງື່ອນໄຂ `if`. ຕ້ອງມີປີກກາ (braces) ສະເໝີ, ແລະ `{` ຕ້ອງຢູ່ແຖວດຽວກັນ.
- Test: ໄຟລ໌ `*_test.go`, ຟັງຊັນ `TestXxx(t *testing.T)`. ໃຊ້ `if` ທຳມະດາ + `t.Errorf`; ບໍ່ມີ `expect()`.
  - `t.Errorf` ໝາຍວ່າ test fail ແຕ່ຍັງເຮັດວຽກຕໍ່; `t.Fatalf` ຢຸດ test ທັນທີ.
  - ໃຊ້ `%q` ໃນຂໍ້ຄວາມ failure ເພື່ອໃຫ້ເຫັນ string ຫວ່າງ ແລະ ຊ່ອງວ່າງ.
- Doc comment ເລີ່ມດ້ວຍຊື່ຂອງສິ່ງທີ່ອະທິບາຍ (`// Hello returns ...`) ແລະ ຂຽນເປັນປະໂຫຍກທຳມະດາ, ບໍ່ແມ່ນ block `@param` ແບບ JSDoc. ບອກວ່າມັນ return ຫຍັງ ແລະ ກໍລະນີພິເສດ. ແຖວ `//` ຫວ່າງ = ຂຶ້ນຫຍໍ້ໜ້າໃໝ່; ແຖວທີ່ຫຍໍ້ເຂົ້າ = code block.

---

## ແບບຝຶກຫັດ 1: greeter

**ໂຈດ:** module `example.com/lesson01`, package `internal/greet` ທີ່ມີ `Hello(name string) string` ແບບ export (`"Hello, <name>!"`, ຫຼື `"Hello, stranger!"` ຖ້າຊື່ຫວ່າງ) ແລະ helper ທີ່ບໍ່ export ໜຶ່ງຕົວ, `cmd/greeter/main.go`, ແລະ test ສອງຕົວ.

**ວິທີແກ້:** ເບິ່ງ [`lesson01/internal/greet/greet.go`](../lesson01/internal/greet/greet.go). `Hello` ສົ່ງຕໍ່ໃຫ້ `displayName` (ບໍ່ export), ເຊິ່ງຕັດຊ່ອງວ່າງອອກ ແລະ ໃຊ້ `"stranger"` ແທນເມື່ອບໍ່ມີຊື່.

**ຜົນລັບ** (ການ run ຄັ້ງທຳອິດ ທີ່ມີ test ສອງຕົວທຳອິດ; test ເລື່ອງຊ່ອງວ່າງສອງຕົວຖືກເພີ່ມຫຼັງການ review, ເບິ່ງ ບັນທຶກການ review)

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

### ຄຳຖາມ ແລະ ຄຳຕອບ

**A. ຈະເກີດຫຍັງຂຶ້ນ ຖ້າ `main.go` ເອີ້ນ helper ທີ່ບໍ່ export?**
Compile ບໍ່ຜ່ານ: `name displayName not exported by package greet`. ຊື່ທີ່ຂຶ້ນຕົ້ນດ້ວຍຕົວພິມນ້ອຍ ເຫັນໄດ້ສະເພາະພາຍໃນ package ຂອງມັນເອງ, ແລະ `main` ເປັນຄົນລະ package. Go ບໍ່ມີ keyword `export` ຫຼື `private`: ຕົວອັກສອນຕົວທຳອິດຂອງຊື່ເປັນຕົວຕັດສິນ, ແລະ compiler ເປັນຜູ້ບັງຄັບ.

**B. Build ໃຫ້ Kubernetes node ທີ່ເປັນ Linux amd64 ແນວໃດ? Node ຕ້ອງມີ Go ບໍ?**
`GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o greeter ./cmd/greeter`. Node ບໍ່ຕ້ອງມີ Go. Binary ມີທຸກຢ່າງໃນຕົວລວມທັງ runtime, ແລະ `CGO_ENABLED=0` ຕັດການເພິ່ງ C library ຂອງລະບົບອອກ.

**C. `go.mod` ກັບ `go.sum` ຕ່າງກັນແນວໃດ? ອັນໃດຄື `package-lock.json`?**
`go.mod` ປະກາດ module path, ເວີຊັນ Go ແລະ ເວີຊັນຕ່ຳສຸດຂອງແຕ່ລະ dependency. `go.sum` ເກັບ checksum ຂອງ module ທີ່ດາວໂຫຼດມາ ເພື່ອກວດຈັບໂຄດທີ່ຖືກປ່ຽນ. ບໍ່ມີອັນໃດເປັນ `package-lock.json` ແທ້. `go.sum` ໃກ້ທີ່ສຸດ ໃນແງ່ທີ່ເປັນໄຟລ໌ທີ່ລະບົບສ້າງໃຫ້ ແລະ ຕ້ອງ commit, ແຕ່ Go ບໍ່ຕ້ອງມີ lockfile: Minimal Version Selection ໃຊ້ເວີຊັນຕ່ຳສຸດໃນ `go.mod` ສະເໝີ, ສະນັ້ນ build ຊ້ຳໄດ້ຜົນຄືເກົ່າຢູ່ແລ້ວ. Commit ທັງສອງໄຟລ໌.

**D. ຈະເກີດຫຍັງຂຶ້ນກັບ import `fmt` ທີ່ບໍ່ໄດ້ໃຊ້? ຕ່າງຈາກ TypeScript ແນວໃດ?**
ເປັນ compile error (`"fmt" imported and not used`), ແລະ ໂປຣແກຣມ build ບໍ່ໄດ້. ໃນ TypeScript import ທີ່ບໍ່ໄດ້ໃຊ້ເປັນພຽງ warning ຈາກ editor ຫຼື linter (ເປັນ error ກໍຕໍ່ເມື່ອເປີດ `noUnusedLocals`), ແລະ ໂຄດຍັງ compile ແລະ run ໄດ້.

### ຂໍ້ຜິດພາດທີ່ເຄີຍເຈີ

| Error | ສາເຫດ | ວິທີແກ້ |
|---|---|---|
| `found packages greeter (greet.go) and main (main.go)` | ມີສອງຊື່ package ໃນໂຟນເດີດຽວ | ໜຶ່ງໂຟນເດີ = ໜຶ່ງ package |
| `import cycle not allowed` | `main.go` import ໂຟນເດີຂອງຕົວເອງ | ຢ່າ import package ຂອງຕົວເອງ; ໄຟລ໌ໃນ package ດຽວກັນເຫັນກັນໄດ້ເລີຍ |
| Editor ເພີ່ມ alias `greeter "…/internal/greet"` ໃຫ້ | `package greeter` ຢູ່ໃນໂຟນເດີ `greet/` | ຊື່ package ຄວນກົງກັບສ່ວນສຸດທ້າຍຂອງຊື່ໂຟນເດີ |
| `docs/` ບໍ່ມີຢູ່ໃນ GitHub | Run `git init` ພາຍໃນ `lesson01/` | ຍ້າຍ `.git` ຂຶ້ນໄປ `go-learning/`; git ບັນທຶກເປັນ rename, ສະນັ້ນ history ຍັງຢູ່ຄົບ |
| Commit binary ຂະໜາດ 2.2 MB ເຂົ້າໄປ | ບໍ່ມີ `.gitignore` ກ່ອນ commit ທຳອິດ | ເພີ່ມ `bin/` ໃນ `.gitignore`, ແລ້ວ `git rm --cached` binary ນັ້ນ |

### ບັນທຶກການ review

- ✅ ໂຄງສ້າງຖືກຕ້ອງ (`cmd/` + `internal/`). Doc comment ເລີ່ມດ້ວຍຊື່ຂອງສິ່ງທີ່ອະທິບາຍ ແລະ ຂຽນເປັນພາສາອັງກິດ ແລະ ລາວ (ຫຍໍ້ໜ້າ ENG/LAO, ຂັ້ນດ້ວຍແຖວ `//` ຫວ່າງ).
- ✅ ໃຊ້ `%q` ໃນຂໍ້ຄວາມ test failure.
- ✅ ການ assign ຄ່າໃໝ່ໃຫ້ parameter (`name = strings.TrimSpace(name)`) ເປັນແບບ Go ທີ່ຖືກຕ້ອງ: parameter ເປັນ local variable, ຈຶ່ງບໍ່ກະທົບຜູ້ເອີ້ນ.
- ✅ ພຶດຕິກຳ `TrimSpace` ທີ່ເພີ່ມມາ (`"   "` ກາຍເປັນ stranger, `"  Ton  "` ກາຍເປັນ `"Ton"`) ມີ test ແລ້ວ (ກໍລະນີ `whitespace only` ແລະ `padded name` ໃນ `TestHello`). ທຸກ branch ທີ່ຂຽນ ຕ້ອງມີ test.
- ✅ ລຶບ block `Parameters:` / `Returns` ອອກແລ້ວ (ເປັນນິໄສຈາກ JSDoc). Doc comment ຂອງ Go ຂຽນເປັນປະໂຫຍກທຳມະດາ.
- ✅ Test ສີ່ຕົວເຄີຍຊ້ຳ block `got` / `want` / `if` ແບບດຽວກັນ. ແກ້ແລ້ວໃນແບບຝຶກຫັດ 1.2 ດ້ວຍ **table-driven test** ຕົວດຽວ.

---

## ແບບຝຶກຫັດ 1.2: table-driven test + CLI argument

1. ແທນ test ສີ່ຕົວດ້ວຍ `TestHello` ແບບ table-driven ຕົວດຽວ ທີ່ໃຊ້ subtest `t.Run`. ໃຫ້ຄອບຄຸມ `"Ton"`, `""`, `"   "` ແລະ `"  Ton  "`.
2. `go run ./cmd/greeter Alice` ທັກທາຍ Alice; ຖ້າບໍ່ມີ argument ໃຫ້ທັກທາຍ stranger. ໃຊ້ `os.Args` ແລະ ກວດຄວາມຍາວກ່ອນເຂົ້າເຖິງ index.
3. Run:
   - `go test -v ./...`
   - `go test -run 'TestHello/<empty-case-name>' -v ./internal/greet`
   - `go test -cover ./...`
   - `GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/greeter-linux ./cmd/greeter && file bin/greeter-linux`
   - ເພີ່ມ `bin/` ໃນ `.gitignore` ກ່ອນ, ເພື່ອບໍ່ໃຫ້ຜົນການ build ຖືກ commit.
4. `os.Args[0]` ແມ່ນຫຍັງ, ແລະ ຈະເກີດຫຍັງຂຶ້ນຖ້າເຂົ້າເຖິງ `os.Args[1]` ໂດຍບໍ່ກວດຄວາມຍາວ? ທຽບກັບ `process.argv[2]` ໃນ Node ແນວໃດ?

**ວິທີແກ້**

- [`greet_test.go`](../lesson01/internal/greet/greet_test.go): `TestHello` ຕົວດຽວ ທີ່ມີ slice ຂອງ anonymous struct (`name`, `input`, `want`) ແລະ subtest `t.Run` ສຳລັບແຕ່ລະກໍລະນີ. ເພີ່ມກໍລະນີໃໝ່ພຽງແຖວດຽວ.
- [`main.go`](../lesson01/cmd/greeter/main.go): `name := ""`, ແລ້ວ `name = os.Args[1]` ສະເພາະເມື່ອ `len(os.Args) > 1`. ຊື່ຫວ່າງຈະໄປໃຊ້ fallback stranger ຂອງ `Hello` ເອງ.

**ຜົນລັບ**

```text
$ go run ./cmd/greeter Alice
Hello, Alice!
$ go run ./cmd/greeter
Hello, stranger!

$ go test -v ./...
?       example.com/lesson01/cmd/greeter        [no test files]
=== RUN   TestHello
=== RUN   TestHello/normal_name
=== RUN   TestHello/empty
=== RUN   TestHello/whitespace_only
=== RUN   TestHello/padded_name
--- PASS: TestHello (0.00s)
    --- PASS: TestHello/normal_name (0.00s)
    --- PASS: TestHello/empty (0.00s)
    --- PASS: TestHello/whitespace_only (0.00s)
    --- PASS: TestHello/padded_name (0.00s)
PASS
ok      example.com/lesson01/internal/greet

$ go test -run 'TestHello/empty' -v ./internal/greet
=== RUN   TestHello
=== RUN   TestHello/empty
--- PASS: TestHello (0.00s)
    --- PASS: TestHello/empty (0.00s)

$ go test -cover ./...
        example.com/lesson01/cmd/greeter                coverage: 0.0% of statements
ok      example.com/lesson01/internal/greet     coverage: 100.0% of statements

$ file bin/greeter-linux
bin/greeter-linux: ELF 64-bit LSB executable, x86-64, statically linked, with debug_info, not stripped
```

- ຊ່ອງວ່າງໃນຊື່ subtest ກາຍເປັນ underscore (`normal name` → `normal_name`).
- `-run 'TestHello/empty'` ຍັງ run `TestHello` ທີ່ເປັນແມ່, ແຕ່ run ພຽງກໍລະນີ `empty` ຢູ່ຂ້າງໃນ.
- `cmd/greeter` ມີ coverage 0% ເປັນເລື່ອງປົກກະຕິ; `main` ບໍ່ມີ test.
- `ELF` + `x86-64` = binary ຂອງ Linux amd64 (ຖ້າ build ໃຫ້ macOS ຈະເປັນ Mach-O arm64). ມັນ run ເທິງ Mac ບໍ່ໄດ້. `-ldflags="-s -w"` ຕັດ debug info ອອກເພື່ອໃຫ້ນ້ອຍລົງ.

### ຄຳຖາມ 4: `os.Args`

**`os.Args[0]`** ແມ່ນ path ທີ່ໃຊ້ເປີດໂປຣແກຣມ, ບໍ່ແມ່ນໄຟລ໌ source:

```text
$ go run ./cmd/greeter Alice
os.Args = [/var/folders/.../go-build1410952333/b001/exe/greeter Alice]
$ ./bin/greeter Alice
os.Args = [./bin/greeter Alice]
```

`go run` compile ໄວ້ໃນໂຟນເດີຊົ່ວຄາວ ແລ້ວ run binary ຈາກບ່ອນນັ້ນ. ຢ່າເພິ່ງ `os.Args[0]`; ມັນຂຶ້ນກັບວ່າເປີດໂປຣແກຣມດ້ວຍວິທີໃດ. Argument ຂອງຜູ້ໃຊ້ເລີ່ມທີ່ `os.Args[1]`.

**`os.Args[1]` ໂດຍບໍ່ກວດຄວາມຍາວ** compile ຜ່ານ ແຕ່ **panic ຕອນ runtime** ເມື່ອບໍ່ມີ argument:

```text
panic: runtime error: index out of range [1] with length 1

goroutine 1 [running]:
main.main()
        .../lesson01/cmd/greeter/main.go:21 +0xb0
exit status 2
```

`[1] with length 1`: ຂໍ index 1, ແຕ່ slice ມີພຽງ `os.Args[0]`. Stack trace ຊີ້ໄປແຖວທີ່ຜິດພໍດີ, ແລະ exit status 2 ໝາຍເຖິງ panic ທີ່ບໍ່ໄດ້ recover.

**Node:** `process.argv[2]` ເມື່ອບໍ່ມີ argument ແມ່ນພຽງ `undefined`. ບໍ່ crash, ແລະ ຄ່າທີ່ຜິດນັ້ນຈະຖືກສົ່ງຕໍ່ໄປຈົນມີບ່ອນອື່ນພັງພາຍຫຼັງ, ໄກຈາກສາເຫດ. Go ລົ້ມເຫຼວທັນທີທີ່ແຖວນັ້ນພໍດີ, ສະນັ້ນໃຫ້ກວດ `len(...)` ກ່ອນເຂົ້າເຖິງ index ສະເໝີ.

**ເປັນຫຍັງ Go ໃຊ້ `[1]` ແຕ່ Node ໃຊ້ `[2]`:** `node app.js Alice` ໄດ້ `[node path, script path, 'Alice']`, ເພາະ runtime ແລະ script ເປັນຄົນລະລາຍການ. Binary ຂອງ Go *ຄື* ໂປຣແກຣມເອງ, ສະນັ້ນມີພຽງລາຍການດຽວຢູ່ທາງໜ້າ. (ກັບ `node -e` ບໍ່ມີ script, ສະນັ້ນ `argv` ມີພຽງ path ຂອງ node.)
