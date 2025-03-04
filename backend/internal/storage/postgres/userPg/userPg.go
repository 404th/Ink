package userPg

import (
	"context"
	"database/sql"
	"strings"

	"github.com/404th/Ink/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userPg struct {
	db *pgxpool.Pool
}

func NewUserPg(db *pgxpool.Pool) *userPg {
	return &userPg{
		db: db,
	}
}

func (u *userPg) SignupUser(ctx context.Context, req *model.SignupUserRequest) (resp *model.SignupUserResponse, err error) {
	resp = &model.SignupUserResponse{}

	var (
		query strings.Builder
		id    string
	)

	query.WriteString(`
		INSERT INTO "users" (
			username,
			email,
			password,
			avatar_url
		) VALUES (
			$1,
			$2,
			$3,
			$4
		) RETURNING id
	`)

	if err = u.db.QueryRow(
		ctx,
		query.String(),
		req.Username,
		req.Email,
		req.Password,
		req.AvatarUrl,
	).Scan(&id); err != nil {
		return
	}

	resp.Id = id
	resp.Username = req.Username
	resp.Email = req.Email
	resp.AvatarUrl = req.AvatarUrl

	return
}

func (u *userPg) LoginUser(ctx context.Context, req *model.LoginUserRequest) (resp *model.LoginUserResponse, err error) {
	resp = &model.LoginUserResponse{}

	var (
		query strings.Builder
	)

	query.WriteString(`
		SELECT 
			"users".id id,
			"users".username username,
			"users".email email,
			"users".avatar_url avatar_url,
			"users".password password
		FROM "users" 
		WHERE "users".username = $1
	`)

	var (
		id_sql, username_sql, email_sql, password_sql, avatar_url_sql sql.NullString
	)

	if err = u.db.QueryRow(ctx, query.String(), req.Username).Scan(
		&id_sql,
		&username_sql,
		&email_sql,
		&avatar_url_sql,
		&password_sql,
	); err != nil {
		return
	}

	if id_sql.Valid {
		resp.Id = id_sql.String
	}

	if username_sql.Valid {
		resp.Username = username_sql.String
	}

	if email_sql.Valid {
		resp.Email = email_sql.String
	}

	if password_sql.Valid {
		resp.Password = password_sql.String
	}

	if avatar_url_sql.Valid {
		resp.AvatarUrl = avatar_url_sql.String
	}

	return
}

func (u *userPg) GetUser(ctx context.Context, req *model.Id) (resp *model.User, err error) {
	resp = &model.User{}

	var (
		query strings.Builder
	)

	query.WriteString(`
		SELECT 
			id,
			username,
			email,
			avatar_url,
			created_at
		FROM 
			"users"
		WHERE id = $1
	`)

	var (
		usr       model.User
		createdAt sql.NullString
	)

	if err = u.db.QueryRow(ctx, query.String(), req.Id).Scan(
		&usr.Id,
		&usr.Username,
		&usr.Email,
		&usr.AvatarUrl,
		&createdAt,
	); err != nil {
		return
	}

	if createdAt.Valid {
		usr.CreatedAt = createdAt.String
	}

	resp = &usr

	return
}

// func (u *userPg) UpdateUser(ctx context.Context, req *model.UpdateUserRequest) (resp *model.Message, err error) {
// 	var (
// 		query, set, filter strings.Builder
// 		params             = make(map[string]interface{})
// 		arr                []interface{}
// 	)

// 	filter.WriteString(` WHERE id = :id`)
// 	params = map[string]interface{}{
// 		"id": req.Id,
// 	}

// 	resp = &model.Message{}

// 	query.WriteString(`UPDATE "user_data" SET `)
// 	set.WriteString(` updated_at = now()`)

// 	p := map[string]interface{}{
// 		// "role_name":   req.RoleName,
// 		// "description": req.Description,
// 	}

// 	for key, val := range p {
// 		if val != "" && key != "active" && key != "id" {
// 			params[key] = val
// 			set.WriteString(", ")
// 			set.WriteString(key)
// 			set.WriteString(" = :")
// 			set.WriteString(key)
// 		}
// 	}

// 	query.WriteString(set.String())
// 	query.WriteString(filter.String())

