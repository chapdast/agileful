package db

import (
	"bytes"
	"context"
	"strconv"
)

func (d database) Read(ctx context.Context, opts ...Option) ([]*Article, error) {
	var input []interface{}
	opt := &option{}
	for _, o := range opts {
		if err := o(opt); err != nil {
			return nil, err
		}
	}
	query := []byte(`select id, writer, body, extract(epoch from time)::bigint, view_count from articles`)
	if len(opt.fillerBy) != 0 {
		var cond [][]byte
		for k, filter := range opt.fillerBy {
			input = append(input, filter.Value)
			switch filter.Operator {
			case OprEqual:
				cond = append(cond, []byte(" "+k+" = $"+strconv.Itoa(len(input))))
			case OprNotEqual:
				cond = append(cond, []byte(" "+k+" != $"+strconv.Itoa(len(input))))
			case OprSubString:
				cond = append(cond, []byte(" "+k+" like '%'||$"+strconv.Itoa(len(input))+"||'%' "))
			}
		}
		query = append(query, []byte(" WHERE")...)
		query = append(query, bytes.Join(cond, []byte(" AND "))...)
	}

	if len(opt.orderBy) > 0 {
		query = append(query, []byte(" ORDER BY ")...)
		query = append(query, bytes.Join(opt.orderBy, []byte(" , "))...)
		if opt.orderDesc {
			query = append(query, []byte(" DESC ")...)
		} else {
			query = append(query, []byte(" ASC ")...)
		}
	}

	if opt.offset > 0 {
		input = append(input, opt.offset)
		query = append(query, []byte(" OFFSET $"+strconv.Itoa(len(input)))...)
	}
	if opt.limit > 0 {
		input = append(input, opt.limit)
		query = append(query, []byte(" LIMIT $"+strconv.Itoa(len(input)))...)
	}

	rows, err := d.conn.Query(ctx, string(query), input...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []*Article
	for rows.Next() {
		article := &Article{}
		if err := rows.Scan(&article.Id, &article.Writer, &article.Body, &article.CreateDate, &article.ViewCount); err != nil {
			return nil, err
		}
		list = append(list, article)
	}
	return list, nil
}
