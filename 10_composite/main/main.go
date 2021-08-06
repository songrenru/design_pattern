package main

import (
	"fmt"
	composite "github.com/songrenru/design_attern/10_composite"
)

func main() {
	fmt.Println("make dir")
	rootDir := composite.NewComponent(composite.TypeComposite, "root", 0)
	binDir := composite.NewComponent(composite.TypeComposite, "bin", 0)
	tmpDir := composite.NewComponent(composite.TypeComposite, "tmp", 0)
	usrDir := composite.NewComponent(composite.TypeComposite, "usr", 0)
	yukiDir := composite.NewComponent(composite.TypeComposite, "yuki", 0)
	hanakoDir := composite.NewComponent(composite.TypeComposite, "hanako", 0)
	tomuraDir := composite.NewComponent(composite.TypeComposite, "tomura", 0)

	rootDir.Add(binDir)
	rootDir.Add(tmpDir)
	rootDir.Add(usrDir)
	usrDir.Add(yukiDir)
	usrDir.Add(hanakoDir)
	usrDir.Add(tomuraDir)

	rootDir.PrintList()

	fmt.Println("make file")
	viFile := composite.NewComponent(composite.TypeLeaf, "vi", 1)
	latexFile := composite.NewComponent(composite.TypeLeaf, "latex", 2)
	diaryFile := composite.NewComponent(composite.TypeLeaf, "diary.html", 3)
	compositFile := composite.NewComponent(composite.TypeLeaf, "composit.java", 4)
	memoFile := composite.NewComponent(composite.TypeLeaf, "memo.txt", 5)
	gameFile := composite.NewComponent(composite.TypeLeaf, "game.doc", 6)
	junkFile := composite.NewComponent(composite.TypeLeaf, "junk.mail", 7)

	binDir.Add(viFile)
	binDir.Add(latexFile)
	yukiDir.Add(diaryFile)
	yukiDir.Add(compositFile)
	hanakoDir.Add(memoFile)
	tomuraDir.Add(gameFile)
	tomuraDir.Add(junkFile)

	rootDir.PrintList()
}
