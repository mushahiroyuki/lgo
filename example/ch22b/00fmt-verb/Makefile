##
## ほかのOS、アーキテクチャ用の実行ファイルを作るためのMakefileです
##  たとえば
##  「make win」 でWindows用のファイルができます。
win:
	GOOS=windows GOARCH=386 go build -o fmt-verb-windows.exe main.go
mac-intel:
	GOOS=darwin GOARC=amd64 go build -o fmt-verb-mac-intel
linux-amd64:
	GOOS=linux GOARCH=amd64 go build

