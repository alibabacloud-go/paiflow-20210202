// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BriefPipelineRun struct {
	Accessibility   *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Duration        *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FinishedAt      *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NodeId          *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	ParentUserId    *string `json:"ParentUserId,omitempty" xml:"ParentUserId,omitempty"`
	PipelineId      *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	PipelineRunId   *string `json:"PipelineRunId,omitempty" xml:"PipelineRunId,omitempty"`
	SourceId        *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType      *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	StartedAt       *string `json:"StartedAt,omitempty" xml:"StartedAt,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s BriefPipelineRun) String() string {
	return tea.Prettify(s)
}

func (s BriefPipelineRun) GoString() string {
	return s.String()
}

func (s *BriefPipelineRun) SetAccessibility(v string) *BriefPipelineRun {
	s.Accessibility = &v
	return s
}

func (s *BriefPipelineRun) SetDuration(v int64) *BriefPipelineRun {
	s.Duration = &v
	return s
}

func (s *BriefPipelineRun) SetFinishedAt(v string) *BriefPipelineRun {
	s.FinishedAt = &v
	return s
}

func (s *BriefPipelineRun) SetGmtCreateTime(v string) *BriefPipelineRun {
	s.GmtCreateTime = &v
	return s
}

func (s *BriefPipelineRun) SetGmtModifiedTime(v string) *BriefPipelineRun {
	s.GmtModifiedTime = &v
	return s
}

func (s *BriefPipelineRun) SetMessage(v string) *BriefPipelineRun {
	s.Message = &v
	return s
}

func (s *BriefPipelineRun) SetName(v string) *BriefPipelineRun {
	s.Name = &v
	return s
}

func (s *BriefPipelineRun) SetNodeId(v string) *BriefPipelineRun {
	s.NodeId = &v
	return s
}

func (s *BriefPipelineRun) SetParentUserId(v string) *BriefPipelineRun {
	s.ParentUserId = &v
	return s
}

func (s *BriefPipelineRun) SetPipelineId(v string) *BriefPipelineRun {
	s.PipelineId = &v
	return s
}

func (s *BriefPipelineRun) SetPipelineRunId(v string) *BriefPipelineRun {
	s.PipelineRunId = &v
	return s
}

func (s *BriefPipelineRun) SetSourceId(v string) *BriefPipelineRun {
	s.SourceId = &v
	return s
}

func (s *BriefPipelineRun) SetSourceType(v string) *BriefPipelineRun {
	s.SourceType = &v
	return s
}

func (s *BriefPipelineRun) SetStartedAt(v string) *BriefPipelineRun {
	s.StartedAt = &v
	return s
}

func (s *BriefPipelineRun) SetStatus(v string) *BriefPipelineRun {
	s.Status = &v
	return s
}

func (s *BriefPipelineRun) SetUserId(v string) *BriefPipelineRun {
	s.UserId = &v
	return s
}

func (s *BriefPipelineRun) SetWorkspaceId(v string) *BriefPipelineRun {
	s.WorkspaceId = &v
	return s
}

type FullPipelineRun struct {
	Accessibility   *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Arguments       *string `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
	Duration        *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FinishedAt      *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Manifest        *string `json:"Manifest,omitempty" xml:"Manifest,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NodeId          *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Options         *string `json:"Options,omitempty" xml:"Options,omitempty"`
	ParentUserId    *string `json:"ParentUserId,omitempty" xml:"ParentUserId,omitempty"`
	PipelineId      *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	PipelineRunId   *string `json:"PipelineRunId,omitempty" xml:"PipelineRunId,omitempty"`
	SourceId        *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType      *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	StartedAt       *string `json:"StartedAt,omitempty" xml:"StartedAt,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s FullPipelineRun) String() string {
	return tea.Prettify(s)
}

func (s FullPipelineRun) GoString() string {
	return s.String()
}

func (s *FullPipelineRun) SetAccessibility(v string) *FullPipelineRun {
	s.Accessibility = &v
	return s
}

func (s *FullPipelineRun) SetArguments(v string) *FullPipelineRun {
	s.Arguments = &v
	return s
}

func (s *FullPipelineRun) SetDuration(v int64) *FullPipelineRun {
	s.Duration = &v
	return s
}

func (s *FullPipelineRun) SetFinishedAt(v string) *FullPipelineRun {
	s.FinishedAt = &v
	return s
}

func (s *FullPipelineRun) SetGmtCreateTime(v string) *FullPipelineRun {
	s.GmtCreateTime = &v
	return s
}

func (s *FullPipelineRun) SetGmtModifiedTime(v string) *FullPipelineRun {
	s.GmtModifiedTime = &v
	return s
}

func (s *FullPipelineRun) SetManifest(v string) *FullPipelineRun {
	s.Manifest = &v
	return s
}

func (s *FullPipelineRun) SetMessage(v string) *FullPipelineRun {
	s.Message = &v
	return s
}

func (s *FullPipelineRun) SetName(v string) *FullPipelineRun {
	s.Name = &v
	return s
}

func (s *FullPipelineRun) SetNodeId(v string) *FullPipelineRun {
	s.NodeId = &v
	return s
}

func (s *FullPipelineRun) SetOptions(v string) *FullPipelineRun {
	s.Options = &v
	return s
}

func (s *FullPipelineRun) SetParentUserId(v string) *FullPipelineRun {
	s.ParentUserId = &v
	return s
}

func (s *FullPipelineRun) SetPipelineId(v string) *FullPipelineRun {
	s.PipelineId = &v
	return s
}

func (s *FullPipelineRun) SetPipelineRunId(v string) *FullPipelineRun {
	s.PipelineRunId = &v
	return s
}

func (s *FullPipelineRun) SetSourceId(v string) *FullPipelineRun {
	s.SourceId = &v
	return s
}

func (s *FullPipelineRun) SetSourceType(v string) *FullPipelineRun {
	s.SourceType = &v
	return s
}

func (s *FullPipelineRun) SetStartedAt(v string) *FullPipelineRun {
	s.StartedAt = &v
	return s
}

func (s *FullPipelineRun) SetStatus(v string) *FullPipelineRun {
	s.Status = &v
	return s
}

func (s *FullPipelineRun) SetUserId(v string) *FullPipelineRun {
	s.UserId = &v
	return s
}

func (s *FullPipelineRun) SetWorkspaceId(v string) *FullPipelineRun {
	s.WorkspaceId = &v
	return s
}

type Node struct {
	ApiVersion *string         `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	Metadata   *NodeMetadata   `json:"Metadata,omitempty" xml:"Metadata,omitempty" type:"Struct"`
	Spec       *NodeSpec       `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
	StatusInfo *NodeStatusInfo `json:"StatusInfo,omitempty" xml:"StatusInfo,omitempty" type:"Struct"`
}

func (s Node) String() string {
	return tea.Prettify(s)
}

func (s Node) GoString() string {
	return s.String()
}

func (s *Node) SetApiVersion(v string) *Node {
	s.ApiVersion = &v
	return s
}

func (s *Node) SetMetadata(v *NodeMetadata) *Node {
	s.Metadata = v
	return s
}

func (s *Node) SetSpec(v *NodeSpec) *Node {
	s.Spec = v
	return s
}

func (s *Node) SetStatusInfo(v *NodeStatusInfo) *Node {
	s.StatusInfo = v
	return s
}

type NodeMetadata struct {
	DisplayName    *string   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Identifier     *string   `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	Name           *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NodeId         *string   `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeType       *string   `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Provider       *string   `json:"Provider,omitempty" xml:"Provider,omitempty"`
	RelatedNodeIds []*string `json:"RelatedNodeIds,omitempty" xml:"RelatedNodeIds,omitempty" type:"Repeated"`
	Version        *string   `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s NodeMetadata) String() string {
	return tea.Prettify(s)
}

func (s NodeMetadata) GoString() string {
	return s.String()
}

func (s *NodeMetadata) SetDisplayName(v string) *NodeMetadata {
	s.DisplayName = &v
	return s
}

func (s *NodeMetadata) SetIdentifier(v string) *NodeMetadata {
	s.Identifier = &v
	return s
}

func (s *NodeMetadata) SetName(v string) *NodeMetadata {
	s.Name = &v
	return s
}

func (s *NodeMetadata) SetNodeId(v string) *NodeMetadata {
	s.NodeId = &v
	return s
}

func (s *NodeMetadata) SetNodeType(v string) *NodeMetadata {
	s.NodeType = &v
	return s
}

func (s *NodeMetadata) SetProvider(v string) *NodeMetadata {
	s.Provider = &v
	return s
}

func (s *NodeMetadata) SetRelatedNodeIds(v []*string) *NodeMetadata {
	s.RelatedNodeIds = v
	return s
}

func (s *NodeMetadata) SetVersion(v string) *NodeMetadata {
	s.Version = &v
	return s
}

type NodeSpec struct {
	Dependencies []*string             `json:"Dependencies,omitempty" xml:"Dependencies,omitempty" type:"Repeated"`
	HasPipelines *bool                 `json:"HasPipelines,omitempty" xml:"HasPipelines,omitempty"`
	Inputs       *NodeIO               `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	Outputs      *NodeIO               `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	Parallelism  *int64                `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	Pipelines    []*Node               `json:"Pipelines,omitempty" xml:"Pipelines,omitempty" type:"Repeated"`
	When         *string               `json:"When,omitempty" xml:"When,omitempty"`
	WithItems    []*string             `json:"WithItems,omitempty" xml:"WithItems,omitempty" type:"Repeated"`
	WithParam    *string               `json:"WithParam,omitempty" xml:"WithParam,omitempty"`
	WithSequence *NodeSpecWithSequence `json:"WithSequence,omitempty" xml:"WithSequence,omitempty" type:"Struct"`
}

func (s NodeSpec) String() string {
	return tea.Prettify(s)
}

func (s NodeSpec) GoString() string {
	return s.String()
}

func (s *NodeSpec) SetDependencies(v []*string) *NodeSpec {
	s.Dependencies = v
	return s
}

func (s *NodeSpec) SetHasPipelines(v bool) *NodeSpec {
	s.HasPipelines = &v
	return s
}

func (s *NodeSpec) SetInputs(v *NodeIO) *NodeSpec {
	s.Inputs = v
	return s
}

func (s *NodeSpec) SetOutputs(v *NodeIO) *NodeSpec {
	s.Outputs = v
	return s
}

func (s *NodeSpec) SetParallelism(v int64) *NodeSpec {
	s.Parallelism = &v
	return s
}

func (s *NodeSpec) SetPipelines(v []*Node) *NodeSpec {
	s.Pipelines = v
	return s
}

func (s *NodeSpec) SetWhen(v string) *NodeSpec {
	s.When = &v
	return s
}

func (s *NodeSpec) SetWithItems(v []*string) *NodeSpec {
	s.WithItems = v
	return s
}

func (s *NodeSpec) SetWithParam(v string) *NodeSpec {
	s.WithParam = &v
	return s
}

func (s *NodeSpec) SetWithSequence(v *NodeSpecWithSequence) *NodeSpec {
	s.WithSequence = v
	return s
}

type NodeSpecWithSequence struct {
	End    *int64  `json:"End,omitempty" xml:"End,omitempty"`
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	Start  *int64  `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s NodeSpecWithSequence) String() string {
	return tea.Prettify(s)
}

