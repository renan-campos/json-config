package configuration

type Configuration struct {
	ProgramName  string `json:"program_name"`
	HelloMessage string `json:"hello_message"`
	ByeMessage   string `json:"bye_message"`
}
