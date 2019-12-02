# GOLANG .vscode/settings.json

This is an example of setting.json for golang workspace using vscode.
You can copy following lines or copy folder `.vscode` and paste it on your vscode workspace folder.
This setting includes auto-format, 
## settings.json
```json
{
    "go.buildOnSave": "off",
    "go.lintOnSave": "workspace",
    "go.vetOnSave": "workspace",
    "go.buildTags": "",
    "go.buildFlags": [],
    "go.lintTool": "golint",
    "go.lintFlags": [],
    "go.vetFlags": [],
    "go.testOnSave": false,
    "go.coverOnSave": false,
    "go.useCodeSnippetsOnFunctionSuggest": true,
    "editor.formatOnSave": true,
    "go.formatTool": "goreturns",
    "go.formatFlags": [],
    "go.goroot": "",
    "go.gopath": "",
    "go.inferGopath": true,
    "go.gocodeAutoBuild": false,
    "go.testFlags": [
        "-v"
    ],
}

```


## References:
<https://github.com/takecy/vscode.go.setting/blob/master/settings.json>