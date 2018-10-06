package languages

// GermanTranslation is the German translation of Wrap
var GermanTranslation = Translation{
	Language: German,

	SceneTags: []string{
		"int", "ext", "est", "int./ext", "int/ext", "i/e",
		"etabl",
	},

	StageplaySceneTags: []string{
		"szene",
	},

	TransitionTags: []string{
		"TO:", "ZU:",
	},

	BeginActTags: []string{
		"AKT",
	},

	EndActTags: []string{
		"ENDE DES AKTS", "ENDE DES AKTES",
	},

	More: "(MEHR)",

	Contd: "(fortges.)",

	Continued: "FORTGESETZT",
}
