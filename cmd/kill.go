package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func killByPort(port string) {
	var cmd *exec.Cmd
	var pid string

	switch runtime.GOOS {
	case "linux", "darwin":
		// macOS / Linux
		cmd = exec.Command("lsof", "-i", fmt.Sprintf(":%s", port), "-t")
	case "windows":
		// Windows: netstat + findstr
		cmd = exec.Command("powershell", "-Command",
			fmt.Sprintf("(netstat -ano | findstr :%s) -match '\\d+$' | %% { $_.Split()[-1] } | Select-Object -Unique", port))
	default:
		fmt.Println("❌ 不支持的系统：", runtime.GOOS)
		return
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("⚠️ 未找到占用端口 %s 的进程。\n", port)
		return
	}

	pid = strings.TrimSpace(out.String())
	if pid == "" {
		fmt.Printf("⚠️ 端口 %s 没有被任何进程占用。\n", port)
		return
	}

	fmt.Printf("🔍 找到进程 PID: %s\n", pid)

	// 执行 kill
	var killCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		killCmd = exec.Command("taskkill", "/PID", pid, "/F")
	} else {
		killCmd = exec.Command("kill", "-9", pid)
	}

	if err := killCmd.Run(); err != nil {
		fmt.Printf("❌ 结束进程失败: %v\n", err)
		return
	}

	fmt.Printf("✅ 已结束占用端口 %s 的进程 (PID: %s)\n", port, pid)
}
