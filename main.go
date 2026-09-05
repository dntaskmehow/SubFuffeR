//go:build !linux

package main

import (
 "bufio"
 "context"
 "encoding/json"
 "errors"
 "fmt"
 "io"
 "net/http"
 "net/url"
 "os"
 "os/exec"
 "path/filepath"
 "regexp"
 "strings"
 "sync"
 "fyne.io/fyne/v2"
 "fyne.io/fyne/v2/app"
 "fyne.io/fyne/v2/container"
 "fyne.io/fyne/v2/widget"
)

type ScopeItem struct { Asset string `json:"asset_identifier"`; Type string `json:"asset_type"`; Eligible bool `json:"eligible_for_submission"` }
type scopeResponse struct { Data []struct { Attributes ScopeItem `json:"attributes"` } `json:"data"` }
type Runner struct { mu sync.Mutex; cancel context.CancelFunc; running bool }
var hostRE = regexp.MustCompile(`(?i)^(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z]{2,63}$`)

func main() {
 a := app.NewWithID("com.subfuffer.desktop"); w := a.NewWindow("SubFuffeR | HackerOne Scope Guard"); w.Resize(fyne.NewSize(1120, 760))
 target := widget.NewEntry(); target.SetPlaceHolder("Attack target, for example app.example.com")
 h1link := widget.NewEntry(); h1link.SetPlaceHolder("HackerOne program URL, for example https://hackerone.com/example")
 sub := widget.NewCheck("subfinder", nil); dns := widget.NewCheck("dnsx", nil); fuzz := widget.NewCheck("ffuf", nil)
 wordlist := widget.NewEntry(); wordlist.SetPlaceHolder("ffuf wordlist path (required only for ffuf)")
 concurrency := widget.NewSlider(1, 50); concurrency.Value = 10; concurrencyLabel := widget.NewLabel("HTTP concurrency: 10")
 concurrency.OnChanged = func(v float64) { concurrencyLabel.SetText(fmt.Sprintf("HTTP concurrency: %d", int(v))) }
 log := widget.NewMultiLineEntry(); log.Wrapping = fyne.TextWrapWord; log.Disable(); appendLog := func(s string) { log.SetText(log.Text + s + "\n") }
 runner := &Runner{}
 load := widget.NewButton("Load HackerOne scope", func() { go func() { items, err := loadScope(h1link.Text); if err != nil { appendLog("[scope] " + err.Error()); return }; appendLog(fmt.Sprintf("[scope] loaded %d structured assets", len(items))); for _, item := range items { appendLog(fmt.Sprintf("[scope] %s %s eligible=%t", item.Type, item.Asset, item.Eligible)) } }() })
 start := widget.NewButton("Run selected tools", func() { if runner.running { return }; if !sub.Checked && !dns.Checked && !fuzz.Checked { appendLog("[guard] select at least one tool"); return }; if !authorizedHost(target.Text) { appendLog("[guard] target must be a plain hostname in the entered HackerOne scope"); return }; if fuzz.Checked && strings.TrimSpace(wordlist.Text) == "" { appendLog("[guard] ffuf requires an explicit wordlist path"); return }; ctx, cancel := context.WithCancel(context.Background()); runner.mu.Lock(); runner.cancel, runner.running = cancel, true; runner.mu.Unlock(); go func() { defer func() { runner.mu.Lock(); runner.running = false; runner.mu.Unlock() }(); runPipeline(ctx, target.Text, h1link.Text, sub.Checked, dns.Checked, fuzz.Checked, wordlist.Text, int(concurrency.Value), appendLog) }() })
 stop := widget.NewButton("Cancel", func() { runner.mu.Lock(); if runner.cancel != nil { runner.cancel() }; runner.mu.Unlock() })
 top := container.NewVBox(widget.NewLabelWithStyle("SubFuffeR", fyne.TextAlignLeading, fyne.TextStyle{Bold:true}), widget.NewLabel("HackerOne constrained reconnaissance workspace"))
 form := container.NewVBox(widget.NewLabel("Attack target"), target, widget.NewLabel("HackerOne scope URL"), h1link, load, widget.NewLabel("Tools"), container.NewHBox(sub, dns, fuzz), widget.NewLabel("ffuf wordlist"), wordlist, concurrencyLabel, concurrency, widget.NewSeparator(), widget.NewLabel("Structured event stream"), log)
 w.SetContent(container.NewBorder(top, container.NewVBox(start, stop), nil, nil, container.NewVScroll(form))); w.ShowAndRun()
}

