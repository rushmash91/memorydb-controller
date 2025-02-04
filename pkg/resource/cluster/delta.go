// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package cluster

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.ACLName, b.ko.Spec.ACLName) {
		delta.Add("Spec.ACLName", a.ko.Spec.ACLName, b.ko.Spec.ACLName)
	} else if a.ko.Spec.ACLName != nil && b.ko.Spec.ACLName != nil {
		if *a.ko.Spec.ACLName != *b.ko.Spec.ACLName {
			delta.Add("Spec.ACLName", a.ko.Spec.ACLName, b.ko.Spec.ACLName)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.ACLRef, b.ko.Spec.ACLRef) {
		delta.Add("Spec.ACLRef", a.ko.Spec.ACLRef, b.ko.Spec.ACLRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.AutoMinorVersionUpgrade, b.ko.Spec.AutoMinorVersionUpgrade) {
		delta.Add("Spec.AutoMinorVersionUpgrade", a.ko.Spec.AutoMinorVersionUpgrade, b.ko.Spec.AutoMinorVersionUpgrade)
	} else if a.ko.Spec.AutoMinorVersionUpgrade != nil && b.ko.Spec.AutoMinorVersionUpgrade != nil {
		if *a.ko.Spec.AutoMinorVersionUpgrade != *b.ko.Spec.AutoMinorVersionUpgrade {
			delta.Add("Spec.AutoMinorVersionUpgrade", a.ko.Spec.AutoMinorVersionUpgrade, b.ko.Spec.AutoMinorVersionUpgrade)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Description, b.ko.Spec.Description) {
		delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
	} else if a.ko.Spec.Description != nil && b.ko.Spec.Description != nil {
		if *a.ko.Spec.Description != *b.ko.Spec.Description {
			delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Engine, b.ko.Spec.Engine) {
		delta.Add("Spec.Engine", a.ko.Spec.Engine, b.ko.Spec.Engine)
	} else if a.ko.Spec.Engine != nil && b.ko.Spec.Engine != nil {
		if *a.ko.Spec.Engine != *b.ko.Spec.Engine {
			delta.Add("Spec.Engine", a.ko.Spec.Engine, b.ko.Spec.Engine)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EngineVersion, b.ko.Spec.EngineVersion) {
		delta.Add("Spec.EngineVersion", a.ko.Spec.EngineVersion, b.ko.Spec.EngineVersion)
	} else if a.ko.Spec.EngineVersion != nil && b.ko.Spec.EngineVersion != nil {
		if *a.ko.Spec.EngineVersion != *b.ko.Spec.EngineVersion {
			delta.Add("Spec.EngineVersion", a.ko.Spec.EngineVersion, b.ko.Spec.EngineVersion)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.KMSKeyID, b.ko.Spec.KMSKeyID) {
		delta.Add("Spec.KMSKeyID", a.ko.Spec.KMSKeyID, b.ko.Spec.KMSKeyID)
	} else if a.ko.Spec.KMSKeyID != nil && b.ko.Spec.KMSKeyID != nil {
		if *a.ko.Spec.KMSKeyID != *b.ko.Spec.KMSKeyID {
			delta.Add("Spec.KMSKeyID", a.ko.Spec.KMSKeyID, b.ko.Spec.KMSKeyID)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.MaintenanceWindow, b.ko.Spec.MaintenanceWindow) {
		delta.Add("Spec.MaintenanceWindow", a.ko.Spec.MaintenanceWindow, b.ko.Spec.MaintenanceWindow)
	} else if a.ko.Spec.MaintenanceWindow != nil && b.ko.Spec.MaintenanceWindow != nil {
		if *a.ko.Spec.MaintenanceWindow != *b.ko.Spec.MaintenanceWindow {
			delta.Add("Spec.MaintenanceWindow", a.ko.Spec.MaintenanceWindow, b.ko.Spec.MaintenanceWindow)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Name, b.ko.Spec.Name) {
		delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
	} else if a.ko.Spec.Name != nil && b.ko.Spec.Name != nil {
		if *a.ko.Spec.Name != *b.ko.Spec.Name {
			delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.NodeType, b.ko.Spec.NodeType) {
		delta.Add("Spec.NodeType", a.ko.Spec.NodeType, b.ko.Spec.NodeType)
	} else if a.ko.Spec.NodeType != nil && b.ko.Spec.NodeType != nil {
		if *a.ko.Spec.NodeType != *b.ko.Spec.NodeType {
			delta.Add("Spec.NodeType", a.ko.Spec.NodeType, b.ko.Spec.NodeType)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.NumReplicasPerShard, b.ko.Spec.NumReplicasPerShard) {
		delta.Add("Spec.NumReplicasPerShard", a.ko.Spec.NumReplicasPerShard, b.ko.Spec.NumReplicasPerShard)
	} else if a.ko.Spec.NumReplicasPerShard != nil && b.ko.Spec.NumReplicasPerShard != nil {
		if *a.ko.Spec.NumReplicasPerShard != *b.ko.Spec.NumReplicasPerShard {
			delta.Add("Spec.NumReplicasPerShard", a.ko.Spec.NumReplicasPerShard, b.ko.Spec.NumReplicasPerShard)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.NumShards, b.ko.Spec.NumShards) {
		delta.Add("Spec.NumShards", a.ko.Spec.NumShards, b.ko.Spec.NumShards)
	} else if a.ko.Spec.NumShards != nil && b.ko.Spec.NumShards != nil {
		if *a.ko.Spec.NumShards != *b.ko.Spec.NumShards {
			delta.Add("Spec.NumShards", a.ko.Spec.NumShards, b.ko.Spec.NumShards)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ParameterGroupName, b.ko.Spec.ParameterGroupName) {
		delta.Add("Spec.ParameterGroupName", a.ko.Spec.ParameterGroupName, b.ko.Spec.ParameterGroupName)
	} else if a.ko.Spec.ParameterGroupName != nil && b.ko.Spec.ParameterGroupName != nil {
		if *a.ko.Spec.ParameterGroupName != *b.ko.Spec.ParameterGroupName {
			delta.Add("Spec.ParameterGroupName", a.ko.Spec.ParameterGroupName, b.ko.Spec.ParameterGroupName)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.ParameterGroupRef, b.ko.Spec.ParameterGroupRef) {
		delta.Add("Spec.ParameterGroupRef", a.ko.Spec.ParameterGroupRef, b.ko.Spec.ParameterGroupRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Port, b.ko.Spec.Port) {
		delta.Add("Spec.Port", a.ko.Spec.Port, b.ko.Spec.Port)
	} else if a.ko.Spec.Port != nil && b.ko.Spec.Port != nil {
		if *a.ko.Spec.Port != *b.ko.Spec.Port {
			delta.Add("Spec.Port", a.ko.Spec.Port, b.ko.Spec.Port)
		}
	}
	if len(a.ko.Spec.SecurityGroupIDs) != len(b.ko.Spec.SecurityGroupIDs) {
		delta.Add("Spec.SecurityGroupIDs", a.ko.Spec.SecurityGroupIDs, b.ko.Spec.SecurityGroupIDs)
	} else if len(a.ko.Spec.SecurityGroupIDs) > 0 {
		if !ackcompare.SliceStringPEqual(a.ko.Spec.SecurityGroupIDs, b.ko.Spec.SecurityGroupIDs) {
			delta.Add("Spec.SecurityGroupIDs", a.ko.Spec.SecurityGroupIDs, b.ko.Spec.SecurityGroupIDs)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.SecurityGroupRefs, b.ko.Spec.SecurityGroupRefs) {
		delta.Add("Spec.SecurityGroupRefs", a.ko.Spec.SecurityGroupRefs, b.ko.Spec.SecurityGroupRefs)
	}
	if len(a.ko.Spec.SnapshotARNs) != len(b.ko.Spec.SnapshotARNs) {
		delta.Add("Spec.SnapshotARNs", a.ko.Spec.SnapshotARNs, b.ko.Spec.SnapshotARNs)
	} else if len(a.ko.Spec.SnapshotARNs) > 0 {
		if !ackcompare.SliceStringPEqual(a.ko.Spec.SnapshotARNs, b.ko.Spec.SnapshotARNs) {
			delta.Add("Spec.SnapshotARNs", a.ko.Spec.SnapshotARNs, b.ko.Spec.SnapshotARNs)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.SnapshotName, b.ko.Spec.SnapshotName) {
		delta.Add("Spec.SnapshotName", a.ko.Spec.SnapshotName, b.ko.Spec.SnapshotName)
	} else if a.ko.Spec.SnapshotName != nil && b.ko.Spec.SnapshotName != nil {
		if *a.ko.Spec.SnapshotName != *b.ko.Spec.SnapshotName {
			delta.Add("Spec.SnapshotName", a.ko.Spec.SnapshotName, b.ko.Spec.SnapshotName)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.SnapshotRef, b.ko.Spec.SnapshotRef) {
		delta.Add("Spec.SnapshotRef", a.ko.Spec.SnapshotRef, b.ko.Spec.SnapshotRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.SnapshotRetentionLimit, b.ko.Spec.SnapshotRetentionLimit) {
		delta.Add("Spec.SnapshotRetentionLimit", a.ko.Spec.SnapshotRetentionLimit, b.ko.Spec.SnapshotRetentionLimit)
	} else if a.ko.Spec.SnapshotRetentionLimit != nil && b.ko.Spec.SnapshotRetentionLimit != nil {
		if *a.ko.Spec.SnapshotRetentionLimit != *b.ko.Spec.SnapshotRetentionLimit {
			delta.Add("Spec.SnapshotRetentionLimit", a.ko.Spec.SnapshotRetentionLimit, b.ko.Spec.SnapshotRetentionLimit)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.SnapshotWindow, b.ko.Spec.SnapshotWindow) {
		delta.Add("Spec.SnapshotWindow", a.ko.Spec.SnapshotWindow, b.ko.Spec.SnapshotWindow)
	} else if a.ko.Spec.SnapshotWindow != nil && b.ko.Spec.SnapshotWindow != nil {
		if *a.ko.Spec.SnapshotWindow != *b.ko.Spec.SnapshotWindow {
			delta.Add("Spec.SnapshotWindow", a.ko.Spec.SnapshotWindow, b.ko.Spec.SnapshotWindow)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.SNSTopicARN, b.ko.Spec.SNSTopicARN) {
		delta.Add("Spec.SNSTopicARN", a.ko.Spec.SNSTopicARN, b.ko.Spec.SNSTopicARN)
	} else if a.ko.Spec.SNSTopicARN != nil && b.ko.Spec.SNSTopicARN != nil {
		if *a.ko.Spec.SNSTopicARN != *b.ko.Spec.SNSTopicARN {
			delta.Add("Spec.SNSTopicARN", a.ko.Spec.SNSTopicARN, b.ko.Spec.SNSTopicARN)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.SNSTopicRef, b.ko.Spec.SNSTopicRef) {
		delta.Add("Spec.SNSTopicRef", a.ko.Spec.SNSTopicRef, b.ko.Spec.SNSTopicRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.SubnetGroupName, b.ko.Spec.SubnetGroupName) {
		delta.Add("Spec.SubnetGroupName", a.ko.Spec.SubnetGroupName, b.ko.Spec.SubnetGroupName)
	} else if a.ko.Spec.SubnetGroupName != nil && b.ko.Spec.SubnetGroupName != nil {
		if *a.ko.Spec.SubnetGroupName != *b.ko.Spec.SubnetGroupName {
			delta.Add("Spec.SubnetGroupName", a.ko.Spec.SubnetGroupName, b.ko.Spec.SubnetGroupName)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.SubnetGroupRef, b.ko.Spec.SubnetGroupRef) {
		delta.Add("Spec.SubnetGroupRef", a.ko.Spec.SubnetGroupRef, b.ko.Spec.SubnetGroupRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.TLSEnabled, b.ko.Spec.TLSEnabled) {
		delta.Add("Spec.TLSEnabled", a.ko.Spec.TLSEnabled, b.ko.Spec.TLSEnabled)
	} else if a.ko.Spec.TLSEnabled != nil && b.ko.Spec.TLSEnabled != nil {
		if *a.ko.Spec.TLSEnabled != *b.ko.Spec.TLSEnabled {
			delta.Add("Spec.TLSEnabled", a.ko.Spec.TLSEnabled, b.ko.Spec.TLSEnabled)
		}
	}
	if !ackcompare.MapStringStringEqual(ToACKTags(a.ko.Spec.Tags), ToACKTags(b.ko.Spec.Tags)) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}

	return delta
}
