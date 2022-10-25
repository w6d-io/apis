package v1alpha1_test

import (
	"fmt"
	apis "github.com/w6d-io/apis/pipeline/v1alpha1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/w6d-io/x/errorx"
)

var _ = Describe("Pipelines", func() {
	Context("Helper", func() {
		It("Succeed the validation", func() {
			p := &apis.Pipeline{
				Triggers: []apis.Trigger{{ID: "trigger-1"}},
				Stages: []apis.Stage{
					{
						ID: "stage-1",
						Tasks: []apis.Task{
							{
								ID:         "task-1",
								Conditions: [][]apis.Condition{{{Ref: "trigger-1"}}},
							},
						},
					},
				},
			}
			Expect(p.ConditionValidation()).To(Succeed())
		})
		It("Failed the validation", func() {
			p := &apis.Pipeline{
				Triggers: []apis.Trigger{{ID: "trigger-1"}},
				Stages: []apis.Stage{
					{
						ID: "stage-1",
						Tasks: []apis.Task{
							{
								ID:         "task-1",
								Conditions: [][]apis.Condition{{{Ref: "task-2"}}},
							},
						},
					},
				},
			}
			Expect(p.ConditionValidation()).To(Equal(errorx.New(nil, "task-2")))
		})
		It("Succeed the validation on empty pipeline", func() {
			p := &apis.Pipeline{}
			Expect(p.ConditionValidation()).To(Succeed())
		})
		It("Failed multiple validation", func() {
			p := &apis.Pipeline{
				Triggers: []apis.Trigger{{ID: "trigger-1"}},
				Stages: []apis.Stage{
					{
						ID: "stage-1",
						Tasks: []apis.Task{
							{
								ID:         "task-1",
								Conditions: [][]apis.Condition{{{Ref: "task-20"}}},
							},
							{
								ID:         "task-2",
								Conditions: [][]apis.Condition{{{Ref: "task-30"}}},
							},
							{
								ID:         "task-3",
								Conditions: [][]apis.Condition{{{Ref: "trigger-1"}}},
							},
						},
					},
					{
						ID: "stage-2",
						Tasks: []apis.Task{
							{
								ID:         "task-1",
								Conditions: [][]apis.Condition{{{Ref: "stage-1"}}},
							},
						},
					},
				},
			}
			fmt.Println(p.ConditionValidation())
			Expect(p.ConditionValidation()).To(Equal(errorx.New(nil, "task-20, task-30")))
		})
	})
})
