// arch.go

package arch

import ( 
    "os" 
    "runtime" 
)

// GetArchitecture returns the target architecture and operating system.
func GetArchitecture() (string, string) {
    arch := runtime.GOARCH  // target architecture
    os := runtime.GOOS      // target operating system
    return arch, os
}

// You can use this function during compilation to detect the target architecture and OS automatically.
