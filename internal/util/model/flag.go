package model

// Username flag short
var Usrn_s string = "-u"

// Username flag full
var Usrn_f string = "--username"

// Password flag short
var Pass_s string = "-p"

// Password flag short
var Pass_f string = "--password"

// Получить значение ключа
func GetFlagValue(args []string, shortFlagFormat, fullFlagFormat, defaultValue string) string {
	for i, arg := range args {
		if arg == shortFlagFormat || arg == fullFlagFormat {
			return args[i+1]
		}
	}
	return defaultValue
}
