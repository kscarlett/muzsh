package session

import (
	"bufio"
	"errors"
	"os"

	"github.com/kelindar/binary"
)

// Save saves the given session to a file
func Save(data *Session, filename string) (err error) {

	encoded, err := binary.Marshal(data)
	if err != nil {
		return errors.New("failed to encode save data")
	}

	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return errors.New("failed to create save file")
	}

	w := bufio.NewWriter(file)
	_, err = w.Write(encoded)
	if err != nil {
		return errors.New("failed to write to save file")
	}

	return
}

// Load loads the given save file into a session
func Load(filename string) (data *Session, err error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("failed to open save file")
	}

	r := bufio.NewReader(file)
	var encoded []byte
	_, err = r.Read(encoded)
	if err != nil {
		return nil, errors.New("failed to read save data")
	}

	err = binary.Unmarshal(encoded, &data)
	if err != nil {
		return nil, errors.New("failed to decode save data")
	}

	return
}

// ExistingSave checks whether the given file exists
func ExistingSave(filename string) bool {

	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
