// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package chap1

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createExerciseStmt, err = db.PrepareContext(ctx, createExercise); err != nil {
		return nil, fmt.Errorf("error preparing query CreateExercise: %w", err)
	}
	if q.createSetStmt, err = db.PrepareContext(ctx, createSet); err != nil {
		return nil, fmt.Errorf("error preparing query CreateSet: %w", err)
	}
	if q.createUserImageStmt, err = db.PrepareContext(ctx, createUserImage); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUserImage: %w", err)
	}
	if q.createUsersStmt, err = db.PrepareContext(ctx, createUsers); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUsers: %w", err)
	}
	if q.createWorkoutStmt, err = db.PrepareContext(ctx, createWorkout); err != nil {
		return nil, fmt.Errorf("error preparing query CreateWorkout: %w", err)
	}
	if q.deleteExerciseStmt, err = db.PrepareContext(ctx, deleteExercise); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteExercise: %w", err)
	}
	if q.deleteSetsStmt, err = db.PrepareContext(ctx, deleteSets); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteSets: %w", err)
	}
	if q.deleteUserImageStmt, err = db.PrepareContext(ctx, deleteUserImage); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUserImage: %w", err)
	}
	if q.deleteUserWorkoutsStmt, err = db.PrepareContext(ctx, deleteUserWorkouts); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUserWorkouts: %w", err)
	}
	if q.deleteUsersStmt, err = db.PrepareContext(ctx, deleteUsers); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUsers: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, getUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.getUserImageStmt, err = db.PrepareContext(ctx, getUserImage); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserImage: %w", err)
	}
	if q.getUserSetsStmt, err = db.PrepareContext(ctx, getUserSets); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserSets: %w", err)
	}
	if q.getUserWorkoutStmt, err = db.PrepareContext(ctx, getUserWorkout); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserWorkout: %w", err)
	}
	if q.listExercisesStmt, err = db.PrepareContext(ctx, listExercises); err != nil {
		return nil, fmt.Errorf("error preparing query ListExercises: %w", err)
	}
	if q.listImagesStmt, err = db.PrepareContext(ctx, listImages); err != nil {
		return nil, fmt.Errorf("error preparing query ListImages: %w", err)
	}
	if q.listSetsStmt, err = db.PrepareContext(ctx, listSets); err != nil {
		return nil, fmt.Errorf("error preparing query ListSets: %w", err)
	}
	if q.listUsersStmt, err = db.PrepareContext(ctx, listUsers); err != nil {
		return nil, fmt.Errorf("error preparing query ListUsers: %w", err)
	}
	if q.listWorkoutsStmt, err = db.PrepareContext(ctx, listWorkouts); err != nil {
		return nil, fmt.Errorf("error preparing query ListWorkouts: %w", err)
	}
	if q.updateSetStmt, err = db.PrepareContext(ctx, updateSet); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateSet: %w", err)
	}
	if q.upsertExerciseStmt, err = db.PrepareContext(ctx, upsertExercise); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertExercise: %w", err)
	}
	if q.upsertUserImageStmt, err = db.PrepareContext(ctx, upsertUserImage); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertUserImage: %w", err)
	}
	if q.upsertWorkoutStmt, err = db.PrepareContext(ctx, upsertWorkout); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertWorkout: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createExerciseStmt != nil {
		if cerr := q.createExerciseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createExerciseStmt: %w", cerr)
		}
	}
	if q.createSetStmt != nil {
		if cerr := q.createSetStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createSetStmt: %w", cerr)
		}
	}
	if q.createUserImageStmt != nil {
		if cerr := q.createUserImageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserImageStmt: %w", cerr)
		}
	}
	if q.createUsersStmt != nil {
		if cerr := q.createUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUsersStmt: %w", cerr)
		}
	}
	if q.createWorkoutStmt != nil {
		if cerr := q.createWorkoutStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createWorkoutStmt: %w", cerr)
		}
	}
	if q.deleteExerciseStmt != nil {
		if cerr := q.deleteExerciseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteExerciseStmt: %w", cerr)
		}
	}
	if q.deleteSetsStmt != nil {
		if cerr := q.deleteSetsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteSetsStmt: %w", cerr)
		}
	}
	if q.deleteUserImageStmt != nil {
		if cerr := q.deleteUserImageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserImageStmt: %w", cerr)
		}
	}
	if q.deleteUserWorkoutsStmt != nil {
		if cerr := q.deleteUserWorkoutsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserWorkoutsStmt: %w", cerr)
		}
	}
	if q.deleteUsersStmt != nil {
		if cerr := q.deleteUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUsersStmt: %w", cerr)
		}
	}
	if q.getUserStmt != nil {
		if cerr := q.getUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStmt: %w", cerr)
		}
	}
	if q.getUserImageStmt != nil {
		if cerr := q.getUserImageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserImageStmt: %w", cerr)
		}
	}
	if q.getUserSetsStmt != nil {
		if cerr := q.getUserSetsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserSetsStmt: %w", cerr)
		}
	}
	if q.getUserWorkoutStmt != nil {
		if cerr := q.getUserWorkoutStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserWorkoutStmt: %w", cerr)
		}
	}
	if q.listExercisesStmt != nil {
		if cerr := q.listExercisesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listExercisesStmt: %w", cerr)
		}
	}
	if q.listImagesStmt != nil {
		if cerr := q.listImagesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listImagesStmt: %w", cerr)
		}
	}
	if q.listSetsStmt != nil {
		if cerr := q.listSetsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listSetsStmt: %w", cerr)
		}
	}
	if q.listUsersStmt != nil {
		if cerr := q.listUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listUsersStmt: %w", cerr)
		}
	}
	if q.listWorkoutsStmt != nil {
		if cerr := q.listWorkoutsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listWorkoutsStmt: %w", cerr)
		}
	}
	if q.updateSetStmt != nil {
		if cerr := q.updateSetStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateSetStmt: %w", cerr)
		}
	}
	if q.upsertExerciseStmt != nil {
		if cerr := q.upsertExerciseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertExerciseStmt: %w", cerr)
		}
	}
	if q.upsertUserImageStmt != nil {
		if cerr := q.upsertUserImageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertUserImageStmt: %w", cerr)
		}
	}
	if q.upsertWorkoutStmt != nil {
		if cerr := q.upsertWorkoutStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertWorkoutStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                     DBTX
	tx                     *sql.Tx
	createExerciseStmt     *sql.Stmt
	createSetStmt          *sql.Stmt
	createUserImageStmt    *sql.Stmt
	createUsersStmt        *sql.Stmt
	createWorkoutStmt      *sql.Stmt
	deleteExerciseStmt     *sql.Stmt
	deleteSetsStmt         *sql.Stmt
	deleteUserImageStmt    *sql.Stmt
	deleteUserWorkoutsStmt *sql.Stmt
	deleteUsersStmt        *sql.Stmt
	getUserStmt            *sql.Stmt
	getUserImageStmt       *sql.Stmt
	getUserSetsStmt        *sql.Stmt
	getUserWorkoutStmt     *sql.Stmt
	listExercisesStmt      *sql.Stmt
	listImagesStmt         *sql.Stmt
	listSetsStmt           *sql.Stmt
	listUsersStmt          *sql.Stmt
	listWorkoutsStmt       *sql.Stmt
	updateSetStmt          *sql.Stmt
	upsertExerciseStmt     *sql.Stmt
	upsertUserImageStmt    *sql.Stmt
	upsertWorkoutStmt      *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                     tx,
		tx:                     tx,
		createExerciseStmt:     q.createExerciseStmt,
		createSetStmt:          q.createSetStmt,
		createUserImageStmt:    q.createUserImageStmt,
		createUsersStmt:        q.createUsersStmt,
		createWorkoutStmt:      q.createWorkoutStmt,
		deleteExerciseStmt:     q.deleteExerciseStmt,
		deleteSetsStmt:         q.deleteSetsStmt,
		deleteUserImageStmt:    q.deleteUserImageStmt,
		deleteUserWorkoutsStmt: q.deleteUserWorkoutsStmt,
		deleteUsersStmt:        q.deleteUsersStmt,
		getUserStmt:            q.getUserStmt,
		getUserImageStmt:       q.getUserImageStmt,
		getUserSetsStmt:        q.getUserSetsStmt,
		getUserWorkoutStmt:     q.getUserWorkoutStmt,
		listExercisesStmt:      q.listExercisesStmt,
		listImagesStmt:         q.listImagesStmt,
		listSetsStmt:           q.listSetsStmt,
		listUsersStmt:          q.listUsersStmt,
		listWorkoutsStmt:       q.listWorkoutsStmt,
		updateSetStmt:          q.updateSetStmt,
		upsertExerciseStmt:     q.upsertExerciseStmt,
		upsertUserImageStmt:    q.upsertUserImageStmt,
		upsertWorkoutStmt:      q.upsertWorkoutStmt,
	}
}