func (s NodeSpecWithSequence) GoString() string {
	return s.String()
}

func (s *NodeSpecWithSequence) SetEnd(v int64) *NodeSpecWithSequence {
	s.End = &v
	return s
}

func (s *NodeSpecWithSequence) SetFormat(v string) *NodeSpecWithSequence {
	s.Format = &v
	return s
}

func (s *NodeSpecWithSequence) SetStart(v int64) *NodeSpecWithSequence {
	s.Start = &v
	return s
}

type NodeStatusInfo struct {
	Conditions []*NodeStatusInfoConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	FinishedAt *string                     `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	Progress   *string                     `json:"Progress,omitempty" xml:"Progress,omitempty"`
	StartedAt  *string                     `json:"StartedAt,omitempty" xml:"StartedAt,omitempty"`
	Status     *string                     `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s NodeStatusInfo) String() string {
	return tea.Prettify(s)
}

func (s NodeStatusInfo) GoString() string {
	return s.String()
}

func (s *NodeStatusInfo) SetConditions(v []*NodeStatusInfoConditions) *NodeStatusInfo {
	s.Conditions = v
	return s
}

func (s *NodeStatusInfo) SetFinishedAt(v string) *NodeStatusInfo {
	s.FinishedAt = &v
	return s
}

func (s *NodeStatusInfo) SetProgress(v string) *NodeStatusInfo {
	s.Progress = &v
	return s
}

func (s *NodeStatusInfo) SetStartedAt(v string) *NodeStatusInfo {
	s.StartedAt = &v
	return s
}

func (s *NodeStatusInfo) SetStatus(v string) *NodeStatusInfo {
	s.Status = &v
	return s
}

type NodeStatusInfoConditions struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s NodeStatusInfoConditions) String() string {
	return tea.Prettify(s)
}

func (s NodeStatusInfoConditions) GoString() string {
	return s.String()
}

func (s *NodeStatusInfoConditions) SetStatus(v string) *NodeStatusInfoConditions {
	s.Status = &v
	return s
}

func (s *NodeStatusInfoConditions) SetType(v string) *NodeStatusInfoConditions {
	s.Type = &v
	return s
}

type NodeIO struct {
	Artifacts  []map[string]interface{} `json:"Artifacts,omitempty" xml:"Artifacts,omitempty" type:"Repeated"`
	Parameters []map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s NodeIO) String() string {
	return tea.Prettify(s)
}

func (s NodeIO) GoString() string {
	return s.String()
}

func (s *NodeIO) SetArtifacts(v []map[string]interface{}) *NodeIO {
	s.Artifacts = v
	return s
}

func (s *NodeIO) SetParameters(v []map[string]interface{}) *NodeIO {
	s.Parameters = v
	return s
}

type Pipeline struct {
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Identifier      *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	PipelineId      *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	Provider        *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	Uuid            *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Pipeline) String() string {
	return tea.Prettify(s)
}

func (s Pipeline) GoString() string {
	return s.String()
}

func (s *Pipeline) SetGmtCreateTime(v string) *Pipeline {
	s.GmtCreateTime = &v
	return s
}

func (s *Pipeline) SetGmtModifiedTime(v string) *Pipeline {
	s.GmtModifiedTime = &v
	return s
}

func (s *Pipeline) SetIdentifier(v string) *Pipeline {
	s.Identifier = &v
	return s
}

func (s *Pipeline) SetPipelineId(v string) *Pipeline {
	s.PipelineId = &v
	return s
}

func (s *Pipeline) SetProvider(v string) *Pipeline {
	s.Provider = &v
	return s
}

func (s *Pipeline) SetUuid(v string) *Pipeline {
	s.Uuid = &v
	return s
}

func (s *Pipeline) SetVersion(v string) *Pipeline {
	s.Version = &v
	return s
}

func (s *Pipeline) SetWorkspaceId(v string) *Pipeline {
	s.WorkspaceId = &v
	return s
}

type Run struct {
	Accessibility   *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Duration        *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ExperimentId    *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	FinishedAt      *int64  `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NodeId          *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	ParentUserId    *string `json:"ParentUserId,omitempty" xml:"ParentUserId,omitempty"`
	RunId           *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	Source          *string `json:"Source,omitempty" xml:"Source,omitempty"`
	StartedAt       *int64  `json:"StartedAt,omitempty" xml:"StartedAt,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Run) String() string {
	return tea.Prettify(s)
}

func (s Run) GoString() string {
	return s.String()
}

func (s *Run) SetAccessibility(v string) *Run {
	s.Accessibility = &v
	return s
}

func (s *Run) SetDuration(v int64) *Run {
	s.Duration = &v
	return s
}

func (s *Run) SetExperimentId(v string) *Run {
	s.ExperimentId = &v
	return s
}

func (s *Run) SetFinishedAt(v int64) *Run {
	s.FinishedAt = &v
	return s
}

func (s *Run) SetGmtCreateTime(v string) *Run {
	s.GmtCreateTime = &v
	return s
}

func (s *Run) SetGmtModifiedTime(v string) *Run {
	s.GmtModifiedTime = &v
	return s
}

func (s *Run) SetMessage(v string) *Run {
	s.Message = &v
	return s
}

func (s *Run) SetName(v string) *Run {
	s.Name = &v
	return s
}

func (s *Run) SetNodeId(v string) *Run {
	s.NodeId = &v
	return s
}

func (s *Run) SetParentUserId(v string) *Run {
	s.ParentUserId = &v
	return s
}

func (s *Run) SetRunId(v string) *Run {
	s.RunId = &v
	return s
}

func (s *Run) SetSource(v string) *Run {
	s.Source = &v
	return s
}

func (s *Run) SetStartedAt(v int64) *Run {
	s.StartedAt = &v
	return s
}

func (s *Run) SetStatus(v string) *Run {
	s.Status = &v
	return s
}

func (s *Run) SetUserId(v string) *Run {
	s.UserId = &v
	return s
}

func (s *Run) SetWorkspaceId(v string) *Run {
	s.WorkspaceId = &v
	return s
}

type CreatePipelineRequest struct {
	Manifest    *string `json:"Manifest,omitempty" xml:"Manifest,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreatePipelineRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequest) SetManifest(v string) *CreatePipelineRequest {
	s.Manifest = &v
	return s
}

func (s *CreatePipelineRequest) SetWorkspaceId(v string) *CreatePipelineRequest {
	s.WorkspaceId = &v
	return s
}

type CreatePipelineResponseBody struct {
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePipelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePipelineResponseBody) SetPipelineId(v string) *CreatePipelineResponseBody {
	s.PipelineId = &v
	return s
}

func (s *CreatePipelineResponseBody) SetRequestId(v string) *CreatePipelineResponseBody {
	s.RequestId = &v
	return s
}

type CreatePipelineResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePipelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePipelineResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineResponse) GoString() string {
	return s.String()
}

func (s *CreatePipelineResponse) SetHeaders(v map[string]*string) *CreatePipelineResponse {
	s.Headers = v
	return s
}

func (s *CreatePipelineResponse) SetStatusCode(v int32) *CreatePipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePipelineResponse) SetBody(v *CreatePipelineResponseBody) *CreatePipelineResponse {
	s.Body = v
	return s
}

type CreatePipelineRunRequest struct {
	Accessibility     *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Arguments         *string `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NoConfirmRequired *bool   `json:"NoConfirmRequired,omitempty" xml:"NoConfirmRequired,omitempty"`
	Options           *string `json:"Options,omitempty" xml:"Options,omitempty"`
	PipelineId        *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	PipelineManifest  *string `json:"PipelineManifest,omitempty" xml:"PipelineManifest,omitempty"`
	SourceId          *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType        *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	WorkspaceId       *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreatePipelineRunRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineRunRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineRunRequest) SetAccessibility(v string) *CreatePipelineRunRequest {
	s.Accessibility = &v
	return s
}

func (s *CreatePipelineRunRequest) SetArguments(v string) *CreatePipelineRunRequest {
	s.Arguments = &v
	return s
}

func (s *CreatePipelineRunRequest) SetName(v string) *CreatePipelineRunRequest {
	s.Name = &v
	return s
}

func (s *CreatePipelineRunRequest) SetNoConfirmRequired(v bool) *CreatePipelineRunRequest {
	s.NoConfirmRequired = &v
	return s
}

func (s *CreatePipelineRunRequest) SetOptions(v string) *CreatePipelineRunRequest {
	s.Options = &v
	return s
}

func (s *CreatePipelineRunRequest) SetPipelineId(v string) *CreatePipelineRunRequest {
	s.PipelineId = &v
	return s
}

func (s *CreatePipelineRunRequest) SetPipelineManifest(v string) *CreatePipelineRunRequest {
	s.PipelineManifest = &v
	return s
}

func (s *CreatePipelineRunRequest) SetSourceId(v string) *CreatePipelineRunRequest {
	s.SourceId = &v
	return s
}

func (s *CreatePipelineRunRequest) SetSourceType(v string) *CreatePipelineRunRequest {
	s.SourceType = &v
	return s
}

func (s *CreatePipelineRunRequest) SetWorkspaceId(v string) *CreatePipelineRunRequest {
	s.WorkspaceId = &v
	return s
}

