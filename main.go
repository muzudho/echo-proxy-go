package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fullCmdLine := strings.Join(os.Args, " ") // 1. コマンドラインを文字列として取得
	fmt.Printf("Full command line: [%s]\n", fullCmdLine)

	//fs1 := flag.CommandLine                              // 2. コマンドラインに紐づいたフラグセットを取得
	fs1 := flag.NewFlagSet(os.Args[0], flag.ExitOnError) // 2. 新規フラグセットを作成（エラー時はプログラムを終了）

	parameters := os.Args[1:] // 5. コマンドライン引数をすべて取得
	pArgsMap := fromFilesetToPArgsMap(fs1, parameters)
	fmt.Printf("exe=%s\n", *pArgsMap["exe"]) // デバッグ出力。ヌルを指していれば、空文字列になるだけ。問題ない。

	if *pArgsMap["exe"] == "" { // 7. （あれば）必須のコマンドライン引数の確認
		panic(fmt.Errorf("--exe <Executable file path>"))
	}

	// FIXME: ここの parameters は、外部プロセスに渡す引数として適切か？
	//        例えば、--exe 自身や、その値が含まれているとまずい。
	//        なので、parameters を再構築する必要があるかもしれない。
	externalProcess := exec.Command(*pArgsMap["exe"], parameters...) // 外部プロセスコマンド作成

	stdin, err := externalProcess.StdinPipe() // 外部プロセス標準入力パイプ取得
	if err != nil {
		panic(err)
	}
	defer stdin.Close() // stdin を使い終わったら、外部プロセス標準入力パイプクローズ

	stdout, err := externalProcess.StdoutPipe() // 外部プロセス標準出力パイプ取得
	if err != nil {
		panic(err)
	}
	defer stdout.Close() // stdout を使い終わったら、外部プロセス標準出力パイプクローズ

	err = externalProcess.Start() // 外部プロセス起動
	if err != nil {
		panic(fmt.Errorf("cmd.Start() --> [%s]", err))
	}

	go receiveStdout(stdout) // 外部プロセスの標準出力受信開始

	// Go言語では標準出力のUTF-8に対応していますが、VSCodeのターミナルはUTF-8に対応していないようです。
	// `chcp 65001`
	// そのため、外部プロセスの標準出力を受信しても、正しく表示されない場合があります。
	// その場合は、WindowsのコマンドプロンプトやPowerShellなど、UTF-8に対応したターミナルで実行してください。

	go receiveStdin(stdin) // 外部プロセスの標準入力送信開始

	fmt.Print("[main.go] Waiting for external process to finish...\n")
	externalProcess.Wait()

	fmt.Print("[main.go] process finished.\n")
}

func fromFilesetToPArgsMap(fs1 *flag.FlagSet, arguments []string) map[string]*string {
	pArgsMap := make(map[string]*string)                               // 3. コマンドライン引数名と、その値が入る変数へのポインターを紐づけるマップ
	pArgsMap["exe"] = fs1.String("exe", "", "Working directory path.") // 4. コマンドライン引数を登録し、後でその値が入る変数へのポインターを取得

	fs1.Parse(arguments) // 6. コマンドライン引数の解析
	return pArgsMap
}

// receiveStdin - 標準入力受信
// `epStdin` - External process stdin
func receiveStdin(epStdin io.WriteCloser) {
	scanner := bufio.NewScanner(os.Stdin) // 標準入力を読み取るスキャナ作成
	for scanner.Scan() {
		command := scanner.Text() // １行読み取り。UTF-8文字列。
		epStdin.Write([]byte(command))
		epStdin.Write([]byte("\n"))
	}
}

// receiveStdout - 標準出力受信
// `epStdout` - External process stdout
func receiveStdout(epStdout io.ReadCloser) {
	var buffer [1]byte // これが満たされるまで待つ。1バイト。

	p := buffer[:]

	for {
		n, err := epStdout.Read(p)

		if nil != err {
			if fmt.Sprintf("%s", err) == "EOF" {
				// End of file
				return
			}

			panic(err)
		}

		if 0 < n {
			bytes := p[:n]

			// 思考エンジンから１文字送られてくるたび、表示。
			print(string(bytes))
		}
	}
}
