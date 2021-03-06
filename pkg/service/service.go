package service

import (
	"github.com/ecnuvj/vhoj_db/pkg/dao/mapper/problem_mapper"
	"github.com/ecnuvj/vhoj_db/pkg/dao/mapper/submission_mapper"
	"github.com/ecnuvj/vhoj_db/pkg/dao/mapper/user_mapper"
	"github.com/ecnuvj/vhoj_db/pkg/dao/model"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
	"github.com/ecnuvj/vhoj_submitter/pkg/manager"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/adapter/holder"
	"github.com/ecnuvj/vhoj_submitter/pkg/sdk/submitterpb"
	"github.com/ecnuvj/vhoj_submitter/pkg/util"
)

type SubmitService struct {
}

func (ss *SubmitService) SubmitCode(submission *model.Submission) (*model.Submission, error) {
	submission, info, err := ss.chooseSuitableRemoteOJ(submission)
	if err != nil {
		return nil, err
	}
	err = user_mapper.UserMapper.AddUserSubmitCountById(submission.UserId)
	if err != nil {
		return nil, err
	}
	manager.SubmitCode(info)
	return submission, nil
}

func (ss *SubmitService) ReSubmitCode(submissionId uint) error {
	_ = submission_mapper.SubmissionMapper.ResetSubmissionById(submissionId)
	submission, err := submission_mapper.SubmissionMapper.FindSubmissionById(submissionId)
	if err != nil {
		return err
	}
	_, err = ss.SubmitCode(submission)
	if err != nil {
		return err
	}
	return nil
}

func (ss *SubmitService) chooseSuitableRemoteOJ(submission *model.Submission) (*model.Submission, *common.SubmissionInfo, error) {
	problemGroups, err := problem_mapper.ProblemMapper.FindGroupProblemsById(submission.ProblemId)
	if err != nil {
		return nil, nil, err
	}
	//todo 选择最优oj
	remoteOJ := problemGroups[0].RemoteOJ
	remoteProblemId := problemGroups[0].RemoteProblemId

	submission.RemoteOJ = remoteOJ
	submission, err = submission_mapper.SubmissionMapper.AddOrModifySubmission(submission)
	if err != nil {
		return nil, nil, err
	}
	return submission, &common.SubmissionInfo{
		SubmissionID:    submission.ID,
		ProblemID:       submission.ProblemId,
		RemoteOJ:        remoteOJ,
		RemoteProblemId: remoteProblemId,
		RemoteLanguage:  holder.Adapters[remoteOJ].ToOJLanguage(submission.Language),
		SourceCode:      submission.SubmissionCode.SourceCode,
	}, nil
}

func (SubmitService) ListSubmissions(pageNo int32, pageSize int32, condition *common.SubmissionSearchCondition) ([]*submitterpb.Submission, *common.PageInfo, error) {
	submissions, count, err := submission_mapper.SubmissionMapper.FindSubmissions(pageNo, pageSize, &submission_mapper.SearchSubmissionCondition{
		Username:  condition.Username,
		ProblemId: condition.ProblemId,
		Status:    condition.Result,
		Language:  condition.Language,
	})
	if err != nil {
		return nil, nil, err
	}
	return util.ModelSubmissionsToRpcSubmissions(submissions), &common.PageInfo{
		TotalPages: (count + pageSize - 1) / pageSize,
		TotalCount: count,
	}, nil
}
