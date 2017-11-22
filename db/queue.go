package db

import (
	"github.com/bcspragu/Radiotation/music"
)

type QueueID struct {
	RoomID RoomID
	UserID UserID
}

func (id QueueID) String() string {
	return string(id.RoomID) + id.UserID.String()
}

type Queue struct {
	ID     QueueID
	Offset int
	Tracks []music.Track
}

func nextTrack(q *Queue) (music.Track, error) {
	if q.Offset < len(q.Tracks) {
		track := q.Tracks[q.Offset]
		q.Offset++
		return track, nil
	}
	return music.Track{}, ErrNoTracksInQueue
}

func (q *Queue) RemoveTrack(delTrack music.Track) {
}

func (q *Queue) HasTracks() bool {
	return len(q.Tracks) > q.Offset
}

func (q *Queue) NumTracks() int {
	return len(q.Tracks)
}

func (q *Queue) HasTrack(track music.Track) bool {
	for _, t := range q.Tracks {
		if t.ID == track.ID {
			return true
		}
	}
	return false
}

func CountTracks(queues []*Queue) int {
	c := 0
	for _, q := range queues {
		c += q.NumTracks()
	}
	return c
}