// 	q, arr := helper.ReplaceQueryParams(query.String(), params)
// 	if err != nil {
// 		return
// 	}

// 	result, err := u.db.Exec(ctx, q, arr...)
// 	if err != nil {
// 		return
// 	}

// 	ra := result.RowsAffected()
// 	if ra < 1 {
// 		err = errors.New("sql: no rows in result set")
// 		resp.Description = "Kiritilgan ma'lumot bo'yicha o'zgartirish amalga oshmadi"
// 		return
// 	}

// 	resp.Description = "Muvaffaqiyatli o'zgartirildi"
// 	return
// }

// func (u *userPg) GetAllUsers(ctx context.Context, req *model.GetAllUsersRequest) (resp *model.GetAllUsersResponse, err error) {
// 	var (
// 		query, cQuery, filter, offset, limit, arrangement, order, groupBy strings.Builder
// 		arr                                                               []interface{}
// 	)

// 	resp = &model.GetAllUsersResponse{}
// 	metadata := model.Metadata{}

// 	metadata.Page = req.Metadata.Page
// 	metadata.Limit = req.Metadata.Limit

// 	params := map[string]interface{}{}

// 	offset.WriteString(` OFFSET :offset`)
// 	params["offset"] = (req.Metadata.Page - 1) * req.Metadata.Limit
// 	limit.WriteString(` LIMIT :limit`)
// 	params["limit"] = req.Metadata.Limit
// 	arrangement.WriteString(" DESC")
// 	order.WriteString(` ORDER BY "user".created_at`)

// 	groupBy.WriteString(`
// 		GROUP BY
// 			"user".id,
// 			"user".user_role_id,
// 			"user_role".role_name,
// 			"user_data".username,
// 			"user_data".first_name,
// 			"user_data".middle_name,
// 			"user_data".last_name,
// 			"user_data".passport_number,
// 			"user_data".passport_pinfl,
// 			"user_data".phone_number
// 	`)

// 	query.WriteString(`
// 		select
// 			"user".id,
// 			"user".user_role_id,
// 			"user_role".role_name,
// 			"user".user_status_id,
// 			"user_data".username,
// 			"user_data".first_name,
// 			"user_data".middle_name,
// 			"user_data".last_name,
// 			"user_data".passport_number,
// 			"user_data".passport_pinfl,
// 			"user_data".phone_number
// 		from "user"
// 			join "user_data" on "user".user_data_id = "user_data".id
// 			join "user_role" on "user".user_role_id = "user_role".id
// 	`)

// 	cQuery.WriteString(`
// 		select
// 			count(*)
// 		from "user"
// 			join "user_data" on "user".user_data_id = "user_data".id
// 			join "user_role" on "user".user_role_id = "user_role".id
// 	`)

// 	filter.WriteString(` WHERE "user".deleted_at IS NULL `)

// 	if len(req.Username) > 0 {
// 		filter.WriteString(`AND "user_data".username ILIKE ('%' || :username || '%')`)
// 		params["username"] = req.Username
// 	}

// 	if len(req.PassportPinfl) > 0 {
// 		filter.WriteString(`AND "user_data".passport_pinfl ILIKE ('%' || :passport_pinfl || '%')`)
// 		params["passport_pinfl"] = req.PassportPinfl
// 	}

// 	if len(req.PassportNumber) > 0 {
// 		filter.WriteString(`AND "user_data".passport_number ILIKE ('%' || :passport_number || '%')`)
// 		params["passport_number"] = req.PassportNumber
// 	}

// 	if len(req.Id) > 0 {
// 		filter.WriteString(`AND "user".id = :id`)
// 		params["id"] = req.Id
// 	}

// 	if len(req.UserStatusId) > 0 {
// 		filter.WriteString(`AND "user".user_status_id = :user_status_id`)
// 		params["user_status_id"] = req.UserStatusId
// 	}

// 	if len(req.UserRoleId) > 0 {
// 		filter.WriteString(`AND "user".user_role_id = :user_role_id`)
// 		params["user_role_id"] = req.UserRoleId
// 	}

// 	if req.Metadata.Page > 0 {
// 		req.Metadata.Page = (req.Metadata.Page - 1) * req.Metadata.Limit
// 	} else {
// 		err = errors.New("Noto'g'ri ma'lumot kitildi. Ichki xatolik")
// 		return
// 	}

