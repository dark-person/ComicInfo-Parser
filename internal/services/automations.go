package services

type AutomationResp struct {
	jobId        int    // Job ID for tracking automation progress
	errorMessage string // Error message if any
}

func (s *ComicService) StartAutomation(allowOverwrite bool) {

}

func (s *ComicService) GetAutomationProgress(jobId int) {

}

func (s *ComicService) CancelAutomation(jobId int) {

}
