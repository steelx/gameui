package ui

import (
	"image/color"
	"testing"
)

func TestWindow(t *testing.T) {
	w, h := 30, 20
	wnd := NewWindow(w, h)
	wnd.SetTitleColor(color.Black)
	wnd.SetBackgroundColor(color.Black)
	wnd.SetBorderColor(color.White)
	wnd.HideCloseButton(true)

	btn := NewButton(20, 14).SetText("HI")
	btn.Position = Point{X: 5, Y: 3}
	wnd.AddChild(btn)

	// make sure same frame is delivered each time
	for i := 0; i < 10; i++ {
		im := wnd.Draw(0, 0)
		testCompareRender(t, []string{
			"##############################",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,####################,,,,#",
			"#,,,,#,,,,,,,,,,,,,,,,,,#,,,,#",
			"#,,,,#,,,,,,,,,,,,,,,,,,#,,,,#",
			"#,,,,#,66+,56o,,666666o,#,,,,#",
			"#,,,,#,##+,6#5,,######5,#,,,,#",
			"#,,,,#,##+,6#5,,550#65+,#,,,,#",
			"#,,,,#,##+,6#5,,,,O#o,,,#,,,,#",
			"#,,,,#,##5oO#5,,,,O#o,,,#,,,,#",
			"#,,,,#,######5,,,,O#o,,,#,,,,#",
			"#,,,,#,##0O0#5,,++0#5+,,#,,,,#",
			"#,,,,#,##+,6#5,,######5,#,,,,#",
			"#,,,,#,##+,6#5,,######5,#,,,,#",
			"#,,,,#,,,,,,,,,,,,,,,,,,#,,,,#",
			"#,,,,####################,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"##############################",
		}, renderAsText(im))
	}
}

func TestWindowWithTitle(t *testing.T) {
	w, h := 30, 20
	wnd := NewWindow(w, h)
	wnd.SetTitle("WOA")
	wnd.SetTitleColor(color.Black)
	wnd.SetBackgroundColor(color.Black)
	wnd.HideCloseButton(true)

	// make sure same frame is delivered each time
	for i := 0; i < 10; i++ {
		im := wnd.Draw(0, 0)
		testCompareRender(t, []string{
			"##############################",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"##,,,,,,##,,,,####,,,,,,##,,,#",
			"##,,,,,,##,,,,####,,,,,,##,,,#",
			"##,,##,,##,,##,,,,##,,##,,##,#",
			"##,,##,,##,,##,,,,##,,##,,##,#",
			"##,,##,,##,,##,,,,##,,######,#",
			"##,,##,,##,,##,,,,##,,######,#",
			"#,##,,##,,,,,,####,,,,##,,##,#",
			"#,##,,##,,,,,,####,,,,##,,##,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"##############################",
		}, renderAsText(im))
	}
}
