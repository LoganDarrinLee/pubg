package common

import (
  "context"
  "net/http"
  "log"
  "os"
  "os/signal"
  "syscall"
  "time"
)

func GracefulShutdown(ctx context.Context,server *http.Server) {
  sigChan := make(chan os.Signal, 1)
  signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
  <-sigChan

  shutDownCtx, shutDownRelease := context.WithTimeout(ctx, 10*time.Second)
  defer shutDownRelease()

  if err := server.Shutdown(shutDownCtx); err != nil {
    log.Fatalf("HTTP shutdown error: %v", err)
  } 

  log.Println("Graceful shutdown complete.")
}
