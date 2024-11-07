package entity

import "fmt"

type Role string
type Department string

const (
	RoleAdmin       Role = "admin"
	RoleDeveloper   Role = "developer"
	RoleDevOps      Role = "devops"
	RoleDataAnalyst Role = "data-analyst"
	RoleManager     Role = "manager"
)

const (
	DepartmentData  Department = "data-analysis"
	DepartmentTI    Department = "TI"
	DepartmentAdmin Department = "admin"
)

func ParseRole(s string) (Role, error) {
    role := Role(s)
    if !role.IsValid() {
        return "", fmt.Errorf("invalid role: %s", s)
    }
    return role, nil
}

func ParseDepartment(s string) (Department, error) {
    dept := Department(s)
    if !dept.IsValid() {
        return "", fmt.Errorf("invalid department: %s", s)
    }
    return dept, nil
}

func (r Role) IsValid() bool {
	switch r {
	case RoleAdmin, RoleDeveloper, RoleDevOps, RoleDataAnalyst, RoleManager:
		return true
	default:
		return false
	}
}

func (d Department) IsValid() bool {
	switch d {
	case DepartmentData, DepartmentTI, DepartmentAdmin:
		return true
	default:
		return false
	}
}
