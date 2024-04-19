package terminal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MyListItemWidget struct {
	widget.Entry
	Title   *widget.Label
	Comment *widget.Label
}

func NewMyListItemWidget(title, comment string) *MyListItemWidget {
	item := &MyListItemWidget{
		Title:   widget.NewLabel(title),
		Comment: widget.NewLabel(comment),
	}
	item.Title.Truncation = fyne.TextTruncateEllipsis
	item.ExtendBaseWidget(item)

	return item
}

func (item *MyListItemWidget) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewBorder(nil, nil, nil, item.Comment, item.Title)
	return widget.NewSimpleRenderer(c)
}
