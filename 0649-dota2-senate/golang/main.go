func predictPartyVictory(senate string) string {
    radiant, dire := 0, 0
    bannedRadiant, bannedDire := 0, 0

    // Contar la cantidad inicial de senadores en cada partido
    for _, s := range senate {
        if s == 'R' {
            radiant++
        } else {
            dire++
        }
    }

    for radiant > 0 && dire > 0 {
        newSenate := ""
        for _, s := range senate {
            if s == 'R' {
                // Si es senador Radiant y hay senadores Dire para banear
                if bannedRadiant > 0 {
                    bannedRadiant--
                    radiant--
                } else {
                    bannedDire++
                    newSenate += string(s)
                }
            } else {
                // Si es senador Dire y hay senadores Radiant para banear
                if bannedDire > 0 {
                    bannedDire--
                    dire--
                } else {
                    bannedRadiant++
                    newSenate += string(s)
                }
            }
        }
        senate = newSenate
    }

    if radiant > 0 {
        return "Radiant"
    }
    return "Dire"
}
