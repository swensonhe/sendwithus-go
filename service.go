package swu

// Service provides an interface for the Sendwithus service.
// This allows for mocking.
type Service interface {
	Templates() ([]*SWUTemplate, error)
	Emails() ([]*SWUTemplate, error)
	GetTemplate(id string) (*SWUTemplate, error)
	GetTemplateVersion(id, version string) (*SWUVersion, error)
	UpdateTemplateVersion(id, version string, template *SWUVersion) (*SWUVersion, error)
	CreateTemplate(template *SWUVersion) (*SWUTemplate, error)
	CreateTemplateVersion(id string, template *SWUVersion) (*SWUTemplate, error)
	Send(email *SWUEmail) error
	GetLogs(q *SWULogQuery) ([]*SWULog, error)
	GetLog(id string) (*SWULog, error)
	GetLogEvents(id string) (*SWULogEvent, error)
	ResendLog(id string) (*SWULogResend, error)
}
