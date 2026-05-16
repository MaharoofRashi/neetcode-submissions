func isAnagram(s string, t string) bool {
    runeS := []rune(s)
    runeT := []rune(t)

    sort.Slice(runeS, func (i, j int) bool {
        return runeS[i] < runeS[j]
    })

    sort.Slice(runeT, func (i, j int) bool {
        return runeT[i] < runeT[j]
    })

    stringS := string(runeS)
    stringT := string(runeT)

    return stringS == stringT
}
