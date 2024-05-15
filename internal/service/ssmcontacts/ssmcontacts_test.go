// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmcontacts_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// SSMContacts resources depend on a replication set existing and
// only one replication set resource can be active at once, so we must have serialised tests
func TestAccSSMContacts_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]map[string]func(t *testing.T){
		"Contact Resource Tests": {
			acctest.CtBasic:     testContact_basic,
			"disappears":        testContact_disappears,
			"updateAlias":       testContact_updateAlias,
			"updateDisplayName": testContact_updateDisplayName,
			"updateTags":        testContact_updateTags,
			"updateType":        testContact_updateType,
		},
		"Contact Data Source Tests": {
			acctest.CtBasic: testContactDataSource_basic,
		},
		"Contact Channel Resource Tests": {
			acctest.CtBasic:   testContactChannel_basic,
			"contactId":       testContactChannel_contactID,
			"deliveryAddress": testContactChannel_deliveryAddress,
			"disappears":      testContactChannel_disappears,
			names.AttrName:    testContactChannel_name,
			names.AttrType:    testContactChannel_type,
		},
		"Contact Channel Data Source Tests": {
			acctest.CtBasic: testContactChannelDataSource_basic,
		},
		"Plan Resource Tests": {
			acctest.CtBasic:           testPlan_basic,
			"disappears":              testPlan_disappears,
			"updateChannelTargetInfo": testPlan_updateChannelTargetInfo,
			"updateContactId":         testPlan_updateContactId,
			"updateContactTargetInfo": testPlan_updateContactTargetInfo,
			"updateDurationInMinutes": testPlan_updateDurationInMinutes,
			"updateStages":            testPlan_updateStages,
			"updateTargets":           testPlan_updateTargets,
		},
		"Plan Data Source Tests": {
			acctest.CtBasic:     testPlanDataSource_basic,
			"channelTargetInfo": testPlanDataSource_channelTargetInfo,
		},
		"RotationResource": {
			acctest.CtBasic: testRotation_basic,
			"disappears":    testRotation_disappears,
			"update":        testRotation_updateRequiredFields,
			"startTime":     testRotation_startTime,
			"contactIds":    testRotation_contactIds,
			"recurrence":    testRotation_recurrence,
			names.AttrTags:  testRotation_tags,
		},
		"RotationDataSource": {
			acctest.CtBasic:   testRotationDataSource_basic,
			"dailySettings":   testRotationDataSource_dailySettings,
			"monthlySettings": testRotationDataSource_monthlySettings,
		},
	}

	acctest.RunSerialTests2Levels(t, testCases, 0)
}
