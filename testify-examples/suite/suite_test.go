package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExampleSuite struct {
	suite.Suite
	indent int
}

func (suite *ExampleSuite) indents() (result string) {
	for i := 0; i < suite.indent; i++ {
		result += "----"
	}
	return
}

func (suite *ExampleSuite) HandleStats(suiteName string, stats *suite.SuiteInformation) {
	fmt.Println(suiteName, *stats)
}

func (suite *ExampleSuite) SetupSuite() {
	fmt.Println("Suite setup")
}

func (suite *ExampleSuite) TearDownSuite() {
	fmt.Println("Suite teardown")
}

func (suite *ExampleSuite) SetupTest() {
	suite.indent++
	fmt.Println(suite.indents(), "Test setup")
}

func (suite *ExampleSuite) TearDownTest() {
	fmt.Println(suite.indents(), "Test teardown")
	suite.indent--
}

func (suite *ExampleSuite) BeforeTest(suiteName, testName string) {
	suite.indent++
	fmt.Printf("%sBefore %s.%s\n", suite.indents(), suiteName, testName)
}

func (suite *ExampleSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("%sAfter %s.%s\n", suite.indents(), suiteName, testName)
	suite.indent--
}

func (suite *ExampleSuite) SetupSubTest() {
	suite.indent++
	fmt.Println(suite.indents(), "SubTest setup")
}

func (suite *ExampleSuite) TearDownSubTest() {
	fmt.Println(suite.indents(), "SubTest teardown")
	suite.indent--
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *ExampleSuite) TestCase1() {
	suite.indent++
	defer func() {
		fmt.Println(suite.indents(), "End TestCase1")
		suite.indent--
	}()

	fmt.Println(suite.indents(), "Begin TestCase1")

	suite.Run("case1-subtest1", func() {
		suite.indent++
		fmt.Println(suite.indents(), "Begin TestCase1.Subtest1")
		fmt.Println(suite.indents(), "End TestCase1.Subtest1")
		suite.indent--
	})
	suite.Run("case1-subtest2", func() {
		suite.indent++
		fmt.Println(suite.indents(), "Begin TestCase1.Subtest2")
		fmt.Println(suite.indents(), "End TestCase1.Subtest2")
		suite.indent--
	})
}

func (suite *ExampleSuite) TestCase2() {
	suite.indent++
	defer func() {
		fmt.Println(suite.indents(), "End TestCase2")
		suite.indent--
	}()
	fmt.Println(suite.indents(), "Begin TestCase2")

	suite.Run("case2-subtest1", func() {
		suite.indent++
		fmt.Println(suite.indents(), "Begin TestCase2.Subtest1")
		fmt.Println(suite.indents(), "End TestCase2.Subtest1")
		suite.indent--
	})
}

func TestExampleSuite(t *testing.T) {
	suite.Run(t, new(ExampleSuite))
}