// 	var count int32

// 	cQuery.WriteString(filter.String())
// 	cQ, arr := helper.ReplaceQueryParams(cQuery.String(), params)
// 	err = u.db.QueryRow(ctx, cQ, arr...).Scan(&count)

// 	metadata.Count = count

// 	// don't change order here
// 	query.WriteString(filter.String())
// 	query.WriteString(groupBy.String())
// 	query.WriteString(order.String())
// 	query.WriteString(arrangement.String())
// 	query.WriteString(limit.String())
// 	query.WriteString(offset.String())

// 	q, arr := helper.ReplaceQueryParams(query.String(), params)
// 	rows, err := u.db.Query(ctx, q, arr...)
// 	if err != nil {
// 		err = errors.New("Noto'g'ri ma'lumot kitilgan. Ichki xatolik")
// 		return
// 	}
// 	defer rows.Close()

// 	var urs []*model.UserBody = []*model.UserBody{}

// 	for rows.Next() {
// 		var (
// 			newRow model.UserBody
// 		)

// 		if err = rows.Scan(
// 			&newRow.Id,
// 			&newRow.UserRoleId,
// 			&newRow.UserRoleName,
// 			&newRow.UserStatusId,
// 			&newRow.Username,
// 			&newRow.Firstname,
// 			&newRow.Middlename,
// 			&newRow.Lastname,
// 			&newRow.PassportNumber,
// 			&newRow.PassportPinfl,
// 			&newRow.PhoneNumber,
// 		); err != nil {
// 			rows.Close()
// 			return
// 		}

// 		urs = append(urs, &newRow)
// 	}
// 	rows.Close()

// 	resp.Users = urs

// 	resp.Metadata = &metadata

// 	return
// }

// func (u *userPg) DeactivateUser() {
// 	return
// }

// // user role pg
// func (r *userPg) CreateUserRole(ctx context.Context, req *model.CreateUserRoleRequest) (resp *model.Id, err error) {
// 	var query strings.Builder

// 	resp = &model.Id{}

// 	query.WriteString(`
// 		INSERT INTO "user_role" (
// 			role_name,
// 			description
// 		) VALUES (
// 			$1,
// 			$2
// 		) RETURNING id
// 	`)

// 	var id string

// 	if err = r.db.QueryRow(ctx, query.String(), req.RoleName, req.Description).Scan(&id); err != nil {
// 		return
// 	}

// 	resp.Id = id

// 	return
// }

// func (r *userPg) UpdateUserRole(ctx context.Context, req *model.UpdateUserRoleRequest) (resp *model.Message, err error) {
// 	var (
// 		query, set, filter strings.Builder
// 		params             = make(map[string]interface{})
// 		arr                []interface{}
// 	)

// 	filter.WriteString(` WHERE id = :id`)
// 	params = map[string]interface{}{
// 		"id": req.Id,
// 	}

// 	resp = &model.Message{}

// 	query.WriteString(`UPDATE "user_role" SET `)
// 	set.WriteString(` updated_at = now()`)

// 	p := map[string]interface{}{
// 		"role_name":   req.RoleName,
// 		"description": req.Description,
// 	}

// 	for key, val := range p {
// 		if val != "" && key != "active" && key != "id" {
// 			params[key] = val
// 			set.WriteString(", ")
// 			set.WriteString(key)
// 			set.WriteString(" = :")
// 			set.WriteString(key)
// 		}
// 	}

// 	query.WriteString(set.String())
// 	query.WriteString(filter.String())

// 	q, arr := helper.ReplaceQueryParams(query.String(), params)
// 	if err != nil {
// 		return
// 	}

// 	result, err := r.db.Exec(ctx, q, arr...)
// 	if err != nil {
// 		return
// 	}

// 	ra := result.RowsAffected()
// 	if ra < 1 {
// 		err = errors.New("sql: no rows in result set")
// 		resp.Description = "Kiritilgan ma'lumot bo'yicha o'zgartirish amalga oshmadi"
// 		return
// 	}

// 	resp.Description = "Muvaffaqiyatli o'zgartirildi"
// 	return
// }

