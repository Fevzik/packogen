package constants

const (
	InternalDirectory   = "internal"
	ServiceDirectory    = "service"
	HandlerDirectory    = "handler"
	RepositoryDirectory = "repository"
	ConstantsDirectory  = "constants"
	ErrorsDirectory     = "errors"
	DomainDirectory     = "domain"
	FormDirectory       = "form"
	QueueDirectory      = "queue"
	ResponseDirectory   = "response"
	GrpcDirectory       = "grpc"
	CmdDirectory        = "cmd"
)

var (
	InternalDirectories = []string{
		ServiceDirectory,
		HandlerDirectory,
		RepositoryDirectory,
		ConstantsDirectory,
		ErrorsDirectory,
		DomainDirectory,
		FormDirectory,
		QueueDirectory,
		ResponseDirectory,
		GrpcDirectory,
	}
)
