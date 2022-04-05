// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
	"net"
	"net/http"
	"net/netip"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	"go.opentelemetry.io/otel/metric/nonrecording"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// No-op definition for keeping imports.
var (
	_ = bytes.NewReader
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = io.Copy
	_ = math.Mod
	_ = big.Rat{}
	_ = bits.LeadingZeros64
	_ = net.IP{}
	_ = http.MethodGet
	_ = netip.Addr{}
	_ = url.URL{}
	_ = regexp.MustCompile
	_ = sort.Ints
	_ = strconv.ParseInt
	_ = strings.Builder{}
	_ = sync.Pool{}
	_ = time.Time{}

	_ = errors.Is
	_ = jx.Null
	_ = uuid.UUID{}
	_ = otel.GetTracerProvider
	_ = attribute.KeyValue{}
	_ = codes.Unset
	_ = metric.MeterConfig{}
	_ = syncint64.Counter(nil)
	_ = nonrecording.NewNoopMeterProvider
	_ = trace.TraceIDFromHex

	_ = conv.ToInt32
	_ = ht.NewRequest
	_ = json.Marshal
	_ = otelogen.Version
	_ = uri.PathEncoder{}
	_ = validate.Int{}
)

// setDefaults set default value of fields.
func (s *ActionsCreateSelfHostedRunnerGroupForOrgReq) setDefaults() {
	{
		val := ActionsCreateSelfHostedRunnerGroupForOrgReqVisibility("all")

		s.Visibility.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ActivitySetThreadSubscriptionReq) setDefaults() {
	{
		val := bool(false)

		s.Ignored.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ChecksSetSuitesPreferencesReqAutoTriggerChecksItem) setDefaults() {
	{
		val := bool(true)

		s.Setting = val
	}
}

// setDefaults set default value of fields.
func (s *DeploymentStatus) setDefaults() {
	{
		val := string("")

		s.Description = val
	}
	{
		val := string("")

		s.Environment.SetTo(val)
	}
	{
		val, _ := json.DecodeURI(jx.DecodeStr("\"\""))

		s.TargetURL = val
	}
	{
		val, _ := json.DecodeURI(jx.DecodeStr("\"\""))

		s.EnvironmentURL.SetTo(val)
	}
	{
		val, _ := json.DecodeURI(jx.DecodeStr("\"\""))

		s.LogURL.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseReq) setDefaults() {
	{
		val := EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseReqVisibility("all")

		s.Visibility.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *FullRepository) setDefaults() {
	{
		val := bool(true)

		s.AnonymousAccessEnabled.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *GitCreateBlobReq) setDefaults() {
	{
		val := string("utf-8")

		s.Encoding.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *GitUpdateRefReq) setDefaults() {
	{
		val := bool(false)

		s.Force.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *IssuesCreateMilestoneReq) setDefaults() {
	{
		val := IssuesCreateMilestoneReqState("open")

		s.State.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *IssuesUpdateMilestoneReq) setDefaults() {
	{
		val := IssuesUpdateMilestoneReqState("open")

		s.State.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *MigrationsStartForOrgReq) setDefaults() {
	{
		val := bool(false)

		s.LockRepositories.SetTo(val)
	}
	{
		val := bool(false)

		s.ExcludeAttachments.SetTo(val)
	}
	{
		val := bool(false)

		s.ExcludeReleases.SetTo(val)
	}
	{
		val := bool(false)

		s.ExcludeOwnerProjects.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *Milestone) setDefaults() {
	{
		val := MilestoneState("open")

		s.State = val
	}
}

// setDefaults set default value of fields.
func (s *NullableMilestone) setDefaults() {
	{
		val := NullableMilestoneState("open")

		s.State = val
	}
}

// setDefaults set default value of fields.
func (s *NullableRepository) setDefaults() {
	{
		val := bool(false)

		s.Private = val
	}
	{
		val := bool(false)

		s.IsTemplate.SetTo(val)
	}
	{
		val := bool(true)

		s.HasIssues = val
	}
	{
		val := bool(true)

		s.HasProjects = val
	}
	{
		val := bool(true)

		s.HasWiki = val
	}
	{
		val := bool(true)

		s.HasDownloads = val
	}
	{
		val := bool(false)

		s.Archived = val
	}
	{
		val := string("public")

		s.Visibility.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowRebaseMerge.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowSquashMerge.SetTo(val)
	}
	{
		val := bool(false)

		s.AllowAutoMerge.SetTo(val)
	}
	{
		val := bool(false)

		s.DeleteBranchOnMerge.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowMergeCommit.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *OrgsCreateInvitationReq) setDefaults() {
	{
		val := OrgsCreateInvitationReqRole("direct_member")

		s.Role.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *OrgsCreateWebhookReq) setDefaults() {
	{
		s.Events = make([]string, 1)
		{
			val := string("push")

			s.Events[0] = val
		}
	}
	{
		val := bool(true)

		s.Active.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *OrgsSetMembershipForUserReq) setDefaults() {
	{
		val := OrgsSetMembershipForUserReqRole("member")

		s.Role.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *OrgsUpdateWebhookReq) setDefaults() {
	{
		s.Events = make([]string, 1)
		{
			val := string("push")

			s.Events[0] = val
		}
	}
	{
		val := bool(true)

		s.Active.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *Page) setDefaults() {
	{
		val := bool(false)

		s.Custom404 = val
	}
}

// setDefaults set default value of fields.
func (s *ProjectsAddCollaboratorReq) setDefaults() {
	{
		val := ProjectsAddCollaboratorReqPermission("write")

		s.Permission.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *PullRequestReviewComment) setDefaults() {
	{
		val := PullRequestReviewCommentStartSide("RIGHT")

		s.StartSide.SetTo(val)
	}
	{
		val := PullRequestReviewCommentSide("RIGHT")

		s.Side.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReposAddCollaboratorReq) setDefaults() {
	{
		val := ReposAddCollaboratorReqPermission("push")

		s.Permission.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReposCreateCommitStatusReq) setDefaults() {
	{
		val := string("default")

		s.Context.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReposCreateDeploymentReq) setDefaults() {
	{
		val := string("deploy")

		s.Task.SetTo(val)
	}
	{
		val := bool(true)

		s.AutoMerge.SetTo(val)
	}
	{
		val := string("production")

		s.Environment.SetTo(val)
	}
	{
		val := string("")

		s.Description.SetTo(val)
	}
	{
		val := bool(false)

		s.TransientEnvironment.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReposCreateDeploymentStatusReq) setDefaults() {
	{
		val := string("")

		s.TargetURL.SetTo(val)
	}
	{
		val := string("")

		s.LogURL.SetTo(val)
	}
	{
		val := string("")

		s.Description.SetTo(val)
	}
	{
		val := string("")

		s.EnvironmentURL.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReposCreateForAuthenticatedUserReq) setDefaults() {
	{
		val := bool(false)

		s.Private.SetTo(val)
	}
	{
		val := bool(true)

		s.HasIssues.SetTo(val)
	}
	{
		val := bool(true)

		s.HasProjects.SetTo(val)
	}
	{
		val := bool(true)

		s.HasWiki.SetTo(val)
	}
	{
		val := bool(false)

		s.AutoInit.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowSquashMerge.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowMergeCommit.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowRebaseMerge.SetTo(val)
	}
	{
		val := bool(false)

		s.AllowAutoMerge.SetTo(val)
	}
	{
		val := bool(false)

		s.DeleteBranchOnMerge.SetTo(val)
	}
	{
		val := bool(true)

		s.HasDownloads.SetTo(val)
	}
	{
		val := bool(false)

		s.IsTemplate.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReposCreateInOrgReq) setDefaults() {
	{
		val := bool(false)

		s.Private.SetTo(val)
	}
	{
		val := bool(true)

		s.HasIssues.SetTo(val)
	}
	{
		val := bool(true)

		s.HasProjects.SetTo(val)
	}
	{
		val := bool(true)

		s.HasWiki.SetTo(val)
	}
	{
		val := bool(false)

		s.IsTemplate.SetTo(val)
	}
	{
		val := bool(false)

		s.AutoInit.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowSquashMerge.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowMergeCommit.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowRebaseMerge.SetTo(val)
	}
	{
		val := bool(false)

		s.AllowAutoMerge.SetTo(val)
	}
	{
		val := bool(false)

		s.DeleteBranchOnMerge.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReposCreatePagesSiteReqSource) setDefaults() {
	{
		val := ReposCreatePagesSiteReqSourcePath("/")

		s.Path.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReposCreateReleaseReq) setDefaults() {
	{
		val := bool(false)

		s.Draft.SetTo(val)
	}
	{
		val := bool(false)

		s.Prerelease.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReposCreateUsingTemplateReq) setDefaults() {
	{
		val := bool(false)

		s.IncludeAllBranches.SetTo(val)
	}
	{
		val := bool(false)

		s.Private.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReposCreateWebhookReq) setDefaults() {
	{
		s.Events = make([]string, 1)
		{
			val := string("push")

			s.Events[0] = val
		}
	}
	{
		val := bool(true)

		s.Active.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReposUpdateReq) setDefaults() {
	{
		val := bool(false)

		s.Private.SetTo(val)
	}
	{
		val := bool(true)

		s.HasIssues.SetTo(val)
	}
	{
		val := bool(true)

		s.HasProjects.SetTo(val)
	}
	{
		val := bool(true)

		s.HasWiki.SetTo(val)
	}
	{
		val := bool(false)

		s.IsTemplate.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowSquashMerge.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowMergeCommit.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowRebaseMerge.SetTo(val)
	}
	{
		val := bool(false)

		s.AllowAutoMerge.SetTo(val)
	}
	{
		val := bool(false)

		s.DeleteBranchOnMerge.SetTo(val)
	}
	{
		val := bool(false)

		s.Archived.SetTo(val)
	}
	{
		val := bool(false)

		s.AllowForking.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReposUpdateWebhookReq) setDefaults() {
	{
		s.Events = make([]string, 1)
		{
			val := string("push")

			s.Events[0] = val
		}
	}
	{
		val := bool(true)

		s.Active.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *Repository) setDefaults() {
	{
		val := bool(false)

		s.Private = val
	}
	{
		val := bool(false)

		s.IsTemplate.SetTo(val)
	}
	{
		val := bool(true)

		s.HasIssues = val
	}
	{
		val := bool(true)

		s.HasProjects = val
	}
	{
		val := bool(true)

		s.HasWiki = val
	}
	{
		val := bool(true)

		s.HasDownloads = val
	}
	{
		val := bool(false)

		s.Archived = val
	}
	{
		val := string("public")

		s.Visibility.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowRebaseMerge.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowSquashMerge.SetTo(val)
	}
	{
		val := bool(false)

		s.AllowAutoMerge.SetTo(val)
	}
	{
		val := bool(false)

		s.DeleteBranchOnMerge.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowMergeCommit.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReviewComment) setDefaults() {
	{
		val := ReviewCommentSide("RIGHT")

		s.Side.SetTo(val)
	}
	{
		val := ReviewCommentStartSide("RIGHT")

		s.StartSide.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TeamMembership) setDefaults() {
	{
		val := TeamMembershipRole("member")

		s.Role = val
	}
}

// setDefaults set default value of fields.
func (s *TeamRepository) setDefaults() {
	{
		val := bool(false)

		s.Private = val
	}
	{
		val := bool(false)

		s.IsTemplate.SetTo(val)
	}
	{
		val := bool(true)

		s.HasIssues = val
	}
	{
		val := bool(true)

		s.HasProjects = val
	}
	{
		val := bool(true)

		s.HasWiki = val
	}
	{
		val := bool(true)

		s.HasDownloads = val
	}
	{
		val := bool(false)

		s.Archived = val
	}
	{
		val := string("public")

		s.Visibility.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowRebaseMerge.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowSquashMerge.SetTo(val)
	}
	{
		val := bool(false)

		s.AllowAutoMerge.SetTo(val)
	}
	{
		val := bool(false)

		s.DeleteBranchOnMerge.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowMergeCommit.SetTo(val)
	}
	{
		val := bool(false)

		s.AllowForking.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TeamsAddOrUpdateMembershipForUserInOrgReq) setDefaults() {
	{
		val := TeamsAddOrUpdateMembershipForUserInOrgReqRole("member")

		s.Role.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TeamsAddOrUpdateMembershipForUserLegacyReq) setDefaults() {
	{
		val := TeamsAddOrUpdateMembershipForUserLegacyReqRole("member")

		s.Role.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TeamsCreateDiscussionInOrgReq) setDefaults() {
	{
		val := bool(false)

		s.Private.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TeamsCreateDiscussionLegacyReq) setDefaults() {
	{
		val := bool(false)

		s.Private.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TeamsCreateReq) setDefaults() {
	{
		val := TeamsCreateReqPermission("pull")

		s.Permission.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TeamsUpdateInOrgReq) setDefaults() {
	{
		val := TeamsUpdateInOrgReqPermission("pull")

		s.Permission.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TeamsUpdateLegacyReq) setDefaults() {
	{
		val := TeamsUpdateLegacyReqPermission("pull")

		s.Permission.SetTo(val)
	}
}