// func (r *userPg) GetAllUserRoles(ctx context.Context, req *model.GetAllUserRolesRequest) (resp *model.GetAllUserRolesResponse, err error) {
// 	var (
// 		query, cQuery, filter, offset, limit, arrangement, order strings.Builder
// 		arr                                                      []interface{}
// 	)

// 	resp = &model.GetAllUserRolesResponse{}
// 	metadata := model.Metadata{}

// 	metadata.Page = req.Metadata.Page
// 	metadata.Limit = req.Metadata.Limit

// 	params := map[string]interface{}{}

// 	offset.WriteString(` OFFSET :offset`)
// 	params["offset"] = (req.Metadata.Page - 1) * req.Metadata.Limit
// 	limit.WriteString(` LIMIT :limit`)
// 	params["limit"] = req.Metadata.Limit
// 	arrangement.WriteString(" DESC")
// 	order.WriteString(` ORDER BY created_at`)

// 	query.WriteString(`
// 		SELECT
// 			id,
// 			role_name,
// 			description,
// 			created_at,
// 			updated_at
// 		FROM
// 			"user_role"
// 	`)

// 	cQuery.WriteString(`
// 		SELECT count(*) FROM "user_role"
// 	`)

// 	filter.WriteString(` WHERE deleted_at IS NULL `)

// 	if len(req.RoleName) > 0 {
// 		filter.WriteString(`AND role_name ILIKE ('%' || :role_name || '%')`)
// 		params["role_name"] = req.RoleName
// 	}

// 	if len(req.Id) > 0 {
// 		filter.WriteString(`AND id = :id`)
// 		params["id"] = req.Id
// 	}

// 	if req.Metadata.Page > 0 {
// 		req.Metadata.Page = (req.Metadata.Page - 1) * req.Metadata.Limit
// 	} else {
// 		err = errors.New("Noto'g'ri ma'lumot kitildi. Ichki xatolik")
// 		return
// 	}

// 	var count int32

// 	cQuery.WriteString(filter.String())
// 	cQ, arr := helper.ReplaceQueryParams(cQuery.String(), params)
// 	err = r.db.QueryRow(ctx, cQ, arr...).Scan(&count)

// 	metadata.Count = count

// 	// don't change order here
// 	query.WriteString(filter.String())
// 	query.WriteString(order.String())
// 	query.WriteString(arrangement.String())
// 	query.WriteString(limit.String())
// 	query.WriteString(offset.String())

// 	q, arr := helper.ReplaceQueryParams(query.String(), params)
// 	rows, err := r.db.Query(ctx, q, arr...)
// 	if err != nil {
// 		err = errors.New("Noto'g'ri ma'lumot kitilgan. Ichki xatolik")
// 		return
// 	}
// 	defer rows.Close()

// 	var urs []*model.UserRole = []*model.UserRole{}

// 	for rows.Next() {
// 		var (
// 			newRow         model.UserRole
// 			descriptionSql sql.NullString

// 			createdAtSql sql.NullString
// 			updatedAtSql sql.NullString
// 		)

// 		if err = rows.Scan(
// 			&newRow.Id,
// 			&newRow.RoleName,
// 			&descriptionSql,
// 			&createdAtSql,
// 			&updatedAtSql,
// 		); err != nil {
// 			rows.Close()
// 			return
// 		}

// 		if descriptionSql.Valid {
// 			newRow.Description = descriptionSql.String
// 		}

// 		if createdAtSql.Valid {
// 			newRow.CreatedAt = createdAtSql.String
// 		}

// 		if updatedAtSql.Valid {
// 			newRow.UpdatedAt = updatedAtSql.String
// 		}

// 		urs = append(urs, &newRow)
// 	}

// 	resp.UserRoles = urs

// 	resp.Metadata = &metadata

// 	return
// }

// func (r *userPg) DeleteUserRole(ctx context.Context, req *model.Id) (resp *model.Message, err error) {
// 	resp = &model.Message{}

// 	var (
// 		query strings.Builder
// 	)

// 	query.WriteString(`
// 		UPDATE "user_role"
// 		SET deleted_at = NOW()
// 		WHERE deleted_at IS NULL AND id = $1
// 	`)

// 	ra, err := r.db.Exec(ctx, query.String(), req.Id)
// 	if err != nil {
// 		println(err)
// 		return
// 	}

