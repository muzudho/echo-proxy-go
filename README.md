# go-echo-proxy

標準入力を外部プロセスAの標準入力へ渡し、外部プロセスAの標準出力を標準出力へ渡すぜ（＾～＾）


## Set up

VSCode の 📄 `settings.json` ファイルを開き、以下の行を追加（抜粋）：  

```json
{
    "terminal.integrated.defaultProfile.windows": "Command Prompt",
    "terminal.integrated.profiles.windows": {
        "Command Prompt": {
            "args": ["/k", "chcp 65001"]
        }
    }
}
```

👆 これで、VSCode のターミナルのデフォルトはコマンド・プロンプトになり、文字エンコーディングを UTF-8 に設定します。  

VSCode のターミナルを開きなおす。  

```shell
go build
```


## Run

```shell
## Format:
## go-echo-proxy --exe <ExecutableFilePath>
go-echo-proxy --exe Z:/muzudho-github.com/muzudho/go-practice/go-practice.exe
```
