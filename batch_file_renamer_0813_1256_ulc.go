// 代码生成时间: 2025-08-13 12:56:36
package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"

    "github.com/labstack/echo"
)

// RenameRule 定义了重命名规则的结构
type RenameRule struct {
    Pattern  string
    Replacer string
}

// Renamer 定义了批量重命名文件的工具结构
type Renamer struct {
    Dir      string
    Rules    []RenameRule
    DryRun   bool
}

// NewRenamer 创建一个新的Renamer实例
func NewRenamer(dir string, rules []RenameRule, dryRun bool) *Renamer {
    return &Renamer{
        Dir:      dir,
        Rules:    rules,
        DryRun:   dryRun,
    }
}

// RenameFiles 执行文件重命名操作
func (r *Renamer) RenameFiles() error {
    err := filepath.Walk(r.Dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            for _, rule := range r.Rules {
                if matched, err := filepath.Match(rule.Pattern, filepath.Base(path)); err != nil {
                    return err
                } else if matched {
                    newBase := strings.ReplaceAll(filepath.Base(path), rule.Pattern, rule.Replacer)
                    newPath := filepath.Join(filepath.Dir(path), newBase)
                    if !r.DryRun {
                        if err := os.Rename(path, newPath); err != nil {
                            return err
                        }
                    }
                    fmt.Printf("Renamed '%s' to '%s'
", path, newPath)
                }
            }
        }
        return nil
    })
    return err
}

func main() {
    var (
        dir      string
        pattern  string
        replacer string
        dryRun   bool
    )
    flag.StringVar(&dir, "dir", "./", "Directory to rename files in")
    flag.StringVar(&pattern, "pattern", "", "Pattern to match filenames")
    flag.StringVar(&replacer, "replacer", "", "String to replace matched pattern")
    flag.BoolVar(&dryRun, "dry-run", false, "Simulate rename without actually renaming files")
    flag.Parse()

    if pattern == "" || replacer == "" {
        log.Fatalf("Both pattern and replacer must be provided")
    }

    rules := []RenameRule{RenameRule{Pattern: pattern, Replacer: replacer}}
    renamer := NewRenamer(dir, rules, dryRun)

    if err := renamer.RenameFiles(); err != nil {
        log.Fatalf("Failed to rename files: %v", err)
    }
}
