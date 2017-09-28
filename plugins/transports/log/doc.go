/*

Package log provides a transport to log information with alerts in JSON.

The plugin takes there construction parameters:

  Name      Type    Default  Description
  lebel     string  ""       Arbitrary string label.
  severity  string  "info"   Log lebel.

Example snippet for TOML configuration:

  [[route.notify]]
    type     = "log"
    label    = "log transport"
    severity = "info"

Available `severity` types are "critical", "error", "warn", "info", and "debug".
For details, see https://godoc.org/github.com/cybozu-go/log#pkg-constants
*/
package log