type CreatePipelineRunResponseBody struct {
	PipelineRunId *string `json:"PipelineRunId,omitempty" xml:"PipelineRunId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePipelineRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePipelineRunResponseBody) SetPipelineRunId(v string) *CreatePipelineRunResponseBody {
	s.PipelineRunId = &v
	return s
}

func (s *CreatePipelineRunResponseBody) SetRequestId(v string) *CreatePipelineRunResponseBody {
	s.RequestId = &v
	return s
}

type CreatePipelineRunResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePipelineRunResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineRunResponse) GoString() string {
	return s.String()
}

func (s *CreatePipelineRunResponse) SetHeaders(v map[string]*string) *CreatePipelineRunResponse {
	s.Headers = v
	return s
}

func (s *CreatePipelineRunResponse) SetStatusCode(v int32) *CreatePipelineRunResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePipelineRunResponse) SetBody(v *CreatePipelineRunResponseBody) *CreatePipelineRunResponse {
	s.Body = v
	return s
}

type DeletePipelineResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePipelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePipelineResponseBody) SetRequestId(v string) *DeletePipelineResponseBody {
	s.RequestId = &v
	return s
}

type DeletePipelineResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePipelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePipelineResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePipelineResponse) GoString() string {
	return s.String()
}

func (s *DeletePipelineResponse) SetHeaders(v map[string]*string) *DeletePipelineResponse {
	s.Headers = v
	return s
}

func (s *DeletePipelineResponse) SetStatusCode(v int32) *DeletePipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePipelineResponse) SetBody(v *DeletePipelineResponseBody) *DeletePipelineResponse {
	s.Body = v
	return s
}

type DeletePipelineRunResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePipelineRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePipelineRunResponseBody) SetRequestId(v string) *DeletePipelineRunResponseBody {
	s.RequestId = &v
	return s
}

type DeletePipelineRunResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePipelineRunResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePipelineRunResponse) GoString() string {
	return s.String()
}

func (s *DeletePipelineRunResponse) SetHeaders(v map[string]*string) *DeletePipelineRunResponse {
	s.Headers = v
	return s
}

func (s *DeletePipelineRunResponse) SetStatusCode(v int32) *DeletePipelineRunResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePipelineRunResponse) SetBody(v *DeletePipelineRunResponseBody) *DeletePipelineRunResponse {
	s.Body = v
	return s
}

type GetPipelineResponseBody struct {
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Identifier      *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	Manifest        *string `json:"Manifest,omitempty" xml:"Manifest,omitempty"`
	PipelineId      *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	Provider        *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Uuid            *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetPipelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineResponseBody) SetGmtCreateTime(v string) *GetPipelineResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetPipelineResponseBody) SetGmtModifiedTime(v string) *GetPipelineResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetPipelineResponseBody) SetIdentifier(v string) *GetPipelineResponseBody {
	s.Identifier = &v
	return s
}

func (s *GetPipelineResponseBody) SetManifest(v string) *GetPipelineResponseBody {
	s.Manifest = &v
	return s
}

func (s *GetPipelineResponseBody) SetPipelineId(v string) *GetPipelineResponseBody {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineResponseBody) SetProvider(v string) *GetPipelineResponseBody {
	s.Provider = &v
	return s
}

func (s *GetPipelineResponseBody) SetRequestId(v string) *GetPipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineResponseBody) SetUuid(v string) *GetPipelineResponseBody {
	s.Uuid = &v
	return s
}

func (s *GetPipelineResponseBody) SetVersion(v string) *GetPipelineResponseBody {
	s.Version = &v
	return s
}

func (s *GetPipelineResponseBody) SetWorkspaceId(v string) *GetPipelineResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetPipelineResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPipelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPipelineResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineResponse) SetHeaders(v map[string]*string) *GetPipelineResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineResponse) SetStatusCode(v int32) *GetPipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineResponse) SetBody(v *GetPipelineResponseBody) *GetPipelineResponse {
	s.Body = v
	return s
}

type GetPipelineRunRequest struct {
	ManifestType *string `json:"ManifestType,omitempty" xml:"ManifestType,omitempty"`
	TokenId      *string `json:"TokenId,omitempty" xml:"TokenId,omitempty"`
	Verbose      *bool   `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetPipelineRunRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineRunRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineRunRequest) SetManifestType(v string) *GetPipelineRunRequest {
	s.ManifestType = &v
	return s
}

func (s *GetPipelineRunRequest) SetTokenId(v string) *GetPipelineRunRequest {
	s.TokenId = &v
	return s
}

func (s *GetPipelineRunRequest) SetVerbose(v bool) *GetPipelineRunRequest {
	s.Verbose = &v
	return s
}

