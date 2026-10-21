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

> Scope-first reconnaissance for authorized HackerOne bug bounty testing.

| Flow | Control | Output |
| --- | --- | --- |
| Target + HackerOne URL | Eligible structured-scope gate | Safe target set |
| subfinder | Passive discovery | Streaming findings |
| dnsx | DNS resolution | Resolved hosts |
| ffuf | Explicit user action | Bounded HTTP results |

---

 **Windows GUI** and Kali Linux terminal editions are included. Read the full program policy and exclusions before every run.

---

**SubFuffeR** is a Go desktop coordinator for explicit, authorized security testing. It accepts an attack target and HackerOne program URL, loads structured scope, and lets the user run subfinder, dnsx, and ffuf in a controlled workflow.

---

**The** HackerOne connector is a scope reference and enforcement input. It does not grant permission or prove bounty eligibility; read the program policy. Set `HACKERONE_USERNAME` and `HACKERONE_API_TOKEN` as needed for authenticated scope operations.

---
*Build*

**From** this directory with Go 1.26+, GCC/MSYS2, and Fyne dependencies available:

```powershell
go mod tidy
$env:CGO_ENABLED='1'
go build -ldflags='-H=windowsgui' -o SubFuffeR.exe .
```

Place pinned, user obtained `subfinder.exe`, `dnsx.exe`, and `ffuf.exe` in `tools\`, or put them on `PATH`. The app does not download or embed them. Fyne uses CGO and OpenGL/graphics drivers. No WebView is used.

Tool invocations use `exec.CommandContext` argument arrays and support cancellation. ffuf is explicit and bounded by the UI thread and rate controls. There is no startup task, persistence, firewall change, or network relay beyond the user-initiated execution path.

## Third party notices

The UI uses Fyne v2.6.3 under its upstream license. The separate tools retain their upstream licenses and notices: projectdiscovery/subfinder, projectdiscovery/dnsx, and ffuf/ffuf. Include license files where required by local policy.

HackerOne structured scopes documentation: https://api.hackerone.com/hacker-resources/

---

<div align="center">

<img src="https://readme-typing-svg.demolab.com?font=Orbitron&size=28&pause=1000&color=FF0000&center=true&vCenter=true&width=700&lines=Recon+-+Enumeration+-+Target+Discovery;Researcher+dontaskmehow;Targeted+Host+Mapping" alt="SubFuFFer banner" />

</div>
