package config

import (
	"hello/utils"
	"log"

	"gopkg.in/go-ini/ini.v1"
)

// Configリストの構造体
type ConfigList struct {
	// 各フィールド作成

	// サーバーのポート番号
	Port string
	// 使用するドライバの設定
	SQLDriver string
	// データーベース名
	DbName string
	// ログを残すファイル
	LogFile string
}

// Configリストを外部から呼び出せるように、グローバルに変数宣言
var Config ConfigList

// main関数より前に実行してconfigリストを作成するため
func init() {
	// main関数より前にini関数が呼ばれて,Loadconfigを実行する
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

// iniファイルを読み込んで、configListに値を読み込む
func LoadConfig() {
	// iniファイルを読み込む
	cfg, err := ini.Load("config.ini")
	// エラーを出力
	if err != nil {
		log.Fatalln(err)
	}

	//　グローバルに宣言されたconfigに読み込んだ値を設定していく
	Config = ConfigList{
		// 読み込んだiniファイルをconfigListに設定
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