type GetPipelineRunResponseBody struct {
	Accessibility   *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Arguments       *string `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
	Duration        *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FinishedAt      *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Manifest        *string `json:"Manifest,omitempty" xml:"Manifest,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NodeId          *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Options         *string `json:"Options,omitempty" xml:"Options,omitempty"`
	ParentUserId    *string `json:"ParentUserId,omitempty" xml:"ParentUserId,omitempty"`
	PipelineId      *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	PipelineRunId   *string `json:"PipelineRunId,omitempty" xml:"PipelineRunId,omitempty"`
	PipelineRunUri  *string `json:"PipelineRunUri,omitempty" xml:"PipelineRunUri,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceId        *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType      *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	StartedAt       *string `json:"StartedAt,omitempty" xml:"StartedAt,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetPipelineRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponseBody) SetAccessibility(v string) *GetPipelineRunResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetArguments(v string) *GetPipelineRunResponseBody {
	s.Arguments = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetDuration(v int64) *GetPipelineRunResponseBody {
	s.Duration = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetFinishedAt(v string) *GetPipelineRunResponseBody {
	s.FinishedAt = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetGmtCreateTime(v string) *GetPipelineRunResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetGmtModifiedTime(v string) *GetPipelineRunResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetManifest(v string) *GetPipelineRunResponseBody {
	s.Manifest = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetMessage(v string) *GetPipelineRunResponseBody {
	s.Message = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetName(v string) *GetPipelineRunResponseBody {
	s.Name = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetNodeId(v string) *GetPipelineRunResponseBody {
	s.NodeId = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetOptions(v string) *GetPipelineRunResponseBody {
	s.Options = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetParentUserId(v string) *GetPipelineRunResponseBody {
	s.ParentUserId = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetPipelineId(v string) *GetPipelineRunResponseBody {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetPipelineRunId(v string) *GetPipelineRunResponseBody {
	s.PipelineRunId = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetPipelineRunUri(v string) *GetPipelineRunResponseBody {
	s.PipelineRunUri = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetRequestId(v string) *GetPipelineRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetSourceId(v string) *GetPipelineRunResponseBody {
	s.SourceId = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetSourceType(v string) *GetPipelineRunResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetStartedAt(v string) *GetPipelineRunResponseBody {
	s.StartedAt = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetStatus(v string) *GetPipelineRunResponseBody {
	s.Status = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetUserId(v string) *GetPipelineRunResponseBody {
	s.UserId = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetWorkspaceId(v string) *GetPipelineRunResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetPipelineRunResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPipelineRunResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineRunResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponse) SetHeaders(v map[string]*string) *GetPipelineRunResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineRunResponse) SetStatusCode(v int32) *GetPipelineRunResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineRunResponse) SetBody(v *GetPipelineRunResponseBody) *GetPipelineRunResponse {
	s.Body = v
	return s
}

type GetPipelineRunNodeRequest struct {
	Depth   *int32  `json:"Depth,omitempty" xml:"Depth,omitempty"`
	TokenId *string `json:"TokenId,omitempty" xml:"TokenId,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetPipelineRunNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineRunNodeRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineRunNodeRequest) SetDepth(v int32) *GetPipelineRunNodeRequest {
	s.Depth = &v
	return s
}

func (s *GetPipelineRunNodeRequest) SetTokenId(v string) *GetPipelineRunNodeRequest {
	s.TokenId = &v
	return s
}

func (s *GetPipelineRunNodeRequest) SetType(v string) *GetPipelineRunNodeRequest {
	s.Type = &v
	return s
}

type GetPipelineRunNodeResponseBody struct {
	ApiVersion *string                                   `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	Metadata   *GetPipelineRunNodeResponseBodyMetadata   `json:"Metadata,omitempty" xml:"Metadata,omitempty" type:"Struct"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Spec       *GetPipelineRunNodeResponseBodySpec       `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
	StatusInfo *GetPipelineRunNodeResponseBodyStatusInfo `json:"StatusInfo,omitempty" xml:"StatusInfo,omitempty" type:"Struct"`
}

func (s GetPipelineRunNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineRunNodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineRunNodeResponseBody) SetApiVersion(v string) *GetPipelineRunNodeResponseBody {
	s.ApiVersion = &v
	return s
}

func (s *GetPipelineRunNodeResponseBody) SetMetadata(v *GetPipelineRunNodeResponseBodyMetadata) *GetPipelineRunNodeResponseBody {
	s.Metadata = v
	return s
}

func (s *GetPipelineRunNodeResponseBody) SetRequestId(v string) *GetPipelineRunNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineRunNodeResponseBody) SetSpec(v *GetPipelineRunNodeResponseBodySpec) *GetPipelineRunNodeResponseBody {
	s.Spec = v
	return s
}

func (s *GetPipelineRunNodeResponseBody) SetStatusInfo(v *GetPipelineRunNodeResponseBodyStatusInfo) *GetPipelineRunNodeResponseBody {
	s.StatusInfo = v
	return s
}

type GetPipelineRunNodeResponseBodyMetadata struct {
	DisplayName    *string   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Identifier     *string   `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	Name           *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NodeId         *string   `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeType       *string   `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Provider       *string   `json:"Provider,omitempty" xml:"Provider,omitempty"`
	RelatedNodeIds []*string `json:"RelatedNodeIds,omitempty" xml:"RelatedNodeIds,omitempty" type:"Repeated"`
	Version        *string   `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetPipelineRunNodeResponseBodyMetadata) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineRunNodeResponseBodyMetadata) GoString() string {
	return s.String()
}

func (s *GetPipelineRunNodeResponseBodyMetadata) SetDisplayName(v string) *GetPipelineRunNodeResponseBodyMetadata {
	s.DisplayName = &v
	return s
}

func (s *GetPipelineRunNodeResponseBodyMetadata) SetIdentifier(v string) *GetPipelineRunNodeResponseBodyMetadata {
	s.Identifier = &v
	return s
}

func (s *GetPipelineRunNodeResponseBodyMetadata) SetName(v string) *GetPipelineRunNodeResponseBodyMetadata {
	s.Name = &v
	return s
}

func (s *GetPipelineRunNodeResponseBodyMetadata) SetNodeId(v string) *GetPipelineRunNodeResponseBodyMetadata {
	s.NodeId = &v
	return s
}

func (s *GetPipelineRunNodeResponseBodyMetadata) SetNodeType(v string) *GetPipelineRunNodeResponseBodyMetadata {
	s.NodeType = &v
	return s
}

func (s *GetPipelineRunNodeResponseBodyMetadata) SetProvider(v string) *GetPipelineRunNodeResponseBodyMetadata {
	s.Provider = &v
	return s
}

func (s *GetPipelineRunNodeResponseBodyMetadata) SetRelatedNodeIds(v []*string) *GetPipelineRunNodeResponseBodyMetadata {
	s.RelatedNodeIds = v
	return s
}

func (s *GetPipelineRunNodeResponseBodyMetadata) SetVersion(v string) *GetPipelineRunNodeResponseBodyMetadata {
	s.Version = &v
	return s
}

type GetPipelineRunNodeResponseBodySpec struct {
	Dependencies []*string                                       `json:"Dependencies,omitempty" xml:"Dependencies,omitempty" type:"Repeated"`
	HasPipelines *bool                                           `json:"HasPipelines,omitempty" xml:"HasPipelines,omitempty"`
	Inputs       *GetPipelineRunNodeResponseBodySpecInputs       `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Struct"`
	Outputs      *GetPipelineRunNodeResponseBodySpecOutputs      `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Struct"`
	Parallelism  *int32                                          `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	Pipelines    []map[string]interface{}                        `json:"Pipelines,omitempty" xml:"Pipelines,omitempty" type:"Repeated"`
	When         *string                                         `json:"When,omitempty" xml:"When,omitempty"`
	WithItems    []*string                                       `json:"WithItems,omitempty" xml:"WithItems,omitempty" type:"Repeated"`
	WithParam    *string                                         `json:"WithParam,omitempty" xml:"WithParam,omitempty"`
	WithSequence *GetPipelineRunNodeResponseBodySpecWithSequence `json:"WithSequence,omitempty" xml:"WithSequence,omitempty" type:"Struct"`
}

func (s GetPipelineRunNodeResponseBodySpec) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineRunNodeResponseBodySpec) GoString() string {
	return s.String()
}

func (s *GetPipelineRunNodeResponseBodySpec) SetDependencies(v []*string) *GetPipelineRunNodeResponseBodySpec {
	s.Dependencies = v
	return s
}

func (s *GetPipelineRunNodeResponseBodySpec) SetHasPipelines(v bool) *GetPipelineRunNodeResponseBodySpec {
	s.HasPipelines = &v
	return s
}

func (s *GetPipelineRunNodeResponseBodySpec) SetInputs(v *GetPipelineRunNodeResponseBodySpecInputs) *GetPipelineRunNodeResponseBodySpec {
	s.Inputs = v
	return s
}

func (s *GetPipelineRunNodeResponseBodySpec) SetOutputs(v *GetPipelineRunNodeResponseBodySpecOutputs) *GetPipelineRunNodeResponseBodySpec {
	s.Outputs = v
	return s
}

func (s *GetPipelineRunNodeResponseBodySpec) SetParallelism(v int32) *GetPipelineRunNodeResponseBodySpec {
	s.Parallelism = &v
	return s
}

func (s *GetPipelineRunNodeResponseBodySpec) SetPipelines(v []map[string]interface{}) *GetPipelineRunNodeResponseBodySpec {
	s.Pipelines = v
	return s
}

func (s *GetPipelineRunNodeResponseBodySpec) SetWhen(v string) *GetPipelineRunNodeResponseBodySpec {
	s.When = &v
	return s
}

func (s *GetPipelineRunNodeResponseBodySpec) SetWithItems(v []*string) *GetPipelineRunNodeResponseBodySpec {
	s.WithItems = v
	return s
}

func (s *GetPipelineRunNodeResponseBodySpec) SetWithParam(v string) *GetPipelineRunNodeResponseBodySpec {
	s.WithParam = &v
	return s
}

func (s *GetPipelineRunNodeResponseBodySpec) SetWithSequence(v *GetPipelineRunNodeResponseBodySpecWithSequence) *GetPipelineRunNodeResponseBodySpec {
	s.WithSequence = v
	return s
}

type GetPipelineRunNodeResponseBodySpecInputs struct {
	Artifacts  []map[string]interface{} `json:"Artifacts,omitempty" xml:"Artifacts,omitempty" type:"Repeated"`
	Parameters []map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s GetPipelineRunNodeResponseBodySpecInputs) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineRunNodeResponseBodySpecInputs) GoString() string {
	return s.String()
}

func (s *GetPipelineRunNodeResponseBodySpecInputs) SetArtifacts(v []map[string]interface{}) *GetPipelineRunNodeResponseBodySpecInputs {
	s.Artifacts = v
	return s
}

func (s *GetPipelineRunNodeResponseBodySpecInputs) SetParameters(v []map[string]interface{}) *GetPipelineRunNodeResponseBodySpecInputs {
	s.Parameters = v
	return s
}

type GetPipelineRunNodeResponseBodySpecOutputs struct {
	Artifacts  []map[string]interface{} `json:"Artifacts,omitempty" xml:"Artifacts,omitempty" type:"Repeated"`
	Parameters []map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s GetPipelineRunNodeResponseBodySpecOutputs) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineRunNodeResponseBodySpecOutputs) GoString() string {
	return s.String()
}

func (s *GetPipelineRunNodeResponseBodySpecOutputs) SetArtifacts(v []map[string]interface{}) *GetPipelineRunNodeResponseBodySpecOutputs {
	s.Artifacts = v
	return s
}

func (s *GetPipelineRunNodeResponseBodySpecOutputs) SetParameters(v []map[string]interface{}) *GetPipelineRunNodeResponseBodySpecOutputs {
	s.Parameters = v
	return s
}

type GetPipelineRunNodeResponseBodySpecWithSequence struct {
	End    *int32  `json:"End,omitempty" xml:"End,omitempty"`
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	Start  *int32  `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetPipelineRunNodeResponseBodySpecWithSequence) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineRunNodeResponseBodySpecWithSequence) GoString() string {
	return s.String()
}

func (s *GetPipelineRunNodeResponseBodySpecWithSequence) SetEnd(v int32) *GetPipelineRunNodeResponseBodySpecWithSequence {
	s.End = &v
	return s
}

func (s *GetPipelineRunNodeResponseBodySpecWithSequence) SetFormat(v string) *GetPipelineRunNodeResponseBodySpecWithSequence {
	s.Format = &v
	return s
}

func (s *GetPipelineRunNodeResponseBodySpecWithSequence) SetStart(v int32) *GetPipelineRunNodeResponseBodySpecWithSequence {
	s.Start = &v
	return s
}

type GetPipelineRunNodeResponseBodyStatusInfo struct {
	Conditions []map[string]interface{} `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	FinishedAt *string                  `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	Progress   *string                  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	StartedAt  *string                  `json:"StartedAt,omitempty" xml:"StartedAt,omitempty"`
	Status     *string                  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPipelineRunNodeResponseBodyStatusInfo) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineRunNodeResponseBodyStatusInfo) GoString() string {
	return s.String()
}

func (s *GetPipelineRunNodeResponseBodyStatusInfo) SetConditions(v []map[string]interface{}) *GetPipelineRunNodeResponseBodyStatusInfo {
	s.Conditions = v
	return s
}

func (s *GetPipelineRunNodeResponseBodyStatusInfo) SetFinishedAt(v string) *GetPipelineRunNodeResponseBodyStatusInfo {
	s.FinishedAt = &v
	return s
}

func (s *GetPipelineRunNodeResponseBodyStatusInfo) SetProgress(v string) *GetPipelineRunNodeResponseBodyStatusInfo {
	s.Progress = &v
	return s
}

func (s *GetPipelineRunNodeResponseBodyStatusInfo) SetStartedAt(v string) *GetPipelineRunNodeResponseBodyStatusInfo {
	s.StartedAt = &v
	return s
}

func (s *GetPipelineRunNodeResponseBodyStatusInfo) SetStatus(v string) *GetPipelineRunNodeResponseBodyStatusInfo {
	s.Status = &v
	return s
}

type GetPipelineRunNodeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPipelineRunNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPipelineRunNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineRunNodeResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineRunNodeResponse) SetHeaders(v map[string]*string) *GetPipelineRunNodeResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineRunNodeResponse) SetStatusCode(v int32) *GetPipelineRunNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineRunNodeResponse) SetBody(v *GetPipelineRunNodeResponseBody) *GetPipelineRunNodeResponse {
	s.Body = v
	return s
}

type ListPipelineRunNodeLogsRequest struct {
	FromTimeInSeconds *int64  `json:"FromTimeInSeconds,omitempty" xml:"FromTimeInSeconds,omitempty"`
	Keyword           *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Offset            *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Reverse           *bool   `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	ToTimeInSeconds   *int64  `json:"ToTimeInSeconds,omitempty" xml:"ToTimeInSeconds,omitempty"`
	TokenId           *string `json:"TokenId,omitempty" xml:"TokenId,omitempty"`
}

func (s ListPipelineRunNodeLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunNodeLogsRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineRunNodeLogsRequest) SetFromTimeInSeconds(v int64) *ListPipelineRunNodeLogsRequest {
	s.FromTimeInSeconds = &v
	return s
}

func (s *ListPipelineRunNodeLogsRequest) SetKeyword(v string) *ListPipelineRunNodeLogsRequest {
	s.Keyword = &v
	return s
}

func (s *ListPipelineRunNodeLogsRequest) SetOffset(v int32) *ListPipelineRunNodeLogsRequest {
	s.Offset = &v
	return s
}

func (s *ListPipelineRunNodeLogsRequest) SetPageSize(v int32) *ListPipelineRunNodeLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPipelineRunNodeLogsRequest) SetReverse(v bool) *ListPipelineRunNodeLogsRequest {
	s.Reverse = &v
	return s
}

func (s *ListPipelineRunNodeLogsRequest) SetToTimeInSeconds(v int64) *ListPipelineRunNodeLogsRequest {
	s.ToTimeInSeconds = &v
	return s
}

func (s *ListPipelineRunNodeLogsRequest) SetTokenId(v string) *ListPipelineRunNodeLogsRequest {
	s.TokenId = &v
	return s
}

type ListPipelineRunNodeLogsResponseBody struct {
	Logs       []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPipelineRunNodeLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunNodeLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineRunNodeLogsResponseBody) SetLogs(v []*string) *ListPipelineRunNodeLogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListPipelineRunNodeLogsResponseBody) SetRequestId(v string) *ListPipelineRunNodeLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineRunNodeLogsResponseBody) SetTotalCount(v int64) *ListPipelineRunNodeLogsResponseBody {
	s.TotalCount = &v
	return s
}

type ListPipelineRunNodeLogsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPipelineRunNodeLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPipelineRunNodeLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunNodeLogsResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineRunNodeLogsResponse) SetHeaders(v map[string]*string) *ListPipelineRunNodeLogsResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineRunNodeLogsResponse) SetStatusCode(v int32) *ListPipelineRunNodeLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineRunNodeLogsResponse) SetBody(v *ListPipelineRunNodeLogsResponseBody) *ListPipelineRunNodeLogsResponse {
	s.Body = v
	return s
}

type ListPipelineRunNodeOutputsRequest struct {
	Depth      *int32  `json:"Depth,omitempty" xml:"Depth,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order      *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	TokenId    *string `json:"TokenId,omitempty" xml:"TokenId,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPipelineRunNodeOutputsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunNodeOutputsRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineRunNodeOutputsRequest) SetDepth(v int32) *ListPipelineRunNodeOutputsRequest {
	s.Depth = &v
	return s
}

func (s *ListPipelineRunNodeOutputsRequest) SetName(v string) *ListPipelineRunNodeOutputsRequest {
	s.Name = &v
	return s
}

func (s *ListPipelineRunNodeOutputsRequest) SetOrder(v string) *ListPipelineRunNodeOutputsRequest {
	s.Order = &v
	return s
}

func (s *ListPipelineRunNodeOutputsRequest) SetPageNumber(v int32) *ListPipelineRunNodeOutputsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPipelineRunNodeOutputsRequest) SetPageSize(v int32) *ListPipelineRunNodeOutputsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPipelineRunNodeOutputsRequest) SetSortBy(v string) *ListPipelineRunNodeOutputsRequest {
	s.SortBy = &v
	return s
}

func (s *ListPipelineRunNodeOutputsRequest) SetTokenId(v string) *ListPipelineRunNodeOutputsRequest {
	s.TokenId = &v
	return s
}

func (s *ListPipelineRunNodeOutputsRequest) SetType(v string) *ListPipelineRunNodeOutputsRequest {
	s.Type = &v
	return s
}

type ListPipelineRunNodeOutputsResponseBody struct {
	Outputs    []*ListPipelineRunNodeOutputsResponseBodyOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPipelineRunNodeOutputsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunNodeOutputsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineRunNodeOutputsResponseBody) SetOutputs(v []*ListPipelineRunNodeOutputsResponseBodyOutputs) *ListPipelineRunNodeOutputsResponseBody {
	s.Outputs = v
	return s
}

func (s *ListPipelineRunNodeOutputsResponseBody) SetRequestId(v string) *ListPipelineRunNodeOutputsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineRunNodeOutputsResponseBody) SetTotalCount(v int64) *ListPipelineRunNodeOutputsResponseBody {
	s.TotalCount = &v
	return s
}

type ListPipelineRunNodeOutputsResponseBodyOutputs struct {
	ExpandableArtifactName *string                `json:"ExpandableArtifactName,omitempty" xml:"ExpandableArtifactName,omitempty"`
	ExpandedArtifactIndex  *int64                 `json:"ExpandedArtifactIndex,omitempty" xml:"ExpandedArtifactIndex,omitempty"`
	GmtCreateTime          *string                `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	Id                     *string                `json:"Id,omitempty" xml:"Id,omitempty"`
	Info                   map[string]interface{} `json:"Info,omitempty" xml:"Info,omitempty"`
	Name                   *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	NodeId                 *string                `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Producer               *string                `json:"Producer,omitempty" xml:"Producer,omitempty"`
	Type                   *string                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPipelineRunNodeOutputsResponseBodyOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunNodeOutputsResponseBodyOutputs) GoString() string {
	return s.String()
}

func (s *ListPipelineRunNodeOutputsResponseBodyOutputs) SetExpandableArtifactName(v string) *ListPipelineRunNodeOutputsResponseBodyOutputs {
	s.ExpandableArtifactName = &v
	return s
}

func (s *ListPipelineRunNodeOutputsResponseBodyOutputs) SetExpandedArtifactIndex(v int64) *ListPipelineRunNodeOutputsResponseBodyOutputs {
	s.ExpandedArtifactIndex = &v
	return s
}

func (s *ListPipelineRunNodeOutputsResponseBodyOutputs) SetGmtCreateTime(v string) *ListPipelineRunNodeOutputsResponseBodyOutputs {
	s.GmtCreateTime = &v
	return s
}

func (s *ListPipelineRunNodeOutputsResponseBodyOutputs) SetId(v string) *ListPipelineRunNodeOutputsResponseBodyOutputs {
	s.Id = &v
	return s
}

func (s *ListPipelineRunNodeOutputsResponseBodyOutputs) SetInfo(v map[string]interface{}) *ListPipelineRunNodeOutputsResponseBodyOutputs {
	s.Info = v
	return s
}

func (s *ListPipelineRunNodeOutputsResponseBodyOutputs) SetName(v string) *ListPipelineRunNodeOutputsResponseBodyOutputs {
	s.Name = &v
	return s
}

func (s *ListPipelineRunNodeOutputsResponseBodyOutputs) SetNodeId(v string) *ListPipelineRunNodeOutputsResponseBodyOutputs {
	s.NodeId = &v
	return s
}

func (s *ListPipelineRunNodeOutputsResponseBodyOutputs) SetProducer(v string) *ListPipelineRunNodeOutputsResponseBodyOutputs {
	s.Producer = &v
	return s
}

func (s *ListPipelineRunNodeOutputsResponseBodyOutputs) SetType(v string) *ListPipelineRunNodeOutputsResponseBodyOutputs {
	s.Type = &v
	return s
}

type ListPipelineRunNodeOutputsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPipelineRunNodeOutputsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPipelineRunNodeOutputsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunNodeOutputsResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineRunNodeOutputsResponse) SetHeaders(v map[string]*string) *ListPipelineRunNodeOutputsResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineRunNodeOutputsResponse) SetStatusCode(v int32) *ListPipelineRunNodeOutputsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineRunNodeOutputsResponse) SetBody(v *ListPipelineRunNodeOutputsResponseBody) *ListPipelineRunNodeOutputsResponse {
	s.Body = v
	return s
}

