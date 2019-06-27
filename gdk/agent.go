/* This is free and unencumbered software released into the public domain. */

package gdk

// Agent
type Agent interface {
	IsPlayer() bool
	IsRobot() bool
	CanFly() bool
	CanMove() bool
	HasLegs() bool
	HasWings() bool
	HasWheels() bool
	Name() string
}

// AgentPredicate
type AgentPredicate func(Agent) bool

// AgentProperty
type AgentProperty func(Agent) interface{}

// AgentPredicates
func AgentPredicates() map[string]AgentPredicate {
	return map[string]AgentPredicate{
		"is_player": func(agent Agent) bool {
			return agent.IsPlayer()
		},
		"is_robot": func(agent Agent) bool {
			return agent.IsRobot()
		},
		"can_fly": func(agent Agent) bool {
			return agent.CanFly()
		},
		"can_move": func(agent Agent) bool {
			return agent.CanMove()
		},
		"has_legs": func(agent Agent) bool {
			return agent.HasLegs()
		},
		"has_wings": func(agent Agent) bool {
			return agent.HasWings()
		},
		"has_wheels": func(agent Agent) bool {
			return agent.HasWheels()
		},
	}
}

// AgentProperties
func AgentProperties() map[string]AgentProperty {
	return map[string]AgentProperty{
		"name": func(agent Agent) interface{} {
			return agent.Name()
		},
	}
}
