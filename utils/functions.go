package utils

import "github.com/oliverperboni/GoApi/schemas"

func BookConvert(b schemas.Book) schemas.BookJSON {
	return schemas.BookJSON{
		ID:              b.ID,
		Title:           b.Title,
		Author:          b.Author,
		Genre:           b.Genre,
		PublicationDate: b.PublicationDate,
		Synopsis:        b.Synopsis,
		CoverImageURL:   b.CoverImageURL,
	}
}
