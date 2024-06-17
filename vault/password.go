package vault

func (vault *Vault) Sync() (err error) {
	_, err = bw("sync")
	if err != nil {
		return
	}

	err = vault.reload()
	return
}

func Search(query string) (err error) { return }
func Get(id string) (err error)       { return }
