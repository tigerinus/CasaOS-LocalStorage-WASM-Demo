# CasaOS-LocalStorage-WASM-Demo

just a crappy demo of consuming events from LocalStorage and transcribing each to something useful for CasaOS-UI

## IDE Setup

Include following settings in `.vscode/settings.json`

```json
{
    "go.toolsEnvVars": {
        "GOOS": "js",
        "GOARCH": "wasm"
    }
}
```

## Build

Artifacts will be included in the `.tar.gz` file built by following command

```bash
goreleaser release --rm-dist --snapshot
```
