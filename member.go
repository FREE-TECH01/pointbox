package main

// Member represents a member in the PointBox application.
//
// A Member contains the basic information we need to keep track of:
//
// - Name: the member's name.
// - Points: the number of points the member currently has.
// - LastUpdated: when the member's information was last updated.
//
// This struct is primarily a data model. It describes the shape of a
// member but does not contain the logic for creating, deleting, or
// modifying members. That logic lives in the MemberStore.
type Member struct {
	Name        string
	Points      int
	LastUpdated string
}
