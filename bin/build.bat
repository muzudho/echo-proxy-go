@echo off

rem 文字化け対策。コマンドプロンプトがデフォルトで Shift-JIS なので、その文字コードを消すことで、UTF-8 にする。
chcp 65001 >nul

echo 外部ドライブにexeファイル置いても実行権限無いし、`C:/MuzudhoWorks/echo-proxy-go` に `echo-proxy-go.exe` ファイルを作るぜ（＾～＾）...
cd ..
go build -o C:/MuzudhoWorks/echo-proxy-go
cd ./bin
echo go build したぜ（＾～＾）！
pause
