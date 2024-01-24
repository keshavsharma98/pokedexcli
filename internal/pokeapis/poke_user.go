package pokeapis

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/keshavsharma98/pokedexcli/internal/common"
)

func (c *Client) SaveGame() error {

	saveData, err := json.Marshal(c.user)
	if err != nil {
		return err
	}
	secret_key := os.Getenv("KEY")
	encrypted_data := common.Encrypt(saveData, secret_key)
	f, err := os.Create(os.Getenv("FILE_PATH"))
	if err != nil {
		return err
	}

	_, err = f.Write(encrypted_data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) LoadGame() error {
	f_name := os.Getenv("FILE_PATH")
	i, _ := os.Stat(f_name)

	if i == nil {
		c.SaveGame()
		return nil
	}

	f, err := os.Open(f_name)
	if err != nil {
		fmt.Println("Error opening saved file")
		return err
	}

	defer f.Close()

	saveData, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	secret_key := os.Getenv("KEY")
	decrypted_data := common.Decrypt(saveData, secret_key)
	err = json.Unmarshal(decrypted_data, &c.user)
	if err != nil {
		return err
	}

	return nil
}
