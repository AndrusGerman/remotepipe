package connection

import (
	"encoding/json"
	"io"
)

type ConnectionType string

const (
	ConnectionTypeStdout = "stdout"
	ConnectionTypeStder  = "stder"
	ConnectionTypeStdin  = "stdin"
	ConnectionTypeCreate = "create"
)

type ConnectionMetadata struct {
	Type ConnectionType
	ID   string
}

func ConnectionSend(ID string, typeConn ConnectionType, conn io.Writer) error {
	var cm = new(ConnectionMetadata)
	cm.Type = typeConn
	cm.ID = ID
	return json.NewEncoder(conn).Encode(cm)
}

func ConnectionRead(conn io.Reader) (*ConnectionMetadata, error) {
	var cm = new(ConnectionMetadata)
	return cm, json.NewDecoder(conn).Decode(cm)
}
