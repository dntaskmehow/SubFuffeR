# Architecture

SubFuffeR has two platform entry points selected at build time.

The Windows build uses Fyne and `main.go` for a native desktop interface. The Linux build uses `linux_main.go` for a colored ANSI terminal interface. A build tag keeps the UI implementations separate while retaining the same product boundary.

Each run accepts an explicit hostname and HackerOne program URL. The connector requests the program's structured scopes endpoint, keeps eligible assets, and checks the target before launching tools. Subfinder output feeds dnsx, and resolved hosts can feed ffuf only when ffuf is explicitly selected and a wordlist is supplied.

External tools are started with `exec.CommandContext` and argument slices. Their stdout and stderr are streamed concurrently. The Windows UI presents the stream incrementally; ffuf concurrency and rate are bounded by user controls. The application has no startup registration, persistence, firewall modification, privilege escalation, or shell command construction.

Tool binaries are intentionally supplied separately so operators can pin, verify, and license each upstream release. See the README for build commands and dependencies.
