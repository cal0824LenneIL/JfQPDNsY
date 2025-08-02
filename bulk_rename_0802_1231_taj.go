// 代码生成时间: 2025-08-02 12:31:57
@author 你的名字
@date 2023-05-06
*/

package main

import (
    "flag"
    "fmt"
    "io/fs"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"
)

// Renamer 定义批量重命名工具的结构
type Renamer struct {
    dryRun bool
    pattern string
    watchDirs []string
}

// NewRenamer 创建并初始化 Renamer 实例
func NewRenamer(dryRun bool, pattern string, watchDirs []string) *Renamer {
    return &Renamer{
        dryRun:    dryRun,
        pattern:   pattern,
        watchDirs: watchDirs,
    }
}

// Run 执行批量文件重命名
func (r *Renamer) Run() error {
    for _, dir := range r.watchDirs {
        err := r.renameFilesInDir(dir)
        if err != nil {
            return err
        }
    }
    return nil
}

// renameFilesInDir 重命名指定目录下的所有文件
func (r *Renamer) renameFilesInDir(dir string) error {
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        return fmt.Errorf("read dir failed: %w", err)
    }

    for _, file := range files {
        if file.IsDir() {
            continue
        }
        if !strings.HasSuffix(file.Name(), r.pattern) {
            continue
        }

        oldPath := filepath.Join(dir, file.Name())
        newPath := filepath.Join(dir, fmt.Sprintf("%s-%s%s",
            file.Name()[:len(file.Name())-len(r.pattern)],
            time.Now().Format("20060102-150405"),
            r.pattern,
        ))

        if r.dryRun {
            fmt.Printf("dry run: rename %s to %s
", oldPath, newPath)
            continue
        }

        err = os.Rename(oldPath, newPath)
        if err != nil {
            return fmt.Errorf("rename file failed: %w", err)
        }
        fmt.Printf("renamed %s to %s
", oldPath, newPath)
    }
    return nil
}

func main() {
    dryRun := flag.Bool("dry-run", false, "dry run mode")
    pattern := flag.String("pattern", ".txt", "file pattern to rename")
    watchDirs := flag.String("dirs", ".", "directories to watch")
    flag.Parse()

    dirs := strings.Split(*watchDirs, ",")
    renamer := NewRenamer(*dryRun, *pattern, dirs)
    err := renamer.Run()
    if err != nil {
        log.Fatalf("error occurred: %s
", err)
    }
}
