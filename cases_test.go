package garabic

//removeHarakatTestCases contains all test cases for TestRemoveHarakat function
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

//normalizeTestCases contains all test cases for TestNormalize function
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

//spellNumberTestCases contains all test cases for reading a number in arabic
var spellNumberTestCases = []struct {
	input    int
	expected string
}{
	{
		0,
		"صفر",
	},
	{
		1,
		"واحد",
	},
	{
		-1,
		"سالب واحد",
	},
	{
		2,
		"اثنان",
	},
	{
		3,
		"ثلاثة",
	},
	{
		4,
		"أربعة",
	},
	{
		5,
		"خمسة",
	},
	{
		6,
		"ستة",
	},
	{
		7,
		"سبعة",
	},
	{
		8,
		"ثمانية",
	},
	{
		9,
		"تسعة",
	},
	{
		10,
		"عشرة",
	},
	{
		11,
		"أحد عشر",
	},
	{
		12,
		"اثنا عشر",
	},
	{
		20,
		"عشرون",
	},
	{
		50,
		"خمسون",
	},
	{
		90,
		"تسعون",
	},
	{
		100,
		"مئة",
	},
	{
		1000,
		"ألف",
	},
	{
		1250,
		"ألف و مئتان و خمسون",
	},
	{
		11225,
		"أحد عشر ألف و مئتان و خمسة و عشرون",
	},
	{
		100000,
		"مئة ألف",
	},
	{
		1000000,
		"مليون",
	},
	{
		-2000000,
		"سالب اثنان مليون",
	},
	{
		141592653589,
		"مئة و واحد و أربعون مليار و خمسمئة و اثنان و تسعون مليون و ستمئة و ثلاثة و خمسون ألف و خمسمئة و تسعة و ثمانون",
	},
	{
		141592653589,
		"مئة و واحد و أربعون مليار و خمسمئة و اثنان و تسعون مليون و ستمئة و ثلاثة و خمسون ألف و خمسمئة و تسعة و ثمانون",
	},
}

//tashkeelTestCases contains all test cases for adding tashkeel to arabic text
var tashkeelTestCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		"Adding Kasrah after 'من'",
		"يقرأ محمد مِنَ الكتاب",
		"يقرأ محمد مِنَ الكتابِ",
	},
}

//shapingTestCases contains all test cases for shaping arabic text
var shapingTestCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		"Shaping 1 word without tashkeel",
		"بالعربي",
		"ﻲﺑﺮﻌﻟﺎﺑ",
	},
	{
		"Shaping 1 word with tashkeel",
		"بِالعَرَبِّي",
		"ﻲِّﺑَﺮَﻌﻟﺎِﺑ",
	},
	{
		"Shaping 1 word with tashkeel",
		"فَحَومَلِ",
		"ِﻞَﻣﻮَﺤَﻓ",
	},
	{
		"Shaping  1 sentence with tashkeel",
		"قِفا نَبكِ مِن ذِكرى حَبيبٍ وَمَنزِلِ   ****   بِسِقطِ اللِوى بَينَ الدَخولِ فَحَومَلِ",
		"ِﻞَﻣﻮَﺤَﻓ ِﻝﻮﺧَﺪﻟا َﻦﻴَﺑ ﻯﻮِﻠﻟا ِﻂﻘِﺴِﺑ **** ِﻝِﺰﻨَﻣَو ٍﺐﻴﺒَﺣ ﻯﺮﻛِذ ﻦِﻣ ِﻚﺒَﻧ ﺎﻔِﻗ",
	},

	{
		"Shaping 1 word without tashkeel",
		"المصفوفة (Multidimentional Array) هي",
		"ﻲﻫ (Multidimentional Array) ﺔﻓﻮﻔﺼﻤﻟا",
	},
}

//arabicLetterTestCases
var arabicLetterTestCases = []struct {
	description string
	input       rune
	expected    bool
}{
	{
		"Checking arabic letter",
		'ص',
		true,
	},
	{
		"Checking english letter",
		's',
		false,
	},
	{
		"Checking tashkeel as arabic letter",
		'ُ',
		true,
	},
}

// /arabicWordsTestCases
var arabicTextTestCases = []struct {
	description string
	input       string
	expected    bool
}{
	{
		"Checking arabic text",
		"نص عربي",
		true,
	},
	{
		"Checking arabic/english mixed text",
		"نص عربي with english",
		false,
	},
}
