package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ScalingSchedule describes a time based metric to be used in
// autoscaling operations.
// +k8s:deepcopy-gen=true
// +resource:path=scalingschedules
type ScalingSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ScalingScheduleSpec `json:"spec"`
}

// ScalingScheduleSpec is the spec part of the ScalingSchedule.
// +k8s:deepcopy-gen=true
type ScalingScheduleSpec struct {
	// Schedules is the list of schedules for this ScalingSchedule
	// resource. All the schedules defined here will result on the value
	// to the same metric. New metrics require a new ScalingSchedule
	// resource.
	Schedules []Schedule `json:"schedules"`
}

// Schedule is the schedule details to be used inside a ScalingSchedule.
// +k8s:deepcopy-gen=true
type Schedule struct {
	// Defines if the schedule is a OneTime schedule or
	// Repeating one. If OneTime, date has to be defined. If Repeating,
	// Period has to be defined.
	// +kubebuilder:validation:Enum=OneTime;Repeating
	ScheduleType string `json:"type"`
	// Defines the details of a Repeating schedule.
	// +optional
	Period SchedulePeriod `json:"period"`
	// Defines the starting date of a OneTime schedule. It has to
	// be a RFC3339 formated date.
	// +optional
	Date ScheduleDate `json:"date"`
	// The duration in minutes that the configured value will be
	// returned for the defined schedule.
	DurationMinutes int `json:"durationMinutes"`
	// The metric value that will be returned for the defined schedule.
	Value int `json:"value"`
}

// SchedulePeriod is the details to be used for a Schedule of the
// Repeating type.
// +k8s:deepcopy-gen=true
type SchedulePeriod struct {
	// The startTime has the format HH:MM
	StartTime string `json:"startTime"`
	// The days that this schedule will be active.
	Days []ScheduleDay `json:"days"`
	// The location name corresponding to a file in the IANA
	// Time Zone database, like Europe/Berlin.
	Timezone string `json:"timezone"`
}

// ScheduleDay represents the valid inputs for days in a SchedulePeriod.
// +kubebuilder:validation:Enum=Mon;Tue;Wed;Thu;Fru;Sat;Sun
// +k8s:deepcopy-gen=true
type ScheduleDay string

// ScheduleDate is a RFC3339 representation of the date for a Schedule
// of the OneTime type.
// +kubebuilder:validation:Format="date-time"
// +k8s:deepcopy-gen=true
type ScheduleDate string

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ScalingScheduleList is a list of scaling schedules.
// +k8s:deepcopy-gen=true
type ScalingScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ScalingSchedule `json:"items"`
}