// 	if ra.RowsAffected() < 1 {
// 		err = errors.New("sql: no rows in result set")
// 		return
// 	}

// 	resp.Description = "Muvaffaqiyatli o'chirildi"

// 	return
// }

// // user data storage
// func (r *userPg) CreateUserData(ctx context.Context, req *model.CreateUserDataRequest) (resp *model.Id, err error) {
// 	var query strings.Builder

// 	resp = &model.Id{}

// 	query.WriteString(`
// 		INSERT INTO "user_data" (
// 			username,
// 			password,
// 			first_name,
// 			middle_name,
// 			last_name,
// 			passport_number,
// 			passport_pinfl,
// 			phone_number
// 		) VALUES (
// 			$1,
// 			$2,
// 			$3,
// 			$4,
// 			$5,
// 			$6,
// 			$7,
// 			$8
// 		) RETURNING id
// 	`)

// 	var id string

// 	if err = r.db.QueryRow(ctx, query.String(),
// 		req.Username,
// 		req.Password,
// 		req.Firstname,
// 		req.Middlename,
// 		req.Lastname,
// 		req.PassportNumber,
// 		req.PassportPinfl,
// 		req.PhoneNumber,
// 	).Scan(&id); err != nil {
// 		return
// 	}

// 	resp.Id = id

// 	return
// }

// func (r *userPg) UpdateUserData(ctx context.Context, req *model.UpdateUserDataRequest) (resp *model.Message, err error) {
// 	var (
// 		query, set, filter strings.Builder
// 		params             = make(map[string]interface{})
// 		arr                []interface{}
// 	)

// 	filter.WriteString(` WHERE id = :id`)
// 	params = map[string]interface{}{
// 		"id": req.Id,
// 	}

// 	resp = &model.Message{}

// 	query.WriteString(`UPDATE "user_data" SET `)
// 	set.WriteString(` updated_at = now()`)

// 	p := map[string]interface{}{
// 		"username":        req.Username,
// 		"password":        req.Password,
// 		"first_name":      req.Firstname,
// 		"middle_name":     req.Middlename,
// 		"last_name":       req.Lastname,
// 		"passport_number": req.PassportNumber,
// 		"passport_pinfl":  req.PassportPinfl,
// 		"phone_number":    req.PhoneNumber,
// 	}

// 	for key, val := range p {
// 		if val != "" && key != "id" {
// 			params[key] = val
// 			set.WriteString(", ")
// 			set.WriteString(key)
// 			set.WriteString(" = :")
// 			set.WriteString(key)
// 		}
// 	}

// 	query.WriteString(set.String())
// 	query.WriteString(filter.String())

// 	q, arr := helper.ReplaceQueryParams(query.String(), params)
// 	if err != nil {
// 		return
// 	}

// 	result, err := r.db.Exec(ctx, q, arr...)
// 	if err != nil {
// 		return
// 	}

// 	ra := result.RowsAffected()
// 	if ra < 1 {
// 		err = errors.New("sql: no rows in result set")
// 		resp.Description = "Kiritilgan ma'lumot bo'yicha o'zgartirish amalga oshmadi"
// 		return
// 	}

// 	resp.Description = "Muvaffaqiyatli o'zgartirildi"
// 	return
// }

// func (r *userPg) GetAllUserDatas(ctx context.Context, req *model.GetAllUserDatasRequest) (resp *model.GetAllUserDatasResponse, err error) {
// 	var (
// 		query, cQuery, filter, offset, limit, arrangement, order strings.Builder
// 		arr                                                      []interface{}
// 	)

// 	resp = &model.GetAllUserDatasResponse{}
// 	metadata := &model.Metadata{}

// 	metadata.Page = req.Metadata.Page
// 	metadata.Limit = req.Metadata.Limit

// 	params := map[string]interface{}{}

// 	offset.WriteString(` OFFSET :offset`)
// 	params["offset"] = (req.Metadata.Page - 1) * req.Metadata.Limit
// 	limit.WriteString(` LIMIT :limit`)
// 	params["limit"] = req.Metadata.Limit
// 	arrangement.WriteString(" DESC")
// 	order.WriteString(` ORDER BY created_at`)

