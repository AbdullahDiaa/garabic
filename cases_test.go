package arabic

var removeHarakatTestCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		description: "Removing Alif Khanjariyah",
		input:       "رَحْمَٰن",
		expected:    "رحمن",
	},
	{
		description: "Removing Waslah",
		input:       "ٱمْشُوا",
		expected:    "امشوا",
	},

	{
		description: "Removing all harakat 1",
		input:       "يَا أَيُّهَا الَّذِينَ آمَنُوا أَوْفُوا بِالْعُقُودِ",
		expected:    "يا أيها الذين آمنوا أوفوا بالعقود",
	},
	{
		description: "Removing all harakat 2",
		input:       "سَنواتٌ قَليلةٌ وأَدْخُلُ الجامِعَةَ، يا لها مِنْ رِحْلةٍ شاقَّةٍ طَويلَةٍ، ما أصعبَ أيامَ الدِّراسَةِ",
		expected:    "سنوات قليلة وأدخل الجامعة، يا لها من رحلة شاقة طويلة، ما أصعب أيام الدراسة",
	},
	{
		description: "Removing all harakat 3",
		input:       "إِنِّني أَشْكُرُ رَبِّي دائماً، لكنني مُنْزعجةٌ.. أَلاَ يَحِقُّ لِيَ التعبيرُ عن ضِيقِ صَدْري؟!",
		expected:    "إنني أشكر ربي دائما، لكنني منزعجة.. ألا يحق لي التعبير عن ضيق صدري؟!",
	},
}

var normalizeTestCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		description: "AlefHamzaAbove removal from beginning of word",
		input:       "أحمد",
		expected:    "احمد",
	},
	{
		description: "AlefHamzaAbove removal from middle of word",
		input:       "مأمون",
		expected:    "مامون",
	},
	{
		description: "AlefHamzaAbove removal from end of word",
		input:       "نمأ",
		expected:    "نما",
	},
	{
		description: "Replacing DotlessYae with Yae",
		input:       "منى",
		expected:    "مني",
	},
	{
		description: "Replacing TehMarbuta with Hae",
		input:       "مكتبة",
		expected:    "مكتبه",
	},
	{
		description: "Trimming Tatweel",
		input:       "بريـــــــــد",
		expected:    "بريد",
	},
	{
		description: "Removing Maddah",
		input:       "قُرْآن",
		expected:    "قران",
	},
	{
		description: "Normalizing Alif Waslah",
		input:       "ٱمْشُوا",
		expected:    "امشوا",
	},
	{
		description: "Removing all harakat combined with text normalization 1",
		input:       "يَا أَيُّهَا الَّذِينَ آمَنُوا أَوْفُوا بِالْعُقُودِ",
		expected:    "يا ايها الذين امنوا اوفوا بالعقود",
	},
	{
		description: "Removing all harakat combined with text normalization 2",
		input:       "سَنواتٌ قَليلةٌ وأَدْخُلُ الجامِعَةَ، يا لها مِنْ رِحْلةٍ شاقَّةٍ طَويلَةٍ، ما أصعبَ أيامَ الدِّراسَةِ",
		expected:    "سنوات قليله وادخل الجامعه، يا لها من رحله شاقه طويله، ما اصعب ايام الدراسه",
	},
	{
		description: "Removing all harakat combined with text normalization 3",
		input:       "إِنِّني أَشْكُرُ رَبِّي دائماً، لكنني مُنْزعجةٌ.. أَلاَ يَحِقُّ لِيَ التعبيرُ عن ضِيقِ صَدْري؟!",
		expected:    "انني اشكر ربي دائما، لكنني منزعجه.. الا يحق لي التعبير عن ضيق صدري؟!",
	},
}
