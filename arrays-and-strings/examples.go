package arraysandstrings

// --------------------------------------------------------------------------------------------------------------

// 1.1: Implement an algorithm to determine if a string has all unique characters, what if you cannot use additional data structures?

// We can use a hashmap to create a frequency analysis, by iterating on each letter and see if it exists in the
// map, if yes increment the counter, if not assign it to the map. Without additional DS I'd need to go with
// O(N2) since I need to check each pair one by one.

func HasAllUniqueChars(input string) bool {
	set := make(map[rune]bool, 0)

	for _, ch := range input {
		if _, ok := set[ch]; ok {
			return false
		}
		set[ch] = true
	}

	return true
}

// What I missed from the Book?

/*
	- Asking whether string is ASCII or Unicode to the interviewer
		- This effects the storage size of the alphabet, which is 128 possible values in ASCII, and 256 in extended ASCII
	- I couldn't think about returning directly false if the string length exceeds the possible number of unique chars
	- Asking about whether we're allowed to modify the string to sort it and linearly checking neighboring chars
*/

// --------------------------------------------------------------------------------------------------------------

// 1.2 Check Permutation: Given two strings, write a method to decide if one is a permutation of the other

// abac , baac , caab
// In this case permutations are same length, so we can immediately return false if lengths are not equal
// We can make use of a hashmap to count the occurrences and then see if they match in terms of number of occurrences and the characters

// Time complexity is O(A + B * N), dropping the constants yields O(N)

func CheckIfPermutation(stringOne, stringTwo string) bool {
	if len(stringOne) != len(stringTwo) {
		return false
	}

	// O(A + B) where A is stringOne's length and B is stringTwo's length
	frequencyMapForOne := getFrequencyMap(stringOne)
	frequencyMapForTwo := getFrequencyMap(stringTwo)

	// O(N)
	for keyOne, valOne := range frequencyMapForOne {
		if valTwo, ok := frequencyMapForTwo[keyOne]; ok {
			if valOne != valTwo {
				return false
			}
		}
	}

	return true
}

func getFrequencyMap(input string) map[rune]int {
	frequencyMap := make(map[rune]int, 0)

	for _, ch := range input {
		if val, ok := frequencyMap[ch]; ok {
			frequencyMap[ch] = val + 1
			continue
		}

		frequencyMap[ch] = 1
	}

	return frequencyMap
}

// --------------------------------------------------------------------------------------------------------------
