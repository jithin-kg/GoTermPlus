package terminal

import (
	"fyne.io/fyne/v2"
)

// This struct will implement the fyne.Layout interface.
type LastItemFullheightVBoxLayout struct{}

func (v *LastItemFullheightVBoxLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	width := float32(0)
	height := float32(0)

	for _, o := range objects {
		childSize := o.MinSize()
		height += childSize.Height
		width = fyne.Max(width, childSize.Width)
	}

	return fyne.NewSize(width, height)
}

func (v *LastItemFullheightVBoxLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	posY := float32(0)
	for _, o := range objects[:len(objects)-1] { // layout all but the last item
		o.Resize(fyne.NewSize(size.Width, o.MinSize().Height))
		o.Move(fyne.NewPos(0, posY))
		posY += o.MinSize().Height
	}

	// Now lay out the last item which should take the remaining space
	last := objects[len(objects)-1]
	last.Resize(fyne.NewSize(size.Width, size.Height-posY))
	last.Move(fyne.NewPos(0, posY))
}

// NewLastItemFullheightVBoxLayout creates a new vertical box layout
func NewLastItemFullheightVBoxLayout() fyne.Layout {
	return &LastItemFullheightVBoxLayout{}
}
