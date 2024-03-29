package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"

	"github.com/dbut2/butla/shortener/pkg/models"
	"github.com/dbut2/butla/shortener/pkg/store"
)

type Config struct {
	Hostname string `yaml:"hostname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Database struct {
	wg sync.WaitGroup
	db *sql.DB
}

var _ store.LinkStore = new(Database)
var _ store.UserStore = new(Database)

func New(c *Config) (*Database, error) {
	db := &Database{}

	connStr := fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true", c.Username, c.Password, c.Hostname, c.Database)
	db.wg.Add(1)
	go db.openConn(connStr)

	return db, nil
}

func (d *Database) openConn(connStr string) {
	conn, err := sql.Open("mysql", connStr)
	if err != nil {
		panic(err.Error())
	}

	d.db = conn

	d.wg.Done()
}

func (d *Database) SetLink(ctx context.Context, link models.Link) error {
	d.wg.Wait()

	stmt, err := d.db.PrepareContext(ctx, "INSERT INTO links (code, url, expiry, ip) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	dbl := dbLink{
		code: link.Code,
		url:  link.Url,
		expiry: sql.NullTime{
			Time:  link.Expiry.Value,
			Valid: link.Expiry.Valid,
		},
		ip: sql.NullString{
			String: link.IP.Value,
			Valid:  link.IP.Valid,
		},
	}

	res, err := stmt.Exec(dbl.code, dbl.url, dbl.expiry, dbl.ip)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return errors.New("not 1 row affected")
	}

	return nil
}

func (d *Database) GetLink(ctx context.Context, code string) (models.Link, bool, error) {
	d.wg.Wait()

	rows, err := d.db.QueryContext(ctx, "SELECT code, url, expiry, ip FROM links WHERE code = ?", code)
	if err != nil {
		return models.Link{}, false, err
	}

	if !rows.Next() {
		return models.Link{}, false, nil
	}

	var dbl dbLink
	err = rows.Scan(&dbl.code, &dbl.url, &dbl.expiry, &dbl.ip)
	if err != nil {
		return models.Link{}, false, err
	}

	link := models.Link{
		Code: dbl.code,
		Url:  dbl.url,
		Expiry: models.NullTime{
			Valid: dbl.expiry.Valid,
			Value: dbl.expiry.Time,
		},
		IP: models.NullString{
			Valid: dbl.ip.Valid,
			Value: dbl.ip.String,
		},
	}

	return link, true, nil
}

func (d *Database) SetUser(ctx context.Context, User models.User) error {
	d.wg.Wait()

	stmt, err := d.db.PrepareContext(ctx, "INSERT INTO users (username, email, name, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	dbu := dbUser{
		username: User.Username,
		email:    User.Email,
		name:     User.Name,
		password: User.Password,
	}

	res, err := stmt.Exec(dbu.username, dbu.email, dbu.name, dbu.password)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return errors.New("not 1 row affected")
	}

	return nil
}

func (d *Database) GetUser(ctx context.Context, username string) (models.User, bool, error) {
	d.wg.Wait()

	rows, err := d.db.QueryContext(ctx, "SELECT username, email, name, password FROM users WHERE username = ?", username)
	if err != nil {
		return models.User{}, false, err
	}

	if !rows.Next() {
		return models.User{}, false, nil
	}

	var dbu dbUser
	err = rows.Scan(&dbu.username, &dbu.email, &dbu.name, &dbu.password)
	if err != nil {
		return models.User{}, false, err
	}

	user := models.User{
		Username: dbu.username,
		Email:    dbu.email,
		Name:     dbu.name,
		Password: dbu.password,
	}

	return user, true, nil
}

func (d *Database) GetAll(ctx context.Context) ([]models.Link, error) {
	d.wg.Wait()

	rows, err := d.db.QueryContext(ctx, "SELECT code, url, expiry, ip FROM links")
	if err != nil {
		return nil, err
	}

	var links []models.Link

	for rows.Next() {
		var dbl dbLink
		err = rows.Scan(&dbl.code, &dbl.url, &dbl.expiry, &dbl.ip)
		if err != nil {
			return nil, err
		}
		link := models.Link{
			Code: dbl.code,
			Url:  dbl.url,
			Expiry: models.NullTime{
				Valid: dbl.expiry.Valid,
				Value: dbl.expiry.Time,
			},
			IP: models.NullString{
				Valid: dbl.ip.Valid,
				Value: dbl.ip.String,
			},
		}
		links = append(links, link)
	}

	return links, nil
}

type dbLink struct {
	code   string
	url    string
	expiry sql.NullTime
	ip     sql.NullString
}

type dbUser struct {
	username string
	email    string
	name     string
	password string
}
