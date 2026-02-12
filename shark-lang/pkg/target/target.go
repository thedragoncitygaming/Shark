// target.go

package target

// Target structure to hold target-related configurations

// Target represents the configuration for a specific architecture

type Target struct {
	Name string  // Name of the target architecture
	Triple string // LLVM target triple
}

// SupportedTargets holds all the supported architectures
var SupportedTargets = []Target{
	{Name: "amd64", Triple: "x86_64-unknown-linux-gnu"},
	{Name: "arm", Triple: "arm-unknown-linux-gnueabi"},
	{Name: "arm64", Triple: "aarch64-unknown-linux-gnu"},
	{Name: "386", Triple: "i386-unknown-linux-gnu"},
}

// A function to return the target triple for a given architecture
func GetTargetTriple(name string) string {
	for _, target := range SupportedTargets {
		if target.Name == name {
			return target.Triple
		}
	}
	return "unsupported target" // Return a message for unsupported architectures
}
