package publish

import (
	scenarios "bigwhite/autotester/scenarios"
	"testing"
)

func Test_Publish_S0001_PublishOK(t *testing.T) {
	client, testCaseTeardown, err := scenarios.TestCaseSetup(addr, nil)
	if err != nil {
		t.Errorf("want ok, got %v", err)
		return
	}
	defer testCaseTeardown()

	//TBD:xxx
	_ = client
	_ = err
}

func Test_Publish_S0002_PublishFail(t *testing.T) {
}
