package errs

import "errors"

var (
	UnknownGroup = errors.New("unknown group")
	UnknownRepo  = errors.New("unknown repo")

	UnsupportedPlatform = errors.New("unsupported platform")

	CloneRepo    = errors.New("clone repo")
	CheckoutRepo = errors.New("checkout repo")
	BranchRepo   = errors.New("branch repo")
	PullRepo     = errors.New("pull repo")

	PreBuild     = errors.New("pre-build")
	LibBuild     = errors.New("lib ubild")
	ServiceBuild = errors.New("service build")
	ServiceRun   = errors.New("service run")
	ServiceStop  = errors.New("service stop")

	InvalidCommand    = errors.New("invalid command")
	InvalidLib        = errors.New("invalid lib")
	InvalidService    = errors.New("invalid service")
	ReadingConfigData = errors.New("reading file with config data")
	InvalidUrl        = errors.New("invalid url")
	InvalidPath       = errors.New("invalid path")
)
