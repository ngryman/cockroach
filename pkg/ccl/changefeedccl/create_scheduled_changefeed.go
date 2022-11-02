// Copyright 2022 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

package changefeedccl

import (
	"context"

	"github.com/cockroachdb/cockroach/pkg/ccl/changefeedccl/changefeedvalidators"
	"github.com/cockroachdb/cockroach/pkg/sql"
	"github.com/cockroachdb/cockroach/pkg/sql/catalog/colinfo"
	"github.com/cockroachdb/cockroach/pkg/sql/exprutil"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/errors"
)

const scheduleChangefeedOp = "CREATE SCHEDULE FOR CHANGEFEED"

const (
	optOnExecFailure     = "on_execution_failure"
	optOnPreviousRunning = "on_previous_running"
)

var scheduledChangefeedOptionExpectValues = map[string]exprutil.KVStringOptValidate{
	optOnExecFailure:     exprutil.KVStringOptRequireValue,
	optOnPreviousRunning: exprutil.KVStringOptRequireValue,
}

var scheduledChangefeedHeader = colinfo.ResultColumns{
	{Name: "schedule_id", Typ: types.Int},
	{Name: "label", Typ: types.String},
	{Name: "status", Typ: types.String},
	{Name: "schedule", Typ: types.String},
	{Name: "changefeed_stmt", Typ: types.String},
}

func createChangefeedScheduleHook(
	ctx context.Context, stmt tree.Statement, p sql.PlanHookState,
) (sql.PlanHookRowFn, colinfo.ResultColumns, []sql.PlanNode, bool, error) {
	//schedule, ok := stmt.(*tree.ScheduledExport)
	//if !ok {
	//	return nil, nil, nil, false, nil
	//}
	//
	//eval, err := makeScheduledExportEval(ctx, p, schedule)
	//if err != nil {
	//	return nil, nil, nil, false, err
	//}
	//
	//fn := func(ctx context.Context, _ []sql.PlanNode, resultsCh chan<- tree.Datums) error {
	//	err := doCreateExportSchedule(ctx, p, eval, resultsCh)
	//	if err != nil {
	//		telemetry.Count("scheduled-export.create.failed")
	//		return err
	//	}
	//
	//	return nil
	//}
	//return fn, schedulebase.CreateScheduleHeader, nil, false, nil
	return nil, nil, nil, false, errors.Newf("not implemented")
}

func createChangefeedScheduleTypeCheck(
	ctx context.Context, stmt tree.Statement, p sql.PlanHookState,
) (matched bool, header colinfo.ResultColumns, _ error) {
	schedule, ok := stmt.(*tree.ScheduledChangefeed)
	if !ok {
		return false, nil, nil
	}

	changefeedStmt := schedule.CreateChangefeed
	if changefeedStmt == nil {
		return false, nil, nil
	}
	if err := exprutil.TypeCheck(ctx, scheduleChangefeedOp, p.SemaCtx(),
		exprutil.Strings{
			changefeedStmt.SinkURI,
			schedule.Recurrence,
			schedule.ScheduleLabelSpec.Label,
		},
		&exprutil.KVOptions{
			KVOptions:  changefeedStmt.Options,
			Validation: changefeedvalidators.CreateOptionValidations,
		},
		&exprutil.KVOptions{
			KVOptions:  schedule.ScheduleOptions,
			Validation: scheduledChangefeedOptionExpectValues,
		},
	); err != nil {
		return false, nil, err
	}

	return true, scheduledChangefeedHeader, nil
}

func init() {
	sql.AddPlanHook("schedule changefeed", createChangefeedScheduleHook, createChangefeedScheduleTypeCheck)
}
