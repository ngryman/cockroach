// Code generated by "stringer"; DO NOT EDIT.

package clusterversion

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[invalidVersionKey - -1]
	_ = x[V22_1-0]
	_ = x[V22_2Start-1]
	_ = x[V22_2LocalTimestamps-2]
	_ = x[V22_2PebbleFormatSplitUserKeysMarkedCompacted-3]
	_ = x[V22_2EnsurePebbleFormatVersionRangeKeys-4]
	_ = x[V22_2EnablePebbleFormatVersionRangeKeys-5]
	_ = x[V22_2TrigramInvertedIndexes-6]
	_ = x[V22_2RemoveGrantPrivilege-7]
	_ = x[V22_2MVCCRangeTombstones-8]
	_ = x[V22_2UpgradeSequenceToBeReferencedByID-9]
	_ = x[V22_2SampledStmtDiagReqs-10]
	_ = x[V22_2AddSSTableTombstones-11]
	_ = x[V22_2SystemPrivilegesTable-12]
	_ = x[V22_2EnablePredicateProjectionChangefeed-13]
	_ = x[V22_2AlterSystemSQLInstancesAddLocality-14]
	_ = x[V22_2SystemExternalConnectionsTable-15]
	_ = x[V22_2AlterSystemStatementStatisticsAddIndexRecommendations-16]
	_ = x[V22_2RoleIDSequence-17]
	_ = x[V22_2AddSystemUserIDColumn-18]
	_ = x[V22_2SystemUsersIDColumnIsBackfilled-19]
	_ = x[V22_2SetSystemUsersUserIDColumnNotNull-20]
	_ = x[V22_2SQLSchemaTelemetryScheduledJobs-21]
	_ = x[V22_2SchemaChangeSupportsCreateFunction-22]
	_ = x[V22_2DeleteRequestReturnKey-23]
	_ = x[V22_2PebbleFormatPrePebblev1Marked-24]
	_ = x[V22_2RoleOptionsTableHasIDColumn-25]
	_ = x[V22_2RoleOptionsIDColumnIsBackfilled-26]
	_ = x[V22_2SetRoleOptionsUserIDColumnNotNull-27]
	_ = x[V22_2UseDelRangeInGCJob-28]
	_ = x[V22_2WaitedForDelRangeInGCJob-29]
	_ = x[V22_2RangefeedUseOneStreamPerNode-30]
	_ = x[V22_2NoNonMVCCAddSSTable-31]
	_ = x[V22_2GCHintInReplicaState-32]
	_ = x[V22_2UpdateInvalidColumnIDsInSequenceBackReferences-33]
	_ = x[V22_2TTLDistSQL-34]
	_ = x[V22_2PrioritizeSnapshots-35]
	_ = x[V22_2EnableLeaseUpgrade-36]
	_ = x[V22_2SupportAssumeRoleAuth-37]
	_ = x[V22_2FixUserfileRelatedDescriptorCorruption-38]
	_ = x[V22_2-39]
	_ = x[V23_1Start-40]
	_ = x[V23_1TenantNames-41]
	_ = x[V23_1DescIDSequenceForSystemTenant-42]
}

const _Key_name = "invalidVersionKeyV22_1V22_2StartV22_2LocalTimestampsV22_2PebbleFormatSplitUserKeysMarkedCompactedV22_2EnsurePebbleFormatVersionRangeKeysV22_2EnablePebbleFormatVersionRangeKeysV22_2TrigramInvertedIndexesV22_2RemoveGrantPrivilegeV22_2MVCCRangeTombstonesV22_2UpgradeSequenceToBeReferencedByIDV22_2SampledStmtDiagReqsV22_2AddSSTableTombstonesV22_2SystemPrivilegesTableV22_2EnablePredicateProjectionChangefeedV22_2AlterSystemSQLInstancesAddLocalityV22_2SystemExternalConnectionsTableV22_2AlterSystemStatementStatisticsAddIndexRecommendationsV22_2RoleIDSequenceV22_2AddSystemUserIDColumnV22_2SystemUsersIDColumnIsBackfilledV22_2SetSystemUsersUserIDColumnNotNullV22_2SQLSchemaTelemetryScheduledJobsV22_2SchemaChangeSupportsCreateFunctionV22_2DeleteRequestReturnKeyV22_2PebbleFormatPrePebblev1MarkedV22_2RoleOptionsTableHasIDColumnV22_2RoleOptionsIDColumnIsBackfilledV22_2SetRoleOptionsUserIDColumnNotNullV22_2UseDelRangeInGCJobV22_2WaitedForDelRangeInGCJobV22_2RangefeedUseOneStreamPerNodeV22_2NoNonMVCCAddSSTableV22_2GCHintInReplicaStateV22_2UpdateInvalidColumnIDsInSequenceBackReferencesV22_2TTLDistSQLV22_2PrioritizeSnapshotsV22_2EnableLeaseUpgradeV22_2SupportAssumeRoleAuthV22_2FixUserfileRelatedDescriptorCorruptionV22_2V23_1StartV23_1TenantNamesV23_1DescIDSequenceForSystemTenant"

var _Key_index = [...]uint16{0, 17, 22, 32, 52, 97, 136, 175, 202, 227, 251, 289, 313, 338, 364, 404, 443, 478, 536, 555, 581, 617, 655, 691, 730, 757, 791, 823, 859, 897, 920, 949, 982, 1006, 1031, 1082, 1097, 1121, 1144, 1170, 1213, 1218, 1228, 1244, 1278}

func (i Key) String() string {
	i -= -1
	if i < 0 || i >= Key(len(_Key_index)-1) {
		return "Key(" + strconv.FormatInt(int64(i+-1), 10) + ")"
	}
	return _Key_name[_Key_index[i]:_Key_index[i+1]]
}
