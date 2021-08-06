package main

import (
	"fmt"
	visitor "github.com/songrenru/design_attern/13_visitor/exercise"
)

func main() {
	fmt.Println("make dir")
	rootDir := visitor.NewEntry(visitor.TypeComposite, "root", 0)
	binDir := visitor.NewEntry(visitor.TypeComposite, "bin", 0)
	tmpDir := visitor.NewEntry(visitor.TypeComposite, "tmp", 0)
	usrDir := visitor.NewEntry(visitor.TypeComposite, "usr", 0)
	yukiDir := visitor.NewEntry(visitor.TypeComposite, "yuki", 0)
	hanakoDir := visitor.NewEntry(visitor.TypeComposite, "hanako", 0)
	tomuraDir := visitor.NewEntry(visitor.TypeComposite, "tomura", 0)

	rootDir.Add(binDir)
	rootDir.Add(tmpDir)
	rootDir.Add(usrDir)
	usrDir.Add(yukiDir)
	usrDir.Add(hanakoDir)
	usrDir.Add(tomuraDir)

	fmt.Println("make file")
	viFile := visitor.NewEntry(visitor.TypeLeaf, "vi", 1)
	latexFile := visitor.NewEntry(visitor.TypeLeaf, "latex", 2)
	diaryFile := visitor.NewEntry(visitor.TypeLeaf, "diary.html", 3)
	compositFile := visitor.NewEntry(visitor.TypeLeaf, "composit.java", 4)
	memoFile := visitor.NewEntry(visitor.TypeLeaf, "memo.txt", 5)
	gameFile := visitor.NewEntry(visitor.TypeLeaf, "game.doc", 6)
	junkFile := visitor.NewEntry(visitor.TypeLeaf, "junk.mail", 7)

	binDir.Add(viFile)
	binDir.Add(latexFile)
	yukiDir.Add(diaryFile)
	yukiDir.Add(compositFile)
	hanakoDir.Add(memoFile)
	tomuraDir.Add(gameFile)
	tomuraDir.Add(junkFile)

	var v visitor.Visitor = &visitor.ListVisitor{}
	rootDir.Accept(v)
}
