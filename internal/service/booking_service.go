package service

import (
	"context"
	"log/slog"
	"time"

	"aviasales/internal/models"
	"aviasales/internal/repository"
)

type BookingService struct {
	repo         *repository.BookingRepo
	flightsRepo  *repository.FlightsRepo
	segmentsRepo *repository.SegmentsRepo
	logger       *slog.Logger
}

func NewBookingService(repo *repository.BookingRepo, flights *repository.FlightsRepo, segments *repository.SegmentsRepo, logger *slog.Logger) *BookingService {
	return &BookingService{
		repo:         repo,
		logger:       logger,
		segmentsRepo: segments,
		flightsRepo:  flights,
	}
}

func (s *BookingService) CreateBooking(ctx context.Context, ticketNo string) (*models.Booking, error) {
	s.logger.Info("creating booking", slog.String("ticket_no", ticketNo))
	return s.repo.CreateFromTicketInfo(ctx, ticketNo)
}

func (s *BookingService) UpdateBooking(ctx context.Context, ticketNo, seat string, price string) (*models.Booking, error) {
	s.logger.Info("updating booking", slog.String("ticket_no", ticketNo))
	return s.repo.Update(ctx, ticketNo, seat, price)
}

func (s *BookingService) DeleteBooking(ctx context.Context, ticketNo string) error {
	s.logger.Info("deleting booking", slog.String("ticket_no", ticketNo))
	return s.repo.Delete(ctx, ticketNo)
}

func (s *BookingService) ListAvailableFlights(ctx context.Context, start, end time.Time, depCity, depCountry string, arrCountries []string) ([]models.Flight, error) {
	return s.flightsRepo.ListAvailable(ctx, start, end, depCity, depCountry, arrCountries)
}

func (s *BookingService) ListFreeSeats(ctx context.Context, flightID int, fare string) ([]models.Segment, error) {
	return s.segmentsRepo.ListFree(ctx, flightID, fare)
}
