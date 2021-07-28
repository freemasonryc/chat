// Copyright 2021 The Matrix.org Foundation C.I.C.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package postgres

import (
	"database/sql"

	"github.com/matrix-org/dendrite/keyserver/storage/tables"
)

var crossSigningStreamsSchema = `
CREATE TABLE IF NOT EXISTS keyserver_cross_signing_streams (
    stream_id BIGINT NOT NULL,
	from_user_id TEXT NOT NULL,
	user_ids TEXT NOT NULL,
	PRIMARY KEY(stream_id)
);
`

type crossSigningStreamsStatements struct {
	db *sql.DB
}

func NewPostgresCrossSigningStreamsTable(db *sql.DB) (tables.CrossSigningStreams, error) {
	s := &crossSigningStreamsStatements{
		db: db,
	}
	_, err := db.Exec(crossSigningStreamsSchema)
	if err != nil {
		return nil, err
	}
	return s, nil
}