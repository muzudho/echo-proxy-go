# go-echo-proxy

標準入力を外部プロセスAの標準入力へ渡し、外部プロセスAの標準出力を標準出力へ渡すぜ（＾～＾）


## Set up

VSCode の 📄 `settings.json` ファイルを開き、以下の行を追加（抜粋）：  

📄 `settings.json` 抜粋:  

```json
{
    "go.toolsEnvVars": {
        "LANG": "ja_JP.UTF-8"
    },

    "terminal.integrated.defaultProfile.windows": "Command Prompt",
    "terminal.integrated.profiles.windows": {
        "Command Prompt": {
            "args": ["/k", "chcp 65001"]
        }
    }
}
```

👆 これで、VSCode のターミナルのデフォルトはコマンド・プロンプトになり、文字エンコーディングを UTF-8 に設定します。  

* Windows のスタート・ボタンの横の検索ボックスに［control］と打鍵。
* ［コントロールパネル］をクリック。
* ［時計と地域　＞　地域　＞　管理　＞　システムロケールの変更］をクリック。
* ［ワールドワイド言語サポートで...］チェックボックスをチェック。
* ［OK］ボタン・クリック。
* PCを再起動。

👆 これで、VSCode のターミナルが UTF-8 対応になるはず？ ならない。  

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
