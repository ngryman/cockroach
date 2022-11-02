// Copyright 2021 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package schedulebase

import (
	"context"
	"fmt"

	"github.com/cockroachdb/cockroach/pkg/scheduledjobs"
	"github.com/cockroachdb/cockroach/pkg/security/username"
	"github.com/cockroachdb/cockroach/pkg/sql"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/sessiondata"
)

// CheckScheduleAlreadyExists returns true if a schedule with the same label already exists.
func CheckScheduleAlreadyExists(
	ctx context.Context, p sql.PlanHookState, scheduleLabel string,
) (bool, error) {

	row, err := p.ExecCfg().InternalExecutor.QueryRowEx(ctx, "check-sched",
		p.Txn(), sessiondata.InternalExecutorOverride{User: username.RootUserName()},
		fmt.Sprintf("SELECT count(schedule_name) FROM %s WHERE schedule_name = '%s'",
			scheduledjobs.ProdJobSchedulerEnv.ScheduledJobsTableName(), scheduleLabel))

	if err != nil {
		return false, err
	}
	return int64(tree.MustBeDInt(row[0])) != 0, nil
}
