package opa.examples

import input as i

# 定义策略
allow {
  i.subject.role == "manager"
  i.object == "employee_info"
  i.action == "read" 
}

allow {
  i.subject.role == "manager"
  i.object == "employee_info"
  i.action == "write"
}

allow {
  i.subject.role == "employee"
  i.object == "employee_info"
  i.action == "read"
}

allow {
  i.subject.role == "employee"
  i.object == "employee_info" 
  i.action == "write"
}

allow {
  i.subject.role == "hr"
  i.object == "employee_info"
  i.action == "read"
}

allow {
  i.subject.role == "hr"
  i.object == "employee_info"
  i.action == "write" 
}

allow {
  i.subject.role == "finance"
  i.object == "employee_salary"
  i.action == "read"
}
