package main

import (
    "fmt"
)
import (_ "github.com/go-sql-driver/mysql"
    "github.com/ume-gear/core"
    "runtime/debug"
)

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("panic recover! p: %v", r)
            debug.PrintStack()
        }
    }()
    // 测试
    fmt.Println("Test start")

    testProp()

    fmt.Println("Test finished")

}

func testIni() {
    cfg := core.NewIniConfig("src/github.com/ume-gear/core/test/config/application.ini")
    fmt.Println(cfg.ParagraphSet())
    fmt.Println(cfg.Paragraph("JDBC basic"))
}

func testProp() {
    cfg := core.NewPropertyConfig("src/github.com/ume-gear/core/test/config/application.properties")
    fmt.Println(cfg.KeySet())
    fmt.Println(cfg.Get("ume.jdbc.url"))
}


