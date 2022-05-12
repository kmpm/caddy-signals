package signals

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
)

func init() {
	httpcaddyfile.RegisterGlobalOption("signals", parseGlobalOption)
}

// Gizmo is an example; put your own type here.
type Gizmo struct {
	Sighup string `json:"sighup,omitempty"`
}

func parseGlobalOption(d *caddyfile.Dispenser, existingVal interface{}) (interface{}, error) {
	fmt.Println("parseCaddyfile: existingVal=", existingVal)
	var m Gizmo
	err := m.UnmarshalCaddyfile(d)
	fmt.Println("unmarshalled", m)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP)
	go func() {
		fmt.Println("awaiting signal")
		sig := <-sigs
		fmt.Println("sig", sig)
	}()
	return m, err
}

// // UnmarshalCaddyfile implements caddyfile.Unmarshaler.
func (m *Gizmo) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	// fmt.Println("Unmarshal", d)
	for d.Next() {
		switch d.Val() {
		case "sigusr1":
			if !d.Args(&m.Sighup) {
				return d.ArgErr()
			}
		}
		// if !d.Args(&m.SigUSR1) {
		// 	return d.ArgErr()
		// }
	}
	return nil
}

// Interface guards
var (
	// _ caddy.Provisioner     = (*Gizmo)(nil)
	_ caddyfile.Unmarshaler = (*Gizmo)(nil)
)
