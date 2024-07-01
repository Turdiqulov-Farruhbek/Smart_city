package postgres

import (
	"context"
	"log/slog"

	citi "citizen/genproto/citizen"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/rs/zerolog/log"
)

type Document struct {
	db *pgx.Conn
}

func NewDocument(db *pgx.Conn) *Document {
	return &Document{db: db}
}

// UploadDocument inserts a new document into the database
func (r *Document) UploadDocument(ctx context.Context, req *citi.DocumentCreate) (*citi.Document, error) {
	DocumentId := uuid.NewString()
	query := `
    INSERT INTO documents (document_id, citizen_id, document_type, file_path, create_at)
    VALUES ($1, $2, $3, $4, NOW())
    RETURNING document_id, citizen_id, document_type, file_path, create_at, update_at, delete_at`

	docRes := citi.Document{}
	err := r.db.QueryRow(ctx, query,
		DocumentId,
		req.CitizenId,
		req.DocumentType,
		req.FilePath,
	).Scan(
		&docRes.DocumentId,
		&docRes.CitizenId,
		&docRes.DocumentType,
		&docRes.FilePath,
		&docRes.CreateAt,
		&docRes.UpdateAt,
		&docRes.DeleteAt,
	)

	if err != nil {
		log.Error().Err(err).Msg("Error uploading document")
		slog.Error("Error uploading document", err)
		return nil, err
	}

	return &docRes, nil
}

// GetCitizenDocuments retrieves documents by Citizen ID
func (r *Document) GetCitizenDocuments(ctx context.Context, req *citi.ById) (*citi.DocumentList, error) {
	query := `
    SELECT 
        document_id, citizen_id, document_type, file_path, create_at, update_at, delete_at
    FROM documents
    WHERE citizen_id = $1 AND deleted_at = 0`

	rows, err := r.db.Query(ctx, query, req.Id)
	if err != nil {
		log.Error().Err(err).Msg("Error retrieving citizen documents")
		slog.Error("Error retrieving citizen documents", err)
		return nil, err
	}
	defer rows.Close()

	docList := citi.DocumentList{}
	for rows.Next() {
		doc := citi.Document{}
		if err := rows.Scan(
			&doc.DocumentId,
			&doc.CitizenId,
			&doc.DocumentType,
			&doc.FilePath,
			&doc.CreateAt,
			&doc.UpdateAt,
			&doc.DeleteAt,
		); err != nil {
			log.Error().Err(err).Msg("Error scanning document row")
			slog.Error("Error scanning document row", err)
			return nil, err
		}
		docList.Model = append(docList.Model, &doc)
	}

	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating over document rows")
		slog.Error("Error iterating over document rows", err)
		return nil, err
	}

	return &docList, nil
}
