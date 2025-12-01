@echo off

rem 文字化け対策。コマンドプロンプトがデフォルトで Shift-JIS なので、その文字コードを消すことで、UTF-8 にする。
chcp 65001 >nul

echo `C:/MuzudhoWorks/echo-proxy-go/echo-proxy-go.exe` ファイルを実行するぜ（＾～＾）...
C:/MuzudhoWorks/echo-proxy-go/echo-proxy-go.exe
echo 実行したぜ（＾～＾）！
pause
