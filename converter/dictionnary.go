package converter

var reference = map[rune]int{
	'A':0, 'B':1, 'C':2, 'D':3, 'E':4, '0':5, '1':6, '2':7, '3':8, '4':9,
	'F':0, 'G':1, 'H':2, 'I':3, 'J':4, '5':5, '6':6, '7':7, '8':8, '9':9,
	'K':0, 'L':1, 'M':2, 'N':3, 'O':4,				 '!':7,
	'P':0, 'Q':1, 'R':2, 'S':3, 'T':4,				 '?':7,
	'U':0, 'V':1, 'W':2, 'X':3, 'Y':4,				 '-':7,
	'Z':0,											 '&':7,
}

var ItemsReferences = map[int]rune{
	0: 'A',
	1: 'B',
	2: 'C',
	3: 'D',
	4: 'E',
	5: 'F',
	6: 'G',
	7: 'H',
	8: 'I',
	9: 'J',
}

var Items = map[rune]string{
	'A':    "Medal",
	'B':    "Boots of Speed",
	'C':    "Headgear Armor Item",
	'D':    "Ring of Protection",
	'E':    "Ring of Fire Resistance",
	'F':    "Ring of Spell Turning",
	'G':    "Gauntlets of Ogre Power",
	'H':    "Anklet",
	'I':    "Brooch",
	'J':    "Orb",
}