// Copyright 2011 John Asmuth. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"path/filepath"
	"os"
	"go/ast"
	"go/printer"
	"go/token"
)

func Copy(srcpath, dstpath string) (err os.Error) {
	var srcFile *os.File
	srcFile, err = os.Open(srcpath)
	if err != nil {
		return
	}

	var dstFile *os.File
	dstFile, err = os.Create(dstpath)
	if err != nil {
		return
	}

	io.Copy(dstFile, srcFile)

	dstFile.Close()
	srcFile.Close()

	return
}

func BackupSource(fpath string) (err os.Error) {
	dir, name := filepath.Split(fpath)
	backup := "."+name+".0.gorf"
	err = Copy(fpath, filepath.Join(dir, backup))
	return
}

func Touch(fpath string) (err os.Error) {
	f, err := os.Create(fpath)
	f.Close()
	return
}

func MoveSource(oldpath, newpath string) (err os.Error) {
	fmt.Printf("Moving %s to %s\n", oldpath, newpath)
	
	if _, e := os.Stat(newpath); e == nil {
		BackupSource(newpath)
	}
	
	err = BackupSource(oldpath)
	if err != nil {
		return
	}
	
	dir, file := filepath.Split(newpath)
	
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return
	}
	
	err = Touch(filepath.Join(dir, "."+file+".0.gorfn"))
	if err != nil {
		return
	}
	
	err = Copy(oldpath, newpath)
	if err != nil {
		return
	}
	
	err = os.Remove(oldpath)
	
	return
}

func NewSource(fpath string, file *ast.File) (err os.Error) {
	fmt.Printf("Creating %s\n", fpath)
	
	dir, name := filepath.Split(fpath)
	
	err = Touch(filepath.Join(dir, "."+name+".0.gorfn"))
	if err != nil {
		return
	}
	
	var out io.Writer
	out, err = os.Create(fpath)
	if err != nil {
		return
	}
	
	err = printer.Fprint(out, token.NewFileSet(), file)
	
	return
}

func RewriteSource(fpath string, file *ast.File) (err os.Error) {
	fmt.Printf("Rewriting %s\n", fpath)

	err = BackupSource(fpath)
	if err != nil {
		return
	}

	var out io.Writer
	out, err = os.Create(fpath)
	if err != nil {
		return
	}
	
	err = printer.Fprint(out, token.NewFileSet(), file)
	
	return
}
