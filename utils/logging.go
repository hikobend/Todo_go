package utils

import (
	"io"
	"log"
	"os"
)

// logの設定の関数
func LoggingSettings(logFile string) {
	// ログファイルの読み込み
	// 引数のログファイルを読み込む。読み書き、ファイルがなければ作成、ファイルの追記。パーミッションは0666
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// エラーハンドリング
	if err != nil {
		log.Fatalln(err)
	}
	// ログの書き込み先を標準出力とログファイルに指定
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	//フォーマット指定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