type ListPipelineRunNodeStatusRequest struct {
	Depth   *int64  `json:"Depth,omitempty" xml:"Depth,omitempty"`
	TokenId *string `json:"TokenId,omitempty" xml:"TokenId,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPipelineRunNodeStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunNodeStatusRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineRunNodeStatusRequest) SetDepth(v int64) *ListPipelineRunNodeStatusRequest {
	s.Depth = &v
	return s
}

func (s *ListPipelineRunNodeStatusRequest) SetTokenId(v string) *ListPipelineRunNodeStatusRequest {
	s.TokenId = &v
	return s
}

func (s *ListPipelineRunNodeStatusRequest) SetType(v string) *ListPipelineRunNodeStatusRequest {
	s.Type = &v
	return s
}

type ListPipelineRunNodeStatusResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    []*ListPipelineRunNodeStatusResponseBodyStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
}

func (s ListPipelineRunNodeStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunNodeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineRunNodeStatusResponseBody) SetRequestId(v string) *ListPipelineRunNodeStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineRunNodeStatusResponseBody) SetStatus(v []*ListPipelineRunNodeStatusResponseBodyStatus) *ListPipelineRunNodeStatusResponseBody {
	s.Status = v
	return s
}

type ListPipelineRunNodeStatusResponseBodyStatus struct {
	FinishedAt  *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	NodeId      *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName    *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	RuntimeInfo *string `json:"RuntimeInfo,omitempty" xml:"RuntimeInfo,omitempty"`
	StartedAt   *string `json:"StartedAt,omitempty" xml:"StartedAt,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListPipelineRunNodeStatusResponseBodyStatus) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunNodeStatusResponseBodyStatus) GoString() string {
	return s.String()
}

func (s *ListPipelineRunNodeStatusResponseBodyStatus) SetFinishedAt(v string) *ListPipelineRunNodeStatusResponseBodyStatus {
	s.FinishedAt = &v
	return s
}

func (s *ListPipelineRunNodeStatusResponseBodyStatus) SetNodeId(v string) *ListPipelineRunNodeStatusResponseBodyStatus {
	s.NodeId = &v
	return s
}

func (s *ListPipelineRunNodeStatusResponseBodyStatus) SetNodeName(v string) *ListPipelineRunNodeStatusResponseBodyStatus {
	s.NodeName = &v
	return s
}

func (s *ListPipelineRunNodeStatusResponseBodyStatus) SetRuntimeInfo(v string) *ListPipelineRunNodeStatusResponseBodyStatus {
	s.RuntimeInfo = &v
	return s
}

func (s *ListPipelineRunNodeStatusResponseBodyStatus) SetStartedAt(v string) *ListPipelineRunNodeStatusResponseBodyStatus {
	s.StartedAt = &v
	return s
}

func (s *ListPipelineRunNodeStatusResponseBodyStatus) SetStatus(v string) *ListPipelineRunNodeStatusResponseBodyStatus {
	s.Status = &v
	return s
}

type ListPipelineRunNodeStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPipelineRunNodeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPipelineRunNodeStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunNodeStatusResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineRunNodeStatusResponse) SetHeaders(v map[string]*string) *ListPipelineRunNodeStatusResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineRunNodeStatusResponse) SetStatusCode(v int32) *ListPipelineRunNodeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineRunNodeStatusResponse) SetBody(v *ListPipelineRunNodeStatusResponseBody) *ListPipelineRunNodeStatusResponse {
	s.Body = v
	return s
}

type ListPipelineRunsRequest struct {
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order         *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PipelineIds   *string `json:"PipelineIds,omitempty" xml:"PipelineIds,omitempty"`
	PipelineRunId *string `json:"PipelineRunId,omitempty" xml:"PipelineRunId,omitempty"`
	SortBy        *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	SourceId      *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType    *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListPipelineRunsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunsRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsRequest) SetName(v string) *ListPipelineRunsRequest {
	s.Name = &v
	return s
}

func (s *ListPipelineRunsRequest) SetOrder(v string) *ListPipelineRunsRequest {
	s.Order = &v
	return s
}

func (s *ListPipelineRunsRequest) SetPageNumber(v int32) *ListPipelineRunsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPipelineRunsRequest) SetPageSize(v int32) *ListPipelineRunsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPipelineRunsRequest) SetPipelineIds(v string) *ListPipelineRunsRequest {
	s.PipelineIds = &v
	return s
}

func (s *ListPipelineRunsRequest) SetPipelineRunId(v string) *ListPipelineRunsRequest {
	s.PipelineRunId = &v
	return s
}

func (s *ListPipelineRunsRequest) SetSortBy(v string) *ListPipelineRunsRequest {
	s.SortBy = &v
	return s
}

func (s *ListPipelineRunsRequest) SetSourceId(v string) *ListPipelineRunsRequest {
	s.SourceId = &v
	return s
}

func (s *ListPipelineRunsRequest) SetSourceType(v string) *ListPipelineRunsRequest {
	s.SourceType = &v
	return s
}

func (s *ListPipelineRunsRequest) SetStatus(v string) *ListPipelineRunsRequest {
	s.Status = &v
	return s
}

