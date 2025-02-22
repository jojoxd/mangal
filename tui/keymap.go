package tui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/metafates/mangal/style"
)

type statefulKeymap struct {
	state state

	quit, forceQuit,
	selectOne, selectAll, clearSelection,
	remove,
	redownloadFailed,
	confirm,
	openURL,
	read,
	openFolder,
	back,
	filter,
	up, down, left, right,
	top, bottom,
	showHelp key.Binding
}

func (k *statefulKeymap) setState(newState state) {
	k.state = newState
}

func newStatefulKeymap() *statefulKeymap {
	k := key.NewBinding
	keys := key.WithKeys
	help := key.WithHelp

	return &statefulKeymap{
		quit: k(
			keys("q"),
			help("q", "quit"),
		),
		forceQuit: k(
			keys("ctrl+c", "ctrl+d"),
			help("ctrl+c", "quit"),
		),
		remove: k(
			keys("d"),
			help("d", "remove"),
		),
		selectOne: k(
			keys(" "),
			help("space", "select one"),
		),
		selectAll: k(
			keys("ctrl+a", "tab", "*"),
			help("tab", "select all"),
		),
		clearSelection: k(
			keys("backspace"),
			help("backspace", "clear selection"),
		),
		confirm: k(
			keys("enter"),
			help("enter", "confirm"),
		),
		openURL: k(
			keys("o"),
			help("o", "open url"),
		),
		read: k(
			keys("r"),
			help(style.Yellow("r"), style.Yellow("read")),
		),
		redownloadFailed: k(
			keys("r"),
			help("r", "redownload failed"),
		),
		openFolder: k(
			keys("o"),
			help("o", "open folder"),
		),
		back: k(
			keys("esc"),
			help("esc", "back"),
		),
		filter: k(
			keys("/"),
			help("/", "filter"),
		),
		up: k(
			keys("up", "k"),
			help("↑", "up"),
		),
		down: k(
			keys("down", "j"),
			help("↓", "down"),
		),
		left: k(
			keys("left", "h"),
			help("←", "left"),
		),
		right: k(
			keys("right", "l"),
			help("→", "right"),
		),
		top: k(
			keys("g"),
			help("g", "top"),
		),
		bottom: k(
			keys("G"),
			help("G", "bottom"),
		),
		showHelp: k(
			keys("?", "h"),
			help("?", "help"),
		),
	}
}

// help returns short and full help for the state
// TODO: add more information for full help
func (k *statefulKeymap) help() ([]key.Binding, []key.Binding) {
	h := func(bindings ...key.Binding) []key.Binding {
		return bindings
	}

	to2 := func(a []key.Binding) ([]key.Binding, []key.Binding) {
		return a, a
	}

	switch k.state {
	case scrapersInstallState:
		return to2(h(k.confirm, k.openURL))
	case loadingState:
		return to2(h(k.forceQuit, k.back))
	case historyState:
		return to2(h(k.confirm, k.remove, k.back, k.openURL))
	case sourcesState:
		return to2(h(k.confirm))
	case searchState:
		return to2(h(k.confirm, k.forceQuit))
	case mangasState:
		return to2(h(k.confirm, k.back, k.openURL))
	case chaptersState:
		return h(k.read, k.selectOne, k.selectAll, k.confirm, k.back), h(k.read, k.selectOne, k.selectAll, k.clearSelection, k.openURL, k.confirm, k.back)
	case confirmState:
		return to2(h(k.confirm, k.back, k.quit))
	case readState:
		return to2(h(k.back, k.forceQuit))
	case downloadState:
		return to2(h(k.back, k.forceQuit))
	case downloadDoneState:
		return to2(h(k.back, k.quit, k.openFolder, k.redownloadFailed))
	case errorState:
		return to2(h(k.back, k.quit))
	default:
		return to2(h())
	}
}

func (k *statefulKeymap) ShortHelp() []key.Binding {
	short, _ := k.help()
	return short
}

func (k *statefulKeymap) FullHelp() [][]key.Binding {
	_, full := k.help()
	return [][]key.Binding{full}
}

func (k *statefulKeymap) forList() list.KeyMap {
	return list.KeyMap{
		CursorUp:             k.up,
		CursorDown:           k.down,
		NextPage:             k.right,
		PrevPage:             k.left,
		GoToStart:            k.top,
		GoToEnd:              k.bottom,
		Filter:               k.filter,
		ClearFilter:          k.back,
		CancelWhileFiltering: k.back,
		AcceptWhileFiltering: k.confirm,
		ShowFullHelp:         k.showHelp,
		CloseFullHelp:        k.showHelp,
		Quit:                 k.quit,
		ForceQuit:            k.forceQuit,
	}
}
