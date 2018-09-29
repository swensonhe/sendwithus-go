package mock

import "github.com/swensonhe/sendwithus-go"

type Service struct {
	TemplatesFn func() ([]*swu.SWUTemplate, error)
	TemplatesInvoked bool

	EmailsFn func() ([]*swu.SWUTemplate, error)
	EmailsInvoked bool

	GetTemplateFn func(id string) (*swu.SWUTemplate, error)
	GetTemplateInvoked bool

	GetTemplateVersionFn func(id, version string) (*swu.SWUVersion, error)
	GetTemplateVersionInvoked bool

	UpdateTemplateVersionFn func(id, version string, template *swu.SWUVersion) (*swu.SWUVersion, error)
	UpdateTemplateVersionInvoked bool

	CreateTemplateFn func(template *swu.SWUVersion) (*swu.SWUTemplate, error)
	CreateTemplateInvoked bool

	CreateTemplateVersionFn func(id string, template *swu.SWUVersion) (*swu.SWUTemplate, error)
	CreateTemplateVersionInvoked bool

	SendFn func(email *swu.SWUEmail) error
	SendInvoked bool

	GetLogsFn func(q *swu.SWULogQuery) ([]*swu.SWULog, error)
	GetLogsInvoked bool

	GetLogFn func(id string) (*swu.SWULog, error)
	GetLogInvoked bool

	GetLogEventsFn func(id string) (*swu.SWULogEvent, error)
	GetLogEventsInvoked bool

	ResendLogFn func(id string) (*swu.SWULogResend, error)
	ResendLogInvoked bool
}

func NewService() *Service {
	return &Service{
		TemplatesFn: func() ([]*swu.SWUTemplate, error) {
			return []*swu.SWUTemplate{}, nil
		},
		EmailsFn: func() ([]*swu.SWUTemplate, error) {
			return []*swu.SWUTemplate{}, nil
		},
		GetTemplateFn: func(id string) (*swu.SWUTemplate, error) {
			return &swu.SWUTemplate{}, nil
		},
		GetTemplateVersionFn: func(id, version string) (*swu.SWUVersion, error) {
			return &swu.SWUVersion{}, nil
		},
		UpdateTemplateVersionFn: func(id, version string, template *swu.SWUVersion) (*swu.SWUVersion, error) {
			return &swu.SWUVersion{}, nil
		},
		CreateTemplateFn: func(template *swu.SWUVersion) (*swu.SWUTemplate, error) {
			return &swu.SWUTemplate{}, nil
		},
		CreateTemplateVersionFn: func(id string, template *swu.SWUVersion) (*swu.SWUTemplate, error) {
			return &swu.SWUTemplate{}, nil
		},
		SendFn: func(email *swu.SWUEmail) error {
			return nil
		},
		GetLogsFn: func(q *swu.SWULogQuery) ([]*swu.SWULog, error) {
			return []*swu.SWULog{}, nil
		},
		GetLogEventsFn: func(id string) (*swu.SWULogEvent, error) {
			return &swu.SWULogEvent{}, nil
		},
		ResendLogFn: func(id string) (*swu.SWULogResend, error) {
			return &swu.SWULogResend{}, nil
		},
	}
}

func (s *Service) Templates() ([]*swu.SWUTemplate, error) {
	s.TemplatesInvoked = true
	return s.TemplatesFn()
}

func (s *Service) Emails() ([]*swu.SWUTemplate, error) {
	s.EmailsInvoked = true
	return s.EmailsFn()
}

func (s *Service) GetTemplate(id string) (*swu.SWUTemplate, error) {
	s.GetTemplateInvoked = true
	return s.GetTemplateFn(id)
}

func (s *Service) GetTemplateVersion(id, version string) (*swu.SWUVersion, error) {
	s.GetTemplateVersionInvoked = true
	return s.GetTemplateVersionFn(id, version)
}

func (s *Service) UpdateTemplateVersion(id, version string, template *swu.SWUVersion) (*swu.SWUVersion, error) {
	s.UpdateTemplateVersionInvoked = true
	return s.UpdateTemplateVersionFn(id, version, template)
}

func (s *Service) CreateTemplate(template *swu.SWUVersion) (*swu.SWUTemplate, error) {
	s.CreateTemplateInvoked = true
	return s.CreateTemplateFn(template)
}

func (s *Service) CreateTemplateVersion(id string, template *swu.SWUVersion) (*swu.SWUTemplate, error) {
	s.CreateTemplateVersionInvoked = true
	return s.CreateTemplateVersionFn(id, template)
}

func (s *Service) Send(email *swu.SWUEmail) error {
	s.SendInvoked = true
	return s.SendFn(email)
}

func (s *Service) GetLogs(q *swu.SWULogQuery) ([]*swu.SWULog, error) {
	s.GetLogsInvoked = true
	return s.GetLogsFn(q)
}

func (s *Service) GetLog(id string) (*swu.SWULog, error) {
	s.GetLogInvoked = true
	return s.GetLogFn(id)
}

func (s *Service) GetLogEvents(id string) (*swu.SWULogEvent, error) {
	s.GetLogEventsInvoked = true
	return s.GetLogEventsFn(id)
}

func (s *Service) ResendLog(id string) (*swu.SWULogResend, error) {
	s.ResendLogInvoked = true
	return s.ResendLogFn(id)
}
