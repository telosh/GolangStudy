package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
	// 画面クリアの方法に応じて必要なパッケージをインポート
	// "os"
	// "os/exec"
	// "runtime"
)

// 画面クリア関数
func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Print("\033[H\033[2J")
	}
}

func main() {
	// 1. アニメーションのコマを用意
	// 各フレームは複数行の文字列なので、[]string で表現
	// それをさらにスライスやマップでまとめる
	animationFrames := [][]string{
		{ // フレーム1
			" (\\___/) ",
			" (='.'=) ",
			" (\\\"_(\")_(\")",
		},
		{ // フレーム2
			" (\\___/) ",
			" (='.'=) ",
			" /)_(\")_(\") ", // 少し手が動いた感じ
		},
		{ // フレーム3
			" (\\___/) ",
			" (-'.'-) ", // 目を閉じた感じ
			" (\\\"_(\")_(\")",
		},
		{ // フレーム4 (フレーム2と同じ)
			" (\\___/) ",
			" (='.'=) ",
			" /)_(\")_(\") ",
		},
	}

	// アニメーションのループ
	for { // 無限ループ (Ctrl+Cで終了)
		for _, frame := range animationFrames {
			// 2. 画面をクリア
			clearScreen()

			// 3. 現在のコマを表示
			for _, line := range frame {
				fmt.Println(line)
			}

			// 4. 待機 (アニメーション速度を調整)
			time.Sleep(200 * time.Millisecond) // 0.2秒待機
		}
	}
}