package graph_algorithms

import (
	"bufio"
	"log"
	"os"
)

type PrettyPrintGraphvizFile struct {
	PrettyPintGraphviz PrettyPintGraphviz
}

func (f PrettyPrintGraphvizFile) PrettyPrint(graph Graph) {
	graphviz := PrettyPintGraphviz{Graph: graph}
	graphviz.PrettyPrint()

}

func (f PrettyPrintGraphvizFile) Write(outputFilename string) error {
	file, err := os.Create(outputFilename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("error closing file: %v", err)
		}
	}(file)

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(f.PrettyPintGraphviz.PrettyPrint())
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
