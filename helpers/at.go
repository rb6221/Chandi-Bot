package helpers

func AtRole(role string) string {
	return "<@&" + role + ">"
}

func AtUser(user string) string {
	return "<@" + user + ">"
}
