package commands

// Commands is the set of root command groups.
type All struct {
	POM POM `command:"pom" alias:"p" description:"Apply the provided migration to a set of POM files."`
	//Settings Settings `command:"settings" alias:"s" description:"Apply the provided migration to a settings.xml file."`
}
