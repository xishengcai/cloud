package errcode

import "github.com/aagu/go-i18n/pkg/translation"

var (
	successMessage                   = translation.Message{ID: "success", Text: "request succeed"}
	serverErrorMessage               = translation.Message{ID: "server_error", Text: "internal server error"}
	invalidParamMessage              = translation.Message{ID: "invalid_param", Text: "invalid parameter: {{stringer .}}"}
	notfoundMessage                  = translation.Message{ID: "not_found", Text: "resource not found"}
	unsupported                      = translation.Message{ID: "unsupported", Text: "action unsupported"}
	tokenNotFoundMessage             = translation.Message{ID: "token_not_found", Text: "token not found in request"}
	tokenValidErrorMessage           = translation.Message{ID: "token_valid_error", Text: "token valid error: {{stringer .}}"}
	partialFailedMessage             = translation.Message{ID: "partial_failed", Text: "request partial failed"}
	badRequestMessage                = translation.Message{ID: "bad_request", Text: "unknown client error: {{stringer .}}"}
	notPresentMessage                = translation.Message{ID: "not_present", Text: "not present project"}
	noPermissionMessage              = translation.Message{ID: "no_permission", Text: "no permission to operate:{{stringer .}}"}
	duplicateNameMessage             = translation.Message{ID: "duplicate_name", Text: "duplicate {{stringer .}} name"}
	projectResourceOccupiedMessage   = translation.Message{ID: "resource_occupied", Text: "{{stringer .}}"}
	gitRepoMaxLimitExceed            = translation.Message{ID: "git_repo_max_limit_exceed", Text: "number of monitored git exceed max limit, limit {{.Limit}}, current {{.Current}}"}
	resourceNotFoundMessage          = translation.Message{ID: "resource_not_found", Text: "{{stringer .}} resource not found"}
	tapdNotBindMessage               = translation.Message{ID: "tapd_not_bind", Text: "bind tapd company not found"}
	environmentLimitMessage          = translation.Message{ID: "environment_limit", Text: "environment num limit ({{.}}/20)"}
	pipelineLimitMessage             = translation.Message{ID: "pipeline_limit", Text: "pipeline exist"}
	haveApprovedMessage              = translation.Message{ID: "have_approved", Text: "already approved"}
	conflict                         = translation.Message{ID: "conflict", Text: "resource already exist"}
	unavailable                      = translation.Message{ID: "unavailable", Text: "remote resource unavailable"}
	cicdEnvUnsupportedMessage        = translation.Message{ID: "cicd_env_unsupported", Text: "流水线环境不能删除"}
	delUpdatingEnvUnsupportedMessage = translation.Message{ID: "del_updating_env_unsupported", Text: "环境正在更新中,不能删除"}
	delDeletedEnvUnsupportedMessage  = translation.Message{ID: "del_deleted_env_unsupported", Text: "已提交删除环境审批,不能再次提交"}
	delDeletingEnvUnsupportedMessage = translation.Message{ID: "del_deleting_env_unsupported", Text: "环境已删除"}
	remoteServerError                = translation.Message{ID: "remote_server_error", Text: "remote server error"}
	clusterErrorMessage              = translation.Message{ID: "cluster_error", Text: "集群存在问题,请先修复"}
	unauthorizedMessage              = translation.Message{ID: "unauthorized", Text: "incorrect username or password"}
)

var (
	Success                   = NewError(0, &successMessage)
	PartialFailed             = NewError(00000007, &partialFailedMessage)
	ServerError               = NewError(10000000, &serverErrorMessage)
	InvalidParams             = NewError(10000001, &invalidParamMessage)
	NotFound                  = NewError(10000002, &notfoundMessage)
	TokenNotFound             = NewError(10000003, &tokenNotFoundMessage)
	TokenValidError           = NewError(10000004, &tokenValidErrorMessage)
	BadRequest                = NewError(10000008, &badRequestMessage)
	NotPresent                = NewError(10000009, &notPresentMessage)
	NoPermission              = NewError(10000010, &noPermissionMessage)
	DuplicateName             = NewError(10000011, &duplicateNameMessage)
	GitRepoMaxLimitExceed     = NewError(10000012, &gitRepoMaxLimitExceed)
	ProjectResourceOccupied   = NewError(10000013, &projectResourceOccupiedMessage)
	ResourceNotFound          = NewError(10000014, &resourceNotFoundMessage)
	TapdNotBind               = NewError(10000015, &tapdNotBindMessage)
	Unsupported               = NewError(10000016, &unsupported)
	EnvironmentLimit          = NewError(10000017, &environmentLimitMessage)
	PipelineLimit             = NewError(10000018, &pipelineLimitMessage)
	HaveApproved              = NewError(10000019, &haveApprovedMessage)
	Conflict                  = NewError(10000020, &conflict)
	Unavailable               = NewError(10000021, &unavailable)
	CicdEnvUnsupported        = NewError(10000022, &cicdEnvUnsupportedMessage)
	DelUpdatingEnvUnsupported = NewError(10000023, &delUpdatingEnvUnsupportedMessage)
	DelDeletedEnvUnsupported  = NewError(10000024, &delDeletedEnvUnsupportedMessage)
	DelDeletingEnvUnsupported = NewError(10000025, &delDeletingEnvUnsupportedMessage)
	RemoteServerError         = NewError(10000026, &remoteServerError)
	ClusterError              = NewError(10000027, &clusterErrorMessage)
	Unauthorized              = NewError(10000028, &unauthorizedMessage)
)
