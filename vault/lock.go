package vault

import "strings"

func (vault *Vault) Unlock(password string) (err error) {
	res, err := bw("unlock " + password)
	if err != nil {
		return
	}

	// get the BW_SESSION token
	token := res["message"].(string)
	token = strings.Split(token, "\n")[2]
	token = strings.Split(token, "BW_SESSION=")[1]
	token = strings.TrimPrefix(token, "\"")
	token = strings.TrimSuffix(token, "\"")
	vault.SessionKey = token

	vault.Status = "unlocked"
	return
}

func (vault *Vault) Lock() (err error) {
	_, err = bw("lock")
	if err != nil {
		return
	}

	vault.Status = "locked"
	return
}
