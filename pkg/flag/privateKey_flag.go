package flag

import (
	"flag"
	"fmt"
)

//PrivateKeyFlag ecdsa private key, you need set this with private key
//wallet file or account file
//wallet
//file path
//private key hex string
type PrivateKeyFlag struct {
	Name        string
	Aliases     []string
	Usage       string
	EnvVars     []string
	FilePath    string
	Required    bool
	Hidden      bool
	TakesFile   bool
	Value       privateValue
	DefaultText string
	HasBeenSet  bool
}

// IsSet returns whether or not the flag has been set through env or file
func (p *PrivateKeyFlag) IsSet() bool {
	return p.HasBeenSet
}

// String returns a readable representation of this value
// (for usage defaults)
func (p *PrivateKeyFlag) String() string {
	fmt.Println(p)
	return p.Value.String()
}

// Names returns the names of the flag
func (p *PrivateKeyFlag) Names() []string {
	return append([]string{p.Name}, p.Aliases...)
}

// IsRequired returns whether or not the flag is required
func (p *PrivateKeyFlag) IsRequired() bool {
	return p.Required
}

// Apply populates the flag given the flag set and environment
func (p *PrivateKeyFlag) Apply(set *flag.FlagSet) error {
	if val, ok := flagFromEnvOrFile(p.EnvVars, p.FilePath); ok {
		err := p.Value.Set(val)
		if err != nil {
			return err
		}
		p.HasBeenSet = true
	}
	for _, name := range p.Names() {
		pV := NewPrivateKeyValue()
		*pV = p.Value
		set.Var(pV, name, p.Usage)
	}

	return nil
}