func (s *ListPipelineRunsRequest) SetUserId(v string) *ListPipelineRunsRequest {
	s.UserId = &v
	return s
}

func (s *ListPipelineRunsRequest) SetWorkspaceId(v string) *ListPipelineRunsRequest {
	s.WorkspaceId = &v
	return s
}

type ListPipelineRunsResponseBody struct {
	PipelineRuns []*ListPipelineRunsResponseBodyPipelineRuns `json:"PipelineRuns,omitempty" xml:"PipelineRuns,omitempty" type:"Repeated"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPipelineRunsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsResponseBody) SetPipelineRuns(v []*ListPipelineRunsResponseBodyPipelineRuns) *ListPipelineRunsResponseBody {
	s.PipelineRuns = v
	return s
}

func (s *ListPipelineRunsResponseBody) SetRequestId(v string) *ListPipelineRunsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineRunsResponseBody) SetTotalCount(v int64) *ListPipelineRunsResponseBody {
	s.TotalCount = &v
	return s
}

type ListPipelineRunsResponseBodyPipelineRuns struct {
	Accessibility   *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Duration        *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FinishedAt      *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NodeId          *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	ParentUserId    *string `json:"ParentUserId,omitempty" xml:"ParentUserId,omitempty"`
	PipelineId      *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	PipelineRunId   *string `json:"PipelineRunId,omitempty" xml:"PipelineRunId,omitempty"`
	PipelineRunUri  *string `json:"PipelineRunUri,omitempty" xml:"PipelineRunUri,omitempty"`
	SourceId        *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType      *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	StartedAt       *string `json:"StartedAt,omitempty" xml:"StartedAt,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListPipelineRunsResponseBodyPipelineRuns) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunsResponseBodyPipelineRuns) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetAccessibility(v string) *ListPipelineRunsResponseBodyPipelineRuns {
	s.Accessibility = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetDuration(v int32) *ListPipelineRunsResponseBodyPipelineRuns {
	s.Duration = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetFinishedAt(v string) *ListPipelineRunsResponseBodyPipelineRuns {
	s.FinishedAt = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetGmtCreateTime(v string) *ListPipelineRunsResponseBodyPipelineRuns {
	s.GmtCreateTime = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetGmtModifiedTime(v string) *ListPipelineRunsResponseBodyPipelineRuns {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetMessage(v string) *ListPipelineRunsResponseBodyPipelineRuns {
	s.Message = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetName(v string) *ListPipelineRunsResponseBodyPipelineRuns {
	s.Name = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetNodeId(v string) *ListPipelineRunsResponseBodyPipelineRuns {
	s.NodeId = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetParentUserId(v string) *ListPipelineRunsResponseBodyPipelineRuns {
	s.ParentUserId = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetPipelineId(v string) *ListPipelineRunsResponseBodyPipelineRuns {
	s.PipelineId = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetPipelineRunId(v string) *ListPipelineRunsResponseBodyPipelineRuns {
	s.PipelineRunId = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetPipelineRunUri(v string) *ListPipelineRunsResponseBodyPipelineRuns {
	s.PipelineRunUri = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetSourceId(v string) *ListPipelineRunsResponseBodyPipelineRuns {
	s.SourceId = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetSourceType(v string) *ListPipelineRunsResponseBodyPipelineRuns {
	s.SourceType = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetStartedAt(v string) *ListPipelineRunsResponseBodyPipelineRuns {
	s.StartedAt = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetStatus(v string) *ListPipelineRunsResponseBodyPipelineRuns {
	s.Status = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetUserId(v string) *ListPipelineRunsResponseBodyPipelineRuns {
	s.UserId = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetWorkspaceId(v string) *ListPipelineRunsResponseBodyPipelineRuns {
	s.WorkspaceId = &v
	return s
}

type ListPipelineRunsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPipelineRunsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPipelineRunsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunsResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsResponse) SetHeaders(v map[string]*string) *ListPipelineRunsResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineRunsResponse) SetStatusCode(v int32) *ListPipelineRunsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineRunsResponse) SetBody(v *ListPipelineRunsResponseBody) *ListPipelineRunsResponse {
	s.Body = v
	return s
}

type ListPipelineRunsStatusRequest struct {
	Nodes        []*ListPipelineRunsStatusRequestNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	OutputType   *string                               `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	PipelineRuns []*string                             `json:"PipelineRuns,omitempty" xml:"PipelineRuns,omitempty" type:"Repeated"`
	WorkspaceId  *string                               `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListPipelineRunsStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunsStatusRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsStatusRequest) SetNodes(v []*ListPipelineRunsStatusRequestNodes) *ListPipelineRunsStatusRequest {
	s.Nodes = v
	return s
}

func (s *ListPipelineRunsStatusRequest) SetOutputType(v string) *ListPipelineRunsStatusRequest {
	s.OutputType = &v
	return s
}

func (s *ListPipelineRunsStatusRequest) SetPipelineRuns(v []*string) *ListPipelineRunsStatusRequest {
	s.PipelineRuns = v
	return s
}

func (s *ListPipelineRunsStatusRequest) SetWorkspaceId(v string) *ListPipelineRunsStatusRequest {
	s.WorkspaceId = &v
	return s
}

type ListPipelineRunsStatusRequestNodes struct {
	NodeId        *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	PipelineRunId *string `json:"PipelineRunId,omitempty" xml:"PipelineRunId,omitempty"`
}

func (s ListPipelineRunsStatusRequestNodes) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunsStatusRequestNodes) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsStatusRequestNodes) SetNodeId(v string) *ListPipelineRunsStatusRequestNodes {
	s.NodeId = &v
	return s
}

func (s *ListPipelineRunsStatusRequestNodes) SetPipelineRunId(v string) *ListPipelineRunsStatusRequestNodes {
	s.PipelineRunId = &v
	return s
}

type ListPipelineRunsStatusResponseBody struct {
	Nodes        []*ListPipelineRunsStatusResponseBodyNodes        `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	Outputs      []*ListPipelineRunsStatusResponseBodyOutputs      `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	PipelineRuns []*ListPipelineRunsStatusResponseBodyPipelineRuns `json:"PipelineRuns,omitempty" xml:"PipelineRuns,omitempty" type:"Repeated"`
	RequestId    *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPipelineRunsStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunsStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsStatusResponseBody) SetNodes(v []*ListPipelineRunsStatusResponseBodyNodes) *ListPipelineRunsStatusResponseBody {
	s.Nodes = v
	return s
}

func (s *ListPipelineRunsStatusResponseBody) SetOutputs(v []*ListPipelineRunsStatusResponseBodyOutputs) *ListPipelineRunsStatusResponseBody {
	s.Outputs = v
	return s
}

func (s *ListPipelineRunsStatusResponseBody) SetPipelineRuns(v []*ListPipelineRunsStatusResponseBodyPipelineRuns) *ListPipelineRunsStatusResponseBody {
	s.PipelineRuns = v
	return s
}

func (s *ListPipelineRunsStatusResponseBody) SetRequestId(v string) *ListPipelineRunsStatusResponseBody {
	s.RequestId = &v
	return s
}

type ListPipelineRunsStatusResponseBodyNodes struct {
	FinishedAt             *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	InputArtifactArchived  *bool   `json:"InputArtifactArchived,omitempty" xml:"InputArtifactArchived,omitempty"`
	NodeId                 *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName               *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	OutputArtifactArchived *bool   `json:"OutputArtifactArchived,omitempty" xml:"OutputArtifactArchived,omitempty"`
	PipelineRunId          *string `json:"PipelineRunId,omitempty" xml:"PipelineRunId,omitempty"`
	StartedAt              *string `json:"StartedAt,omitempty" xml:"StartedAt,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListPipelineRunsStatusResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunsStatusResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsStatusResponseBodyNodes) SetFinishedAt(v string) *ListPipelineRunsStatusResponseBodyNodes {
	s.FinishedAt = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyNodes) SetInputArtifactArchived(v bool) *ListPipelineRunsStatusResponseBodyNodes {
	s.InputArtifactArchived = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyNodes) SetNodeId(v string) *ListPipelineRunsStatusResponseBodyNodes {
	s.NodeId = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyNodes) SetNodeName(v string) *ListPipelineRunsStatusResponseBodyNodes {
	s.NodeName = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyNodes) SetOutputArtifactArchived(v bool) *ListPipelineRunsStatusResponseBodyNodes {
	s.OutputArtifactArchived = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyNodes) SetPipelineRunId(v string) *ListPipelineRunsStatusResponseBodyNodes {
	s.PipelineRunId = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyNodes) SetStartedAt(v string) *ListPipelineRunsStatusResponseBodyNodes {
	s.StartedAt = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyNodes) SetStatus(v string) *ListPipelineRunsStatusResponseBodyNodes {
	s.Status = &v
	return s
}

type ListPipelineRunsStatusResponseBodyOutputs struct {
	ExpandableArtifactName *string                `json:"ExpandableArtifactName,omitempty" xml:"ExpandableArtifactName,omitempty"`
	ExpandedArtifactIndex  *int32                 `json:"ExpandedArtifactIndex,omitempty" xml:"ExpandedArtifactIndex,omitempty"`
	GmtCreateTime          *string                `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	Id                     *string                `json:"Id,omitempty" xml:"Id,omitempty"`
	Metadata               map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	Name                   *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	NodeId                 *string                `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	PipelineRunId          *string                `json:"PipelineRunId,omitempty" xml:"PipelineRunId,omitempty"`
	Producer               *string                `json:"Producer,omitempty" xml:"Producer,omitempty"`
	Type                   *string                `json:"Type,omitempty" xml:"Type,omitempty"`
	Value                  *string                `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPipelineRunsStatusResponseBodyOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunsStatusResponseBodyOutputs) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsStatusResponseBodyOutputs) SetExpandableArtifactName(v string) *ListPipelineRunsStatusResponseBodyOutputs {
	s.ExpandableArtifactName = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyOutputs) SetExpandedArtifactIndex(v int32) *ListPipelineRunsStatusResponseBodyOutputs {
	s.ExpandedArtifactIndex = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyOutputs) SetGmtCreateTime(v string) *ListPipelineRunsStatusResponseBodyOutputs {
	s.GmtCreateTime = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyOutputs) SetId(v string) *ListPipelineRunsStatusResponseBodyOutputs {
	s.Id = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyOutputs) SetMetadata(v map[string]interface{}) *ListPipelineRunsStatusResponseBodyOutputs {
	s.Metadata = v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyOutputs) SetName(v string) *ListPipelineRunsStatusResponseBodyOutputs {
	s.Name = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyOutputs) SetNodeId(v string) *ListPipelineRunsStatusResponseBodyOutputs {
	s.NodeId = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyOutputs) SetPipelineRunId(v string) *ListPipelineRunsStatusResponseBodyOutputs {
	s.PipelineRunId = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyOutputs) SetProducer(v string) *ListPipelineRunsStatusResponseBodyOutputs {
	s.Producer = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyOutputs) SetType(v string) *ListPipelineRunsStatusResponseBodyOutputs {
	s.Type = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyOutputs) SetValue(v string) *ListPipelineRunsStatusResponseBodyOutputs {
	s.Value = &v
	return s
}

type ListPipelineRunsStatusResponseBodyPipelineRuns struct {
	IsDeleted     *bool   `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NodeId        *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	ParentUserId  *string `json:"ParentUserId,omitempty" xml:"ParentUserId,omitempty"`
	PipelineRunId *string `json:"PipelineRunId,omitempty" xml:"PipelineRunId,omitempty"`
	SourceId      *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListPipelineRunsStatusResponseBodyPipelineRuns) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunsStatusResponseBodyPipelineRuns) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsStatusResponseBodyPipelineRuns) SetIsDeleted(v bool) *ListPipelineRunsStatusResponseBodyPipelineRuns {
	s.IsDeleted = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyPipelineRuns) SetName(v string) *ListPipelineRunsStatusResponseBodyPipelineRuns {
	s.Name = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyPipelineRuns) SetNodeId(v string) *ListPipelineRunsStatusResponseBodyPipelineRuns {
	s.NodeId = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyPipelineRuns) SetParentUserId(v string) *ListPipelineRunsStatusResponseBodyPipelineRuns {
	s.ParentUserId = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyPipelineRuns) SetPipelineRunId(v string) *ListPipelineRunsStatusResponseBodyPipelineRuns {
	s.PipelineRunId = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyPipelineRuns) SetSourceId(v string) *ListPipelineRunsStatusResponseBodyPipelineRuns {
	s.SourceId = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyPipelineRuns) SetStatus(v string) *ListPipelineRunsStatusResponseBodyPipelineRuns {
	s.Status = &v
	return s
}

func (s *ListPipelineRunsStatusResponseBodyPipelineRuns) SetUserId(v string) *ListPipelineRunsStatusResponseBodyPipelineRuns {
	s.UserId = &v
	return s
}

type ListPipelineRunsStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPipelineRunsStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPipelineRunsStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineRunsStatusResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsStatusResponse) SetHeaders(v map[string]*string) *ListPipelineRunsStatusResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineRunsStatusResponse) SetStatusCode(v int32) *ListPipelineRunsStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineRunsStatusResponse) SetBody(v *ListPipelineRunsStatusResponseBody) *ListPipelineRunsStatusResponse {
	s.Body = v
	return s
}

type ListPipelinesRequest struct {
	FuzzyMatching      *bool   `json:"FuzzyMatching,omitempty" xml:"FuzzyMatching,omitempty"`
	PageNumber         *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PipelineIdentifier *string `json:"PipelineIdentifier,omitempty" xml:"PipelineIdentifier,omitempty"`
	PipelineProvider   *string `json:"PipelineProvider,omitempty" xml:"PipelineProvider,omitempty"`
	PipelineVersion    *string `json:"PipelineVersion,omitempty" xml:"PipelineVersion,omitempty"`
	WorkspaceId        *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListPipelinesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelinesRequest) GoString() string {
	return s.String()
}

func (s *ListPipelinesRequest) SetFuzzyMatching(v bool) *ListPipelinesRequest {
	s.FuzzyMatching = &v
	return s
}

func (s *ListPipelinesRequest) SetPageNumber(v int32) *ListPipelinesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPipelinesRequest) SetPageSize(v int32) *ListPipelinesRequest {
	s.PageSize = &v
	return s
}

func (s *ListPipelinesRequest) SetPipelineIdentifier(v string) *ListPipelinesRequest {
	s.PipelineIdentifier = &v
	return s
}

func (s *ListPipelinesRequest) SetPipelineProvider(v string) *ListPipelinesRequest {
	s.PipelineProvider = &v
	return s
}

func (s *ListPipelinesRequest) SetPipelineVersion(v string) *ListPipelinesRequest {
	s.PipelineVersion = &v
	return s
}

func (s *ListPipelinesRequest) SetWorkspaceId(v string) *ListPipelinesRequest {
	s.WorkspaceId = &v
	return s
}

type ListPipelinesResponseBody struct {
	Pipelines  []*ListPipelinesResponseBodyPipelines `json:"Pipelines,omitempty" xml:"Pipelines,omitempty" type:"Repeated"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPipelinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBody) SetPipelines(v []*ListPipelinesResponseBodyPipelines) *ListPipelinesResponseBody {
	s.Pipelines = v
	return s
}

func (s *ListPipelinesResponseBody) SetRequestId(v string) *ListPipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelinesResponseBody) SetTotalCount(v int64) *ListPipelinesResponseBody {
	s.TotalCount = &v
	return s
}

type ListPipelinesResponseBodyPipelines struct {
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Identifier      *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	PipelineId      *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	Provider        *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	Uuid            *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListPipelinesResponseBodyPipelines) String() string {
	return tea.Prettify(s)
}

func (s ListPipelinesResponseBodyPipelines) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyPipelines) SetGmtCreateTime(v string) *ListPipelinesResponseBodyPipelines {
	s.GmtCreateTime = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetGmtModifiedTime(v string) *ListPipelinesResponseBodyPipelines {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetIdentifier(v string) *ListPipelinesResponseBodyPipelines {
	s.Identifier = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetPipelineId(v string) *ListPipelinesResponseBodyPipelines {
	s.PipelineId = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetProvider(v string) *ListPipelinesResponseBodyPipelines {
	s.Provider = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetUuid(v string) *ListPipelinesResponseBodyPipelines {
	s.Uuid = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetVersion(v string) *ListPipelinesResponseBodyPipelines {
	s.Version = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetWorkspaceId(v string) *ListPipelinesResponseBodyPipelines {
	s.WorkspaceId = &v
	return s
}

type ListPipelinesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPipelinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPipelinesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPipelinesResponse) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponse) SetHeaders(v map[string]*string) *ListPipelinesResponse {
	s.Headers = v
	return s
}

func (s *ListPipelinesResponse) SetStatusCode(v int32) *ListPipelinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelinesResponse) SetBody(v *ListPipelinesResponseBody) *ListPipelinesResponse {
	s.Body = v
	return s
}

type RerunPipelineRunResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RerunPipelineRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RerunPipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *RerunPipelineRunResponseBody) SetRequestId(v string) *RerunPipelineRunResponseBody {
	s.RequestId = &v
	return s
}

type RerunPipelineRunResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RerunPipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RerunPipelineRunResponse) String() string {
	return tea.Prettify(s)
}

func (s RerunPipelineRunResponse) GoString() string {
	return s.String()
}

func (s *RerunPipelineRunResponse) SetHeaders(v map[string]*string) *RerunPipelineRunResponse {
	s.Headers = v
	return s
}

func (s *RerunPipelineRunResponse) SetStatusCode(v int32) *RerunPipelineRunResponse {
	s.StatusCode = &v
	return s
}

func (s *RerunPipelineRunResponse) SetBody(v *RerunPipelineRunResponseBody) *RerunPipelineRunResponse {
	s.Body = v
	return s
}

type StartPipelineRunResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartPipelineRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartPipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *StartPipelineRunResponseBody) SetRequestId(v string) *StartPipelineRunResponseBody {
	s.RequestId = &v
	return s
}

type StartPipelineRunResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartPipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartPipelineRunResponse) String() string {
	return tea.Prettify(s)
}

func (s StartPipelineRunResponse) GoString() string {
	return s.String()
}

func (s *StartPipelineRunResponse) SetHeaders(v map[string]*string) *StartPipelineRunResponse {
	s.Headers = v
	return s
}

func (s *StartPipelineRunResponse) SetStatusCode(v int32) *StartPipelineRunResponse {
	s.StatusCode = &v
	return s
}

func (s *StartPipelineRunResponse) SetBody(v *StartPipelineRunResponseBody) *StartPipelineRunResponse {
	s.Body = v
	return s
}

type TerminatePipelineRunResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TerminatePipelineRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TerminatePipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *TerminatePipelineRunResponseBody) SetRequestId(v string) *TerminatePipelineRunResponseBody {
	s.RequestId = &v
	return s
}

type TerminatePipelineRunResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TerminatePipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TerminatePipelineRunResponse) String() string {
	return tea.Prettify(s)
}

func (s TerminatePipelineRunResponse) GoString() string {
	return s.String()
}

func (s *TerminatePipelineRunResponse) SetHeaders(v map[string]*string) *TerminatePipelineRunResponse {
	s.Headers = v
	return s
}

func (s *TerminatePipelineRunResponse) SetStatusCode(v int32) *TerminatePipelineRunResponse {
	s.StatusCode = &v
	return s
}

func (s *TerminatePipelineRunResponse) SetBody(v *TerminatePipelineRunResponseBody) *TerminatePipelineRunResponse {
	s.Body = v
	return s
}

type UpdatePipelineRequest struct {
	Manifest *string `json:"Manifest,omitempty" xml:"Manifest,omitempty"`
}

func (s UpdatePipelineRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineRequest) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRequest) SetManifest(v string) *UpdatePipelineRequest {
	s.Manifest = &v
	return s
}

type UpdatePipelineResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePipelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePipelineResponseBody) SetRequestId(v string) *UpdatePipelineResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePipelineResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdatePipelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePipelineResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineResponse) GoString() string {
	return s.String()
}

func (s *UpdatePipelineResponse) SetHeaders(v map[string]*string) *UpdatePipelineResponse {
	s.Headers = v
	return s
}

func (s *UpdatePipelineResponse) SetStatusCode(v int32) *UpdatePipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePipelineResponse) SetBody(v *UpdatePipelineResponseBody) *UpdatePipelineResponse {
	s.Body = v
	return s
}

type UpdatePipelineRunRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdatePipelineRunRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineRunRequest) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRunRequest) SetName(v string) *UpdatePipelineRunRequest {
	s.Name = &v
	return s
}

type UpdatePipelineRunResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePipelineRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRunResponseBody) SetRequestId(v string) *UpdatePipelineRunResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePipelineRunResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdatePipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePipelineRunResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineRunResponse) GoString() string {
	return s.String()
}

