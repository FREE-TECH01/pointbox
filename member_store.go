package main

import "time"

//like the storage or the database of all the members
type MemberStore struct {
	members map[int]Member
	nextID  int
}

//a temporary storage for new members to be poured into the main storage
//its a constructor
func NewMemberStore() *MemberStore {
	return &MemberStore{
		members: make(map[int]Member),
		nextID:  1,
	}
}

//to add the members method
func (s *MemberStore) AddMember(name string) {
	//
	id := s.nextID

	//get the 
	member := Member{
		Name:        name,
		Points:      1500,
		LastUpdated: time.Now().Format("02-01-2006 03:04PM"),
	}

	s.members[id] = member
	s.nextID++
}

func (s *MemberStore) GetMembers() map[int]Member {
	return s.members
}

func(s *MemberStore) DeleteMember (id int) bool {
	//checking for the prescence of a member in the store
	_, ok := s.members[id]

	//its not possible to delete who is not in the store
	if !ok {
		return false
	}
	
	//deleting it from the map store since it exists
	delete(s.members, id)

	return true
}

func(s *MemberStore) AddPoints(id int) bool {
	//check if the member exits
	member, ok := s.members[id]

	//if the member does not exist there is no where to add the point
	if !ok {
		return  false
	}

	//the variable "member" is a copy of the Member struct
	//so i can reference its fields
	member.Points += 10
	member.LastUpdated = time.Now().Format("02-01-2006 03:04PM")

	//i need to update the actual mapstorage which is the Memberstore
	s.members[id] = member

	return true
}

func(s *MemberStore) DeletePoints(id int) bool {
	//check if the member exits
	member, ok := s.members[id]

	//if the member does not exist there is no where to add the point
	if !ok {
		return  false
	}

	//the variable "member" is a copy of the Member struct
	//so i can reference its fields
	//dedect points
	member.Points -= 10
	member.LastUpdated = time.Now().Format("02-01-2006 03:04PM")

	//i need to update the actual mapstorage which is the Memberstore
	s.members[id] = member

	return true
}