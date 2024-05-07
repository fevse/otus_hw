package internalgrpc

import (
	"context"

	"github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/server/grpc/pb"
	"github.com/fevse/otus_hw/hw12_13_14_15_calendar/internal/storage"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *CalendarServer) Create(_ context.Context, r *pb.Event) (*emptypb.Empty, error) {
	e := storage.Event{
		ID:          r.ID,
		Title:       r.Title,
		Date:        r.Date.AsTime(),
		Duration:    r.Duration,
		Description: r.Description,
		UserID:      r.UserID,
		Reminder:    r.Reminder.AsTime(),
	}
	
	err := s.App.Storage.CreateEvent(&e)
	return nil, err
}

func (s *CalendarServer) Update(_ context.Context, r *pb.IDEvent) (*emptypb.Empty, error) {
	e := storage.Event{
		ID:          r.Event.ID,
		Title:       r.Event.Title,
		Date:        r.Event.Date.AsTime(),
		Duration:    r.Event.Duration,
		Description: r.Event.Description,
		UserID:      r.Event.UserID,
		Reminder:    r.Event.Reminder.AsTime(),
	}
	
	err := s.App.Storage.UpdateEvent(r.ID, &e)
	return nil, err
}

func (s *CalendarServer) Delete(_ context.Context, r *pb.ID) (*emptypb.Empty, error) {
	return nil, s.App.Storage.DeleteEvent(r.ID)
}

func (s *CalendarServer) Show(context.Context, *emptypb.Empty) (events *pb.Events, err error) {
	events = new(pb.Events)
	stor, err := s.App.Storage.List()
	if err != nil {
		return nil, err
	}
	events.Events = toPbEvents(stor)
	return

}

func (s *CalendarServer) ShowEventDay(_ context.Context, date *pb.Date) (events *pb.Events, err error) {
	events = new(pb.Events)
	stor, err := s.App.Storage.ListOfEventsDay(date.GetDate().AsTime())
	if err != nil {
		return nil, err
	}
	events.Events = toPbEvents(stor)
	return
}

func (s *CalendarServer) ShowEventWeek(_ context.Context, date *pb.Date) (events *pb.Events, err error) {
	events = new(pb.Events)
	stor, err := s.App.Storage.ListOfEventsWeek(date.GetDate().AsTime())
	if err != nil {
		return nil, err
	}
	events.Events = toPbEvents(stor)
	return
}

func (s *CalendarServer) ShowEventMonth(_ context.Context, date *pb.Date) (events *pb.Events, err error) {
	events = new(pb.Events)
	stor, err := s.App.Storage.ListOfEventsMonth(date.GetDate().AsTime())
	if err != nil {
		return nil, err
	}
	events.Events = toPbEvents(stor)
	return
}

func toPbEvents(s storage.Events) (map[int64]*pb.Event) {
	p := make(map[int64]*pb.Event)
	for _, v := range s {
		ev := &pb.Event{
			ID:          v.ID,
			Title:       v.Title,
			Date:        timestamppb.New(v.Date),
			Duration:    v.Duration,
			Description: v.Description,
			UserID:      v.UserID,
			Reminder:    timestamppb.New(v.Reminder),
		}
		p[v.ID] = ev
	}
	return p
}