func (s *UpdatePipelineRunResponse) SetHeaders(v map[string]*string) *UpdatePipelineRunResponse {
	s.Headers = v
	return s
}

func (s *UpdatePipelineRunResponse) SetStatusCode(v int32) *UpdatePipelineRunResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePipelineRunResponse) SetBody(v *UpdatePipelineRunResponseBody) *UpdatePipelineRunResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("paiflow"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePipelineWithOptions(request *CreatePipelineRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePipelineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Manifest)) {
		body["Manifest"] = request.Manifest
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePipeline"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/pipelines"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePipeline(request *CreatePipelineRequest) (_result *CreatePipelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePipelineResponse{}
	_body, _err := client.CreatePipelineWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePipelineRunWithOptions(request *CreatePipelineRunRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePipelineRunResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.Arguments)) {
		body["Arguments"] = request.Arguments
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NoConfirmRequired)) {
		body["NoConfirmRequired"] = request.NoConfirmRequired
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["Options"] = request.Options
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		body["PipelineId"] = request.PipelineId
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineManifest)) {
		body["PipelineManifest"] = request.PipelineManifest
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		body["SourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePipelineRun"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/pipelineruns"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePipelineRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePipelineRun(request *CreatePipelineRunRequest) (_result *CreatePipelineRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePipelineRunResponse{}
	_body, _err := client.CreatePipelineRunWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePipelineWithOptions(PipelineId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeletePipelineResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePipeline"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(PipelineId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePipeline(PipelineId *string) (_result *DeletePipelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePipelineResponse{}
	_body, _err := client.DeletePipelineWithOptions(PipelineId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePipelineRunWithOptions(PipelineRunId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeletePipelineRunResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePipelineRun"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/pipelineruns/" + tea.StringValue(openapiutil.GetEncodeParam(PipelineRunId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePipelineRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePipelineRun(PipelineRunId *string) (_result *DeletePipelineRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePipelineRunResponse{}
	_body, _err := client.DeletePipelineRunWithOptions(PipelineRunId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPipelineWithOptions(PipelineId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPipelineResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipeline"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(PipelineId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPipeline(PipelineId *string) (_result *GetPipelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPipelineResponse{}
	_body, _err := client.GetPipelineWithOptions(PipelineId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPipelineRunWithOptions(PipelineRunId *string, request *GetPipelineRunRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPipelineRunResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ManifestType)) {
		query["ManifestType"] = request.ManifestType
	}

	if !tea.BoolValue(util.IsUnset(request.TokenId)) {
		query["TokenId"] = request.TokenId
	}

	if !tea.BoolValue(util.IsUnset(request.Verbose)) {
		query["Verbose"] = request.Verbose
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipelineRun"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/pipelineruns/" + tea.StringValue(openapiutil.GetEncodeParam(PipelineRunId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPipelineRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPipelineRun(PipelineRunId *string, request *GetPipelineRunRequest) (_result *GetPipelineRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPipelineRunResponse{}
	_body, _err := client.GetPipelineRunWithOptions(PipelineRunId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPipelineRunNodeWithOptions(PipelineRunId *string, NodeId *string, request *GetPipelineRunNodeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPipelineRunNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Depth)) {
		query["Depth"] = request.Depth
	}

	if !tea.BoolValue(util.IsUnset(request.TokenId)) {
		query["TokenId"] = request.TokenId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipelineRunNode"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/pipelineruns/" + tea.StringValue(openapiutil.GetEncodeParam(PipelineRunId)) + "/nodes/" + tea.StringValue(openapiutil.GetEncodeParam(NodeId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPipelineRunNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPipelineRunNode(PipelineRunId *string, NodeId *string, request *GetPipelineRunNodeRequest) (_result *GetPipelineRunNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPipelineRunNodeResponse{}
	_body, _err := client.GetPipelineRunNodeWithOptions(PipelineRunId, NodeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPipelineRunNodeLogsWithOptions(PipelineRunId *string, NodeId *string, request *ListPipelineRunNodeLogsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPipelineRunNodeLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FromTimeInSeconds)) {
		query["FromTimeInSeconds"] = request.FromTimeInSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Reverse)) {
		query["Reverse"] = request.Reverse
	}

	if !tea.BoolValue(util.IsUnset(request.ToTimeInSeconds)) {
		query["ToTimeInSeconds"] = request.ToTimeInSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.TokenId)) {
		query["TokenId"] = request.TokenId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPipelineRunNodeLogs"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/pipelineruns/" + tea.StringValue(openapiutil.GetEncodeParam(PipelineRunId)) + "/nodes/" + tea.StringValue(openapiutil.GetEncodeParam(NodeId)) + "/logs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPipelineRunNodeLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPipelineRunNodeLogs(PipelineRunId *string, NodeId *string, request *ListPipelineRunNodeLogsRequest) (_result *ListPipelineRunNodeLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelineRunNodeLogsResponse{}
	_body, _err := client.ListPipelineRunNodeLogsWithOptions(PipelineRunId, NodeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPipelineRunNodeOutputsWithOptions(PipelineRunId *string, NodeId *string, request *ListPipelineRunNodeOutputsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPipelineRunNodeOutputsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Depth)) {
		query["Depth"] = request.Depth
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.TokenId)) {
		query["TokenId"] = request.TokenId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPipelineRunNodeOutputs"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/pipelineruns/" + tea.StringValue(openapiutil.GetEncodeParam(PipelineRunId)) + "/nodes/" + tea.StringValue(openapiutil.GetEncodeParam(NodeId)) + "/outputs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPipelineRunNodeOutputsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPipelineRunNodeOutputs(PipelineRunId *string, NodeId *string, request *ListPipelineRunNodeOutputsRequest) (_result *ListPipelineRunNodeOutputsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelineRunNodeOutputsResponse{}
	_body, _err := client.ListPipelineRunNodeOutputsWithOptions(PipelineRunId, NodeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPipelineRunNodeStatusWithOptions(PipelineRunId *string, NodeId *string, request *ListPipelineRunNodeStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPipelineRunNodeStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Depth)) {
		query["Depth"] = request.Depth
	}

	if !tea.BoolValue(util.IsUnset(request.TokenId)) {
		query["TokenId"] = request.TokenId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPipelineRunNodeStatus"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/pipelineruns/" + tea.StringValue(openapiutil.GetEncodeParam(PipelineRunId)) + "/nodes/" + tea.StringValue(openapiutil.GetEncodeParam(NodeId)) + "/status"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPipelineRunNodeStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPipelineRunNodeStatus(PipelineRunId *string, NodeId *string, request *ListPipelineRunNodeStatusRequest) (_result *ListPipelineRunNodeStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelineRunNodeStatusResponse{}
	_body, _err := client.ListPipelineRunNodeStatusWithOptions(PipelineRunId, NodeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPipelineRunsWithOptions(request *ListPipelineRunsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPipelineRunsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineIds)) {
		query["PipelineIds"] = request.PipelineIds
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineRunId)) {
		query["PipelineRunId"] = request.PipelineRunId
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		query["SourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPipelineRuns"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/pipelineruns"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPipelineRunsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPipelineRuns(request *ListPipelineRunsRequest) (_result *ListPipelineRunsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelineRunsResponse{}
	_body, _err := client.ListPipelineRunsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPipelineRunsStatusWithOptions(request *ListPipelineRunsStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPipelineRunsStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Nodes)) {
		body["Nodes"] = request.Nodes
	}

	if !tea.BoolValue(util.IsUnset(request.OutputType)) {
		body["OutputType"] = request.OutputType
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineRuns)) {
		body["PipelineRuns"] = request.PipelineRuns
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPipelineRunsStatus"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/pipelineruns"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPipelineRunsStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPipelineRunsStatus(request *ListPipelineRunsStatusRequest) (_result *ListPipelineRunsStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelineRunsStatusResponse{}
	_body, _err := client.ListPipelineRunsStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPipelinesWithOptions(request *ListPipelinesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPipelinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FuzzyMatching)) {
		query["FuzzyMatching"] = request.FuzzyMatching
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineIdentifier)) {
		query["PipelineIdentifier"] = request.PipelineIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineProvider)) {
		query["PipelineProvider"] = request.PipelineProvider
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineVersion)) {
		query["PipelineVersion"] = request.PipelineVersion
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPipelines"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/pipelines"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPipelinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPipelines(request *ListPipelinesRequest) (_result *ListPipelinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelinesResponse{}
	_body, _err := client.ListPipelinesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RerunPipelineRunWithOptions(PipelineRunId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RerunPipelineRunResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RerunPipelineRun"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/pipelineruns/" + tea.StringValue(openapiutil.GetEncodeParam(PipelineRunId)) + "/rerun"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RerunPipelineRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RerunPipelineRun(PipelineRunId *string) (_result *RerunPipelineRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RerunPipelineRunResponse{}
	_body, _err := client.RerunPipelineRunWithOptions(PipelineRunId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartPipelineRunWithOptions(PipelineRunId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartPipelineRunResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartPipelineRun"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/pipelineruns/" + tea.StringValue(openapiutil.GetEncodeParam(PipelineRunId)) + "/start"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartPipelineRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartPipelineRun(PipelineRunId *string) (_result *StartPipelineRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartPipelineRunResponse{}
	_body, _err := client.StartPipelineRunWithOptions(PipelineRunId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TerminatePipelineRunWithOptions(PipelineRunId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TerminatePipelineRunResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("TerminatePipelineRun"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/pipelineruns/" + tea.StringValue(openapiutil.GetEncodeParam(PipelineRunId)) + "/termination"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TerminatePipelineRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TerminatePipelineRun(PipelineRunId *string) (_result *TerminatePipelineRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TerminatePipelineRunResponse{}
	_body, _err := client.TerminatePipelineRunWithOptions(PipelineRunId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePipelineWithOptions(PipelineId *string, request *UpdatePipelineRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdatePipelineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Manifest)) {
		body["Manifest"] = request.Manifest
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePipeline"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/pipelines/" + tea.StringValue(openapiutil.GetEncodeParam(PipelineId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePipeline(PipelineId *string, request *UpdatePipelineRequest) (_result *UpdatePipelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePipelineResponse{}
	_body, _err := client.UpdatePipelineWithOptions(PipelineId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePipelineRunWithOptions(PipelineRunId *string, request *UpdatePipelineRunRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdatePipelineRunResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePipelineRun"),
		Version:     tea.String("2021-02-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/pipelineruns/" + tea.StringValue(openapiutil.GetEncodeParam(PipelineRunId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePipelineRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePipelineRun(PipelineRunId *string, request *UpdatePipelineRunRequest) (_result *UpdatePipelineRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePipelineRunResponse{}
	_body, _err := client.UpdatePipelineRunWithOptions(PipelineRunId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
