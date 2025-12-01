@echo off

rem 文字化け対策。コマンドプロンプトがデフォルトで Shift-JIS なので、その文字コードを消すことで、UTF-8 にする。
chcp 65001 >nul

echo 全部任せろだぜ（＾～＾）...

rem プロジェクト・ルートで作業する。
cd ..

rem Go言語を使っています。
rem GO111MODULE を on に設定します。（モジュール・モード） 
set GO111MODULE=on



