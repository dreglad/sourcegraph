package database

import (
	"context"
	"database/sql"

	"github.com/sourcegraph/log"

	"github.com/sourcegraph/sourcegraph/internal/database/basestore"
	"github.com/sourcegraph/sourcegraph/internal/types"
	"github.com/sourcegraph/sourcegraph/lib/errors"
)

type RepoFileStore interface {
	CreateIfNotExists(ctx context.Context, f types.RepoFile) (*types.RepoFile, error)
}

var _ RepoFileStore = (*repoFileStore)(nil)

// repoFileStore handles access to the repo_files table
type repoFileStore struct {
	logger log.Logger
	*basestore.Store
}

func RepoFilesWith(logger log.Logger, other basestore.ShareableStore) RepoFileStore {
	return &repoFileStore{
		logger: logger,
		Store:  basestore.NewWithHandle(other.Handle()),
	}
}

// CreateIfNotExists
func (s *repoFileStore) CreateIfNotExists(ctx context.Context, f types.RepoFile) (*types.RepoFile, error) {
	var id int
	row := s.Handle().QueryRowContext(
		ctx,
		`INSERT INTO repo_files(directory_id, version_id, topological_order, base_name, content_id)
		VALUES($1, $2, $3, $4, $5)
		ON CONFLICT ("directory_id", "version_id", "base_name") DO NOTHING
		RETURNING id`,
		f.DirectoryID, f.VersionID, f.TopologicalOrder, f.BaseName, f.ContentID,
	)
	err := row.Scan(&id)
	if err == sql.ErrNoRows {
		ff, err := s.FindUnique(ctx, f.DirectoryID, f.VersionID, f.BaseName)
		if err != nil {
			return nil, errors.Wrap(err, "file already exists, but encountered errors retrieving it")
		}
		if ff == nil {
			return nil, errors.New("this is weird, file cannot be uniquely found, but unique constraint fails")
		}
		return ff, nil
	}
	if err != nil {
		return nil, err
	}
	f.ID = id
	return &f, nil
}

func (s *repoFileStore) FindUnique(ctx context.Context, directoryID int, versionID int, baseName string) (*types.RepoFile, error) {
	var f types.RepoFile
	row := s.Handle().QueryRowContext(ctx,
		`SELECT
			f.id,
			f.directory_id,
			f.version_id,
			f.topological_order,
			f.base_name,
			f.content_id
		FROM repo_files AS f
		WHERE
			f.directory_id = $1
			AND f.version_id = $2
			AND f.base_name = $3`,
		directoryID, versionID, baseName)
	if row == nil {
		return nil, nil
	}
	if err := row.Scan(&f.ID, &f.DirectoryID, &f.VersionID, &f.TopologicalOrder, &f.BaseName, &f.ContentID); err != nil {
		return nil, err
	}
	return &f, nil
}