func authorizedHost(input string) bool { h := strings.ToLower(strings.TrimSpace(input)); return !strings.Contains(h, "/") && hostRE.MatchString(h) }
func programHandle(raw string) (string, error) { u, err := url.Parse(strings.TrimSpace(raw)); if err != nil || u.Host != "hackerone.com" { return "", errors.New("use a HackerOne program URL on hackerone.com") }; p := strings.Split(strings.Trim(u.Path, "/"), "/"); if len(p) != 1 || p[0] == "" { return "", errors.New("HackerOne URL must identify one program handle") }; return p[0], nil }
func loadScope(raw string) ([]ScopeItem, error) { h, err := programHandle(raw); if err != nil { return nil, err }; req, _ := http.NewRequest(http.MethodGet, "https://api.hackerone.com/v1/hackers/programs/"+url.PathEscape(h)+"/structured_scopes", nil); req.Header.Set("Accept", "application/json"); if user, token := os.Getenv("HACKERONE_USERNAME"), os.Getenv("HACKERONE_API_TOKEN"); user != "" && token != "" { req.SetBasicAuth(user, token) }; resp, err := http.DefaultClient.Do(req); if err != nil { return nil, err }; defer resp.Body.Close(); if resp.StatusCode < 200 || resp.StatusCode > 299 { return nil, fmt.Errorf("HackerOne API returned %s", resp.Status) }; var page scopeResponse; if err := json.NewDecoder(resp.Body).Decode(&page); err != nil { return nil, err }; out := make([]ScopeItem, 0, len(page.Data)); for _, d := range page.Data { out = append(out, d.Attributes) }; return out, nil }
func runPipeline(ctx context.Context, target, scopeURL string, doSub, doDNS, doFuzz bool, wordlist string, concurrency int, emit func(string)) { emit("[guard] verify HackerOne program rules, exclusions, rate limits, and authorization before proceeding"); scope, err := loadScope(scopeURL); if err != nil { emit("[guard] scope load failed: " + err.Error()); return }; if !inScope(target, scope) { emit("[guard] target is not an eligible asset in the loaded HackerOne structured scope"); return }; items := []string{target}; if doSub { items, err = runTool(ctx, "subfinder", []string{"-d", target, "-silent"}, emit); if err != nil { emit("[subfinder] " + err.Error()); return } }; if doDNS { items, err = runTool(ctx, "dnsx", []string{"-silent", "-l", strings.Join(items, "\n")}, emit); if err != nil { emit("[dnsx] " + err.Error()); return } }; if doFuzz { for _, host := range items { if !authorizedHost(host) || !inScope(host, scope) { continue }; args := []string{"-u", "https://" + host + "/FUZZ", "-w", wordlist, "-t", fmt.Sprint(concurrency), "-rate", fmt.Sprint(concurrency * 5), "-ac", "-json"}; if _, err := runTool(ctx, "ffuf", args, emit); err != nil { emit("[ffuf] " + err.Error()); return } } }; emit("[engine] completed") }
func inScope(host string, scope []ScopeItem) bool { h := strings.ToLower(strings.TrimSuffix(strings.TrimSpace(host), ".")); for _, item := range scope { if !item.Eligible { continue }; asset := strings.ToLower(strings.TrimSpace(item.Asset)); asset = strings.TrimPrefix(asset, "*."); asset = strings.TrimPrefix(asset, "https://"); asset = strings.TrimPrefix(asset, "http://"); asset = strings.Split(asset, "/")[0]; if h == asset || strings.HasSuffix(h, "."+asset) { return true } }; return false }
func runTool(ctx context.Context, name string, args []string, emit func(string)) ([]string, error) { path := filepath.Join("tools", name+".exe"); if _, err := os.Stat(path); err != nil { path = name }; cmd := exec.CommandContext(ctx, path, args...); stdout, err := cmd.StdoutPipe(); if err != nil { return nil, err }; stderr, err := cmd.StderrPipe(); if err != nil { return nil, err }; if err := cmd.Start(); err != nil { return nil, fmt.Errorf("%s unavailable (%v); place pinned executable in tools\\ or add it to PATH", name, err) }; lines := make([]string, 0); var wg sync.WaitGroup; wg.Add(2); read := func(r io.Reader, isErr bool) { defer wg.Done(); sc := bufio.NewScanner(r); for sc.Scan() { s := strings.TrimSpace(sc.Text()); if s == "" { continue }; if !isErr { lines = append(lines, s) }; emit("["+name+"] "+s) } }; go read(stdout, false); go read(stderr, true); wg.Wait(); err = cmd.Wait(); if ctx.Err() != nil { return lines, ctx.Err() }; return lines, err }
