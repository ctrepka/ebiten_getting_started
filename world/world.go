package world

type StoryScene struct {
	Scene string
	Title string
}
type Position struct {
	x int
	y int
}
type Item struct {
	Title       string
	Description string
	Position    Position
}

func (p *Position) moveLeft() {
	p.x = p.x - 1
}
func (p *Position) moveUp() {
	p.y = p.y + 1
}
func (p *Position) moveRight() {
	p.x = p.x + 1
}
func (p *Position) moveDown() {
	p.y = p.y - 1
}
