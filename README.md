# SubFuffeR

```text
  ███████╗██╗   ██╗██████╗ ███████╗██╗   ██╗███████╗██████╗
  ██╔════╝██║   ██║██╔══██╗██╔════╝██║   ██║██╔════╝██╔══██╗
  ███████╗██║   ██║██████╔╝█████╗  ██║   ██║█████╗  ██████╔╝
  ╚════██║██║   ██║██╔══██╗██╔══╝  ╚██╗ ██╔╝██╔══╝  ██╔══██╗
  ███████║╚██████╔╝██████╔╝███████╗ ╚████╔╝ ███████╗██║  ██║
  ╚══════╝ ╚═════╝ ╚═════╝ ╚══════╝  ╚═══╝  ╚══════╝╚═╝  ╚═╝
```

> Scope-first reconnaissance for authorized HackerOne bug bounty testing.

| Flow | Control | Output |
| --- | --- | --- |
| Target + HackerOne URL | Eligible structured-scope gate | Safe target set |
| subfinder | Passive discovery | Streaming findings |
| dnsx | DNS resolution | Resolved hosts |
| ffuf | Explicit user action | Bounded HTTP results |

Windows GUI and Kali Linux terminal editions are included. Read the full program policy and exclusions before every run.

SubFuffeR is a Go desktop coordinator for explicit, authorized security testing. It accepts an attack target and HackerOne program URL, loads structured scope, and lets the user run subfinder, dnsx, ffuf, or a selected sequence.

The HackerOne connector is a scope reference and enforcement input. It does not grant permission or prove bounty eligibility; read the program policy. Set `HACKERONE_USERNAME` and `HACKERONE_API_TOKEN` for authenticated scope data. Credentials are never written by the app.

## Build

From this directory with Go 1.26+, GCC/MSYS2, and Fyne dependencies available:

```powershell
go mod tidy
$env:CGO_ENABLED='1'
go build -ldflags='-H=windowsgui' -o SubFuffeR.exe .
```

Place pinned, user obtained `subfinder.exe`, `dnsx.exe`, and `ffuf.exe` in `tools\`, or put them on `PATH`. The app does not download or embed them. Fyne uses CGO and OpenGL/graphics drivers. No WebView2 runtime is required.

Tool invocations use `exec.CommandContext` argument arrays and support cancellation. ffuf is explicit and bounded by the UI thread and rate controls. There is no startup task, persistence, firewall change, privilege escalation, or shell command concatenation. Do not point development builds at live targets.

## Third party notices

The UI uses Fyne v2.6.3 under its upstream license. The separate tools retain their upstream licenses and notices: projectdiscovery/subfinder, projectdiscovery/dnsx, and ffuf/ffuf. Include license files with `tools\` when distributing them.

HackerOne structured scopes documentation: https://api.hackerone.com/hacker-resources/
