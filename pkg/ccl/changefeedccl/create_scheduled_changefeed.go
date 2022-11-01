// Copyright 2020 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

package changefeedccl

import (
	"context"

	"github.com/cockroachdb/cockroach/pkg/sql"
	"github.com/cockroachdb/cockroach/pkg/sql/catalog/colinfo"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/errors"
)

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
	//changefeedStmt := getChangefeedStatement(stmt)
	//if changefeedStmt == nil {
	//	return false, nil, nil
	//}
	//if err := exprutil.TypeCheck(ctx, `CREATE CHANGEFEED`, p.SemaCtx(),
	//	exprutil.Strings{changefeedStmt.SinkURI},
	//	&exprutil.KVOptions{
	//		KVOptions:  changefeedStmt.Options,
	//		Validation: changefeedvalidators.CreateOptionValidations,
	//	},
	//); err != nil {
	//	return false, nil, err
	//}
	//unspecifiedSink := changefeedStmt.SinkURI == nil
	//if unspecifiedSink {
	//	return true, sinklessHeader, nil
	//}
	//return true, withSinkHeader, nil
	return true, withSinkHeader, nil
}

func init() {
	sql.AddPlanHook("schedule changefeed", createChangefeedScheduleHook, createChangefeedScheduleTypeCheck)
}
