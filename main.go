package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

// resolveCwd はターゲットパスから WezTerm に渡す cwd を決定する。
// ディレクトリならそのまま、ファイルなら親ディレクトリを返す。
// 解決できない場合はカレントディレクトリを返す。
func resolveCwd(target string) string {
	abs, err := filepath.Abs(target)
	if err != nil {
		cwd, _ := os.Getwd()
		return cwd
	}

	info, err := os.Stat(abs)
	if err != nil {
		cwd, _ := os.Getwd()
		return cwd
	}

	if info.IsDir() {
		return abs
	}
	return filepath.Dir(abs)
}

func main() {
	nvimArgs := []string{"nvim"}
	cwd, _ := os.Getwd()

	if len(os.Args) >= 2 && os.Args[1] != "" {
		target := os.Args[1]
		nvimArgs = append(nvimArgs, target)
		cwd = resolveCwd(target)
	}

	args := append([]string{"start", "--cwd", cwd, "--"}, nvimArgs...)
	cmd := exec.Command("wezterm-gui.exe", args...)
	_ = cmd.Start() // 起動だけして即終了
}
