package main

import (
    "context"
    "github.com/chromedp/chromedp"
    "log"
    "net/http"
)

func Screenshot(urlstr string, res *[]byte) chromedp.Tasks {
    return chromedp.Tasks{
        chromedp.Navigate(urlstr),
        chromedp.EmulateViewport(800, 600),
        // these were added while doing performance tests
        chromedp.WaitReady("body"),
        chromedp.FullScreenshot(res, 100),
    }
}

func (app *Config) GetPicture(w http.ResponseWriter, r *http.Request)  {
    // run task list
    allocCtx, allocCancel := chromedp.NewRemoteAllocator(context.Background(), "ws://chromium:9222")
    defer allocCancel()

    ctx, cancel := chromedp.NewContext(allocCtx)
    defer cancel()

//    var buf []byte
    err := chromedp.Run(ctx, Screenshot("https://google.com", &app.buf))
    if err != nil {
        log.Fatal(err)
    }
    w.Write(app.buf)
    app.buf = app.buf[:0]

//    chromedp.Stop()
}