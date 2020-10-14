// Package resources provide common function and interface for cloud resources
package resources

// Cloud interface for structure cloud data
type Cloud interface {
	ReportTitle() string
	String() string
}
