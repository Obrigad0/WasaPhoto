package api

import (
	"strconv"
	"strings"
)

// estrazioneToken() preleva dalla stringa il token
func estrazioneToken(authorization string) int {
	// FORMATO TOKEN: BEARER idUtente
	// ES: BEARER 260513

	// splitto la stringa
	var tokenStr = strings.Split(authorization, " ")
	// se ho due elementi tutto ok
	if len(tokenStr) == 2 {
		// prendo il token in formato stringa
		tokenStr := strings.Trim(tokenStr[1], " ")
		// trasformo il token in un intero
		tokenInt, err := strconv.Atoi(tokenStr)
		if err != nil {
			// C'e' un errore con la trasformazione del token in intero
			return -1
		}
		return tokenInt
	}
	// c'e' un errore con l'estrazione
	return -1

}

// verificaToken() verifica la validita' del token
func verificaToken(uId int, token string) bool {
	// Estraggo il token
	tokenId := estrazioneToken(token)
	if uId != tokenId {
		// token non uguali torno false
		return false
	}
	// tutto ok!
	return true
}

// verificaLogin() verifica semplicemente che l'utente sia loggato
func verificaLogin(token string) bool {
	// Estraggo il token

	tokenId := estrazioneToken(token)

	// se il token non c'e' allora errore
	if tokenId < 0 {
		return false
	}
	// tutto ok token valido!
	return true
}
