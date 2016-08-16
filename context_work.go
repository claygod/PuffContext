package PuffContext

// PuffContext
// Work
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

// PuffContext structure
type PuffContext struct {
	parent *PuffContext
	hash   map[interface{}]interface{}
}

// New - create a new PuffContext-struct
func New() *PuffContext {
	pc := &PuffContext{}
	pc.hash = make(map[interface{}]interface{})
	return pc
}

// Get - get variable.
func (pc *PuffContext) Get(key interface{}) interface{} {
	// fmt.Print("\n key: ", key)
	if value, ok := pc.hash[key]; ok {
		return value
	}
	if pc.parent != nil {
		return pc.parent.Get(key)
	}
	return nil
}

// Set - add the variable in context.
func (pc *PuffContext) Set(key interface{}, value interface{}) {
	pc.hash[key] = value
	return
}

// Fix - fix the state to create a branch (context).
func (pc *PuffContext) Fix() *PuffContext {
	x := New()
	x.parent = pc
	return x
}

// Del - remove variable (can only be from the current branch!).
func (pc *PuffContext) Del(key interface{}) bool {
	if _, ok := pc.hash[key]; ok {
		delete(pc.hash, key)
		return true
	} else {
		return false
	}
}
