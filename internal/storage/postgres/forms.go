package postgres

import (
	"time"
)

// Form — структура для работы с таблицей forms в БД
type Form struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Schema    string    `json:"schema"`
	CreatedAt time.Time `json:"created_at"`
}

// Response — структура для работы с таблицей responses в БД
type Response struct {
	ID        int       `json:"id"`
	FormID    int       `json:"form_id"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
	Status    string    `json:"status"`
}

// CreateForm — добавляет новую форму в таблицу forms
// Принимает имя и схему формы, возвращает ошибку при неудаче
func (s *Storagepostgres) CreateForm(name, schema string) error {
	_, err := s.db.Exec(`INSERT INTO forms (name, schema, created_at) VALUES ($1, $2, now())`, name, schema)
	return err
}

// GetForms — получает список всех форм из таблицы forms
// Возвращает слайс структур Form и ошибку
func (s *Storagepostgres) GetForms() ([]Form, error) {
	rows, err := s.db.Query(`SELECT id, name, schema, created_at FROM forms`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var forms []Form
	for rows.Next() {
		var f Form
		if err := rows.Scan(&f.ID, &f.Name, &f.Schema, &f.CreatedAt); err != nil {
			return nil, err
		}
		forms = append(forms, f)
	}
	return forms, nil
}

// GetFormByID — получает одну форму по её ID из таблицы forms
// Возвращает структуру Form и ошибку
func (s *Storagepostgres) GetFormByID(id int) (Form, error) {
	var f Form
	err := s.db.QueryRow(`SELECT id, name, schema, created_at FROM forms WHERE id = $1`, id).
		Scan(&f.ID, &f.Name, &f.Schema, &f.CreatedAt)
	return f, err
}

// SaveResponse — сохраняет новый ответ на форму в таблицу responses
// Устанавливает статус по умолчанию 'new'

// TODO
func (s *Storagepostgres) SaveResponse(formID int, data string) error {
	_, err := s.db.Exec(
		`INSERT INTO feedback (form_id, data, created_at, status) VALUES ($1, $2, now(), 'new')`,
		formID, data,
	)
	return err
}

// GetResponsesByFormID — получает все ответы для формы по её ID из таблицы responses
// Возвращает слайс структур Response и ошибку
func (s *Storagepostgres) GetResponsesByFormID(formID int) ([]Response, error) {
	rows, err := s.db.Query(
		`SELECT id, form_id, data, created_at, status FROM feedback WHERE form_id = $1`,
		formID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var responses []Response
	for rows.Next() {
		var r Response
		if err := rows.Scan(&r.ID, &r.FormID, &r.Data, &r.CreatedAt, &r.Status); err != nil {
			return nil, err
		}
		responses = append(responses, r)
	}
	return responses, nil
}

// UpdateResponseStatus — обновляет статус конкретного ответа по его ID в таблице responses.
func (s *Storagepostgres) UpdateResponseStatus(id int, status string) error {
	_, err := s.db.Exec(`UPDATE feedback SET status = $1 WHERE id = $2`, status, id)
	return err
}
