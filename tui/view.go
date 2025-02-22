package tui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/metafates/mangal/icon"
	"github.com/metafates/mangal/style"
	"github.com/metafates/mangal/util"
	"math/rand"
	"strconv"
	"strings"
)

func (b *statefulBubble) View() string {
	switch b.state {
	case scrapersInstallState:
		return b.viewScrapersInstallState()
	case loadingState:
		return b.viewLoading()
	case historyState:
		return b.viewHistory()
	case sourcesState:
		return b.viewSources()
	case searchState:
		return b.viewSearch()
	case mangasState:
		return b.viewMangas()
	case chaptersState:
		return b.viewChapters()
	case confirmState:
		return b.viewConfirm()
	case readState:
		return b.viewRead()
	case downloadState:
		return b.viewDownload()
	case downloadDoneState:
		return b.viewDownloadDone()
	case errorState:
		return b.viewError()
	}

	panic("unknown state")
}

func (b *statefulBubble) viewLoading() string {
	return b.renderLines(
		true,
		[]string{
			style.Title("Loading"),
			"",
			b.spinnerC.View() + b.progressStatus,
		},
	)
}

func (b *statefulBubble) viewHistory() string {
	return listExtraPaddingStyle.Render(b.historyC.View())
}

func (b *statefulBubble) viewSources() string {
	return listExtraPaddingStyle.Render(b.sourcesC.View())
}

func (b *statefulBubble) viewSearch() string {
	return b.renderLines(
		true,
		[]string{
			style.Title("Search Manga"),
			"",
			b.inputC.View(),
		},
	)
}

func (b *statefulBubble) viewMangas() string {
	return listExtraPaddingStyle.Render(b.mangasC.View())
}

func (b *statefulBubble) viewChapters() string {
	return listExtraPaddingStyle.Render(b.chaptersC.View())
}

func (b *statefulBubble) viewConfirm() string {
	return b.renderLines(
		true,
		[]string{
			style.Title("Confirm"),
			"",
			fmt.Sprintf("%s Download %s?", icon.Get(icon.Question), util.Quantity(len(b.selectedChapters), "chapter")),
		},
	)
}

func (b *statefulBubble) viewRead() string {
	var chapterName string

	chapter := b.currentDownloadingChapter
	if chapter != nil {
		chapterName = chapter.Name
	}

	return b.renderLines(
		true,
		[]string{
			style.Title("Reading"),
			"",
			style.Truncate(b.width)(fmt.Sprintf(icon.Get(icon.Progress)+" Downloading %s", style.Magenta(chapterName))),
			"",
			style.Truncate(b.width)(b.spinnerC.View() + b.progressStatus),
		},
	)
}

func (b *statefulBubble) viewDownload() string {
	var chapterName string

	chapter := b.currentDownloadingChapter
	if chapter != nil {
		chapterName = chapter.Name
	}

	return b.renderLines(
		true,
		[]string{
			style.Title("Downloading"),
			"",
			style.Truncate(b.width)(fmt.Sprintf(icon.Get(icon.Progress)+" Downloading %s", style.Magenta(chapterName))),
			"",
			b.progressC.View(),
			"",
			style.Truncate(b.width)(b.spinnerC.View() + b.progressStatus),
		},
	)
}

func (b *statefulBubble) viewDownloadDone() string {
	failed := len(b.failedChapters)
	succeded := len(b.succededChapters)

	var msg string

	{
		temp := strings.Split(util.Quantity(succeded, "chapter"), " ")
		temp[0] = style.Green(temp[0])
		s := strings.Join(temp, " ") + " downloaded"
		f := fmt.Sprintf("%s failed", style.Red(strconv.Itoa(failed)))

		msg = fmt.Sprintf("%s, %s", s, f)
	}

	return b.renderLines(
		true,
		[]string{
			style.Title("Finish"),
			"",
			msg,
		},
	)
}

func (b *statefulBubble) viewError() string {
	errorMsg := util.Wrap(style.Combined(style.Italic, style.Red)(b.lastError.Error()), b.width)
	return b.renderLines(
		true,
		append([]string{
			style.ErrorTitle("Error"),
			"",
			icon.Get(icon.Fail) + " Uggh, something went wrong. Maybe try again?",
			"",
		},
			strings.Split(util.Wrap(style.Italic(b.errorPlot), b.width)+"\n\n"+errorMsg, "\n")...,
		),
	)
}

func (b *statefulBubble) viewScrapersInstallState() string {
	return listExtraPaddingStyle.Render(b.scrapersInstallC.View())
}

var (
	listExtraPaddingStyle = lipgloss.NewStyle().Padding(1, 2, 1, 0)
	paddingStyle          = lipgloss.NewStyle().Padding(1, 2)
)

func (b *statefulBubble) renderLines(addHelp bool, lines []string) string {
	h := len(lines)
	l := strings.Join(lines, "\n")
	if addHelp {
		l += strings.Repeat("\n", b.height-h) + b.helpC.View(b.keymap)
	}

	return paddingStyle.Render(l)
}

func randomPlot() string {
	plots := []string{
		"The universe is a dangerous place. There are many things that can go wrong. This is one of them:",
		"Fighting an endless army of errors and bugs Mangal died a hero. Their last words were:",
		"I used to download stuff without any errors, then I took an arrow to the knee. By arrow I mean this:",
	}

	return plots[rand.Intn(len(plots))]
}
