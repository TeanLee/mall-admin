export function setUserName(username) {
  window.localStorage.setItem('username', username)
}

export function getUserName() {
  window.localStorage.getItem('username')
}

export function removeUserName() {
  window.localStorage.removeItem('username')
}

export function setRole(role) {
  window.localStorage.getItem('role', role)
}

export function getRole() {
  window.localStorage.getItem('role')
}

export function removeRole() {
  window.localStorage.removeItem('role')
}
