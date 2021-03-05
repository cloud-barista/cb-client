package main

import (
        "log"

        "github.com/cloud-barista/cb-tumblebug/src/api/grpc/cbadm/cmd"
)

// ===== [ Constants and Variables ] =====

// ===== [ Types ] =====

// ===== [ Implementations ] =====

// ===== [ Private Functions ] =====

// main - Entrypoint
func main() {
        rootCmd := cmd.NewRootCmd()
        if err := rootCmd.Execute(); err != nil {
                log.Println("cbctl terminated with error: ", err.Error())
        }
}
