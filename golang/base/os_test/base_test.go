package os_test

import (
	"os"
	"os/exec"
	"testing"
)

func TestOpenFile(t *testing.T) {
	f, err := os.OpenFile("./demo.txt", os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
}

// 获取系统变量
func TestGetEnv(t *testing.T) {
	envVars := os.Environ()
	for _, envVar := range envVars {
		t.Log(envVar)
	}

	path := os.Getenv("path")
	t.Log(path)
}

func TestGetPid(t *testing.T) {
	pid := os.Getpid()
	t.Log(pid)
}

func TestCmd(t *testing.T) {
	cmd := exec.Command("ping", "z.cn")
	err := cmd.Start()
	if err != nil {
		t.Fatal(err)
	}
}

func TestChmod(t *testing.T) {
	err := os.Chmod("./demo.log", 0777)
	if err != nil {
		panic(err)
	}
	f, _ := os.Open("./demo.txt")
	defer f.Close()
	err = f.Chmod(0777)
	if err != nil {
		panic(err)
	}
}
