package client

import (
	"context"
	"time"

	"github.com/dn-github/ratings/pb"
)

// GetRating ...
func GetRating(c pb.RatingServiceClient, book string) (int64, error) {
	ratingRequest := pb.Book{
		Name: book,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	res, err := c.Ratings(ctx, &ratingRequest)
	if err != nil {
		return 0, err
	}

	return res.Rating, nil
}
