// +build ci

package ci

import (
	"io"
	"os"
	"os/exec"
	"testing"
)

type liteproc struct {
	stdin io.WriteCloser
	log   io.Closer
	cmd   *exec.Cmd
}

// open sqlite3 process which you can write to
func openSqlite(t *testing.T, file string) *liteproc {
	t.Helper()
	out, err := os.Create(file + ".out")
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("sqlite3", file)
	cmd.Stdout = out
	cmd.Stderr = out
	w, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	return &liteproc{
		stdin: w,
		log:   out,
		cmd:   cmd,
	}
}

func (l *liteproc) Close() {
	l.stdin.Close()
	l.cmd.Wait()
	l.log.Close()
}

func (l *liteproc) Kill() {
	l.cmd.Process.Kill()
}

func (l *liteproc) write(t *testing.T, sql string) {
	if n, err := l.stdin.Write([]byte(sql)); err != nil || n != len(sql) {
		t.Fatal(err)
	}
}
