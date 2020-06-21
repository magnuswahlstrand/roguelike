package room

import "image"

type Room image.Rectangle

func New(pos image.Point, w, h int) Room {
	return Room(image.Rect(0, 0, w, h).Add(pos))
}

func (r Room) Center() (int, int) {
	centerX := (r.Min.X + r.Max.X) / 2
	centerY := (r.Min.Y + r.Max.Y) / 2
	return centerX, centerY
}

func (r Room) Intersects(other Room) bool {
	return r.Min.X <= other.Max.X && r.Max.X >= other.Min.X &&
		r.Min.Y <= other.Max.Y && r.Max.Y >= other.Min.Y
}