// 	query.WriteString(`
// 		SELECT
// 			id,
// 			username,
// 			password,
// 			first_name,
// 			middle_name,
// 			last_name,
// 			passport_number,
// 			passport_pinfl,
// 			phone_number
// 		FROM
// 			"user_data"
// 	`)

// 	cQuery.WriteString(`
// 		SELECT count(*) FROM "user_data"
// 	`)

// 	filter.WriteString(` WHERE deleted_at IS NULL `)

// 	if len(req.Username) > 0 {
// 		filter.WriteString(`AND username ILIKE ('%' || :username || '%')`)
// 		params["username"] = req.Username
// 	}

// 	if len(req.PassportNumber) > 0 {
// 		filter.WriteString(`AND passport_number ILIKE ('%' || :passport_number || '%')`)
// 		params["passport_number"] = req.PassportNumber
// 	}

// 	if len(req.PassportPinfl) > 0 {
// 		filter.WriteString(`AND passport_pinfl ILIKE ('%' || :passport_pinfl || '%')`)
// 		params["passport_pinfl"] = req.PassportPinfl
// 	}

// 	if len(req.Id) > 0 {
// 		filter.WriteString(`AND id = :id`)
// 		params["id"] = req.Id
// 	}

// 	if req.Metadata.Page > 0 {
// 		req.Metadata.Page = (req.Metadata.Page - 1) * req.Metadata.Limit
// 	} else {
// 		err = errors.New("Noto'g'ri ma'lumot kitildi. Ichki xatolik")
// 		return
// 	}

// 	var count int32

// 	cQuery.WriteString(filter.String())
// 	cQ, arr := helper.ReplaceQueryParams(cQuery.String(), params)
// 	err = r.db.QueryRow(ctx, cQ, arr...).Scan(&count)

// 	metadata.Count = count

// 	// don't change order here
// 	query.WriteString(filter.String())
// 	query.WriteString(order.String())
// 	query.WriteString(arrangement.String())
// 	query.WriteString(limit.String())
// 	query.WriteString(offset.String())

// 	q, arr := helper.ReplaceQueryParams(query.String(), params)
// 	rows, err := r.db.Query(ctx, q, arr...)
// 	if err != nil {
// 		err = errors.New("Noto'g'ri ma'lumot kitilgan. Ichki xatolik")
// 		return
// 	}
// 	defer rows.Close()

// 	var urs []*model.UserData = []*model.UserData{}

// 	for rows.Next() {
// 		var (
// 			newRow model.UserData

// 			createdAtSql sql.NullString
// 			updatedAtSql sql.NullString
// 		)

// 		// to be continues...
// 		if err = rows.Scan(
// 			&newRow.Id,
// 			&newRow.Username,
// 			&newRow.Password,
// 			&newRow.Firstname,
// 			&newRow.Middlename,
// 			&newRow.Lastname,
// 			&newRow.PassportNumber,
// 			&newRow.PassportPinfl,
// 			&newRow.PhoneNumber,
// 			&createdAtSql,
// 			&updatedAtSql,
// 		); err != nil {
// 			rows.Close()
// 			return
// 		}

// 		if updatedAtSql.Valid {
// 			newRow.UpdatedAt = updatedAtSql.String
// 		}

// 		if createdAtSql.Valid {
// 			newRow.CreatedAt = createdAtSql.String
// 		}

// 		urs = append(urs, &newRow)
// 	}

// 	resp.UserDatas = urs
// 	resp.Metadata = metadata

// 	return
// }

// func (r *userPg) DeleteUserData(ctx context.Context, req *model.Id) (resp *model.Message, err error) {
// 	resp = &model.Message{}

// 	var (
// 		query strings.Builder
// 	)

// 	query.WriteString(`
// 		UPDATE "user_data"
// 		SET
// 			deleted_at = now(),
// 			passport_pinfl_previous = passport_pinfl,
// 			passport_pinfl = uuid_generate_v4()
// 		WHERE deleted_at IS NULL AND id = $1
// 	`)

// 	ra, err := r.db.Exec(ctx, query.String(), req.Id)
// 	if err != nil {
// 		return
// 	}

// 	if ra.RowsAffected() < 1 {
// 		err = errors.New("sql: no rows in result set")
// 		return
// 	}

// 	resp.Description = "Muvaffaqiyatli o'chirildi"

// 	return
// }
